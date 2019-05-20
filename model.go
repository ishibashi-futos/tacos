package main

type Response struct {
	ResponseType string `json:"response_type"`
	UserName     string `json:"tacos"`
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
}
