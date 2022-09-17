(ns wemebox.website
  (:require [reagent.core :as r]
            [centipair.components.page :as page]
            [centipair.components.raw :as raw]
            [centipair.components.input :as input] 
            [centipair.validators :as v]
            [centipair.ajax :as ajax] 
            [centipair.spa :as spa]))



(defn save-website 
  [])


(def website-list (r/atom []))
(def domain-name (r/atom {:id "domain-name" :type "text" :class "form-control" :validator v/required}))
(def domain-active (r/atom {:id "domain-active" :type "text" :class "form-control" }))
(def website-button (r/atom {:id "login-button" :label "Save" :on-click save-website}))

(defn website-form
  []
  [:form
   [:div {:class " mb-4"} [:label "Domain Name"]
    (raw/text domain-name)
    [:div {:id "domain-name-error" :class "invalid-field"} (:message @domain-name)]]
   [:div {:class "form-check"} [:label "Active"]
    (raw/checkbox domain-active)
    [:div {:id "phone-error" :class "invalid-field"} (:message @domain-active)]]
   (input/button website-button [domain-name domain-active])])


(defn render-website-form
  [q]
  (println q)
  (page/render "Manage Website" website-form "app" "500px"))


(defn website-list-view
  []
  [:div 
   [:table {:class "table"}]]
  )