(ns centipair.ui
  (:require 
   [reagent.dom :as r]
   [centipair.dom :as dom]
   [centipair.components.notifier :refer [notify]]
   [centipair.control :as d-chan]))

(defn render-ui
  [elements root-id]
  (notify 200)
  (r/render
   [elements] 
   (. js/document (getElementById root-id))))


(defn unauthorized-view
  []
  [:div
   [:h1 "Unauthorized"]
   [:div "Please " [:a {:href "/#/login"} "login"]]])


(defn ui-authorized?
  [permission-key auth]
  (if (nil? permission-key)
    (:LoggedIn auth)
    (permission-key auth)))

(defn auth-render
  [elements root-id permission-key auth]
  (if (ui-authorized? permission-key auth)
    (render-ui elements root-id)
    (render-ui unauthorized-view root-id)))


(defn render
  [elements root-id &[permission-key]]
  (d-chan/auth-function (partial auth-render elements root-id permission-key)))

(defn blank-ui[]
  [:span])

(defn refresh-ui [root-id]
  (r/render
   [blank-ui] 
   (. js/document (getElementById root-id))))

(defn refresh-all
  []
  (refresh-ui "app")
  (refresh-ui "app1"))


(defn page-title
  [value]
  (dom/set-title (str "Wemebox - " value)))

(defn show-loader
  []
  (dom/show-element "loader"))

(defn hide-loader
  []
  (dom/hide-element "loader"))
