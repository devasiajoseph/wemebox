(ns centipair.validators
  (:require [centipair.time :as t]
            [centipair.dom :as dom]))


(defn matches-regex?
  "Returns true if the string matches the given regular expression"
  [v regex]
  (boolean (re-matches regex v)))


(defn has-value?
  "Returns true if v is truthy and not an empty string."
  [v]
  (and v (not= v "")))


(defn not-nil?
  "Returns true if v is not nil"
  [v]
  (boolean (or v (false? v))))


(defn is-email?
  "Returns true if v is an email address"
  [v]
  (if (nil? v)
    false
    (matches-regex? v #"(?i)[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")))


(defn min-length?
  "Returns true if v is greater than or equal to the given len"
  [v len]
  (>= (count v) len))

(defn is-float?
  "Returns true if a valid float"
  [value]
  (not (js/isNaN (js/parseFloat value))))


(defn is-number?
  "Returns true if a valid float"
  [value]
  (not (js/isNaN value)))


(defn is-indian-date?
  "Checks if date format if of type dd/MM/yyyy"
  [date]
  (if (nil? (aget (t/parse-indian-date date) "message"))
    true
    false))



;;error message for required field 
(def required-field-error "This field is required")


;;error message for email field
(def email-field-error "Not a valid email address")


;;error message for password length
(def password-length-error "Minimum 6 characters required")

;;error message for integer
(def integer-error "This value should be an integer")

(def float-error "This value should be a valid number")

;;error message for date
(def date-error "Invalid date")

(defn validation-error [message]
  {:message message :valid false})


(defn validation-success []
  {:valid true})


(defn required [field]
  (if (has-value? (:value field))
    (validation-success)
    (validation-error required-field-error)))

(defn checkbox-required [field]
  (if (= (:checked field) "checked")
    (validation-success)
    (validation-error required-field-error)))

(defn agree-terms [field]
  (if (= (:checked field) "checked")
    (validation-success)
    (validation-error "You have to agree to the terms and conditions")))


(defn email? [field]
  (if (has-value? (:value field))
    (if (is-email? (:value field))
      (validation-success)
      (validation-error email-field-error))
    (validation-success)))


(defn email-required [field]
  (if (has-value? (:value field))
    (if (is-email? (:value field))
      (validation-success)
      (validation-error email-field-error))
    (validation-error required-field-error)))


(defn integer-required [field]
  (if (has-value? (:value field))
    (if (integer? (js/parseInt (:value field)))
      (validation-success)
      (validation-error integer-error))
    (validation-error required-field-error)))


(defn number-required [field]
  (if (has-value? (:value field))
    (if (is-number? (:value field))
      (validation-success)
      (validation-error float-error))
    (validation-error required-field-error)))



(defn number-field [field]
  (if (has-value? (:value field))
    (if (integer? (js/parseInt (:value field)))
      (validation-success)
      (validation-error float-error))
    (validation-success)))


(defn indian-date-required
  [field]
  (let [value (dom/get-value (:id field))]
    (if (has-value? value)
      (if (is-indian-date? value)
        (validation-success)
        (validation-error date-error))
      (validation-error required-field-error))))


(defn valid-date-required
  [field date-format]
  (let [value (dom/get-value (:id field))]
    (if (has-value? value)
      (if (t/valid-date? value date-format)
        (validation-success)
        (validation-error date-error))
      (validation-error required-field-error))))
