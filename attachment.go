package tixmessage

import "encoding/base64"

type Attachment struct {
	Data     string         `json:"data"`
	Filename string         `json:"filename"`
	Type     AttachmentType `json:"type"`
}

// NewFileStreamAttachment returns a new file stream attachment.
func NewFileStreamAttachment(filename string, data []byte) Attachment {
	return Attachment{
		Data:     base64.StdEncoding.EncodeToString(data),
		Filename: filename,
		Type:     AttachmentTypeBase64,
	}
}

// NewURLAttachment returns a new URL attachment.
func NewURLAttachment(filename, url string) Attachment {
	return Attachment{
		Data:     url,
		Filename: filename,
		Type:     AttachmentTypeUrl,
	}
}
