;; Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
;; Unauthorized copying of this file, via any medium is strictly prohibited
;; Proprietary and confidential
;; Written By Devasia Joseph <devasiajoseph@centipair.com>, December 2018

;; ==================
;; Purpose and Design
;; ==================
;; Notification UI
(ns centipair.components.notifier
  (:require [reagent.dom :as r]
            [reagent.core :as reagent]
            ))


(def notifier-state (reagent/atom {:class "notify" :text ""}))
(def alert-state (reagent/atom {:type "" :message ""}))

(defn notifier-component 
  []
  [:div {:class (:class @notifier-state)} (:text @notifier-state)])

(defn render-notifier-component 
  []
  (r/render
   [notifier-component]
   (. js/document (getElementById "notifier"))))


(defn notify [code & [message]]
    (case code
      201 (reset! notifier-state {:class "notify notify-loading" :text (or message "Saved")})
      200 (reset! notifier-state {:class "notify" :text ""})
      204 (reset! notifier-state {:class "notify notify-info" :text "No content"})
      102 (reset! notifier-state {:class "notify notify-loading" :text (or message "Loading")})
      404 (reset! notifier-state {:class "notify notify-error" :text (or message "Resource Unavailable")})
      500 (reset! notifier-state {:class "notify notify-error" :text (or message "Internal Server Error")})
      422 (reset! notifier-state {:class "notify notify-error" :text (or message "Unprocessable Entity")})
      400 (reset! notifier-state {:class "notify notify-error" :text (if (nil? message) "Illegal operation" message)})
      401 (reset! notifier-state {:class "notify notify-error" :text (if (nil? message) "Not Authorized" message)})
      403 (reset! notifier-state {:class "notify notify-error" :text (if (nil? message) "Access Denied" message)})
      405 (reset! notifier-state {:class "notify notify-error" :text (if (nil? message) "Method not allowed" message)})
      (reset! notifier-state {:class "notify" :text ""})))


(defn alert-component
  []
  [:div {:class (str "alert alert-" (:type @alert-state)) :role "alert"} (:message @alert-state)])

(defn alert
  [type message]
  (reset! alert-state {:type type :message message})
  (r/render
   [alert-component]
   (. js/document (getElementById "alert"))))


(def info-data (reagent/atom {:title "" :text ""}))

(defn delete-modal []
  [:div {:class "modal" :role "dialog" :tab-index "-1" :id "info-modal"}
   [:div {:class "modal-dialog" :role "document"}
    [:div {:class "modal-content"}
     [:div {:class "modal-header"}
      [:h5 {:class "modal-title"} (:title @info-data)]
      [:button {:type "button" :class "close" :data-dismiss "modal" :aria-label "Close"}]]
     [:div {:class "modal-body"}
      [:p (:text @info-data)]]
     [:div {:class "modal-footer"}
      [:button {:type "button" :class "btn btn-secondary" :data-dismiss "modal"} "Cancel"]]]]])



(def info-state (reagent/atom {:title "" :body "" }))

(defn hide-info
  [])
(defn info-modal
  []
  (reagent/create-class 
   {:component-did-mount (fn [this] (println "created"))
    :reagent-render   (fn []  [:div {:class "modal" :id "info-modal"}
                              
                               [:div {:class "modal-dialog"}
                                [:div {:class "modal-content"}
                                 [:div {:class "modal-header"}
                                  ;;[:h5 {:class "modal-title"} (:title @info-state)]
                                  [:button {:class "btn-close"
                                            :on-click #(js/hideInfo)

                                            :aria-label "Close"

                                            :type "button"}]]
                                 [:div {:class "modal-body"}
                                  [:div {:class "d-flex justify-content-center align-items-center"}
                                   [:h5 (:body @info-state)]]]
                                 [:div {:class "modal-footer d-flex justify-content-center align-items-center"}
                                  [:button {:class "btn btn-primary" :on-click  #(js/hideInfo) :type "button"}

                                   "Ok"]]]]])}))



(defn show-info
  []
  (r/render
   [info-modal]
   (. js/document (getElementById "info-modal-container")))
  (js/showInfo))

(defn saved
  [info]
  (swap! info-state assoc :title "Info" :body info )
  (set! (.. js/document -documentElement -scrollTop) 0)
  (show-info))