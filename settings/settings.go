package settings

import (
	"log"

	"github.com/devasiajoseph/wemebox/db/postgres"
)

var OTPRegistrationVerificationTemplateID = "rv-otp-temp-id"
var OTPForgotPasswordTemplateID = "fp-otp-temp-id"

var kv = map[string]string{
	"rv-otp-temp-id": "60421ddcb3c91864354e1d75",
	"fp-otp-temp-id": "604768a47842e8322e4fb5f7",
}

type Settings struct {
	SettingsID    int    `db:"settings_id" json:"settings_id"`
	SettingsName  string `db:"settings_name" json:"settings_name"`
	SettingsValue string `db:"settings_value" json:"settings_value"`
}

func GetSettings(key string) (Settings, error) {
	var set Settings
	set.SettingsName = key
	set.SettingsValue = kv[key]
	return set, nil
}

func GetSettingsBkp(key string) (Settings, error) {
	var set Settings
	db := postgres.Db
	err := db.Get(&set, "select * from settings where settings_name=$1", key)
	if err != nil {
		log.Println("Error getting settings")
		log.Println(err)
	}
	return set, err
}
