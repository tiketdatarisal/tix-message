package tixmessage

type Attachment struct {
	Data     string `json:"data"`
	Filename string `json:"filename"`
	Type     string `json:"type"`
}
