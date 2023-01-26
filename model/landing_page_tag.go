package model

type LandingPageTag struct {
	LandingPage          *LandingPage `json:"landingPage,omitempty"`
	LandingPageId        string       `json:"landingPageId,omitempty"`
	LandingPageVersionId string       `json:"landingPageVersionId,omitempty"`
	Tag                  *Tag         `json:"tag,omitempty"`
	TagId                string       `json:"tagId,omitempty"`
}
