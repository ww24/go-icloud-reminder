// +build appengine

package setup

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"datastore"
	"icloud"
)

const (
	tokenCookieKey            = "MSG"
	tokenCookieExpireDuration = 30 * time.Minute
)

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	switch r.Method {
	case http.MethodGet:
		tokenMsg, err := RandomBytes()
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		csrfToken := GenerateToken(tokenMsg)
		http.SetCookie(w, &http.Cookie{
			Name:    tokenCookieKey,
			Value:   base64.RawStdEncoding.EncodeToString(tokenMsg),
			Path:    r.URL.Path,
			Domain:  r.URL.Hostname(),
			Expires: time.Now().Add(tokenCookieExpireDuration),
		})
		err = render(w, &tmplData{
			CSRFToken: csrfToken,
		})
		if err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		form := parseForm(r.PostForm)

		// verify CSRF token
		var tokenMsg []byte
		for _, cookie := range r.Cookies() {
			if cookie.Name == tokenCookieKey {
				var err error
				tokenMsg, err = base64.RawStdEncoding.DecodeString(cookie.Value)
				if err != nil {
					http.Error(w, "Error: "+err.Error(), http.StatusInternalServerError)
					return
				}
				break
			}
		}

		if !VerifyToken(tokenMsg, form.CSRFToken) {
			http.Error(w, "Error: invalid CSRF token", http.StatusBadRequest)
			return
		}

		iCloud, err := icloud.Login(&icloud.Config{
			ID:       form.ICloudID,
			Password: form.ICloudPassword,
			Code:     form.VerificationCode,
			Client:   urlfetch.Client(ctx),
		})
		if err != nil {
			http.Error(w, "Login Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		credential := iCloud.GetCredentials()
		secret := datastore.Secret{
			AppleWebauthUser:  credential.User,
			AppleWebauthToken: credential.Token,
		}
		if err := secret.Save(ctx); err != nil {
			http.Error(w, "Failed to save secret after login: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if !secret.Valid() {
			http.Redirect(w, r, "/setup", http.StatusSeeOther)
		}

		fmt.Fprintf(w, "Success: %#v", form)
	case http.MethodHead, http.MethodOptions:
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
