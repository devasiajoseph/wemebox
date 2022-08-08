(ns centipair.dom
  (:require [goog.dom :as gdom]
            [clojure.string :as string]))


(defn object-exists?
  [element-id]
  (not (nil? (.getElementById js/document element-id))))

(defn get-element
  [id]
  (gdom/getElement id))

(defn get-value
  [id]
  (if (nil? (gdom/getElement id))
    nil
    (.-value (gdom/getElement id))))


(defn set-value [id value]
  (set! (.-value ( gdom/getElement id)) value))


(defn innerHtml [id value]
  (set! (.-innerHTML ( gdom/getElement id)) value))


(defn get-element-by-class
  [class-name]
  (gdom/getElementByClass class-name))


(defn set-inner-html
  [dom-obj value]
  (set! (.-innerHTML dom-obj) value))


(defn show-field
  [field class-key]
  (if (not (nil? (class-key @field)))
    (swap! field assoc class-key (string/replace (class-key @field) #" hide" ""))))


(defn hide-field [field class-key]
  (swap! field assoc class-key (str (class-key @field) " hide")))


(defn show-element
  [element-id]
  (if (object-exists? element-id)
    (set! (.-display (.-style (.getElementById js/document element-id))) "block")))

(defn hide-element
  [element-id]
  (if (object-exists? element-id)
    (set! (.-display (.-style (.getElementById js/document element-id))) "none")))


(defn set-title
  [value]
  (set! (.-title js/document) value))