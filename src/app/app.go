package app

import (
	"log"
	"net/http"

	"goproject/app/handler"
	"goproject/app/model"
	"goproject/config"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	// dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
	// 	config.DB.Username,
	// 	config.DB.Password,
	// 	config.DB.Host,
	// 	config.DB.Port,
	// 	config.DB.Name,
	// 	config.DB.Charset,
	// )

	//db, err := gorm.Open(config.DB.Dialect, dbURI)
	db, err := gorm.Open("postgres", "host=kandula.db.elephantsql.com port=5432 user=plrvuppn password=DyhDQ6VlBGElGdX-qTJSjB5mR1fAvkrd dbname=plrvuppn")
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/persons", a.handleRequest(handler.GetAllPersons))
	a.Post("/persons", a.handleRequest(handler.CreatePerson))
	a.Get("/persons/{PersonID}", a.handleRequest(handler.GetPerson))
	a.Put("/persons/{PersonID}", a.handleRequest(handler.UpdatePerson))
	a.Delete("/persons/{PersonID}", a.handleRequest(handler.DeletePerson))

	a.Get("/genders", a.handleRequest(handler.GetAllGenders))
	a.Post("/genders", a.handleRequest(handler.CreateGender))
	a.Get("/genders/{GenderID}", a.handleRequest(handler.GetGender))
	a.Put("/genders/{GenderID}", a.handleRequest(handler.UpdateGender))
	a.Delete("/genders/{GenderID}", a.handleRequest(handler.DeleteGender))

	a.Get("/cities", a.handleRequest(handler.GetAllCities))
	a.Post("/cities", a.handleRequest(handler.CreateCity))
	a.Get("/cities/{CityID}", a.handleRequest(handler.GetCity))
	a.Put("/cities/{CityID}", a.handleRequest(handler.UpdateCity))
	a.Delete("/cities/{CityID}", a.handleRequest(handler.DeleteCity))

	a.Get("/statuses", a.handleRequest(handler.GetAllStatues))
	a.Post("/statuses", a.handleRequest(handler.CreateStatus))
	a.Get("/statuses/{StatusID}", a.handleRequest(handler.GetStatus))
	a.Put("/statuses/{StatusID}", a.handleRequest(handler.UpdateStatus))
	a.Delete("/statuses/{StatusID}", a.handleRequest(handler.DeleteStatus))

	a.Get("/maritalstatuses", a.handleRequest(handler.GetAllMaritalStatus))
	a.Post("/maritalstatuses", a.handleRequest(handler.CreateMaritalStatus))
	a.Get("/maritalstatuses/{MaritalStatusID}", a.handleRequest(handler.GetMaritalStatus))
	a.Put("/maritalstatuses/{MaritalStatusID}", a.handleRequest(handler.UpdateMaritalStatus))
	a.Delete("/maritalstatuses/{MaritalStatusID}", a.handleRequest(handler.DeleteMaritalStatus))

	a.Get("/nationalities", a.handleRequest(handler.GetAllNationalities))
	a.Post("/nationalities", a.handleRequest(handler.CreateNationality))
	a.Get("/nationalities/{NationalityID}", a.handleRequest(handler.GetNationality))
	a.Put("/nationalities/{NationalityID}", a.handleRequest(handler.UpdateNationality))
	a.Delete("/nationalities/{NationalityID}", a.handleRequest(handler.DeleteNationality))

	a.Get("/users", a.handleRequest(handler.GetAllUsers))
	a.Post("/users", a.handleRequest(handler.CreateUser))
	a.Get("/users/{UserID}", a.handleRequest(handler.GetUser))
	a.Put("/users/{UserID}", a.handleRequest(handler.UpdateUser))
	a.Delete("/users/{UserID}", a.handleRequest(handler.DeleteUser))

}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
