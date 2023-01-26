package model

type ThemeSalesChannel struct {
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
	Theme          *Theme        `json:"theme,omitempty"`
	ThemeId        string        `json:"themeId,omitempty"`
}

type ThemeSalesChannelCollection struct {
	EntityCollection

	Data []ThemeSalesChannel `json:"data"`
}