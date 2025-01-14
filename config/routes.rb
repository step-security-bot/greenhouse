Rails.application.routes.draw do
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Reveal health status on /up that returns 200 if the app boots with no exceptions, otherwise 500.
  # Can be used by load balancers and uptime monitors to verify that the app is live.
  get "up" => "rails/health#show", :as => :rails_health_check

  mount Audits1984::Engine => "/admin/console-audit"
  mount MissionControl::Jobs::Engine, at: "/admin/jobs"

  mount LetterOpenerWeb::Engine, at: "/letter_opener" if Rails.env.development?

  namespace :admin do
    # resources :users, only: %i[index show]
    # resources :active_sessions, only: %i[index destroy]

    # root to: "users#index"
  end

  # Defines the root path route ("/")
  root "root#index"

  resources :search
  # resources :details

  resources :users, only: [:index, :new, :create]

  get "login", to: "authentications#new"
  get "logout", to: "authentications#destroy"

  get "enter", to: "root#enter"

  resources :authentications
end
