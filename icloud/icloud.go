package icloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/ww24/go-icloud-reminder/entity"
)

const (
	iCloudValidateEndpoint = "https://setup.icloud.com/setup/ws/1/validate"
	iCloudClientTimeout    = time.Second * 30

	userTZ = "Asia/Tokyo"
	lang   = "ja-jp"
)

const (
	iCloudAPIStatusActive = "active"
)

var (
	ErrAPIUnavailable = errors.New("iCloud API Unavailable")
)

type XAppleWebauth struct {
	User  string
	Token string
}

type Config struct {
	ID            string
	Password      string
	XAppleWebauth *XAppleWebauth
}

type iCloud struct {
	config     *Config
	dsid       string
	client     *http.Client
	validation *entity.ValidateResponse
}

func (i *iCloud) request(method, uri string, body io.Reader, entity interface{}) error {
	query := url.Values{}
	if i.dsid != "" {
		query.Set("dsid", i.dsid)
	}
	query.Set("usertz", userTZ)
	query.Set("lang", lang)

	uri += "?" + query.Encode()
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
	}

	req.Header.Add("Origin", "https://www.icloud.com")
	req.Header.Add("Content-Type", "application/json")

	resp, err := i.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	respBody := io.TeeReader(resp.Body, buf)

	decoder := json.NewDecoder(respBody)
	err = decoder.Decode(entity)
	if err != nil {
		log.Println("DEBUG:", resp.Status)
		log.Println("ERROR:", err)
		log.Println(string(buf.Bytes()))
		return err
	}
	return nil
}

func (i *iCloud) setCredentialToEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	i.client.Jar.SetCookies(u, []*http.Cookie{{
		Name:  "X-APPLE-WEBAUTH-USER",
		Value: i.config.XAppleWebauth.User,
	}, {
		Name:  "X-APPLE-WEBAUTH-TOKEN",
		Value: i.config.XAppleWebauth.Token,
	}})
	return nil
}

func (i *iCloud) login() error {
	payload, err := json.Marshal(map[string]interface{}{
		"apple_id":       i.config.ID,
		"password":       i.config.Password,
		"extended_login": false,
	})
	if err != nil {
		return err
	}

	body := bytes.NewReader(payload)
	i.validation = new(entity.ValidateResponse)
	err = i.request("POST", iCloudValidateEndpoint, body, i.validation)
	if err != nil {
		return err
	}

	i.dsid = i.validation.DsInfo.Dsid

	return nil
}

func (i *iCloud) NewReminder() (Reminder, error) {
	if i.validation.Webservices.Reminders.Status != iCloudAPIStatusActive {
		log.Printf("iCloud API Reminder Status: %v", i.validation.Webservices.Reminders.Status)
		return nil, ErrAPIUnavailable
	}
	endpoint := i.validation.Webservices.Reminders.URL
	err := i.setCredentialToEndpoint(endpoint)
	if err != nil {
		return nil, err
	}
	api := &reminder{
		endpoint: endpoint,
		i:        i,
	}
	return api, nil
}

// New returns ICloudService.
func New(c *Config) (Service, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	iCloud := &iCloud{
		config: c,
		client: &http.Client{
			Timeout: iCloudClientTimeout,
			Jar:     jar,
		},
	}

	err = iCloud.setCredentialToEndpoint(iCloudValidateEndpoint)
	if err != nil {
		return nil, err
	}

	err = iCloud.login()
	if err != nil {
		return nil, err
	}

	return iCloud, nil
}
