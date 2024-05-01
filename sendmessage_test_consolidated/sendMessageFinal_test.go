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
			Name: "TestSendMessageIfNotBlank",
			F:    TestSendMessageIfNotBlank,
		},
	}
	testing.Main(func(pat, str string) (bool, error) { return pat == str, nil }, tests, nil, nil)
}

// START CUT 1 OMIT
// goodEmailConfig implements MessageSender
type goodEmailConfig struct{}

// Calling sendMessage on the goodEmailConfig type doesn't send any emails!
func (g goodEmailConfig) sendMessage(msg string) error {
	return nil
}

var errBadEmailConfig error = errors.New("Message was not sent")

type badEmailConfig struct{}

// We ALWAYS return a non-nil error here
func (b badEmailConfig) sendMessage(msg string) error {
	return errBadEmailConfig
}

// END CUT 1 OMIT

// START FULL TEST OMIT
func TestSendMessageIfNotBlank(t *testing.T) {
	// START TYPEDEF OMIT
	type testCase struct {
		description string
		m           MessageSender
		msg         string
		expectedErr error
	}
	// END TYPEDEF OMIT

	// START TESTCASES1 OMIT
	testCases := []testCase{
		{
			"Good email config, blank message - should not send",
			goodEmailConfig{},
			"",
			errEmptyMessage,
		},
		{
			"Bad email config, blank message - should not send", // HL
			badEmailConfig{},
			"",
			errEmptyMessage,
		},
		// END TESTCASES1 OMIT
		// START TESTCASES2 OMIT
		{
			"Good email config, non-blank message - should send",
			goodEmailConfig{},
			"This is a test message",
			nil,
		},
		{
			"Bad email config, non-blank message - should not send",
			badEmailConfig{},
			"This is a test message",
			errBadEmailConfig, // HL
		},
	}
	// END TESTCASES2 OMIT

	// Note:  t.Run has the signature
	// func (t *testing.T) Run(string, func(*testing.T)) bool
	// One can check the return value, or simply call T.Fatal or T.Error from
	// inside the test
	// START TEST OMIT
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			err := SendMessageIfNotBlank(test.m, test.msg)
			if !errors.Is(err, test.expectedErr) {
				t.Errorf("We got the wrong error.  Expected %v, got %v instead.",
					test.expectedErr,
					err)
			}
		})
	}
	// END TEST OMIT
}

// END FULL TEST OMIT

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
