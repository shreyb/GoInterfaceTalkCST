package main

import (
	"errors"
	"log"
)

// START CUT 1 OMIT
var errEmptyMessage = errors.New("the message was empty") // HL

func sendEmail(from string, to []string, subj string, msg string) error { // HL
	if msg == "" {
		log.Print("Empty message - not sending anything")
		return errEmptyMessage
	}
	err := sendTheEmail(msg)
	return err // Yes - this could be simply written as return sendTheEmail
}

// END CUT 1 OMIT

// START CUT 2
func main() {
	err := sendEmail( // HL
		"foo@example.com",
		[]string{"bar@example.com", "baz@example.com"},
		"[EXTERNAL] - SUPER important",
		"This is totally not a phishing email.  Click here!",
	)

	if err != nil {
		if !errors.Is(err, errEmptyMessage) {
			log.Fatal("Couldn't send the message")
		}
	}
}

// END CUT 2

func sendTheEmail(msg string) error { return nil }
