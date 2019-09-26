package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	app:=App{}
	app.Init()
	app.Run(":9999")
}

type App struct {
	Router *mux.Router
}
type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Content interface{} `json:content`
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
func ( app *App)Run(address string )  {
	error := http.ListenAndServe(address, app.Router)
	if error!=nil{
		panic(error)
	}
}

func( app *App) initRoute()  {
	app.Router.HandleFunc("/api/shorten",app.createShortLink).Methods("POST")
	app.Router.HandleFunc("/api/info",app.getShortLinkInfo).Methods("GET")
	app.Router.HandleFunc("/{sk:[a-zA-z0-9]{1,11}}",app.redirect).Methods("GET")
}

func (app *App)createShortLink(response http.ResponseWriter, request *http.Request)  {
	fmt.Println("createShortLink")
	var req shortReq
	if error:=json.NewDecoder(request.Body).Decode(&req);error!=nil{
		returnResponseJson(response,500,error)
		return
	}
	returnResponseJson(response,200,&req.URL)
}

func (app *App)getShortLinkInfo(response http.ResponseWriter,request *http.Request)  {
	fmt.Println("getShortLinkInfo")

}

func (app *App)redirect(response http.ResponseWriter,request *http.Request)  {
	fmt.Println("redirect")

}
func returnResponseJson(response http.ResponseWriter,status int ,content interface{})  {
	res, _ := json.Marshal(BaseResponse{
		Code:    status,
		Message: http.StatusText(status),
		Content: content,
	})
	response.Header().Set("Content-type","application/json")
	_, _ = response.Write(res)
}

/*
type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Content interface{} `json:content`
}*/