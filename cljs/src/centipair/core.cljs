(ns centipair.core
  (:require [app.routes :as routes]
            [centipair.components.notifier :as notifier]
            [centipair.control :as control]))

(control/load-auth)
(routes/init!)
(notifier/render-notifier-component)