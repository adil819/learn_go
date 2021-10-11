package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Adil"
	want := regexp.MustCompile(`\b` + name + `b`)
	msg, err := Hello("Adil")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Adil") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TesthelloEmpty calls greetings.Hello with an empty string,
// checking for and error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
