;; Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
;; Unauthorized copying of this file, via any medium is strictly prohibited
;; Proprietary and confidential
;; Written By Devasia Joseph <devasiajoseph@centipair.com>, December 2018

;; ==================
;; Purpose and Design
;; ==================
;; Ajax base functions

(ns centipair.ajax
  (:require [ajax.core :refer [GET POST DELETE json-request-format json-response-format url-request-format]]
            [centipair.dom :as dom]
            [centipair.spa :as spa]
            [centipair.error :refer [append-error remove-error]]
            [centipair.components.notifier :refer [notify]]
            [reagent.core :as reagent]
            [reagent.dom :as r]))


(def csrf-token (atom nil))

(def ajax-cache (atom {}))

(defn set-value-type [field]
  (case (:value-type @field)
    nil (:value @field)
    "string" (:value @field)
    "integer" (js/parseInt  (:value @field))
    "float" (js/parseFloat  (:value @field))
    (:value @field)))


(defn set-select-value-type [field]
  (let [fetched-value (dom/get-value (:id @field))]
    (case (:value-type @field)
    nil fetched-value
    "string" fetched-value
    "integer" (js/parseInt  fetched-value)
    "float" (js/parseFloat  fetched-value)
    fetched-value)))


(defn handle-deleted []
  (println "deleted"))


(defn success-handler
  [function-handler response]
  (notify (:status response))
  (function-handler response))


(defn error-handler [function-handler response]
  (let [status (:status response)]
    (case status
      200 (success-handler function-handler response)
      201 (success-handler function-handler response)
      202 (handle-deleted)
      204 (notify 204)
      400 (notify 400)
      401 (notify 401)
      403 (notify 403)
      404 (notify 404)
      405 (notify 405)
      500 (notify 500)
      422 (notify 422 "The submitted data is not valid")
      (notify 403 "Unknown request"))))


(defn delete-success-handler
  [function-handler response]
  (do 
    (js/hideDeleteModal)
    (success-handler function-handler response)))

(defn text-to-key [previous each]
  (assoc previous (keyword (:id @each)) (set-value-type each)))

(defn check-to-key[previous each]
  (if (= (:checked @each) "checked")
      (assoc previous (keyword (:id @each)) true)
      (assoc previous (keyword (:id @each)) false)))

(defn toggle-to-key [previous each]
  (if (= (:checked @each) "checked")
    (assoc previous (keyword (:id @each)) true)
    (assoc previous (keyword (:id @each)) false)))

(defn image-spec-to-key[previous each]
  (assoc previous 
    (keyword (:id (:width @each))) (js/parseInt (:value (:width @each)))
    (keyword (:id (:height @each))) (js/parseInt (:value (:height @each)))
    (keyword (:id (:crop @each))) (= (:checked (:crop @each)) "checked")))

(defn select-to-key [previous each]
  (assoc previous (keyword (:id @each)) (set-select-value-type each)))

(defn select-text-to-key
  [previous each]
  (assoc previous 
    (keyword (:id (:text @each))) (:value (:text @each))
    (keyword (:id (:select @each))) (dom/get-value (:id (:select @each)))))

(defn datepicker-to-key
  [previous each]
  (assoc previous (keyword (:id @each)) (dom/get-value (:id @each))))

(defn to-key
  "each is an atom"
  [previous each]
  (case (:type @each)
    "text" (text-to-key previous each)
    "checkbox" (check-to-key previous each)
    "toggle" (toggle-to-key previous each)
    "image-spec" (image-spec-to-key previous each)
    "hidden" (text-to-key previous each)
    "select" (select-to-key previous each)
    "datepicker" (datepicker-to-key previous each)
    "select-text" (select-text-to-key previous each)
    (assoc previous (keyword (:id @each)) (set-value-type each))))


(defn handle-form-error [form response]
  (notify 422 "The submitted data is not valid")
  (doseq [each form]
      (append-error (:Errors (:response response)) each)))


(defn execute-ajax-json [url params function-handler]
  (do
    (notify 102 "Loading")
    (GET url
         :params params
         :handler (partial success-handler function-handler)
         :error-handler error-handler
         :format (json-request-format)
         :response-format (json-response-format {:keywords? true})))
  )

(defn get-json [url params function-handler & [cached]]
  (let [request-key (keyword (str url params))]
    (if cached
      (if (nil? (request-key @ajax-cache))
        (execute-ajax-json url params (fn [response]
                                        (do
                                          (swap! ajax-cache assoc request-key response)
                                          (function-handler response))))
        (success-handler function-handler (request-key @ajax-cache)))
      (execute-ajax-json url params function-handler))))



(defn prepare-api [& [try-again ]]
  (.log js/console "preparing api")
  (get-json "/api/prepare" nil
            (fn [response]
              (reset! csrf-token (:csrf-token response))
              (if try-again (try-again)))))

(declare form-post)

