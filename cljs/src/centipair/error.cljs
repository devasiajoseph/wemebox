(ns centipair.error)


(defn parse-error-message
  [key errors]
  (if (vector? (key errors))
    (first (key errors))
    (key errors)))


(defn append-default-error [field error]
  (swap! field assoc
         :message error
         :class-name "form-field-has-error"))

(defn append-checkbox-error [field error]
  (swap! field assoc
         :message error
         :class-name "checkbox-has-error"))

(defn append-field-error
  [field error]
  (case (:type @field)
    "text" (append-default-error field error)
    "checkbox" (append-checkbox-error field error)
    (append-default-error field error)))


(defn append-error [errors field]
  (doseq [error errors]
    (if (= (:Id error) (:id @field))
      (swap! field assoc :message (:Message error)))))

(defn remove-error
  [field]
  (swap! field assoc
         :message ""))