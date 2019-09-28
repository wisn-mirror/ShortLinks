package mux

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	Content interface{} `json:"content"`
}

type shortReq struct {
	URL string  `json:"url" validate:"nonzero"`
	Name string  `json:"name"`
	Password string  `json:"password"`
	List   []listTemp `json:"list"`
	ExpirationInMinute int64 `json:"expiration_inminutes" validate:"min=0"`
}

type listTemp struct {
	Name string  `json:"name"`
}

type SLInfo struct {
	Key    string     `json:"key"`
	Target string     `json:"target"`
	List   []listTemp `json:"list"`
}
var ns NameSpace


type NameSpace struct {
	data map[string]string
}

func (ns *NameSpace)add(key  string,value string )  {
	ns.data[key]=value
}

func (ns *NameSpace)get(key  string )string  {
	return ns.data[key]
}


func( app *App) Init()  {
	app.Router = mux.NewRouter()
	app.initRoute()
	ns = NameSpace{
		data: make(map[string]string),
	}
}
func ( app *App)Run(address string )  {
	error := http.ListenAndServe(address, app.Router)
	if error != nil {
		panic(error)
	}
}

func( app *App) initRoute()  {
	app.Router.HandleFunc("/api/shorten",app.createShortLink).Methods("POST")
	app.Router.HandleFunc("/api/info",app.getShortLinkInfo).Methods("GET")
	app.Router.HandleFunc("/s/{sk:[a-zA-z0-9]{1,11}}",app.redirect).Methods("GET")
}

func (app *App)createShortLink(response http.ResponseWriter, request *http.Request)  {
	defer request.Body.Close()
	var req shortReq
	if error:=json.NewDecoder(request.Body).Decode(&req);error!=nil{
		returnResponseError(response,500,nil,error.Error())
		return
	}
	fmt.Println("requestjson:",request.Body)
	fmt.Println("createShortLink",ns," ",req)
	ns.add(strconv.Itoa(len(ns.data)),req.URL)
	returnResponseJson(response,&req.URL)
}

func (app *App)getShortLinkInfo(response http.ResponseWriter,request *http.Request)  {
	values := request.URL.Query()
	key := values.Get("key")
	fmt.Println("getShortLinkInfo:",key," ",ns)
	target:=ns.get(key)
	if target == "" {
		returnResponseError(response, 500, nil, "no key")
		return
	}
	temp:=make([]listTemp,0)
	temp=append(temp,listTemp{
		Name:"a",
	})
	temp=append(temp,listTemp{
		Name:"b",
	})
	shopinfo:=SLInfo{
		Key:    key,
		Target: target,
		List: temp,
	}
	fmt.Println(shopinfo)
	returnResponseJson(response,shopinfo)
}

func (app *App)redirect(response http.ResponseWriter,request *http.Request)  {
	fmt.Println("redirect",ns)
	vars := mux.Vars(request)
	fmt.Println(vars)
	key :=vars["sk"]
	target:=ns.get(key)
	http.Redirect(response,request,target,http.StatusTemporaryRedirect)
	//http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}
func returnResponseError(response http.ResponseWriter,status int ,content interface{},message string)  {
	if message == "" {
		message = http.StatusText(status)
	}
	res, _ := json.Marshal(BaseResponse{
		Code:    status,
		Message: message,
		Content: content,
	})
	response.Header().Set("Content-Type","application/json;charset=utf-8")
	response.Header().Set("AAA","12345")
	response.WriteHeader(status)
	_, _ = response.Write(res)
}

func returnResponseJson(response http.ResponseWriter,content interface{})  {
	returnResponseError(response,200,content,"处理成功")
}
