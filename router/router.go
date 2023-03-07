package router

import (
	"net/http"

	"github.com/MrAndreiL/go-rest-api/controllers"
	"github.com/MrAndreiL/go-rest-api/utils"
)

func handleStudent(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controllers.GetStudentEntityRequest(w, r)
	case "POST":
		utils.SendBadRequestGeneric(w, "Cannot create resource at given location")
	case "PUT":
		controllers.PutStudentEntityRequest(w, r)
	case "DELETE":
		controllers.DeleteStudentEntity(w, r)
	default:
		utils.SendBadRequestGeneric(w, "Invalid request.")
	}
}

func handleStudents(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		controllers.GetStudentCollection(w, r)
	case "POST":
		controllers.PostStudentCollection(w, r)
	case "PUT":
		controllers.PutStudentCollection(w)
	case "DELETE":
		controllers.DeleteStudentCollection(w)
	default:
		utils.SendBadRequestGeneric(w, "Invalid request.")
	}
}

func handleDoctor(w http.ResponseWriter, r *http.Request) {

}

func handleDoctors(w http.ResponseWriter, r *http.Request) {

}

func HandleRequests() {
	http.HandleFunc("/api/students", handleStudents)

	http.HandleFunc("/api/students/", handleStudent)

	http.HandleFunc("/api/doctors", handleDoctors)

	http.HandleFunc("/api/doctors/", handleDoctor)
}
