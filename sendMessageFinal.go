package main

// START CUT 1 OMIT
import (
	"errors"
	"log"
)

var errEmptyMessage = errors.New("the message was empty")

type MessageSender interface {
	sendMessage(string) error
}

type emailConfig struct {
	from string
	to   []string
	subj string
}

func (e *emailConfig) sendMessage(msg string) {
	err := sendTheEmail()
	return err
}

// END CUT 1 OMIT

// START CUT 2 OMIT
func SendMessageIfNotBlank(m MessageSender, msg string) error {
	if msg == "" {
		log.Info("Empty message - not sending anything")
		return errEmptyMessage
	}
	return m.sendMessage(msg)
}

func main() {
	e := &emailConfig{
		from: "foo@example.com",
		to:   []string{"bar@example.com", "baz@example.com"},
		subj: "[EXTERNAL] - SUPER important",
	}

	err := SendMessageIfNotBlank(e, "This is totally not a phishing email.  Click here!")
	if err != nil {
		if !errors.Is(err, errEmptyMessage) {
			log.Fatal("Couldn't send the message")
		}
	}
}

// END CUT 2 OMIT
