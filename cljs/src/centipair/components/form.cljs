(ns centipair.components.form)



(defn card-form
  [width elements]
  [:div {:class "row d-flex justify-content-center"}
   [:div {:class (str "col-md-" width)}
    [:div {:class "card px-5 py-5"} (elements)]]])