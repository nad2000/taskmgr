package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"taskmgr/common"
	"taskmgr/models"

	"github.com/gorilla/mux"
)

var Tasks = new(taskController)

type taskController struct{}

func (tc *taskController) Create(w http.ResponseWriter, r *http.Request) {

	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	task, err := models.Tasks.Create(t.Name, t.Desc)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err == nil {
		common.JsonOk(w, res, http.StatusCreated)
		return
	}
}

func (tc *taskController) Get(w http.ResponseWriter, r *http.Request) {

	tasks, err := models.Tasks.FindAll()
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(tasks)
	if err == nil {
		common.JsonOk(w, res, http.StatusOK)
		return
	} else {
		common.JsonError(w, err, http.StatusBadRequest)
	}

}

func (tc *taskController) Show(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	log.Printf("INFO: ID: %s\n", id)

	task, err := models.Tasks.FindOne(id)
	if err != nil {
		log.Printf("Failed to retrieve ID: %s\n", id)
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err == nil {
		common.JsonOk(w, res, http.StatusOK)
		return
	} else {
		common.JsonError(w, err, http.StatusBadRequest)
	}

}

func (tc *taskController) Update(w http.ResponseWriter, r *http.Request) {

	var t models.Task
	id := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	if err := models.Tasks.Update(id, t.Name, t.Desc); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonStatus(w, http.StatusNoContent)
}

func (tc *taskController) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := models.Tasks.Delete(id); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	common.JsonStatus(w, http.StatusNoContent)

}
