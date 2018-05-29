package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ww24/go-icloud-reminder/icloud"
)

const (
	iCloudValidateEndpoint = "https://setup.icloud.com/setup/ws/1/validate"
)

var (
	iCloudID           = os.Getenv("ICLOUD_ID")
	iCloudPW           = os.Getenv("ICLOUD_PW")
	xAppleWebauthUser  = os.Getenv("X_APPLE_WEBAUTH_USER")
	xAppleWebauthToken = os.Getenv("X_APPLE_WEBAUTH_TOKEN")
)

func main() {
	iCloud, err := icloud.New(&icloud.Config{
		ID:       iCloudID,
		Password: iCloudPW,
		XAppleWebauth: &icloud.XAppleWebauth{
			User:  xAppleWebauthUser,
			Token: xAppleWebauthToken,
		},
	})
	if err != nil {
		log.Fatalln("Error:", err)
	}

	reminder, err := iCloud.NewReminder()
	if err != nil {
		log.Fatalln("Error:", err)
	}

	startupRes, err := reminder.Startup()
	if err != nil {
		log.Fatalln("Error:", err)
	}

	for _, col := range startupRes.Collections {
		fmt.Printf("\n## title: %s, id: %s\n", col.Title, col.GUID)

		taskRes, err := reminder.GetTasks(col.GUID)
		if err != nil {
			log.Fatalln("Error:", err)
		}
		for _, task := range taskRes.Reminders {
			fmt.Printf("%s: %+v\n\n", task.Title, task)
		}
	}
}
