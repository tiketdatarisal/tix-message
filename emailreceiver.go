package tixmessage

type EmailReceiver struct {
	BCC []string `json:"bcc"`
	CC  []string `json:"cc"`
	To  []string `json:"to"`
}

// IsValid return an error when there is no To, CC, and BCC.
func (e *EmailReceiver) IsValid() error {
	if len(e.To) == 0 && len(e.CC) == 0 && len(e.BCC) == 0 {
		return ErrEmailReceiverRequired
	}

	return nil
}

// AddTo add target recipient.
func (e *EmailReceiver) AddTo(recipient string) {
	e.To = append(e.To, recipient)
}

// AddCC add target carbon copy.
func (e *EmailReceiver) AddCC(recipient string) {
	e.CC = append(e.CC, recipient)
}

// AddBCC add target blank carbon copy.
func (e *EmailReceiver) AddBCC(recipient string) {
	e.BCC = append(e.BCC, recipient)
}
