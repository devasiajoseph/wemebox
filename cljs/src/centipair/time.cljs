(ns centipair.time
  (:require [cljs-time.format :as fmt]
            [cljs-time.local :as lt]))

(def indian-date-formatter (fmt/formatter "dd-MM-yyyy"))
(def financial-time-format (fmt/formatter "yyyy-MM-dd"))

(defn local-string-date-today
  []
  (fmt/unparse indian-date-formatter (lt/local-now)))

(defn parse-indian-date
  [date]
  (try
    (fmt/parse indian-date-formatter date)
    (catch js/Error e
      e)))

(defn valid-date?
  [date format]
  (try
    (do
      (fmt/parse format date)
      true)
    (catch js/Error e
      false)))