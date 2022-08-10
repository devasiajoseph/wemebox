(ns app.home
   (:require [reagent.core :as r]
             [centipair.ui :as ui]))

(defonce email (r/atom {:type "text" :id "email" :class "form-control"}))

(defn home-page []
  [:div {:class "mt-5"}
   [:h2 "Welcome to your website dashboard "]
   [:p "Use the side menu to navigate"]])


(defn render-home []
  (ui/render-ui home-page "app"))

(defn about-page []
  [:div {:class "mt-5"}
   [:h2 "Welcome to about"]
   [:p "This is about page"]])


(defn render-about [match]
  (println match)
  (ui/render-ui about-page "app"))


