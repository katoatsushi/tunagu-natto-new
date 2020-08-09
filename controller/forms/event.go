package forms

import {
	models,
	api
}

func NewEventForCreate(in *api.EventCreateIn) *models.Event {
	var subTags []models.EventSubTag
	for _, subTagID := range in.EventSubTagIDs {
		subTag := models.SubTag{ID: subTagID}
		subTags = append(subTags, models.EventSubTag{
			SubTagID: subTagID,
			SubTag:   subTag,
		})
	}

	tag := models.Tag{ID: in.TagID}
	model := &models.Event{
		EventName:              in.EventName,
		EventTitle:             in.EventTitle,
		EventDescription:       in.EventDescription,
		EventLocation:          in.EventLocation,
		EventCost:              in.EventCost,
		EventTime:              in.EventTime,
		EventCompanyName:       in.EventCompanyName,
		EventPopulation:        in.EventPopulation,
		Event_president:        in.Event_president,
		EventPresidentBirthDay: in.EventPresidentBirthDay,
		EventMotto:             in.EventMotto,
		Position:               in.Position,
		GetPoint:               in.GetPoint,
		PayPoint:               in.PayPoint,
		Finish:                 in.Finish,
		ContentOne:             in.ContentOne,
		ContentTwo:             in.ContentTwo,
		ContentThree:           in.ContentThree,
		ContentFour:            in.ContentFour,
		ContentFive:            in.ContentFive,
		Tag:                    tag,
	}

	return model
}
