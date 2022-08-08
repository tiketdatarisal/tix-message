package main

import tixmessage "github.com/tiketdatarisal/tix-message"

func main() {
	client, err := tixmessage.NewClient("https://your_host.com")
	if err != nil {
		panic(err)
	}

	p := tixmessage.Param{
		TemplateName: "your_template_name",
		EmailReceiver: tixmessage.EmailReceiver{
			To: []string{"someone@tiket.com"},
		},
		SubjectParams: map[string]string{
			"subject": "hello world",
		},
		Language: tixmessage.LanguageEnglish,
	}

	err = client.SendEmail(&p)
	if err != nil {
		panic(err)
	}
}
