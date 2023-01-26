package model

import "time"

type MailTemplate struct {
	ContentHtml        string                    `json:"contentHtml,omitempty"`
	ContentPlain       string                    `json:"contentPlain,omitempty"`
	CreatedAt          time.Time                 `json:"createdAt,omitempty"`
	CustomFields       interface{}               `json:"customFields,omitempty"`
	Description        string                    `json:"description,omitempty"`
	Id                 string                    `json:"id,omitempty"`
	MailTemplateType   *MailTemplateType         `json:"mailTemplateType,omitempty"`
	MailTemplateTypeId string                    `json:"mailTemplateTypeId,omitempty"`
	Media              []MailTemplateMedia       `json:"media,omitempty"`
	SenderName         string                    `json:"senderName,omitempty"`
	Subject            string                    `json:"subject,omitempty"`
	SystemDefault      bool                      `json:"systemDefault,omitempty"`
	Translated         interface{}               `json:"translated,omitempty"`
	Translations       []MailTemplateTranslation `json:"translations,omitempty"`
	UpdatedAt          time.Time                 `json:"updatedAt,omitempty"`
}

type MailTemplateCollection struct {
	EntityCollection

	Data []MailTemplate `json:"data"`
}