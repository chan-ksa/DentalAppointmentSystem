package handlers

import (
	"DentalAppointmentSystem/pkg"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Main Page
func Customer(res http.ResponseWriter, req *http.Request) {
	user := pkg.GetUser(res, req)
	data := pkg.UserAppointment{User: user, App: pkg.AppointmentsList}
	tpl.ExecuteTemplate(res, "customer.gohtml", data)
}

// Profile
func Profile(res http.ResponseWriter, req *http.Request) {
	user := pkg.GetUser(res, req)
	tpl.ExecuteTemplate(res, "profile.gohtml", user)
}

// Profile
func UpdateProfile(res http.ResponseWriter, req *http.Request) {
	user := pkg.GetUser(res, req)
	tpl.ExecuteTemplate(res, "updateprofile.gohtml", user)
}

// Profile
func EditUsername(res http.ResponseWriter, req *http.Request) {
	// Process form (post)
	if req.Method == http.MethodPost {
		user := pkg.GetUser(res, req)
		nUsername := req.FormValue("newUsername")
		cookieValue := pkg.GetCurrentSessionID(res, req)

		// Check whether the new username is not empty
		if strings.TrimSpace(nUsername) != "" {
			// check if username exist/ taken
			if pkg.CheckExistingUser(nUsername) {
				http.Error(res, "Username already taken", http.StatusForbidden)
				// to do redirect
				return
			}

			// Update the UsersList to the new username
			updatedUser := pkg.UpdateUserUsername(user, nUsername)

			// Update users.json file
			pkg.UpdateUserFromJson(updatedUser)

			// Update sessions.json file
			updatedSession := pkg.UpdateSessionUsername(cookieValue, nUsername)
			pkg.UpdateSessionsFromJson(updatedSession)

			// redirect to update profile page
			http.Redirect(res, req, "/updateprofile", http.StatusSeeOther)

		} else {
			http.Error(res, "Username cannot be empty", http.StatusForbidden)
			return
		}
	}

	tpl.ExecuteTemplate(res, "editusername.gohtml", nil)
}

// Profile
func EditPassword(res http.ResponseWriter, req *http.Request) {
	// Process form (post)
	if req.Method == http.MethodPost {
		user := pkg.GetUser(res, req)
		cPassword := req.FormValue("currentPassword")
		nPassword := req.FormValue("newPassword")

		// Check whether the new password is not empty
		if strings.TrimSpace(nPassword) != "" {
			// Check if the password key in is correct
			err := bcrypt.CompareHashAndPassword(user.Password, []byte(cPassword))
			if err != nil {
				http.Error(res, "Username and/or password do not match", http.StatusForbidden)
				return
			}

			// BCrypt the new password given
			bNewPassword, err := bcrypt.GenerateFromPassword([]byte(nPassword), bcrypt.MinCost)
			if err != nil {
				http.Error(res, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Update the UsersList to the new password
			updatedUser := pkg.UpdateUserPassword(user, bNewPassword)

			// Update users.json file
			pkg.UpdateUserFromJson(updatedUser)

			// redirect to update profile page
			http.Redirect(res, req, "/updateprofile", http.StatusSeeOther)
		} else {
			http.Error(res, "New Password cannot be empty", http.StatusForbidden)
			return
		}
	}

	tpl.ExecuteTemplate(res, "editpassword.gohtml", nil)
}

// Appointment
func BookAppointment(res http.ResponseWriter, req *http.Request) {
	// Process form (post)
	if req.Method == http.MethodPost {
		user := pkg.GetUser(res, req)
		newId := req.FormValue("newAppointmentId")

		// Convert the id into int
		newAppId, err := strconv.Atoi(newId)
		if err != nil {
			log.Fatal(err)
		}

		// Set the user name into the appointment list
		updatedList := pkg.UpdateAppointmentById(newAppId, user.Username)

		// Add into appointment.json
		pkg.AddToAppointmentJson(updatedList)

		// redirect to book appointment feedback page
		http.Redirect(res, req, "/bookappointmentdone", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(res, "bookappointment.gohtml", pkg.AppointmentsList)
}

// Appointment
func BookAppointmentDone(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "bookappointmentdone.gohtml", nil)
}

// Appointment
func DeleteAppointment(res http.ResponseWriter, req *http.Request) {
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

	tpl.ExecuteTemplate(res, "deleteappointment.gohtml", pkg.AppointmentsList)
}

// Appointment
func EditAppointment(res http.ResponseWriter, req *http.Request) {
	var data pkg.EditAppointment

	// Process form submission
	if req.Method == http.MethodPost {
		AId := req.FormValue("AId")
		id, err := strconv.Atoi(AId)
		if err != nil {
			log.Fatal(err)
		}

		appointment := pkg.GetAppointmentById(id)
		data = pkg.EditAppointment{App: appointment, AllApps: pkg.AppointmentsList}
	}

	tpl.ExecuteTemplate(res, "editappointment.gohtml", data)
}

// Appointment
func EditAppointmentDone(res http.ResponseWriter, req *http.Request) {
	// Process form submission
	if req.Method == http.MethodPost {
		user := pkg.GetUser(res, req)
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
		updatedList := pkg.EditAppointmentById(iOldId, iNewId, user.Username)

		// Add into appointments.json
		pkg.AddToAppointmentJson(updatedList)
	}

	tpl.ExecuteTemplate(res, "editappointmentdone.gohtml", pkg.AppointmentsList)
}
