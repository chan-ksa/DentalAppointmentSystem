package pkg

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Appointment struct {
	Name      string `json:"Name"`
	Date      string `json:"Date"`
	StartTime string `json:"Start Time"`
	EndTime   string `json:"End Time"`
	AppId     int    `json:"AppId"`
}

type BookAppointment struct {
	NewId   int
	AllApps []Appointment
}

type EditAppointment struct {
	App     Appointment
	AllApps []Appointment
}

var (
	AppointmentsList []Appointment
	mutex            sync.Mutex
)

func AddToAppointmentJson(app []Appointment) {
	result, err := json.MarshalIndent(AppointmentsList, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	// Write into json file
	err = os.WriteFile("cmd/json/appointments.json", result, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully added into appointments.json")
}

func RetrieveAppointmentFromJson(applist *[]Appointment, path string, filename string) {
	// Open the file with the given path and file name
	data, err := OpenFile(path, filename)
	if err != nil {
		log.Fatal(err)
	}

	// Store into the list using the data from the json file
	json.Unmarshal(data, &AppointmentsList)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAppointmentById(id int) (app Appointment) {
	// Loop through the list and retrieve the respective appointment
	for i, app := range AppointmentsList {
		if app.AppId == id {
			return AppointmentsList[i]
		}
	}

	// Cannot find the appointment with the given id
	log.Printf("Appointment with id %d does not exist!", id)
	return
}

func DeleteAppointmentById(id int) []Appointment {
	// Remove the name from the appointment list
	for i, app := range AppointmentsList {
		if app.AppId == id {
			AppointmentsList[i].Name = ""
		}
	}

	return AppointmentsList
}

func UpdateAppointmentById(id int, username string) []Appointment {
	// Set the user name into the appointment list
	for i, app := range AppointmentsList {
		if app.AppId == id {
			AppointmentsList[i].Name = username
		}
	}

	return AppointmentsList
}

func EditAppointmentById(oId int, nId int, name string) []Appointment {
	for i, app := range AppointmentsList {
		if app.AppId == oId {
			AppointmentsList[i].Name = ""
		}
		if app.AppId == nId {
			AppointmentsList[i].Name = name
		}
	}

	return AppointmentsList
}
