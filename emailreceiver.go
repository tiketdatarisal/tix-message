package tixmessage

type EmailReceiver struct {
	BCC []string `json:"bcc"`
	CC  []string `json:"cc"`
	To  []string `json:"to"`
}
