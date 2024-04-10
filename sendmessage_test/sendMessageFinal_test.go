package main

import (
	"errors"
	"log"
	"testing"
)

// Adopted from https://stackoverflow.com/a/46206116
func main() {
	tests := []testing.InternalTest{
		{
			Name: "TestSendMessageIfNotBlankEmptyString",
			F:    TestSendMessageIfNotBlankEmptyString,
		},
		{
			Name: "TestSendMessageIfNotBlankWithString",
			F:    TestSendMessageIfNotBlankWithString,
		},
	}
	testing.Main(func(pat, str string) (bool, error) { return pat == str, nil }, tests, nil, nil)
}

// START TYPEDEF OMIT

// fakeEmailConfig implements MessageSender
type fakeEmailConfig struct{}

// Calling sendMessage on the fakeEmail type doesn't send any emails!
func (f *fakeEmailConfig) sendMessage(msg string) error {
	return nil
}

// END TYPEDEF OMIT

// START OMIT
func TestSendMessageIfNotBlankEmptyString(t *testing.T) {
	f := &fakeEmailConfig{}
	err := SendMessageIfNotBlank(f, "")
	if !errors.Is(err, errEmptyMessage) {
		t.Fatal("Should have not sent an email, and gotten errEmptyMessage back")
	}
}

func TestSendMessageIfNotBlankWithString(t *testing.T) {
	f := &fakeEmailConfig{}
	err := SendMessageIfNotBlank(f, "test message")
	if err != nil {
		t.Fatal("Should have gotten a nil error from SendMessageIfNotBlank")
	}
}

// END OMIT

var errEmptyMessage = errors.New("the message was empty")

type MessageSender interface {
	sendMessage(string) error
}

func SendMessageIfNotBlank(m MessageSender, msg string) error {
	if msg == "" {
		log.Print("Empty message - not sending anything")
		return errEmptyMessage
	}
	log.Print("Sending message")
	return m.sendMessage(msg)
}
