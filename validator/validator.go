package validator

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var emailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

//ValidationError validation error structure for json response
type ValidationError struct {
	ID      string `json:"Id"`
	Message string `json:"Message"`
	Label   string `json:"Label"`
}

//ValidatorResponse list of errors in a form
type ValidatorResponse struct {
	Message string            `json:"Message"`
	Valid   bool              `json:"Valid"`
	Errors  []ValidationError `json:"Errors"`
}

//Response list of errors in a form
type Response struct {
	Message string            `json:"Message"`
	Valid   bool              `json:"Valid"`
	Errors  []ValidationError `json:"Errors"`
}

//InitResponse initializes valiadtion error response
func InitResponse() ValidatorResponse {
	var ve = []ValidationError{}
	vr := ValidatorResponse{
		Message: "Valid",
		Valid:   true,
		Errors:  ve}
	return vr
}

//EmptyString checks for empty string
//"" "<space>" etc are empty strings
func EmptyString(v string) bool {
	v = strings.Replace(v, " ", "", -1)
	if len(v) == 0 || v == "null" {
		return true
	}
	return false
}

//InvalidEmail checks for email format
func InvalidEmail(v string) bool {
	if EmptyString(v) {
		return false
	}
	if emailRegexp.MatchString(v) {
		return false
	} else {
		return true
	}
}

//FileExist check if a file exist
func FileExist(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

//ValidFileExtension checks if the extension of a file in file path matches
//parameters : filepath, extension
func ValidFileExtension(fp string, ext string) bool {
	if filepath.Ext(fp) != ext {
		return false
	}
	return true
}

//AppendError appends validator error to Validator response
func AppendError(em string, id string, r *ValidatorResponse) {
	r.Message = "Validation Failed"
	r.Valid = false
	verror := ValidationError{
		ID:      id,
		Message: em}
	r.Errors = append(r.Errors, verror)
}

//EmailValidation adds email validation check to ValidatorResponse
func EmailValidation(v string, id string, r *ValidatorResponse) {
	if InvalidEmail(v) {
		AppendError("Invalid Email", id, r)
	}
}

func FileExtensionValidation(fp string, ext string, id string, r *ValidatorResponse) {
	if !ValidFileExtension(fp, ext) {
		AppendError("Invalid File extension", id, r)
	}
}

func RequiredStringValidation(v string, id string, r *ValidatorResponse) {
	if EmptyString(v) {
		AppendError("This field is required", id, r)
	}
}

func RequiredNumberValidation(v string, id string, r *ValidatorResponse) {
	if EmptyString(v) {
		AppendError("This field is required", id, r)
		return
	}

	_, err := strconv.ParseFloat(v, 64)
	if err != nil {
		AppendError("Invalid Value", id, r)
	}
}

func FileExistValidation(v string, id string, r *ValidatorResponse) {
	if EmptyString(v) {
		AppendError("This field is required", id, r)
		return
	}

	if FileExist(v) {
		AppendError("This file already exist", id, r)
	}
}

func FileNotExistValidation(v string, id string, r *ValidatorResponse) {
	if EmptyString(v) {
		AppendError("This field is required", id, r)
		return
	}

	if !FileExist(v) {
		AppendError("This file doesn't exist", id, r)
	}
}

//IntegerValueValidation checks if the value is a valid integer
func IntegerValueValidation(v string, id string, r *ValidatorResponse) {
	if EmptyString(v) {
		AppendError("This field is required", id, r)
		return
	}
	_, err := strconv.Atoi(v)
	if err != nil {
		AppendError("Invalid Integer", id, r)
	}
}

//DBIntegerValidation checks for a positive id related to an sql row primary key
func DBIntegerValidation(v string, id string, r *ValidatorResponse) {
	if EmptyString(v) {
		AppendError(id+":ID required", id, r)
		return
	}
	ID, err := strconv.Atoi(v)
	if err != nil || ID <= 0 {
		AppendError("Invalid ID for:"+id, id, r)
	}

}

//ValidationErrorResponse generic validation error response
func ValidationErrorResponse(w http.ResponseWriter, vRes ValidatorResponse) {
	w.Header().Set("Content-Type", "application/json")
	responseJSON, _ := json.Marshal(vRes)
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(responseJSON)
	return
}

//EmailValue if not null check for email value
func EmailValue(v string, id string, r *ValidatorResponse) {
	if EmptyString(v) {
		return
	}
	if InvalidEmail(v) {
		AppendError("Invalid Email", id, r)
	}
}

var indianPhoneRegexp = regexp.MustCompile(`^(?:(?:\+|0{0,2})91(\s*[\-]\s*)?|[0]?)?[789]\d{9}$`)

func ValidPhone(v string) bool {
	if indianPhoneRegexp.MatchString(v) {
		return true
	}

	return false
}

func IndianPhoneValidation(v string, id string, r *ValidatorResponse) {
	if ValidPhone(v) {
		return
	} else {
		AppendError("Invalid Phone", id, r)
		return
	}

}
