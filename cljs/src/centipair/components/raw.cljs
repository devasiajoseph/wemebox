(ns centipair.components.raw
  (:require [centipair.components.input :as cin]
            [centipair.components.notifier :as notifier]))



(defn valid-form?
  [form-fields]
  (apply = true (doall (map cin/valid-field? form-fields))))


(defn perform-action
  [action form-fields]
  (notifier/notify 200)
  (if (valid-form? form-fields)
    (action)
    (notifier/notify 422 "Invalid data submitted")))

(defn text [field]
  ^{:key (:id @field)}
  [:input {:class (:class @field)
           :type (:type @field)
           :id (:id @field)
           :value (:value @field)
           :on-change #(cin/update-value field (-> % .-target .-value))
           :maxLength (or (:maxLength @field) 255)
           :disabled (if (:disabled @field) "disabled" "")
           :placeholder (if (nil? (:placeholder @field))
                          ""
                          (:placeholder @field))}])