/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>
 */

package uauth

import (
	"database/sql"
	"errors"
	"fmt"

	//"database/sql"

	"github.com/devasiajoseph/wemebox/core"
	"github.com/devasiajoseph/wemebox/crypt"
	"github.com/devasiajoseph/wemebox/db/postgres"

	//"github.com/jmoiron/sqlx"
	"log"
	"time"

	"github.com/devasiajoseph/wemebox/cpmath"
	"github.com/devasiajoseph/wemebox/mail"
	"github.com/devasiajoseph/wemebox/validator"
)

const passwordResetKey = "password-reset"
const registrationKey = "registration-key"
const RoleAdmin = "admin"

//UserAccount User data structure
type UserAccount struct {
	UserAccountID int       `db:"user_account_id" json:"UserAccountID"`
	Email         string    `db:"email" json:"Email"`
	Phone         string    `db:"phone" json:"Phone"`
	Password      string    `db:"password" json:"-"`
	FullName      string    `db:"full_name" json:"FullName"`
	Active        bool      `db:"active" json:"Active"`
	CreatedOn     time.Time `db:"created_on" json:"CreatedOn"`
	LastLogin     time.Time `db:"last_login" json:"LastLogin"`
	Role          string    `db:"role" json:"Role"`
}

//UserKeys has user login creds
type UserKeys struct {
	UserKeysID    int    `db:"user_keys_id" json:"UserKeysID"`
	UserAccountID int    `db:"user_account_id" json:"UserAccountID"`
	KeyName       string `db:"key_name" json:"KeyName"`
	KeyValue      string `db:"key_value" json:"KeyValue"`
	OTP           string `db:"otp" json:"OTP"`
}

type OTPSession struct {
	KeyName  string `db:"key_name" json:"KeyName"`
	KeyValue string `db:"key_value" json:"KeyValue"`
}

//UserSession holds user session creds
type UserSession struct {
	UserSessionID int       `db:"user_session_id" json:"UserSessionID"`
	UserAccountID int       `db:"user_account_id" json:"UserAccountID"`
	UAuthToken    string    `db:"uauth_token" json:"UAuthToken"`
	SessionExpiry time.Time `db:"session_expiry" json:"SessionExpiry"`
}

//UserStatus holds login status
type UserStatus struct {
	Username string `json:"Username"`
	LoggedIn bool   `json:"LoggedIn"`
	Role     string `json:"Role"`
}

//AuthUser current authenticated user
type AuthUser struct {
	UserAccountID int       `db:"user_account_id" json:"UserAccountID"`
	FullName      string    `db:"full_name" json:"FullName"`
	Email         string    `db:"email" json:"Email"`
	Phone         string    `db:"phone" json:"Phone"`
	Active        bool      `db:"active" json:"Active"`
	CreatedOn     time.Time `db:"created_on" json:"CreatedOn"`
	LastLogin     time.Time `db:"last_login" json:"LastLogin"`
	SessionExpiry time.Time `db:"session_expiry" json:"SessionExpiry"`
	Role          string    `db:"role" json:"Role"`
}

type ObjList struct {
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
	Data  []UserAccount `json:"data"`
	Total int           `json:"total"`
	Query string        `json:"query"`
}

//Create creates new user account
func (ua *UserAccount) Create() error {
	db := postgres.Db
	var sqlInsertUser = "INSERT INTO user_account" +
		"(phone,email, password, full_name, active, created_on, last_login)" +
		"VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING user_account_id;"

	err := db.QueryRow(sqlInsertUser,
		ua.Phone,
		ua.Email,
		crypt.HashPassword(ua.Password),
		ua.FullName,
		//ua.Active,
		false,
		time.Now(),
		time.Now()).Scan(&ua.UserAccountID)
	if err != nil {
		log.Println("Error creating user")
		log.Println(err)
		return err
	}

	return err
}

