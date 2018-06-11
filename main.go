// +build appengine

package main

import (
	"net/http"

	"google.golang.org/appengine"

	"cron"
	"setup"
)

func main() {
	http.HandleFunc("/setup", setup.Handler)
	http.HandleFunc("/tasks/sync_reminder", cron.SyncReminder)
	appengine.Main()
}
