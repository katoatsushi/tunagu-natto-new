package models

import (
	"time"
)

type Event struct {
	ID uint64 `json:"id"`
	// イベント名
	EventName string `json:"event_name"`
	//
	EventTitle string `json:"event_title"`
	//
	EventDescription string `json:"event_description"`
	EventLocation    string `json:"event_location"`
	//
	EventCost string `json:"event_cost"`
	//
	EventTime time.Time `json:"event_time"`
	//
	EventCompanyName string `json:"event_company_name"`
	//
	EventPopulation int `json:"event_population"`
	//
	Event_president string `json:"event_president"`
	//
	EventPresidentBirthDay time.Time `json:"event_president_birth_day"`
	//
	EventMotto string `json:"event_motto"`
	//
	Position string `json:"position"`
	//
	GetPoint int `json:"get_point"`
	//
	PayPoint int `json:"pay_point"`
	//
	Finish bool `json:"finish"`
	//
	ContentOne string `json:"content_one"`
	//
	ContentTwo string `json:"content_two"`
	//
	ContentThree string `json:"content_three"`
	//
	ContentFour string `json:"content_four"`
	//
	ContentFive string `json:"content_five"`
	// EventApplies []EventApply `json:"event_Aapplies"`
	Tag       Tag       `json:"tag"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Tag struct {
	ID      uint64 `json:"id"`
	TagName string `json:"tag_name"`
	TagInfo string `json:"tag_info"`
	TagBody string `json:"tag_body"`
}

type EventStyle struct {
	ID      uint64 `json:"id"`
	EventID uint64 `gorm:"not null;unique_index:idx_event_style" json:"event_id"`
	StyleID uint64 `gorm:"not null;unique_index:idx_event_style" json:"style_id"`
}

type EventApply struct {
	ID        uint64 `json:"id"`
	EventID   uint64 `json:"event_id"`
	StudentID uint64 `json:"student_id"`
	// 参加費を払ったか?
	HasPaid bool `json:"has_paid"`
	// 参加費を保持ポイントで払ったか?
	PayPoint bool `json:"pay_point"`
	// 参加費を現金で払ったか?
	PayCash bool `json:"pay_cash"`
	// イベントに参加したか？
	Checke bool `json:"checke "`
}
