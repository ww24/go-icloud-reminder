package setup

import (
	"io"
	"log"
	"net/url"
	"text/template"
)

const (
	csrfTokenKey        = "csrf_token"
	iCloudIDKey         = "icloud_id"
	iCloudPasswordKey   = "icloud_pw"
	verificationCodeKey = "verification_code"
)

const setupTemplate = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>setup</title>
<style>
.setup-form > input {
  display: block;
  width: 20em;
}
</style>
</head>
<body>
  <h1>Setup</h1>
  <form method="POST" class="setup-form">
    <input type="hidden" name="{{ .CSRFTokenKey }}" value="{{ .CSRFToken }}">
    <input type="text" name="{{ .ICloudIDKey }}" placeholder="iCloud ID">
    <input type="password" name="{{ .ICloudPasswordKey }}" placeholder="iCloud Password">
    <input type="text" name="{{ .VerificationCodeKey }}" placeholder="Verification code">
    <input type="submit" value="Save">
  </form>
</body>
</html>`

var (
	tmpl *template.Template
)

func init() {
	var err error
	tmpl, err = template.New("setup").Parse(setupTemplate)
	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}
}

type tmplData struct {
	CSRFToken           string
	CSRFTokenKey        string
	ICloudIDKey         string
	ICloudPasswordKey   string
	VerificationCodeKey string
}

type formData struct {
	CSRFToken        string
	ICloudID         string
	ICloudPassword   string
	VerificationCode string
}

func render(w io.Writer, d *tmplData) error {
	d.CSRFTokenKey = csrfTokenKey
	d.ICloudIDKey = iCloudIDKey
	d.ICloudPasswordKey = iCloudPasswordKey
	d.VerificationCodeKey = verificationCodeKey
	return tmpl.Execute(w, d)
}

func parseForm(v url.Values) *formData {
	f := new(formData)
	f.CSRFToken = v.Get(csrfTokenKey)
	f.ICloudID = v.Get(iCloudIDKey)
	f.ICloudPassword = v.Get(iCloudPasswordKey)
	f.VerificationCode = v.Get(verificationCodeKey)
	return f
}
