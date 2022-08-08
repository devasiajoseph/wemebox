(ns app.init
  (:require [app.routes :as routes]))


(defn init! 
  []
  (routes/init!))