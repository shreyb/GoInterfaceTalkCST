package main

import (
	"errors"
	"testing"
)

// START CUT 1 OMIT
func TestSendEmailEmptyString(t *testing.T) {
	// We expect to get back errEmptyMessage
	err := sendEmail(
		"foo@example.com",
		[]string{"bar@example.com", "baz@example.com"},
		"[EXTERNAL] - SUPER important",
		"",
	)
	if !errors.Is(err, errEmptyMessage) {
		t.Error("We should have not sent the email, and gotten back an errEmptyMessage")
	}
}

// END CUT 1 OMIT

// START CUT 2 OMIT
func TestSendEmail(t *testing.T) {
	// There should be no error, but I also need to check my inbox to see
	// if this test ran!
	err := sendEmail(
		"foo@example.com",
		[]string{"sbhat@fnal.gov"},
		"[EXTERNAL] - SUPER important",
		"This is a test",
	)

	if err != nil {
		t.Error("There should have been no error")
	}
	t.Log("OK, so there was no error.  Go check your inbox to make sure this REALLY passes")
}

// END CUT 2 OMIT
