package router

import (
	"net/http"

	"github.com/MrAndreiL/go-rest-api/controllers"
)

func handleStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		controllers.GetStudentEntityRequest(w, r)
	}
}

func handleStudents(w http.ResponseWriter, r *http.Request) {
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
