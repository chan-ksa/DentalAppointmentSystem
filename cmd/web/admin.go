package handlers

import (
	"DentalAppointmentSystem/pkg"
	"log"
	"net/http"
	"strconv"
)

// Main Page
func Admin(res http.ResponseWriter, req *http.Request) {
	user := pkg.GetUser(res, req)
	tpl.ExecuteTemplate(res, "admin.gohtml", user)
}

// User
func ViewUsers(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "viewusers.gohtml", pkg.UsersList)
}

// User
func DeleteUser(res http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	fname := req.FormValue("fName")
	lname := req.FormValue("lName")
	pkg.DeleteUserFromJson(username)
	data := pkg.User{Username: username, Userinfo: pkg.UserInfo{Firstname: fname, Lastname: lname}}
	tpl.ExecuteTemplate(res, "deleteuser.gohtml", data)
}

// Session
func ViewSessions(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "viewsessions.gohtml", pkg.RecordedSession)
}

// Session
func DeleteSession(res http.ResponseWriter, req *http.Request) {
	sessionID := req.FormValue("sessionID")
	username := req.FormValue("username")
	pkg.DeleteSessionsfromJson(sessionID)
	data := pkg.SessionInfo{SessionID: sessionID, Username: username}
	tpl.ExecuteTemplate(res, "deletesession.gohtml", data)
}

// Appointment
func ViewAppointments(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "viewappointments.gohtml", pkg.AppointmentsList)
}

// Appointment
func EditUserAppointment(res http.ResponseWriter, req *http.Request) {
	var data pkg.EditAppointment

	// Process form submission
	if req.Method == http.MethodPost {
		AId := req.FormValue("AId")
		id, err := strconv.Atoi(AId)
		if err != nil {
			log.Fatal(err)
		}

		// Retrieve the user appointment
		appointment := pkg.GetAppointmentById(id)
		data = pkg.EditAppointment{App: appointment, AllApps: pkg.AppointmentsList}
	}

	tpl.ExecuteTemplate(res, "edituserappointment.gohtml", data)
}

// Appointment
func EditUserAppointmentDone(res http.ResponseWriter, req *http.Request) {
	// Process form submission
	if req.Method == http.MethodPost {
		name := req.FormValue("user")
		newId := req.FormValue("newAppointment")
		oldId := req.FormValue("oldAppointment")

		// Convert the new id to int from string
		iNewId, err := strconv.Atoi(newId)
		if err != nil {
			log.Fatal(err)
		}

		// Convert the old id to int from string
		iOldId, err := strconv.Atoi(oldId)
		if err != nil {
			log.Fatal(err)
		}

		// Update the appointment list
		updatedList := pkg.EditAppointmentById(iOldId, iNewId, name)

		// Add into appointments.json
		pkg.AddToAppointmentJson(updatedList)
	}

	tpl.ExecuteTemplate(res, "edituserappointmentdone.gohtml", pkg.AppointmentsList)
}

// Appointment
func DeleteUserAppointment(res http.ResponseWriter, req *http.Request) {
	// Process form submission
	if req.Method == http.MethodPost {
		AId := req.FormValue("AId")
		iAId, err := strconv.Atoi(AId)
		if err != nil {
			log.Fatal(err)
		}

		// Update the appointment list
		updatedList := pkg.DeleteAppointmentById(iAId)

		// Write to appointments.json file
		pkg.AddToAppointmentJson(updatedList)
	}

	tpl.ExecuteTemplate(res, "deleteuserappointment.gohtml", pkg.AppointmentsList)
}
