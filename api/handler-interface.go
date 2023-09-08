package api

import "net/http"

//Handler ...
type RestfulHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}

type AuthorHandler interface {
	RestfulHandler
}
