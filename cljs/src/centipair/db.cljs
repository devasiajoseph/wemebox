(ns centipair.db)

(def api {:user-login "/api/uauth/login"
          :user-register "/api/uauth/register"
          :user-logout "/api/uauth/logout"
          :user-status "/api/uauth/status"
          :user-profile "/api/uauth/profile"
          :change-my-password "/api/uauth/my-password"
          :user-list "/api/uauth/list"
          :user-delete "/api/uauth/delete"
          :user-create "/api/uauth/create"
          :user-update "/api/uauth/update"
          :user-edit "/api/uauth/edit"
          :change-user-password "/api/uauth/password"
          :save-base-file "/api/write/base"
          :save-static-file "/api/write/static"
          :save-template "/api/write/template"
          :save-markdown "/api/write/markdown"
          :save-dynamic "/api/write/dynamic"
          :pages "/api/pages"
          :search-pages "/api/search/pages"
          :search-images "/api/admin/search/images"
          :base-options "/api/pages/options/base"
          :template-options "/api/pages/options/template"
          :domain "/api/domain"
          :images "/api/admin/images"
          :delete-selected-images "/api/admin/delete/images"
          :open-file "/api/admin/files/open"
          :save-file "/api/admin/files/save"
          })

(defn api-paginate
  [k page]
  (str (k api) "?page=" page))
