package tixmessage

import "errors"

const (
	LanguageEnglish   = `EN`
	LanguageIndonesia = `ID`

	headerContentType    = `Content-Type`
	headerAccept         = `Accept`
	headerAcceptLanguage = `Accept-Language`
	mimeTypeJson         = `application/json`

	sendEmailPath = `tix-message/v2/emails/send-queue`
)

var (
	isValidLanguage = map[string]bool{
		LanguageIndonesia: true,
		LanguageEnglish:   true,
	}

	ErrInvalidHost           = errors.New("invalid host name, please recheck")
	ErrParamRequired         = errors.New("email param is required")
	ErrEmailReceiverRequired = errors.New("at least one email recipient in to/cc/bcc is required")
	ErrTemplateNameRequired  = errors.New("template name is required")
)
