package tixmessage

import "errors"

type Language string
type AttachmentType string

const (
	headerContentType    = `Content-Type`
	headerAccept         = `Accept`
	headerAcceptLanguage = `Accept-Language`
	mimeTypeJson         = `application/json`

	sendEmailPath = `tix-message/v2/emails/send-queue`
)

const (
	LanguageEnglish   Language = `EN`
	LanguageIndonesia Language = `ID`
)

const (
	AttachmentTypeBase64 AttachmentType = "BASE_64"
	AttachmentTypeUrl    AttachmentType = "URL"
)

var (
	isValidLanguage = map[Language]bool{
		LanguageIndonesia: true,
		LanguageEnglish:   true,
	}

	ErrInvalidHost           = errors.New("invalid host name, please recheck")
	ErrParamRequired         = errors.New("email param is required")
	ErrEmailReceiverRequired = errors.New("at least one email recipient in to/cc/bcc is required")
	ErrTemplateNameRequired  = errors.New("template name is required")
)
