(ns app.routes
  (:require [centipair.router :as router]
            [app.home :as home]
            [app.uauth :as uauth]))



(def rts {"/" {:run home/render-home}
          "/login" {:run uauth/render-login-form}
          "/register" {:run uauth/render-register-form}})

(defn init!
  []
  (router/init! rts))