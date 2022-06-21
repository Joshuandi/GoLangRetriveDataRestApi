package http_request

import (
	//"GoLangRetriveDataRestApi/model"
	"GoLangRetriveDataRestApi/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpRequestInterface interface {
	HttpRequest(w http.ResponseWriter, r *http.Request)
}

type HttpRequest struct{}

func NewHttpRequest() HttpRequestInterface {
	return &HttpRequest{}
}

func (h *HttpRequest) HttpRequest(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	switch r.Method {
	case http.MethodGet:
		h.getHttpRequest(w, r)
	case http.MethodPost:
		h.postHttpRequest(w, r, id)
	}

}

func (h *HttpRequest) getHttpRequest(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://random-data-api.com/api/users/random_user?size=10.")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var person []model.Person
	if err := json.NewDecoder(res.Body).Decode(&person); err != nil {
		log.Fatal(err)
	}

	jsonData, _ := json.Marshal(&person)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

}

func (h *HttpRequest) postHttpRequest(w http.ResponseWriter, r *http.Request, id string) {

}
