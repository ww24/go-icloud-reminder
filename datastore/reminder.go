// +build appengine

package datastore

import (
	"context"

	"google.golang.org/appengine/datastore"

	"icloud"
)

const (
	reminderKind = "Reminder"
)

type Reminder struct {
	GUID                string             `datastore:"guid"`
	PGUID               string             `datastore:"pGuid"`
	Etag                string             `datastore:"etag"`
	LastModifiedDate    []int              `datastore:"lastModifiedDate"`
	CreatedDate         []int              `datastore:"createdDate"`
	CreatedDateExtended int                `datastore:"createdDateExtended"`
	Priority            int                `datastore:"priority"`
	CompletedDate       []int              `datastore:"completedDate"`
	Order               interface{}        `datastore:"-"` // datastore: unsupported struct field type: interface{}
	Title               string             `datastore:"title"`
	Description         string             `datastore:"description"`
	DueDate             []int              `datastore:"dueDate"`
	DueDateIsAllDay     bool               `datastore:"dueDateIsAllDay"`
	StartDate           interface{}        `datastore:"-"` // datastore: unsupported struct field type: interface{}
	StartDateIsAllDay   bool               `datastore:"startDateIsAllDay"`
	StartDateTz         interface{}        `datastore:"-"` // datastore: unsupported struct field type: interface{}
	Recurrence          *icloud.Recurrence `datastore:"-"`
	Alarms              []*icloud.Alarm    `datastore:"-"`
}

type Reminders []*Reminder

func (s *Reminder) Save(ctx context.Context) error {
	k := datastore.NewKey(ctx, reminderKind, s.GUID, 0, nil)
	if _, err := datastore.Put(ctx, k, s); err != nil {
		return err
	}
	return nil
}

func (s *Reminders) Get(ctx context.Context) error {
	*s = make(Reminders, 0, 64)
	q := datastore.NewQuery(reminderKind).
		Order("-order")
	for t := q.Run(ctx); ; {
		c := new(Reminder)
		_, err := t.Next(c)
		if err == datastore.Done {
			break
		}
		if err != nil {
			return err
		}
		*s = append(*s, c)
	}
	return nil
}
