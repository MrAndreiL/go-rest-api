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

		if utils.CacheSupport(string(key), w, r) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	w.WriteHeader(code)
	w.Write(response)
}
