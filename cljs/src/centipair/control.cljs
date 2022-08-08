(ns centipair.control
  (:require-macros [cljs.core.async.macros :refer [go]])
  (:require [reagent.core :as reagent]
            [cljs.core.async :refer [chan <! >!]]
            [centipair.ajax :as ajax]
            [centipair.spa :as spa]
            [centipair.db :as db]
            ))


(def auth-waiting-counter (atom 0))
(def auth-channel (chan))

(def navigation-channel (chan))


(def auth (atom {}))


(def dashboard-menu (reagent/atom []))

(def backoffice-links (reagent/atom []))

(defn auth-function
  [f]
  (if (empty? @auth)
    (do
      (swap! auth-waiting-counter inc)
      (go
        (f (<! auth-channel))))
    (do
      (reset! auth-waiting-counter 0)
      (f @auth))))


(defn activate-navigation
  []
  (if (:LoggedIn @auth)
    (let [nav-menu @dashboard-menu]
      (go (>! navigation-channel nav-menu)))))


(defn handle-auth
  [response]
  (if (:LoggedIn response)
    (do
      (reset! auth response)
      ;;(activate-navigation)
      (dotimes [n @auth-waiting-counter]
        (go (>! auth-channel @auth))))
    (spa/render-initial)))


(def auth-url (db/api :user-status))

(defn load-auth [& [auth-handler]]
  (ajax/get-json auth-url nil (or auth-handler handle-auth)))