func (ua *UserAccount) CreateRaw() error {
	db := postgres.Db
	var sqlInsertUser = "INSERT INTO user_account" +
		"(phone,email, password, full_name, active, created_on, last_login)" +
		"VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING user_account_id;"
	err := db.QueryRow(sqlInsertUser,
		ua.Phone,
		ua.Email,
		ua.Password,
		ua.FullName,
		ua.Active,
		time.Now(),
		time.Now()).Scan(&ua.UserAccountID)
	if err != nil {
		log.Println("Error creating user")
		//log.Println(err)
		return err
	}

	return err
}

//Update updates user account
func (ua *UserAccount) Update() error {
	db := postgres.Db
	var sqlUpdateUser = "UPDATE user_account " +
		"SET full_name=$1,active=$2 " +
		"WHERE user_account_id=$3;"
	stmt, err := db.Prepare(sqlUpdateUser)
	if err != nil {
		log.Println("Error preparing statement for user update")
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(
		ua.FullName,
		ua.Active,
		ua.UserAccountID)
	if err != nil {
		log.Println("Error updating user")
		log.Println(err)
	}
	return err
}

//UpdateProfile updates user profile
func (ua *UserAccount) UpdateProfile() error {
	db := postgres.Db
	var sqlUpdateUser = "UPDATE user_account " +
		"SET full_name=$1,email=$2,phone=$3 " +
		"WHERE user_account_id=$4;"
	stmt, err := db.Prepare(sqlUpdateUser)
	if err != nil {
		log.Println("Error preparing statement for user update")
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(
		ua.FullName,
		ua.Email,
		ua.Phone,
		ua.UserAccountID)
	if err != nil {
		log.Println("Error updating user")
		log.Println(err)
	}
	return err
}

//Activate activates user
func (ua *UserAccount) Activate() error {
	err := ua.Data()
	if err != nil {
		fmt.Println("Unable to get user for activation")
		return err
	}

	ua.Active = true
	err = ua.Update()
	if err != nil {
		log.Println("Unable to activate user")
	}
	return err
}

//Save saves or updates user acccount
func (ua *UserAccount) Save() error {
	if ua.UserAccountID == 0 {
		return ua.Create()
	}
	return ua.Update()
}

//Data fetches user account data based on id or phone or email
func (ua *UserAccount) Data() error {
	db := postgres.Db
	bsql := "select * from user_account where "
	var err error
	if ua.UserAccountID > 0 {
		bsql += " user_account_id=$1;"
		err = db.Get(ua, bsql, ua.UserAccountID)
		if err != nil {
			fmt.Println(err)
			log.Println("Unable to fetch user data with id")
		}
		return err
	}
	if ua.Phone != "" {
		bsql += " phone=$1;"
		err := db.Get(ua, bsql, ua.Phone)
		if err != nil {
			log.Println(err)
			log.Println("Unable to fetch user data with phone ")
		}
		return err
	}

	if ua.Email != "" {
		bsql += " email=$1;"
		err := db.Get(ua, bsql, ua.Email)
		if err != nil {
			log.Println(err)
			log.Println("Unable to fetch user data with email ")
		}
		return err
	}

	return errors.New("no fetch parameter provided")
}

//Account gets account data with phone or email
func (ua *UserAccount) Account(pe string) error {
	db := postgres.Db
	bsql := "select * from user_account where phone=$1;"
	err := db.Get(ua, bsql, pe)
	if err != nil {
		fmt.Println(err)
		log.Println("Unable to fetch user data with id")
	}
	return err
}

func (ua *UserAccount) SetLoginTime() {
	db := postgres.Db
	_, err := db.Exec("Update user_account set last_login=$1 where user_account_id=$2", time.Now(), ua.UserAccountID)
	if err != nil {
		log.Println(err)
		log.Println("Error setting last login time")
	}
}

//Login logs in a user
func (ua *UserAccount) Login(password string) (UserSession, error) {
	var us UserSession
	err := ua.Data()
	if err != nil {
		log.Println(err)
		return us, errors.New("user account not found")
	}
	if crypt.CheckPasswordHash(password, ua.Password) {
		us, err = ua.CreateUserSession()
		ua.SetLoginTime()
		return us, err
	}
	return us, errors.New("invalid password")
}

var sqlKeepAlive = "update user_session set session_expiry=$1 where auth_token=$2"

func keepAlive(cookie string) {
	db := postgres.Db
	expiry := time.Now().Add(30 * time.Minute)
	_, err := db.Exec(sqlKeepAlive, expiry, cookie)
	if err != nil {
		log.Println("Error keeping alive")
	}
}

//CreateUserSession creates user session
func (ua *UserAccount) CreateUserSession() (UserSession, error) {
	var us UserSession
	db := postgres.Db
	authToken, err := crypt.RandomString(32)
	if err != nil {
		return us, err
	}
	us.UserAccountID = ua.UserAccountID
	us.UAuthToken = authToken
	us.SessionExpiry = time.Now().AddDate(0, 6, 0)
	_, err = db.Exec(sqlCreateSession, us.UserAccountID, us.UAuthToken, us.SessionExpiry)
	if err != nil {
		log.Println("Error creating session")
	}

	return us, err
}

var sqlDeleteUserSession = "DELETE FROM user_session WHERE auth_token=$1;"

//Logout user
func Logout(uaid string) error {
	db := postgres.Db
	_, err := db.Exec(sqlDeleteUserSession, uaid)
	if err != nil {
		log.Println("Error deleting user session")
	}
	return err
}

//GetUserByAuth gets user by auth token
func GetUserByAuth(uauth string) (AuthUser, error) {
	db := postgres.Db
	var au AuthUser
	err := db.Get(&au, sqlFetchAuthUser, uauth)
	if err != nil {
		log.Println(err)
		return au, err
	}
	if !au.Active {
		log.Println("Inactive user => " + au.Email + ":" + au.Phone)
		return au, errors.New("inactive user")
	}
	return au, err
}

var sqlCreateUserKey = "insert into user_keys (user_account_id, key_name, key_value,otp) values($1,$2,$3,$4);"

//Save saves userkeys
func (uk *UserKeys) Save() error {
	db := postgres.Db
	_, err := db.Exec(sqlCreateUserKey, uk.UserAccountID, uk.KeyName, uk.KeyValue, uk.OTP)
	if err != nil {
		fmt.Println(err)
		log.Println("Unable to create user key")
	}

	return err
}

var sqlFetchUserKey = "select * from user_keys where key_name=$1 and key_value=$2;"

//Data gets userkeys data
func (uk *UserKeys) Data() error {
	db := postgres.Db
	err := db.Get(uk, sqlFetchUserKey, uk.KeyName, uk.KeyValue)
	if err != nil {
		log.Println(err)
		log.Println("Error fetching user keys")

	}

	return err
}

//Delete deletes user keys
func (uk *UserKeys) Delete() error {
	db := postgres.Db
	stmt, err := db.Prepare("delete from user_keys where user_keys_id=$1")
	if err != nil {
		log.Println("Error preparing statement for user keys delete")
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(uk.UserKeysID)
	if err != nil {
		log.Println("Error deleting userkeys")
		log.Println(err)
	}

	return err
}

//Delete deletes user keys
func (uk *UserKeys) DeleteUserKey() error {
	db := postgres.Db
	stmt, err := db.Prepare("delete from user_keys where key_name=$1 and user_account_id=$2")
	if err != nil {
		log.Println("Error preparing statement for user keys delete")
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(uk.KeyName, uk.UserAccountID)
	if err != nil {
		log.Println("Error deleting userkeys")
		log.Println(err)
	}

	return err
}

//VerifyOTP verifies generated otp
func (uk *UserKeys) VerifyOTP(otp string) bool {
	err := uk.Data()
	if err != nil {
		fmt.Println("Unable to get user keys for otp verification")
		fmt.Println(err)
		return false
	}
	if uk.OTP != otp {
		return false
	}
	return true
}

//VerifyOTP verifies generated otp
func (uk *UserKeys) UpdateKey(keyValue string) error {
	db := postgres.Db

	_, err := db.Exec("Update user_keys set key_value = $1 where user_keys_id=$2", keyValue, uk.UserKeysID)
	if err != nil {
		log.Println("Error updating user key")
		log.Println(err)
	}
	uk.KeyValue = keyValue
	return err

}

//VerifyKey verifies a user key
func (uk *UserKeys) VerifyKey(key string, value string) bool {
	err := uk.Data()
	if err != nil {
		fmt.Println("Unable to get user keys for verification")
		fmt.Println(err)
		return false
	}

	return true
}

//SendOTP sends OTP via SMS
func (uk *UserKeys) SendOTP() {
	ua := UserAccount{UserAccountID: uk.UserAccountID}
	err := ua.Data()
	if err != nil {
		log.Println("Error fetching user data to send OTP")
		log.Println(err)
	}

	/*err = sms.SendOTP(uk.OTP, ua.Phone)
	if err != nil {
		log.Println("Error sending OTP")
		log.Println(err)
	}*/

}

//GenerateOTP generates OTP for verification
func (ua *UserAccount) GenerateOTP(keyName string) (UserKeys, error) {
	var uk UserKeys
	authKey := crypt.UID()
	otp, err := crypt.NumericKey(6)
	if err != nil {
		log.Println("Unable to generate otp")
		return uk, err
	}
	uk.KeyName = keyName
	uk.KeyValue = authKey
	uk.UserAccountID = ua.UserAccountID
	uk.OTP = otp
	uk.DeleteUserKey()
	uk.Save()
	return uk, err

}

func (ua *UserAccount) Register() (UserKeys, error) {
	var uk UserKeys
	err := ua.Create()
	if err != nil {
		log.Println("Error registering new user")
		return uk, err
	}

	return uk, err
}

func (ua *UserAccount) Delete() error {
	db := postgres.Db
	stmt, err := db.Prepare(sqlDeleteUser)
	if err != nil {
		log.Println("Error preparing statement for user delete")
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(ua.UserAccountID)
	if err != nil {
		log.Println("Error deleting user")
		log.Println(err)
	}

	return err
}

func (ua *UserAccount) ChangePassword(password string) error {
	db := postgres.Db
	stmt, err := db.Prepare(sqlChangePassword)
	if err != nil {
		log.Println("Error preparing statement for user password")
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(crypt.HashPassword(password), ua.UserAccountID)
	if err != nil {
		log.Println("Error updating password")
		log.Println(err)
	}

	return err
}

func UniqueEmailValidation(v string, id string, r *validator.ValidatorResponse) {
	ua := UserAccount{Email: v}
	err := ua.Data()
	if err != nil {
		return
	}
	if ua.UserAccountID > 0 {
		validator.AppendError("This email already exists", id, r)

	}
}

func EmailExistValidation(v string, id string, r *validator.ValidatorResponse) {
	ua := UserAccount{Email: v}
	err := ua.Data()
	if err != nil {
		validator.AppendError("This account does not exist", id, r)
	}
}

func UniquePhoneValidation(v string, id string, r *validator.ValidatorResponse) {
	ua := UserAccount{Phone: v}
	err := ua.Data()

	if err != nil {
		return
	}
	if ua.UserAccountID > 0 {
		validator.AppendError("This phone already exists", id, r)

	}
}

func PhoneExistValidation(v string, id string, r *validator.ValidatorResponse) {
	ua := UserAccount{Phone: v}
	err := ua.Data()
	if err != nil {
		validator.AppendError("This phone does not exist", id, r)
	}
}

//UpdatePhoneValidation checks for unique phone number while updating
func UpdatePhoneValidation(ua UserAccount, phone string, id string, r *validator.ValidatorResponse) {
	uaCheck := UserAccount{Phone: phone}
	err := uaCheck.Data()

	if err != nil {
		return
	}

	if ua.UserAccountID != uaCheck.UserAccountID {
		validator.AppendError("This phone already exists for another account", id, r)

	}
}

func ChangePasswordValidation(ua AuthUser, oldPassword string, id string, r *validator.ValidatorResponse) {
	userAccount := UserAccount{UserAccountID: ua.UserAccountID}
	err := userAccount.Data()
	if err != nil {
		return
	}
	if !crypt.CheckPasswordHash(oldPassword, userAccount.Password) {
		validator.AppendError("Your old password entered for authentication is wrong", id, r)
	}
}

func (ua *UserAccount) Exist() bool {
	db := postgres.Db
	bsql := "select * from user_account where phone=$1;"
	err := db.Get(ua, bsql, ua.Phone)
	return err == nil
}

//AccountExistValidation Checks if active account exist
func AccountExistValidation(v string, id string, r *validator.ValidatorResponse) {
	var ua UserAccount
	err := ua.Account(v)

	if err != nil || !ua.Active {
		validator.AppendError("An active account with this phone or email doesn't exist.", id, r)
	}
}

func (ol *ObjList) List() error {
	offset := cpmath.Offset(ol.Page, ol.Limit)
	db := postgres.Db
	err := db.Select(&ol.Data, sqlFetchUsers, ol.Limit, offset)
	if err != nil {
		log.Println(err)
		log.Println("Error fetching users for admin")
		return err
	}
	err = db.Get(&ol.Total, sqlTotalUsers)
	if err != nil {
		log.Println(err)
		log.Println("Error fetching total for admin user list")
	}
	return err
}

func (ol *ObjList) Search() error {
	offset := cpmath.Offset(ol.Page, ol.Limit)
	db := postgres.Db
	regexQ := "%" + ol.Query + "%"
	err := db.Select(&ol.Data, sqlSearchUsers, regexQ, ol.Limit, offset)
	if err != nil {
		log.Println(err)
		log.Println("Error searching users in admin")
		return err
	}
	err = db.Get(&ol.Total, sqlSearchUsersTotal, regexQ)

	if err != nil {
		log.Println(err)
		log.Println("Error fetching total")
	}

	return err
}

//RemovePasswordResetKey deletes password reset key
func (ua *UserAccount) RemovePasswordResetKey() error {
	log.Println("REmoving key")
	db := postgres.Db
	stmt, err := db.Prepare(sqlDeletePasswordResetKey)
	if err != nil {
		log.Println("Error preparing statement for user delete")
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(ua.UserAccountID, passwordResetKey)
	if err != nil {
		log.Println("Error deleting user password reset key")
		log.Println(err)
	}

	return err
}

//SetPasswordResetKey sets password reset key for changing password
func (ua *UserAccount) SetPasswordResetKey() (string, error) {
	err := ua.RemovePasswordResetKey()
	if err != nil {
		return "", err
	}
	resetKey, err := crypt.RandomString(32)
	if err != nil {
		log.Println(err)
		return resetKey, err
	}
	db := postgres.Db
	db.MustExec(sqlCreateUserKey,
		ua.UserAccountID,
		passwordResetKey,
		resetKey)

	return resetKey, err

}

var sqlSetUserPassword = "update user_account set password=$1 where user_account_id=$2;"

//SetPassword sets new password for existing user accounr
func (ua *UserAccount) SetPassword(password string) error {
	db := postgres.Db
	res := db.MustExec(sqlSetUserPassword, crypt.HashPassword(password), ua.UserAccountID)
	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if ra < 1 {
		return errors.New("no rows affected")
	}
	return nil
}

//ResetPasswordRequest initializes password reset process
func (ua *UserAccount) ResetPasswordRequest() error {
	resetKey, err := ua.SetPasswordResetKey()
	if err != nil {
		log.Println("Reset password request erro")
		return err
	}
	proto := "http://"

	if core.Config.Ssl {
		proto = "https://"
	}
	m := mail.Mail{
		Sender:    core.Config.AppEmail,
		Recipient: ua.Email,
		Subject:   "Reset your " + core.Config.AppName + " account password",
		Body: "Please click on this link to reset your password: " +
			proto + core.Config.Domain + "/app/uauth/password-reset-confirm/" + resetKey,
	}

	err = m.SendMailGun()
	if err != nil {
		log.Println("Error sending password reset mail")
	}
	return err
}

//ResetPasswordRequest initializes password reset process with OTP
func (ua *UserAccount) ForgotPasswordRequest() error {
	resetKey, err := ua.SetPasswordResetKey()
	if err != nil {
		log.Println("Reset password request erro")
		return err
	}
	proto := "http://"

	if core.Config.Ssl {
		proto = "https://"
	}
	m := mail.Mail{
		Sender:    core.Config.AppEmail,
		Recipient: ua.Email,
		Subject:   "Reset your " + core.Config.AppName + " account password",
		Body: "Please click on this link to reset your password: " +
			proto + core.Config.Domain + "/app/uauth/password-reset-confirm/" + resetKey,
	}

	err = m.SendMailGun()
	if err != nil {
		log.Println("Error sending password reset mail")
	}
	return err
}

//CompletePasswordReset completes password reset process
func (uk *UserKeys) CompletePasswordReset(newPassword string) error {
	err := uk.Data()
	if err != nil {
		log.Println("Error completing password reset:Fetching keys")
		return err
	}

	ua := UserAccount{UserAccountID: uk.UserAccountID}
	err = ua.Data()
	if err != nil {
		log.Println("Error in password reset user fetch")
		return err
	}
	err = ua.SetPassword(newPassword)
	if err != nil {
		log.Println(err)
		log.Println("Error in password reset save password")
		return err
	}

	err = ua.RemovePasswordResetKey()
	if err != nil {
		log.Println(err)
	}
	return nil
}

//GenerateTestUser generates a userd for testing
func GenerateTestUser() (UserAccount, error) {
	ua := UserAccount{
		Email:    "testuser00944127890452834834@mailfromnowhere.com",
		Password: "password",
		Phone:    "12344434343"}
	err := ua.Data()
	if err != nil {
		log.Println(err)
		log.Println("Test user not found..creating one")
	} else {
		return ua, err
	}
	err = ua.Create()
	if err != nil {
		log.Println("Failed to create test user")
	}
	return ua, err
}

type SearchUser struct {
	SearchUserID int    `db:"search_user_id" json:"search_user_id"`
	Phone        string `db:"phone" json:"phone"`
	OTP          string `db:"otp" json:"otp"`
	OTPKey       string `db:"otp_key" json:"otp_key"`
	SearchAuth   string `db:"search_auth" json:"search_auth"`
}

func (su *SearchUser) Create() error {

	su.OTPKey = crypt.UID()
	authKey, err := crypt.RandomString(32)

	if err != nil {
		log.Println(err)
		log.Println("Error creating search auth")
		return err
	}
	su.SearchAuth = authKey

	otp, err := crypt.NumericKey(6)
	if err != nil {
		log.Println(err)
		log.Println("Error creating otp for search auth")
		return err
	}
	su.OTP = otp

	db := postgres.Db

	_, err = db.NamedExec("insert into search_user (phone,otp,otp_key,search_auth) values (:phone,:otp,:otp_key,:search_auth);", su)
	if err != nil {
		log.Println(err)
		log.Println("Error creating search auth")
	}

	return err
}

func GetSearchUser(phone string) (SearchUser, error) {
	db := postgres.Db
	su := SearchUser{}
	err := db.Get(&su, "select * from search_user where phone=$1", phone)
	if err != nil {
		if err == sql.ErrNoRows {
			su.Phone = phone
			err = su.Create()
			if err != nil {
				log.Println(err)
				log.Println("Error creating search user")
			}
		} else {
			fmt.Println(err)

		}
	}

	return su, err
}

func VerifySearchOTP(otp string, otp_key string) bool {
	db := postgres.Db
	su := SearchUser{}
	err := db.Get(&su, "select * from search_user where otp=$1 and otp_key=$2;", otp, otp_key)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
