package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/jimmyhogerty/lenslocked/controllers"
	"github.com/jimmyhogerty/lenslocked/migrations"
	"github.com/jimmyhogerty/lenslocked/models"
	"github.com/jimmyhogerty/lenslocked/templates"
	views "github.com/jimmyhogerty/lenslocked/views"
)

func main() {
	// ~ESTABLISH DB CONNECTION~
	cfg := models.DefaultPostgresConfig()
	fmt.Println(cfg.String())
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// ~SETUP SERVICES~
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	// ~SETUP MIDDLEWARE~
	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}
	var csrfKey = "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect([]byte(csrfKey),
		// INFO
		// If developing locally, set csrf.Secure to false.
		// If deploying to HTTPS, set csrf.Secure to true.
		csrf.Secure(false))

	// ~SETUP CONTROLLERS~
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS, "signin.gohtml", "tailwind.gohtml"))

	// ~SETUP ROUTES~
	r := chi.NewRouter()
	r.Use(csrfMw)
	r.Use(umw.SetUser)
	// ~PUBLIC ROUTES~
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))
	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)
	// ~PRIVATE ROUTES~
	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})
	// ~OOPSIES~
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// ~START SERVER~
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
