package shortlink

import (
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	Router *mux.Router
}

type shortReq struct {
	URL string  `json:"url" validate:nonzero`
	ExpirationInMinute int64 `json:"expiration_inminutes" validate:"min=0"`
}

type shortLinkResponse struct {
	ShortLink string  `json:"short_link"`
}

func( app *App) Init()  {
	app.Router=mux.NewRouter()
	app.initRoute()
}

func( app *App) initRoute()  {
	app.Router.HandleFunc("/api/shorten",app.createShortLink).Methods("POST")
	app.Router.HandleFunc("/api/info",app.getShortLinkInfo).Methods("GET")
	app.Router.HandleFunc("/{sk:[a-zA-z0-9]{1,11}}",app.redirect).Methods("GET")
}

func (app *App)createShortLink(http.ResponseWriter, *http.Request)  {

}

func (app *App)getShortLinkInfo(http.ResponseWriter, *http.Request)  {

}

func (app *App)redirect(http.ResponseWriter, *http.Request)  {

}