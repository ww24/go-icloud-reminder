package icloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const (
	iCloudValidateEndpoint = "https://setup.icloud.com/setup/ws/1/validate"
	iCloudLoginEndpoint    = "https://setup.icloud.com/setup/ws/1/accountLogin"
	iCloudClientTimeout    = time.Second * 30

	userTZ                   = "Asia/Tokyo"
	lang                     = "ja-jp"
	cookieXAppleWebauthUser  = "X-APPLE-WEBAUTH-USER"
	cookieXAppleWebauthToken = "X-APPLE-WEBAUTH-TOKEN"
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
	Code          string // 2FA verification code
	XAppleWebauth *XAppleWebauth
	Client        *http.Client
}

type iCloud struct {
	config     *Config
	dsid       string
	validation *ValidateResponse
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

	resp, err := i.config.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	respBody := io.TeeReader(resp.Body, buf)
	decoder := json.NewDecoder(respBody)

	if resp.StatusCode >= 400 {
		e := new(Error)
		if err = decoder.Decode(e); err != nil {
			return fmt.Errorf("%s: %+v\nraw_body=%s", resp.Status, err, string(buf.Bytes()))
		}
		return fmt.Errorf("%s: %+v\nraw_body=%s", resp.Status, e, string(buf.Bytes()))
	}

	err = decoder.Decode(entity)
	if err != nil {
		return fmt.Errorf("%s: %+v\nraw_body=%s", resp.Status, err, string(buf.Bytes()))
	}

	for _, cookie := range resp.Cookies() {
		switch cookie.Name {
		case cookieXAppleWebauthUser:
			i.config.XAppleWebauth.User = cookie.Value
		case cookieXAppleWebauthToken:
			i.config.XAppleWebauth.Token = cookie.Value
		default:
			continue
		}
	}

	return nil
}

func (i *iCloud) setCredentialToEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	i.config.Client.Jar.SetCookies(u, []*http.Cookie{{
		Name:  cookieXAppleWebauthUser,
		Value: i.config.XAppleWebauth.User,
	}, {
		Name:  cookieXAppleWebauthToken,
		Value: i.config.XAppleWebauth.Token,
	}})
	return nil
}

func (i *iCloud) GetCredentials() *XAppleWebauth {
	return i.config.XAppleWebauth
}

func (i *iCloud) validate() error {
	payload, err := json.Marshal(&loginRequest{
		AppleID:       i.config.ID,
		Password:      i.config.Password,
		ExtendedLogin: false,
	})
	if err != nil {
		return err
	}

	body := bytes.NewReader(payload)
	i.validation = new(ValidateResponse)
	err = i.request("POST", iCloudValidateEndpoint, body, i.validation)
	if err != nil {
		return err
	}

	i.dsid = i.validation.DsInfo.Dsid

	return nil
}

func (i *iCloud) login(code string) error {
	payload, err := json.Marshal(&loginRequest{
		AppleID:       i.config.ID,
		Password:      i.config.Password + code,
		ExtendedLogin: true,
	})
	if err != nil {
		return err
	}

	body := bytes.NewReader(payload)
	i.validation = new(ValidateResponse)
	err = i.request("POST", iCloudLoginEndpoint, body, i.validation)
	if err != nil {
		return err
	}

	i.dsid = i.validation.DsInfo.Dsid

	return nil
}

func (i *iCloud) NewReminder() (ReminderService, error) {
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

func newICloud(c *Config) (*iCloud, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	if c.Client == nil {
		c.Client = &http.Client{}
	}

	c.Client.Timeout = iCloudClientTimeout
	c.Client.Jar = jar

	if c.XAppleWebauth == nil {
		c.XAppleWebauth = &XAppleWebauth{}
	}

	i := &iCloud{
		config: c,
	}

	return i, nil
}

// New returns ICloudService.
func New(c *Config) (Service, error) {
	i, err := newICloud(c)
	if err != nil {
		return nil, err
	}

	err = i.setCredentialToEndpoint(iCloudValidateEndpoint)
	if err != nil {
		return nil, err
	}

	err = i.validate()
	if err != nil {
		return nil, err
	}

	return i, nil
}

// Login returns ICloudService.
func Login(c *Config) (Service, error) {
	i, err := newICloud(c)
	if err != nil {
		return nil, err
	}

	err = i.login(c.Code)
	if err != nil {
		return nil, err
	}

	return i, nil
}
