package uauth

import "time"

type UserAccount struct {
	UserAccountID int       `db:"user_account_id" json:"UserAccountID"`
	Email         string    `db:"email" json:"Email"`
	Password      string    `db:"password" json:"Password"`
	FullName      string    `db:"full_name" json:"FullName"`
	AddressLine1  string    `db:"address_line_1" json:"AddressLine1"`
	AddressLine2  string    `db:"address_line_2" json:"AddressLine2"`
	City          string    `db:"city" json:"City"`
	StateProvince string    `db:"state_province" json:"StateProvince"`
	CountryID     string    `db:"country_id" json:"CountryID"`
	Phone         string    `db:"phone" json:"Phone"`
	PostalCode    string    `db:"postal_code" json:"PostalCode"`
	UserRole      string    `db:"user_role" json:"UserRole"`
	Active        bool      `db:"active" json:"Active"`
	IsManager     bool      `db:"is_manager" json:"IsManager"`
	IsAdmin       bool      `db:"is_admin" json:"IsAdmin"`
	IsClient      bool      `db:"is_client" json:"IsClient"`
	CreatedOn     time.Time `db:"created_on" json:"CreatedOn"`
	LastLogin     time.Time `db:"last_login" json:"LastLogin"`
}

type UserSession struct {
	UserSessionID int       `db:"user_session_id" json:"UserSessionID"`
	UserAccountID int       `db:"user_account_id" json:"UserAccountID"`
	UAuthToken    string    `db:"uauth_token" json:"UAuthToken"`
	SessionExpiry time.Time `db:"session_expiry" json:"SessionExpiry"`
}

type UserStatus struct {
	Username string `json:"Username"`
	LoggedIn bool   `json:"LoggedIn"`
	Role     string `json:"Role"`
}
