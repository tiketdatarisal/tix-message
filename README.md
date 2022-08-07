# TIX-Message Client
TIX-Message client library for Go. This client library used by internal project to access TIX-Message service.

## Feature
* Send email message.

## Installation
```
go get github.com/tiketdatarisal/tix-message
```

## Usage
The following samples will assist you to use this library.
```go
import "github.com/tiketdatarisal/tix-message"
```
Create a new instance of TIX-Message client.
```go
client, err := tixmessage.NewClient("https://your-tix-message-host.com")
if err != nil {
	panic(err)
}
```
Initialize `Param` and call `SendEmail` method. Note that `TemplateName` and `EmailReceiver` are required fields.
```go
p := Param {
	TemplateName: "data_send_sample_email",
	EmailReceiver: {
		To: []string { "somebody@tiket.com" },
    },
}

err = client.SendEmail(&p)
if err != nil {
	panic(err)
}
```
That's all folks, easy peasy right?