(ns centipair.components.page
  (:require [reagent.core :as reagent]
            [centipair.ui :as ui]))


(def page-data (reagent/atom {:header ""}))


(defn card-container
  "Card UI container for dashboard"
  [max-width]
  [:div {:class "card shadow mb-4" :style {:max-width (or max-width "100%")}}
   [:div {:class "card-header"} [:h2 (:header @page-data)]]
   [:div {:class "card-body"}
    ((:components @page-data))]])


(defn page-component 
  [title components]
  (ui/page-title title)
  [:div {:class "card shadow mb-4"}
   [:div {:class "card-header"} [:h2 title]]
   [:div {:class "card-body"}
    (components)]])

(defn render 
  [title components root &[max-width]]
  (ui/page-title title)
  (swap! page-data assoc :header title :components components)
  (ui/render (partial card-container max-width) root))