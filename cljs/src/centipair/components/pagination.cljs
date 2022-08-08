;; Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
;; Unauthorized copying of this file, via any medium is strictly prohibited
;; Proprietary and confidential
;; Written By Devasia Joseph <devasiajoseph@centipair.com>, December 2018
;; Updated By Devasia Joseph <devasiajoseph@centipair.com>, May 2020
(ns centipair.components.pagination)


(defn number-display
  [ds page]
  [:li {:class (str "page-item" (if (= (+ 1 page) (:page @ds)) "" "")) }
   [:a {:href (str "#" (:page-url @ds) "&page=" (+ 1 page))
        :id (str "pagin-"page)
        :class "page-link"
        :key (str "pagin-"page)} (+ 1 page)]])


(defn page-numbers
  [total per-page]
  (let [d (mod total per-page)
        q (quot total per-page)]
    (if (> d 0 )
      (+ 1 q)
      q)))

(defn pagination-view
  [ds]
  (if (not (nil? (:limit @ds)))
    (let [pn (count (:rows @ds))]
      [:nav 
       [:ul {:class "pagination"}
        [:li {:class (str "page-item" (if (<= (:page @ds) 1) " disabled" ""))}
         [:a {:href (if (<= (:page @ds) 1) (str "#"(:page-url @ds) "&page=1" ) (str "#"(:page-url @ds) "&page=" (- (:page @ds) 1)))
            :id (str "pagin-previous")
              :class "page-link"
              :on-click #(swap! ds assoc :page-click true)
              }
        "Previous"]]
        
        [:li {:class (str "page-item" (if (>= (:page @ds) (count (:rows @ds))) " disabled" ""))}
         [:a {:href (if (>= (:page @ds) pn) (str "#"(:page-url @ds) "&page=" (:page @ds) ) (str "#"(:page-url @ds) "&page=" (+ (js/parseInt (:page @ds)) 1)))
              :id (str "pagin-next")
              :class "page-link"
              :on-click #(swap! ds assoc :page-click true)
              } "Next"]]]])))


(defn remote-pagination-view
  [ds]
  (let [pn (page-numbers (:total @ds) (or (:limit @ds) 50))]
      [:nav 
       [:ul {:class "pagination"}
        [:li {:class (str "page-item" (if (<= (:page @ds) 1) " disabled" ""))}
         [:a {:href (if (<= (:page @ds) 1) (str "#"(:page-url @ds) "&page=1" ) (str "#"(:page-url @ds) "&page=" (- (:page @ds) 1)))
            :id (str "pagin-previous")
              :class "page-link"
              :on-click #(swap! ds assoc :page-click true)
              }
        "Previous"]]
        
        [:li {:class (str "page-item" (if (>= (:page @ds) pn) " disabled" ""))}
         [:a {:href (if (>= (:page @ds) pn) (str "#"(:page-url @ds) "&page=" (:page @ds) ) (str "#"(:page-url @ds) "&page=" (+ (js/parseInt (:page @ds)) 1)))
              :id (str "pagin-next")
              :class "page-link"
              :on-click #(swap! ds assoc :page-click true)
              } "Next"]]]]))



(defn prepare-obj-list
  [response obj-list limit &[page-url]]
  (reset! obj-list response)
    (swap! obj-list assoc :limit limit :page-url page-url))



