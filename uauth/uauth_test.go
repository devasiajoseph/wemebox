package uauth

import (
	"fmt"
	"log"
	"testing"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/db/postgres"
)

func TestLogin(t *testing.T) {
	core.Start()
	postgres.InitDB()
	phone := "9539100781"
	password := "password"
	user := UserAccount{
		Phone:    phone,
		Password: password,
	}
	us, err := user.Login(password)
	if err != nil {
		log.Println(err)
		t.Errorf("User Login Failed")
	}

	log.Println(us.UAuthToken)

	au, err := GetUserByAuth(us.UAuthToken)

	if err != nil {
		log.Println(err)
		t.Errorf("Auth User Fetch failed")
	}
	log.Println("Fetched user => " + au.Email)

	err = user.Delete()
	if err != nil {
		log.Println(err)
		t.Errorf("User delete failed")
	}
}

func TestPasswordReset(t *testing.T) {
	core.Start()

	postgres.InitDB()
	ua := UserAccount{Email: "devasiajoseph@gmail.com"}
	err := ua.Data()
	if err != nil {
		t.Errorf("Error fetching user for password reset test")
	}
	ua.ResetPasswordRequest()
}

func TestUser(t *testing.T) {
	core.Start()
	postgres.InitDB()
	ua := UserAccount{Phone: "9539100781", Password: "password"}
	uk, err := ua.Register()
	if err != nil {
		log.Println(err)
		t.Errorf("Error registering user")
	}

	fmt.Println("User created successfully")
	fmt.Println(uk)

	err = ua.Delete()
	if err != nil {
		t.Errorf("Error deleting user")
	}
}

func TestSearchAuth(t *testing.T) {
	core.Start()
	postgres.InitDB()
	su, err := GetSearchUser("9539100781")
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Error creating search user")
	}

	v := VerifySearchOTP(su.OTP, su.OTPKey)
	if !v {
		t.Errorf("Erro verifying otp")
	}
}
