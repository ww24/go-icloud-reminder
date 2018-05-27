package entity

type TasksResponse struct {
	Error

	Reminders []*Reminder `json:"Reminders"`
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
	Recurrence          interface{} `json:"recurrence"`
	Alarms              []*Alarm    `json:"alarms"`
}
