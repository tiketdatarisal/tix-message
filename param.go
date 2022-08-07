package tixmessage

type Param struct {
	TemplateName  string            `json:"templateName"`
	EmailReceiver EmailReceiver     `json:"emailReceiver"`
	SubjectParams map[string]string `json:"subjectParams,omitempty"`
	BodyParams    map[string]string `json:"bodyParams,omitempty"`
	Language      string            `json:"-"`
}
