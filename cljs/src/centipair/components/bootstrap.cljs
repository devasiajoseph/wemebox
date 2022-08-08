(ns centipair.components.bootstrap
  (:require [centipair.components.input :as cin]
            [centipair.ajax :as ajax]))


(defn checkbox
  [field]
  ^{:key (:id @field)}
  [:div {:class "form-check form-check-inline"}
   [:input {:type "checkbox"
            :value (:value @field)
            :on-change #(cin/update-check field (-> % .-target .-checked) )
            :checked (:checked @field)
            :class-name "form-check-input"
            :id (:id @field)}]
   [:label {:class-name "form-check-label" :for (:id @field)}
    
    (str " " (:label @field))]
   [:br]
   [:label {:class-name "checkbox-message-label"} (:message @field)]])



(defn text
  [field]
  ^{:key (:id @field)}
  [:div {:class ""}
   [:label {:for (:id @field) :class "form-label"} (:label @field)]
   [:input {:class "form-control"
             :type (:type @field)
             :id (:id @field)
             :placeholder
             (if (nil? (:placeholder @field))
               ""
               (:placeholder @field))
             :value (:value @field)
             :on-change #(cin/update-value field (-> % .-target .-value) )
             :maxLength (or (:maxLength @field) 255)
             :disabled (if (:disabled @field) "disabled" "")}]
   
   [:div {:class-name "invalid-field"} (:message @field)]
   ])

(defn date-picker
  [field]
  ^{:key (:id @field)}
  [:div {:class ""}
   [:label {:for (:id @field) :class "form-label"} (:label @field)]
   [:input {:class "form-control"
            :type (:type @field)
            :id (:id @field)
            :placeholder
            (if (nil? (:placeholder @field))
              ""
              (:placeholder @field))
            :value (:value @field)
            :on-change #(cin/update-value field (-> % .-target .-value) )
            :maxLength (or (:maxLength @field) 255)
            :disabled (if (:disabled @field) "disabled" "")
            :data-provide "datepicker"
            :data-date-format (or (:data-date-format @field) "dd-mm-yyyy")
            }]
   
   [:div {:class-name "invalid-field"} (:message @field)]])

(defn textarea
  [field]
  ^{:key (:id @field)}
  [:div {:class "form-group"}
   [:label {:for (:id @field) :class "label"} (:label @field)]
   [:textarea {:class "form-control"
               :id (:id @field)          
               :rows (or (:rows @field) "10")
               :value (:value @field)
               :on-change #(cin/update-value field (-> % .-target .-value) )
               :key (:id @field)
               }]
   [:div {:class-name "invalid-field"} (:message @field)]])


(defn select
  [field]
  ^{:key (:id @field)}
  [:div {:class ""}
   [:label {:for (:id @field) :class "form-label"} (:label @field)]
   [:select {:class "form-control"
             :key (:id @field)
             :id (:id @field)
             :on-change #(cin/select-on-change field (-> % .-target .-value) )
             :value (or (:value @field) 0)
             :disabled (if (:disabled @field) "disabled" "")}
    (doall (map (partial cin/select-option field) (:options @field)))]
   [:div {:class-name "invalid-field"} (:message @field)]])


(defn autocomplete-search
  [field]
  (if (> (count (:value @field)) 2)
    (ajax/get-json (:remote @field)
                 (assoc (:remote-params @field) :q (:value @field))
                 (fn [response]
                   (swap! field assoc :options response)))))


(defn autocomplete-option [field index option]
  [:option {:key (str (:id @field) "-" index)
            :value (:value option)
            } (:label option)])


(defn autocomplete
  [field]
  ^{:key (:id @field)}
  [:div {:class "form-group"}
   [:label {:for (:id @field) :class "form-label"} (:label @field)]
   [:input {:class "form-control"
            :type (:type @field)
            :id (:id @field)
            :placeholder
            (if (nil? (:placeholder @field))
              ""
              (:placeholder @field))
            :value (:value @field)
            :on-change #(do
                          (cin/update-value field (-> % .-target .-value))
                          (autocomplete-search field))
            :maxLength (or (:maxLength @field) 255)
            :disabled (if (:disabled @field) "disabled" "")
            :list (:data-list-id @field)}]
   
   [:div {:class-name "invalid-field"} (:message @field)]
   [:datalist {:id (:data-list-id @field)
               :key (:data-list-id @field)}
    (doall (map-indexed (partial autocomplete-option field) (:options @field)))]])




(defn button [action-button form-fields]
  [:a {:class "btn btn-primary"
       :on-click #(cin/perform-action (:on-click @action-button) form-fields)
       :disabled (if (nil? (:disabled action-button)) "" "disabled")
       :key (:id @action-button)}
   (:label @action-button)])

(defn input-field [field]
  (case (:type @field)
    "text" (text field)
    ;;"text-button" (text-button field)
    "number" (text field)
    ;;"email" (text field)
    "textarea" (textarea field)
    "password" (text field)
    "date" (text field)
    "datepicker" (date-picker field)
    "checkbox" (checkbox field)
    "select" (select field)
    ;;"select-action" (select-action field)
    ;;"radio" (radio field)
    ;;"hidden" (hidden field)
    ;;"subheading" (subheading field)
    ;;"image-spec" (image-spec field)
    ;;"description" (description field)
    ;;"datepicker" (date-picker field)
    ;;"markdown" (markdown-editor field)
    ;;"select-text" (select-text field)
    "autocomplete" (autocomplete field)
    ;;"ajax-autocomplete" (ajax-autocomplete field)
    ;;"link-text" (link-text field)
    ))



(defn form-default [form-fields action-button]
  [:form
   [:h1 {:class "title"} (or (:form-title @action-button) "")]
   (doall (map input-field form-fields))
   (button action-button form-fields)])
