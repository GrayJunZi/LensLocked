package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/grayjunzi/lenslocked/controllers"
	"github.com/grayjunzi/lenslocked/migrations"
	"github.com/grayjunzi/lenslocked/models"
	"github.com/grayjunzi/lenslocked/templates"
	"github.com/grayjunzi/lenslocked/views"
)

func main() {
	// 设置数据库
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// 设置服务
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	// 设置控制器
	usersController := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersController.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))
	usersController.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))

	// 设置中间件
	userMiddleware := controllers.UserMiddleware{
		SesionService: &sessionService,
	}

	csrfKey := "the lenslocked csrf key"
	csrfMiddleware := csrf.Protect([]byte(csrfKey), csrf.Secure(false))

	// 设置路由
	r := chi.NewRouter()
	r.Use(csrfMiddleware)
	r.Use(userMiddleware.SetUser)
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))
	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"contact.gohtml", "tailwind.gohtml",
	))))
	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(
		templates.FS,
		"faq.gohtml", "tailwind.gohtml",
	))))

	r.Get("/signup", usersController.New)
	r.Post("/users", usersController.Create)
	r.Get("/signin", usersController.SignIn)
	r.Post("/signin", usersController.ProcessSignIn)
	r.Post("/signout", usersController.ProcessSignOut)

	// r.Get("/users/me", usersController.CurrentUser)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(userMiddleware.RequireUser)
		r.Get("/", usersController.CurrentUser)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	// 启动服务
	fmt.Println("Starting the server on :3000 ...")
	http.ListenAndServe(":3000", r)
}
