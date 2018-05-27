package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

const (
	iCloudValidateEndpoint = "https://setup.icloud.com/setup/ws/1/validate"
)

var (
	iCloudID           = os.Getenv("ICLOUD_ID")
	iCloudPW           = os.Getenv("ICLOUD_PW")
	dsid               = os.Getenv("DSID")
	xAppleWebauthUser  = os.Getenv("X_APPLE_WEBAUTH_USER")
	xAppleWebauthToken = os.Getenv("X_APPLE_WEBAUTH_TOKEN")
)

func main() {
	err := login(iCloudID, iCloudPW)
	if err != nil {
		log.Fatalln("Error:", err)
	}
}

func login(id, pw string) error {
	payload := map[string]interface{}{
		"apple_id":       id,
		"password":       pw,
		"extended_login": false,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	u, err := url.Parse(iCloudValidateEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	jar.SetCookies(u, []*http.Cookie{{
		Name:  "X-APPLE-WEBAUTH-USER",
		Value: xAppleWebauthUser,
	}, {
		Name:  "X-APPLE-WEBAUTH-TOKEN",
		Value: xAppleWebauthToken,
	}})

	client := &http.Client{
		Timeout: time.Second * 30,
		Jar:     jar,
	}

	query := url.Values{
		"dsid": {dsid},
	}
	uri := fmt.Sprintf("%s?%s", iCloudValidateEndpoint, query.Encode())
	log.Println("uri:", uri)
	req, err := http.NewRequest("POST", uri, bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Add("Origin", "https://www.icloud.com")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	res := make(map[string]interface{})
	err = decoder.Decode(&res)
	if err != nil {
		return err
	}

	fmt.Println("Header:")
	for k, v := range resp.Header {
		fmt.Printf("%s: %+v\n", k, v)
	}

	fmt.Println(resp.Status)
	fmt.Printf("res: %+v\n", res)

	return nil
}
