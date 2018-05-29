package entity

type TasksResponse struct {
	Error

	Reminders []*Reminder `json:"Reminders"`
}

type StartupResponse struct {
	Error

	Reminders   []*Reminder   `json:"Reminders"`
	InboxItem   []interface{} `json:"InboxItem"`
	Collections []*Collection `json:"Collections"`
}

type Alarm struct {
	MessageType     string      `json:"messageType"`
	OnDate          []int       `json:"onDate"`
	Measurement     interface{} `json:"measurement"`
	Description     string      `json:"description"`
	GUID            string      `json:"guid"`
	IsLocationBased bool        `json:"isLocationBased"`
	Proximity       interface{} `json:"proximity"`
}

type Recurrence struct {
	WeekStart     string      `json:"weekStart"`
	Freq          string      `json:"freq"`
	Count         interface{} `json:"count"`
	Interval      int         `json:"interval"`
	ByDay         interface{} `json:"byDay"`
	ByMonth       interface{} `json:"byMonth"`
	Until         interface{} `json:"until"`
	FrequencyDays interface{} `json:"frequencyDays"`
	WeekDays      interface{} `json:"weekDays"`
}

type Reminder struct {
	GUID                string      `json:"guid"`
	PGUID               string      `json:"pGuid"`
	Etag                string      `json:"etag"`
	LastModifiedDate    []int       `json:"lastModifiedDate"`
	CreatedDate         []int       `json:"createdDate"`
	CreatedDateExtended int         `json:"createdDateExtended"`
	Priority            int         `json:"priority"`
	CompletedDate       []int       `json:"completedDate"`
	Order               interface{} `json:"order"`
	Title               string      `json:"title"`
	Description         string      `json:"description"`
	DueDate             []int       `json:"dueDate"`
	DueDateIsAllDay     bool        `json:"dueDateIsAllDay"`
	StartDate           interface{} `json:"startDate"`
	StartDateIsAllDay   bool        `json:"startDateIsAllDay"`
	StartDateTz         interface{} `json:"startDateTz"`
	Recurrence          *Recurrence `json:"recurrence"`
	Alarms              []*Alarm    `json:"alarms"`
}

type Collection struct {
	Title               string      `json:"title"`
	GUID                string      `json:"guid"`
	Ctag                string      `json:"ctag"`
	Order               int         `json:"order"`
	Color               string      `json:"color"`
	SymbolicColor       string      `json:"symbolicColor"`
	Enabled             bool        `json:"enabled"`
	EmailNotification   interface{} `json:"emailNotification"`
	CreatedDate         []int       `json:"createdDate"`
	IsFamily            bool        `json:"isFamily"`
	CollectionShareType interface{} `json:"collectionShareType"`
	CreatedDateExtended interface{} `json:"createdDateExtended"`
	Participants        interface{} `json:"participants"`
	CompletedCount      int         `json:"completedCount"`
}
