(ns centipair.rest
  (:require [centipair.ajax :as ajax]))



(def url (atom {}))


(defn obj-url
  [obj-key id]
  (str (obj-key url) "/" id))

(defn obj-option-url
  [obj-key id]
  (str (obj-key url) "/" id))

(defn get-obj
  [obj-key id callback]
  (ajax/get-json (obj-url obj-key id) nil callback))


(defn rest-url
  [obj-key id]
  (if (:value @id)
    (str (obj-key url) "/" (:value @id))
    (obj-key url)))

(defn search
  [key q]
  (str "/api/search/" (name key) "?q=" q))


(defn object
  [obj & [id]]
  (if (:value @id)
    (str  "/api/" obj "/" (:value @id))
    (str "/api/" obj)))

(defn object-options
  [obj & [params]]
  (str "/api/options/" obj))
