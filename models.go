package main

type TemplateParameter struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Filename string `json:"filename,omitempty"`
}

type BatchSendMessageRequest struct {
	Number       string              `json:"number"`
	TemplateName string              `json:"templateName"`
	Language     string              `json:"language"`
	Parameters   []TemplateParameter `json:"parameters"`
	Sender       string              `json:"sender,omitempty"`
}

type BatchSendMessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
