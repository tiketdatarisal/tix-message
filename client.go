package tixmessage

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	client *resty.Client
	host   string
}

// NewClient Instantiate a new TIX-Message client.
func NewClient(host string) (*Client, error) {
	if _, err := url.Parse(host); err != nil {
		return nil, ErrInvalidHost
	}

	return &Client{
		client: resty.New(),
		host:   host,
	}, nil
}

func (c *Client) SendEmail(p *Param) error {
	if p == nil {
		return ErrParamRequired
	}

	if p.TemplateName == "" {
		return ErrTemplateNameRequired
	}

	if len(p.EmailReceiver.To) == 0 && len(p.EmailReceiver.CC) == 0 && len(p.EmailReceiver.BCC) == 0 {
		return ErrEmailReceiverRequired
	}

	if !isValidLanguage[p.Language] {
		p.Language = LanguageEnglish
	}

	u, err := url.Parse(c.host)
	if err != nil {
		return ErrInvalidHost
	}

	u.Path = path.Join(u.Path, sendEmailPath)
	targetUrl := u.String()

	res, err := c.client.R().
		SetHeader(headerContentType, mimeTypeJson).
		SetHeader(headerAccept, mimeTypeJson).
		SetHeader(headerAcceptLanguage, p.Language).
		SetBody(p).
		Post(targetUrl)
	if err != nil {
		return err
	}

	if code := res.StatusCode(); code >= 300 {
		return errors.New(http.StatusText(code))
	}

	return nil
}
