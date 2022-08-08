(ns centipair.components.table
  (:require [centipair.ajax :as ajax]
            [centipair.spa :as spa]
            [centipair.dom :as dom]
            [centipair.components.pagination :as pagination]))



(defn set-data
  [field rows]
  (swap! field assoc :rows (partition-all (js/parseInt (:limit @field)) rows)))


(defn table-rows
  [field]
  [:tbody
   (if (> (count (:rows @field)) 0)
     (doall
      (for [each-row (nth (:rows @field) (- (:page @field) 1))]
        ((:row @field) each-row))))])


(defn data-table-view
  [field]
  [:div 
   [:div {:class "row"}
    [:div {:class "col-sm-12 col-md-6"}
     [:label "Show"]
     [:select {:class "form-control"
               :key (:id (str "select-" @field))
               :id (:id (str "number-" @field))
               :on-change #(do
                             (swap! field assoc :limit (-> % .-target .-value))
                             (set-data field (apply concat (:rows @field)))
                             )
               :value (or (:limit @field) 25)}
      [:option {:value 10} 10]
      [:option {:value 25} 25]
      [:option {:value 50} 50]
      [:option {:value 100} 100]]]
    [:div {:class "col-sm-12 col-md-6"}
     ;;[:label "Search"]
     ;;[:input {:type "search" :class "form-control form-control-sm"}]
     ]]
   [:table {:class "table table-bordered"}
    ((:header @field))
    (table-rows field)]
   (pagination/pagination-view field)])


(defn set-page
  [field page]
  (swap! field assoc :page page))



(defn load-remote-data
  [field & [qp]]
  (reset! field (merge @field qp))
  (let [params {:page (:page @field)
                :limit (:limit @field)}
        q-params (spa/params->query-string params)]
    
    (ajax/get-json
     (:api-url @field)
     params
     (fn [response]
       (swap! field assoc :rows (:data response) :total (:total response))))))


(defn remote-table-rows
  [field]
  [:tbody (doall (for [each-row (:rows @field)]
                   ((:row-ui @field) each-row)))])


(defn search-remote
  [field]
  (spa/redirect (str (:search-url @field) "&q=" (dom/get-value "remote-search-string"))))


(defn update-search
  "Updates value of the input element atom"
  [field value]
  (swap! field assoc :q value))


(defn remote-table-view
  ;; field will contain the following data :>
  ;; rows : data 
  ;; row-ui : ui for each row
  ;; header-ui: ui for header
  ;; api-url : url for the data
  ;; page : current page of data rows
  ;; limit : limit the number of rows
  ;; q : search query
  [field]
  [:div
   (if (:disable-search @field)
     [:span]
     [:div {:class "input-group mb-3"}
      [:input {:type "text" :class "form-control form-control-sm" :placeholder (or (:search-placeholder @field) "Search")
               :id "remote-search-string" :value (:q @field) :on-change #(update-search field (-> % .-target .-value))}]
      [:div {:class "input-group-append"}
       [:a  {:class "input-group-text" :href "javascript:void(0)" :on-click #(search-remote field)}
        [:i {:class "fa fa-search"}] ]]])
   [:table {:class "table table-bordered data-table"}
    ((:header-ui @field))
    (remote-table-rows field)]
   (pagination/remote-pagination-view field)])



(defn set-remote-table
  [field page-url header-ui row-ui]
  (swap! field assoc
         :page-url (str page-url "?limit=" (:limit @field))
         :header-ui header-ui
         :row-ui row-ui))
