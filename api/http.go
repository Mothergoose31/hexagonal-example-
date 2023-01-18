package api

import (
	"encoding/json"
	"net/http"
	"tea/domain"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type handler struct {
	storeService domain.Service
}

//NewStoreHandler ...
func NewStoreHandler(storeService domain.Service) StoreHandler {
	return &handler{storeService: storeService}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	store, err := h.storeService.Find(params["id"])
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store)
}

//================================================================================

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := &domain.Store{}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.storeService.Store(s)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(s)
}

//================================================================================

func (h *handler) Put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := &domain.Store{}
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.storeService.Store(s)
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&s)
}

//================================================================================

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := h.storeService.Delete(params["id"])
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stores, err := h.storeService.FindAll()
	if err != nil {
		logrus.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(stores)
}
