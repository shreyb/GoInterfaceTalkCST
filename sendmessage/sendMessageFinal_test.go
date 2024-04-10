package sendmessage

import (
	"errors"
	"testing"
)

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
