package main

import (
	"errors"
	"log"
)

// START CUT 1 OMIT
var errEmptyMessage = errors.New("the message was empty") // HL

// START TYPEDEF OMIT
type emailConfig struct { // HL
	from string
	to   []string
	subj string
}

// END TYPEDEF OMIT

func (e emailConfig) sendMessage(msg string) error { // HL
	if msg == "" {
		log.Print("Empty message - not sending anything")
		return errEmptyMessage
	}
	err := sendTheEmail(e.from, e.to, e.subj, msg)
	return err
}

// END CUT 1 OMIT

// START CUT 2 OMIT
func main() { // HL
	e := emailConfig{
		from: "foo@example.com",
		to:   []string{"bar@example.com", "baz@example.com"},
		subj: "[EXTERNAL] - SUPER important",
	}

	err := e.sendMessage("This is totally not a phishing email.  Click here!")
	if err != nil {
		if !errors.Is(err, errEmptyMessage) {
			log.Fatal("Couldn't send the message")
		}
	}
}

// END CUT 2 OMIT

func sendTheEmail(from string, to []string, subject, msg string) error { return nil }
