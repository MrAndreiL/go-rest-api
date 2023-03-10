package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/MrAndreiL/go-rest-api/models"
	"github.com/MrAndreiL/go-rest-api/utils"
)

func GetStudentEntityRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/students/"))
	if err != nil { // invalid URI
		w.WriteHeader(http.StatusBadRequest)
		w.Write(models.JsonErrorResponseMessage("Invalid resource identifier."))
		return
	}

	response, code := models.GetStudent(id)
	if code == http.StatusOK { // provide caching support
		key := fmt.Sprint(id) + "s"

		if utils.CacheSupport(key, w, r) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	w.WriteHeader(code)
	w.Write(response)
}

func PutStudentEntityRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/students/"))
	if err != nil { // invalid URI
		w.WriteHeader(http.StatusBadRequest)
		w.Write(models.JsonErrorResponseMessage("Invalid resource identifier."))
		return
	}

	response, code := models.PutStudent(id, r.Body)
	w.WriteHeader(code)
	if code != http.StatusNoContent {
		w.Write(response)
	}
}

func DeleteStudentEntity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/students/"))
	if err != nil { // invalid URI
		w.WriteHeader(http.StatusBadRequest)
		w.Write(models.JsonErrorResponseMessage("Invalid resource identifier."))
		return
	}

	response, code := models.DeleteStudent(id)
	w.WriteHeader(code)
	w.Write(response)
}

func PutStudentCollection(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusMethodNotAllowed)

	w.Write(models.JsonErrorResponseMessage("Cannot replace entire collection."))
}

func DeleteStudentCollection(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusMethodNotAllowed)

	w.Write(models.JsonErrorResponseMessage("Cannot delete entire collection."))
}

func PostStudentCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, code, id := models.PostStudentCollection(r.Body)
	if code == http.StatusCreated {
		w.Header().Set("Location", r.URL.String()+"/"+strconv.Itoa(id))
	}
	w.WriteHeader(code)
	w.Write(response)
}

func GetStudentCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// no limit or offset, get the entire collection.
	if r.URL.String() == "/api/students" {
		response, code := models.GetRequestCollection()
		w.WriteHeader(code)
		w.Write(response)
	}
}
