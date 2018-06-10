// +build appengine

package cron

import (
	"context"
	"fmt"
	"net/http"

	"datastore"
	"icloud"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

// SyncReminder dispatch sync reminder event.
func SyncReminder(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	secret := new(datastore.Secret)
	if err := secret.Get(ctx); err != nil {
		http.Error(w, "Failed to get secrets: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if secret == nil || !secret.Valid() {
		http.Error(w, "Failed to get secrets: secret is empty", http.StatusInternalServerError)
		return
	}

	sync(ctx, w, secret)
}

func sync(ctx context.Context, w http.ResponseWriter, secret *datastore.Secret) {
	iCloud, err := icloud.New(&icloud.Config{
		XAppleWebauth: &icloud.XAppleWebauth{
			User:  secret.AppleWebauthUser,
			Token: secret.AppleWebauthToken,
		},
		Client: urlfetch.Client(ctx),
	})
	if err != nil {
		http.Error(w, "Failed to initialize iCloud: "+err.Error(), http.StatusInternalServerError)
		return
	}

	reminder, err := iCloud.NewReminder()
	if err != nil {
		http.Error(w, "Failed to initialize iCloud Reminder: "+err.Error(), http.StatusInternalServerError)
		return
	}

	startupRes, err := reminder.Startup()
	if err != nil {
		http.Error(w, "Failed to startup iCloud Reminder: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for _, col := range startupRes.Collections {
		fmt.Fprintf(w, "\n## title: %s, id: %s\n", col.Title, col.GUID)

		taskRes, err := reminder.GetTasks(col.GUID)
		if err != nil {
			http.Error(w, "Failed to get tasks: "+err.Error(), http.StatusInternalServerError)
			return
		}
		for _, task := range taskRes.Reminders {
			reminder := (*datastore.Reminder)(task)
			if err := reminder.Save(ctx); err != nil {
				http.Error(w, "Failed to save reminder: "+err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "%s: %+v\n\n", task.Title, task)
		}
	}

	credential := iCloud.GetCredentials()
	secret.AppleWebauthUser = credential.User
	secret.AppleWebauthToken = credential.Token
	if err := secret.Save(ctx); err != nil {
		http.Error(w, "Failed to save secret after sync reminder: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
