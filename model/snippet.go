package model

import "time"

type Snippet struct {
	Author         string      `json:"author,omitempty"`
	CreatedAt      time.Time   `json:"createdAt,omitempty"`
	CustomFields   interface{} `json:"customFields,omitempty"`
	Id             string      `json:"id,omitempty"`
	Set            *SnippetSet `json:"set,omitempty"`
	SetId          string      `json:"setId,omitempty"`
	TranslationKey string      `json:"translationKey,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt,omitempty"`
	Value          string      `json:"value,omitempty"`
}
