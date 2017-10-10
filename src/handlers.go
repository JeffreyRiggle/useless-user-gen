package main

import (
	"encoding/json"
	"net/http"
)

var manager UserManager

type CreateRequest struct {
	NumberToCreate int
}

type UpdateRequest struct {
	NumberToUpdate int
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var req CreateRequest

	err := decoder.Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users := manager.CreateUsers(req.NumberToCreate)

	data, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var req UpdateRequest

	err := decoder.Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users := manager.UpdateUsers(req.NumberToUpdate)

	data, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(manager.GetUsers())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}
