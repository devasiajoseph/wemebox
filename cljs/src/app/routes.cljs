(ns app.routes
  (:require [centipair.router :as router]
            [centipair.rest :as rest]
            [app.home :as home]
            [app.uauth :as uauth]))


(def rest-urls {:user-login "/api/uauth/login"
          :user-logout "/api/uauth/logout"
          :user-register "/api/uauth/register"
          :user-status "/api/uauth/status"})


(def rts {"/" {:run home/render-home}
          "/login" {:run uauth/render-login-form}
          "/register" {:run uauth/render-register-form}})

(defn init!
  []
  (rest/init! rest-urls)
  (router/init! rts))