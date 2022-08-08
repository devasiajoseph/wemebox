(ns centipair.components.buttons)



(defn icon-button
  [field]
  [:button {:class (or (:class @field) "btn btn-primary") :on-click (:on-click @field)}
   [:i {:class (str "fas "(:icon @field))}] (str " " (:label @field))])



(defn show-more
  [field]
  ((:on-click @field))
  (swap! field assoc :open (not (:open @field)))
  (if (:open @field)
    (swap! field assoc :icon "fas fa-xmark" :label "close")
    (swap! field assoc :icon "fas fa-chevron-down" :label "show more")))


(defn show-more-button
  [field]
  [:button {:class (or (:class @field) "btn btn-primary") :on-click #(show-more field)}
   [:i {:class (or (:icon @field) "fas fa-chevron-down")}] (str " "(or (:label @field) "show more"))])