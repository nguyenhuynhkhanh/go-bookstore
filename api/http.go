package api

import (
	"encoding/json"
	"log"
	"net/http"

	"bookstore.com/domain"
	entities "bookstore.com/domain/entity"
	"github.com/go-chi/chi"
)

type handler struct {
	authorService domain.AuthorService
}

//NewHandler ...
func NewHandler(authorService domain.AuthorService) AuthorHandler {
	return &handler{authorService: authorService}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := chi.URLParam(r, "code")
	p, err := h.authorService.Find(code)

	if err != nil {

		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	json.NewEncoder(w).Encode(&p)

}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	//requestBody, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")

	p := &entities.Author{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = h.authorService.Store(p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&p)

}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := chi.URLParam(r, "code")
	err := h.authorService.Delete(code)
	if err != nil {
		log.Fatal(err)
	}

}
func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p, err := h.authorService.FindAll()

	if err != nil {

		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&p)

}
