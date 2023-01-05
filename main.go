package main

import (
	handlers "DentalAppointmentSystem/cmd/web"
	pkg "DentalAppointmentSystem/pkg"
	"encoding/xml"
	"net/http"
	"os"
)

type Settings struct {
	NIP   string `xml:"IP"`
	NPort string `xml:"Port"`
}

var settings Settings

// Init function
func init() {
	// To check for the required files
	jsonFiles := []string{"users.json", "sessions.json"}
	requiredFiles := []string{"admin.json", "appointments.json"}
	pkg.CheckIfFilesExistAndCreateFile("cmd/json/", jsonFiles)
	pkg.CheckRequiredFiles("cmd/json/", requiredFiles)

	// To initalise the settings (ip & port) from xml
	settingData, _ := os.ReadFile("cmd/xml/settings.xml")
	_ = xml.Unmarshal([]byte(settingData), &settings)

	// Init the template for the html
	handlers.TemplateInit()

	// Retrieve all the data from json and store into data structure
	pkg.RetrieveUsersFromJson(&pkg.AdminsList, "cmd/json/", "admin.json")
	pkg.RetrieveUsersFromJson(&pkg.UsersList, "cmd/json/", "users.json")
	pkg.RetrieveSessionsFromJson(&pkg.RecordedSession, "cmd/json/", "sessions.json")
	pkg.RetrieveAppointmentFromJson(&pkg.AppointmentsList, "cmd/json/", "appointments.json")
}

func main() {
	http.Handle("/", http.HandlerFunc(handlers.Index))
	http.Handle("/signup", http.HandlerFunc(handlers.Signup))
	http.Handle("/login", http.HandlerFunc(handlers.Login))
	http.Handle("/logout", http.HandlerFunc(handlers.Logout))

	http.Handle("/admin", http.HandlerFunc(handlers.Admin))
	http.Handle("/viewusers", http.HandlerFunc(handlers.ViewUsers))
	http.Handle("/deleteuser", http.HandlerFunc(handlers.DeleteUser))
	http.Handle("/viewsessions", http.HandlerFunc(handlers.ViewSessions))
	http.Handle("/deletesession", http.HandlerFunc(handlers.DeleteSession))
	http.Handle("/viewappointments", http.HandlerFunc(handlers.ViewAppointments))
	http.Handle("/edituserappointment", http.HandlerFunc(handlers.EditUserAppointment))
	http.Handle("/edituserappointmentdone", http.HandlerFunc(handlers.EditUserAppointmentDone))
	http.Handle("/deleteuserappointment", http.HandlerFunc(handlers.DeleteUserAppointment))

	http.Handle("/customer", http.HandlerFunc(handlers.Customer))

	// Profile
	http.Handle("/profile", http.HandlerFunc(handlers.Profile))
	http.Handle("/editusername", http.HandlerFunc(handlers.EditUsername))
	http.Handle("/editpassword", http.HandlerFunc(handlers.EditPassword))
	http.Handle("/updateprofile", http.HandlerFunc(handlers.UpdateProfile))

	// Appointment
	http.Handle("/bookappointment", http.HandlerFunc(handlers.BookAppointment))
	http.Handle("/bookappointmentdone", http.HandlerFunc(handlers.BookAppointmentDone))
	http.Handle("/deleteappointment", http.HandlerFunc(handlers.DeleteAppointment))
	http.Handle("/editappointment", http.HandlerFunc(handlers.EditAppointment))
	http.Handle("/editappointmentdone", http.HandlerFunc(handlers.EditAppointmentDone))

	http.ListenAndServe(settings.NIP+":"+settings.NPort, nil)
}
