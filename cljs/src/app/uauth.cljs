(ns app.uauth
  (:require [reagent.core :as r]
            [centipair.ui :as ui]
            [centipair.components.raw :as raw]
            [centipair.components.input :as input]
            [centipair.components.form :as form]
            [centipair.validators :as v]
            [centipair.ajax :as ajax]
            [centipair.control :as control]
            [centipair.spa :as spa]
            [app.api :as api]))

(def phone (r/atom {:id "phone" :type "text" :class "form-control form-control-lg" :validator v/required}))
(def email (r/atom {:id "email" :type "text" :class "form-control form-control-lg" :validator v/email-required}))
(def password (r/atom {:id "password" :type "password" :class "form-control form-control-lg" :validator v/required}))

(defn login 
  []
  (ajax/form-post (api/url :user-login) [phone password] 
                  (fn [response]
                    (control/handle-auth response)
                    (spa/redirect "/"))))

(def login-button (r/atom {:id "login-button" :label "Login" :on-click login}))

(defn login-form 
  []
  [:form
   [:div {:class " mb-4"} [:label "Phone"]
    (raw/text phone)
    [:div {:id "phone-error" :class "invalid-field"} (:message @phone)]]
   [:div {:class "mb-4"} [:label "Password"]
    (raw/text password)
    [:div {:id "phone-error" :class "invalid-field"} (:message @password)]]
   (input/button login-button [phone password])
   [:div {:class "mb-4"} "Don't have an account? " [:a {:href "#/register"} "Signup"]] ])

(defn login-page
  []
  (form/card-form 6 login-form))


(defn render-login-form
  []
  (ui/render-ui login-page "app"))



(defn register
  []
  (ajax/form-post (api/url :user-register) [phone email password] (fn [response])))

(def register-button (r/atom {:id "login-button" :label "Sign up" :on-click register}))

(defn register-form
  []
  [:form
   [:div {:class " mb-4"} [:label "Phone"]
    (raw/text phone)
    [:div {:id "phone-error" :class "invalid-field"} (:message @phone)]]
   [:div {:class " mb-4"} [:label "Email"]
    (raw/text email)
    [:div {:id "email-error" :class "invalid-field"} (:message @email)]]
   [:div {:class "mb-4"} [:label "Password"]
    (raw/text password)
    [:div {:id "password-error" :class "invalid-field"} (:message @password)]]
   (input/button register-button [phone password email])
   [:div {:class "mb-4"} "Already registered? " [:a {:href "#/login"} "Login"]]])

(defn register-page
  []
  (form/card-form 6 register-form))

(defn render-register-form
  []
  (ui/render-ui register-page "app"))


