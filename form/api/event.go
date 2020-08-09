package api

import "time"

type EventCreateIn struct {
	EventName              string    `json:"event_name"`
	EventTitle             string    `json:"event_title"`
	EventDescription       string    `json:"event_description"`
	EventLocation          string    `json:"event_location"`
	EventCost              string    `json:"event_cost"`
	EventTime              time.Time `json:"event_time"`
	EventCompanyName       string    `json:"event_company_name"`
	EventPopulation        int       `json:"event_population"`
	Event_president        string    `json:"event_president"`
	EventPresidentBirthDay time.Time `json:"event_president_birth_day"`
	EventMotto             string    `json:"event_motto"`
	Position               string    `json:"position"`
	GetPoint               int       `json:"get_point"`
	PayPoint               int       `json:"pay_point"`
	Finish                 bool      `json:"finish"`
	ContentOne             string    `json:"content_one"`
	ContentTwo             string    `json:"content_two"`
	ContentThree           string    `json:"content_three"`
	ContentFour            string    `json:"content_four"`
	ContentFive            string    `json:"content_five"`
	Tag                    Tag       `json:"tag"`
}

type Tag struct {
	TagName string `json:"tag_name"`
	TagInfo string `json:"tag_info"`
	TagBody string `json:"tag_body"`
}
