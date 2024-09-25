package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Deepanshuisjod/rest-api/db"
	"github.com/Deepanshuisjod/rest-api/models"
)

var InformationList = []*models.Information{}

func EmployeeInformation(rw http.ResponseWriter, r *http.Request) {
	// Get Employee information
	InformationList = db.GetEmpInformation()
	log.Println("Employee information fetched:", InformationList)

	// Convert the InformationList to JSON and write it to the response
	err := json.NewEncoder(rw).Encode(InformationList)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}
}

func CreateProject(rw http.ResponseWriter, r *http.Request) {
	var emp = models.Assignees{}

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(rw, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	err = db.SetProjectInformation(&emp)
	if err != nil {
		http.Error(rw, "[Error] Running query in database", http.StatusBadRequest)
		return
	}

	rw.Write([]byte("Done adding the value to the database"))

}
