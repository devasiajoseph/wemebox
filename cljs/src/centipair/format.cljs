;; Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
;; Unauthorized copying of this file, via any medium is strictly prohibited
;; Proprietary and confidential
;; Written By Devasia Joseph <devasiajoseph@centipair.com>, December 2018
;; ======================
;; Purpose and Design
;; ======================
;; Purpose of this file is to add functions to manipulate strings and data structures
;; to a desired format

(ns centipair.format
  (:require [clojure.string :as s]
            [goog.string :as gstring]
            [centipair.validators :as validator]))

(defn url-delimter
  [v]
  (s/replace v #" " "-"))

(defn url-slug
  [v]
  (-> v
      s/trim
      s/lower-case
      url-delimter))

(defn append-zero
  [num]
  (if (< num 10)
    (str "0" num)
    (str "" num)))

(defn format-null-string
  [v]
  (if (= "null" v)
    ""
    v))

(defn dbdate-datepicker
  [db-date]
  (if (= db-date "0001-01-01T00:00:00Z")
    ""
    (let [dt (js/Date. (first (clojure.string/split db-date #"T")))]   
      (str  (append-zero (.getDate dt))
            "-"
            (append-zero (+ 1 (.getMonth dt)))
            "-"
            (.getFullYear dt)))))


(defn round-float
  [v]
  (if (validator/has-value? v)
    (gstring/format "%.2f" v)
    "0.0"))


(defn date-string 
  [date]
  (let [d (new js/Date date)
        dtstr (.toDateString d)]
    dtstr))