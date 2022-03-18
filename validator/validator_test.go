package validator

import (
	"log"
	"testing"
)

var goodEmail = "sdsd@gmail.com"
var badEmail = "sdsd@gmail"

func TestIsValidEmail(t *testing.T) {

	if InvalidEmail(badEmail) {
		log.Printf("Bad email validation test success")
	} else {
		t.Errorf("Email validation failed")
	}

	if InvalidEmail(goodEmail) {
		t.Errorf("Good email validation failed")
	} else {
		log.Printf("Good email validation test success")
	}
}

func TestEmailValidation(t *testing.T) {
	r := InitResponse()
	EmailValidation(badEmail, "Email", &r)
	EmailValidation(badEmail, "AnotherEmail", &r)

	log.Printf("Email validation success")
}
