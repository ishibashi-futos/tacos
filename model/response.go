package model

type Response struct {
	ResponseType string `json:"response_type"`
	UserName     string `json:"user_name"`
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
}
