(ns centipair.router
  (:require [clojure.string :as str]
            [centipair.ui :as ui]
            [centipair.rest :as rest]))


(def routes (atom {}))
(def loaded (atom true))


(defn get-url
  []
  (.. js/window -location -hash))

(defn no-params
  [params] 
  (and (= 1 (count (to-array params))) (= nil (second (first (to-array params))))))

(defn parse-params
  "Parse URL parameters into a hashmap"
  [url]
  (let [param-strs (-> url (str/split #"\?") last (str/split #"\&"))
        params (into {} (for [[k v] (map #(str/split % #"=") param-strs)]
                          [(keyword k) v]))]
    (if (no-params params)
      {}
      params)))


(defn parse-url
  [url]
  (let [curl (str/replace url #"#" "")
        url-obj (new js/URL (str "http://localhost" curl))]
    {:path (.-pathname url-obj)
     :params (parse-params curl)}))

(defn render-default 
  []
  (println "default render"))


(defn run
  [url rt]
  (let [p (if (= url "")
            (get @routes "/")
            (get @routes (:path rt)))] 
    (if (:q p)
      ((:run p) (:params rt))
      ((:run p)))))



(defn load-hash
  [url]
  (ui/refresh-all)
  (let [rt  (parse-url url)
        p (if (= url "")
            (get @routes "/")
            (get @routes (:path rt)))]
    (if (:q p)
      ((:run p) (:params rt))
      ((:run p)))))

(defn window-load
  []
  (.addEventListener
   js/window
   "load"
   (fn [] (let [url (get-url)]
            (reset! loaded false)
            (load-hash url)))))

(defn window-hash-change
  []
  
  (.addEventListener
   js/window
   "hashchange"
   (fn [] (let [url (get-url)] 
            (load-hash url)))))

(defn add-routes 
  [rts]
  (reset! routes rts))

(defn init!
  [rts]
  (add-routes rts)
  (window-load)
  (window-hash-change)
  (if (not @loaded)
   (load-hash (get-url))))


(defn render-menu [])