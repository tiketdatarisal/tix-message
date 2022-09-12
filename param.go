package tixmessage

type Param struct {
	TemplateName  string            `json:"templateName"`
	EmailReceiver EmailReceiver     `json:"emailReceiver"`
	SubjectParams map[string]string `json:"subjectParams,omitempty"`
	BodyParams    map[string]string `json:"bodyParams,omitempty"`
	Attachment    []Attachment      `json:"attachment,omitempty"`
	Language      Language          `json:"-"`
}

// AddSubjectParam add subject param. Warning this will replace old key with new value.
func (p *Param) AddSubjectParam(key, value string) {
	if p.SubjectParams == nil {
		p.SubjectParams = map[string]string{}
	}

	p.SubjectParams[key] = value
}

// AddBodyParam add body param. Warning this will replace old key with new value.
func (p *Param) AddBodyParam(key, value string) {
	if p.BodyParams == nil {
		p.BodyParams = map[string]string{}
	}

	p.BodyParams[key] = value
}

// AddAttachment add an attachment to an email.
func (p *Param) AddAttachment(attachment Attachment) {
	p.Attachment = append(p.Attachment, attachment)
}
