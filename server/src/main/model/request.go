package model

type RequestData struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type APIResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}
