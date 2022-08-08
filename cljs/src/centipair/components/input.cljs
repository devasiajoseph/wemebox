;; Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
;; Unauthorized copying of this file, via any medium is strictly prohibited
;; Proprietary and confidential
;; Written By Devasia Joseph <devasiajoseph@centipair.com>, December 2018

;; ==================
;; Purpose and Design
;; ==================
;; HTML Input components
(ns centipair.components.input
  (:require [centipair.error :refer [append-field-error]]
            [centipair.components.notifier :as notifier]
            [centipair.ajax :as ajax]
            [centipair.dom :as dom]))


(defn update-value
  "Updates value of the input element atom"
  [field value]
  (swap! field assoc :value value)
  (if (:on-change @field)
    ((:on-change @field) field)))

(defn update-check
  "Update value of check field element atom"
  [field checked?]
  (if checked?
    (reset! field (assoc @field :checked "checked"))
    (reset! field (assoc @field :checked ""))))

(defn update-date
  [field value]
  (update-value field value)
  (dom/set-value (:id @field) value))


(defn get-date
  [field]
  (or (dom/get-value (:id @field)) (:value @field)))


(defn make-valid
  [field]
  (swap! field assoc :message "") true)

(defn valid-field?
  [field]
  (if (nil? (:validator @field))
    (do
      (make-valid field)
      true)
    (let [result ((:validator @field) @field)] 
      (if (:valid result)
        (make-valid field)
        (do
          (append-field-error field (:message result))
          false)))))



(defn valid-form?
  [form-fields]
  (apply = true (doall (map valid-field? form-fields))))


(defn perform-action
  [action form-fields]
  (notifier/notify 200)
  (if (valid-form? form-fields)
    (action)
    (notifier/notify 422 "Invalid data submitted")))

(defn text
  [field]
  ^{:key (:id @field)}
  [:div {:class "field"}
   [:label {:for (:id @field) :class "label"} (:label @field)]
   [:div {:class (if (:icon @field) "control has-icons-left" "control")}
    [:input {:class "input"
             :type (:type @field)
             :id (:id @field)
             :placeholder
             (if (nil? (:placeholder @field))
               ""
               (:placeholder @field))
             :value (:value @field)
             :on-change #(update-value field (-> % .-target .-value))
             :maxLength (or (:maxLength @field) 255)
             :disabled (if (:disabled @field) "disabled" "")}]
    (if (:icon @field) [:span {:class "icon is-small is-left"} [:i {:class (str "fas " (:icon @field))}]] [:span])]

   [:p {:class-name "help is-danger"} (:message @field)]])


(defn button [action-button form-fields]
  [:div {:class "mb-3"} [:a {:class "btn btn-primary w-100"
                             :on-click #(perform-action (:on-click @action-button) form-fields)
                             :disabled (if (nil? (:disabled action-button)) "" "disabled")
                             :key (:id @action-button)}
                         (:label @action-button)]])





(defn select-on-change
  [field value]
  (update-value field value)
  (if (not (nil? (:on-change @field)))
    ((:on-change @field) value)))

(defn select-option [field option]
  ^{:key option} [:option {:key (str (:id @field) "-" (:value option))
                           :value (:value option)} (:label option)])

(defn to-select-option
  [map-list label-key value-key]
  (map (fn [each] {:label (label-key each) :value (value-key each)}) map-list))

(defn remote-select-options
  [field & [callback]]
  (ajax/get-json (:remote @field)
                 (:remote-params @field)
                 (fn [response]
                   (if (nil? (:value @field))
                     (swap! field assoc :value ((:value-key @field) (first response))))
                   (swap! field assoc :options
                          (if (:has-all @field)
                            (into [{:label "All" :value 0}] (to-select-option response (:label-key @field) (:value-key @field)))
                            (to-select-option response (:label-key @field) (:value-key @field))))
                   (if callback (callback)))
                 (:remote-cache @field)))