(defn access-denied-handler
  [url form function-handler res]
  (println res)
  (let [response (:response res)
        error-code (aget (clj->js response) "error-code")]
    (case error-code
      "csrf" (prepare-api (partial form-post url form function-handler))
      "invalid-session" (spa/redirect "/logout")
      (notify 403))))


(defn form-error-handler [form function-handler response]
  
  (let [status (:status response)]
    (case status
      200 (success-handler function-handler response)
      201 (success-handler function-handler response)
      401 (notify 401)
      403 (notify 403)
      404 (notify 404)
      405 (notify 405)
      500 (notify 500)
      409 (notify 409)
      422 (handle-form-error form response)
      0 (notify 500 "Network error")
      (notify 500 "Data send error"))))


(defn rest-url
  [url id]
  (if (nil? id)
    url
    (str url "/" id)))

(defn form-to-key
  [form]
  (reduce to-key {} form))


(defn rform-post [url form function-handler & [recap-token]]
  (do
   
    (doseq [each form]
      (remove-error  each))
    (let [params (reduce to-key {} form)]
      (POST url
        :params (assoc params :recap-token recap-token)
        :handler (partial success-handler function-handler)
        :error-handler (partial form-error-handler form function-handler)
        :format (url-request-format)
        :headers {:X-CSRF-Token (dom/get-value "csrf_token")}
        :response-format (json-response-format {:keywords? true})))))


(defn recap-form-post
  [url form function-handler]
    (notify 102 "Loading")
    (js/recap (partial rform-post url form function-handler)))

(defn form-post [url form function-handler ]
  (do
    (notify 102 "Loading")
    (doseq [each form]
      (remove-error  each))
    (POST url
      :params (reduce to-key {} form)
      :handler (partial success-handler function-handler)
      :error-handler (partial form-error-handler form function-handler)
      :format (url-request-format)
      :headers {:X-CSRF-Token (dom/get-value "csrf_token")}
      :response-format (json-response-format {:keywords? true}))))




(defn post-raw [url params function-handler]
  (do
    (notify 102 "Loading")
    (POST url
          :params params
          :handler (partial success-handler function-handler)
          :error-handler (partial error-handler function-handler)
          :format (url-request-format)
          :headers {:X-CSRF-Token (dom/get-value "csrf_token")}
          :response-format (json-response-format {:keywords? true}))))

(defn post-inputs [url inputs function-handler]
  (post-raw url (reduce to-key {} inputs) function-handler))


(defn bget-json
  [url params function-handler]
  (GET url
       :params params
       :handler function-handler
       :error-handler (fn [response] (.log js/console "Error!"))
       :format (json-request-format)
       :response-format (json-response-format {:keywords? true})))


(def delete-obj (reagent/atom {}))
(declare delete-request)


(defn delete-modal []
  [:div {:class "modal" :role "dialog" :tab-index "-1" :id "delete-modal"}
   [:div {:class "modal-dialog" :role "document"}
    [:div {:class "modal-content"}
     [:div {:class "modal-header"}
      [:h5 {:class "modal-title"} "Delete Data"]
      [:button {:type "button" :class "close" :data-dismiss "modal" :aria-label "Close"}]]
     [:div {:class "modal-body"}
      [:p "Are you sure you want to delete this?"]]
     [:div {:class "modal-footer"}
      [:button {:type "button" :class "btn btn-primary"
                :on-click #(delete-request (:url @delete-obj) (:function-handler @delete-obj))} "Delete"]
      [:button {:type "button" :class "btn btn-secondary" :data-dismiss "modal"} "Cancel"]]]]])


(defn render-element
  [elements root-id]
  (r/render
   [elements] 
   (. js/document (getElementById root-id))))


(defn render-delete-modal
  []
  (render-element delete-modal "delete-modal-container")
  (js/showDeleteModal))


(defn delete-access-denied-handler 
  [url function-handler res]
  (let [response (:response res)
        error-code (aget (clj->js response) "error-code")]
    (case error-code
      "csrf" (prepare-api (partial delete-request url function-handler))
      "invalid-session" (spa/redirect "/logout")
      (notify 403))))


(defn delete-error-handler
  [url function-handler response]
  (let [status (:status response)]
    (case status
      400 (notify 400)
      401 (notify 401)
      403 (delete-access-denied-handler url function-handler response)
      404 (notify 404)
      405 (notify 405)
      500 (notify 500)
      422 (notify 422 "Invalid delete request")
      0 (notify 500 "Network error")
      (notify 500 "Data error")))
  )


(defn delete-request [url function-handler]
  (notify 102 "Deleting")
  (DELETE url
          :handler (partial delete-success-handler function-handler)
          :headers {:X-CSRF-Token (dom/get-value "csrf_token")}
          :error-handler (partial delete-error-handler url function-handler)))


(defn delete
  [url function-handler]
  (swap! delete-obj assoc :url url :function-handler function-handler)
  (render-delete-modal))
