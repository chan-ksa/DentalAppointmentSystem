package handlers

import (
	"DentalAppointmentSystem/pkg"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(res http.ResponseWriter, req *http.Request) {
	// Check if already logged in
	if pkg.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	// Process to form submission
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")

		// Check if the user exist
		if !pkg.CheckExistingUser(username) {
			http.Error(res, "Username and/or password do not match", http.StatusUnauthorized)
			return
		}

		// Check if there's a matching of password entered
		userPassword, admin := pkg.GetUserPassword(username)
		err := bcrypt.CompareHashAndPassword(userPassword, []byte(password))
		if err != nil {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// Successfully log in
		log.Println("Creating session now...")

		// Create the session cookie
		id := uuid.NewV4()
		userCookie := &http.Cookie{
			Name:  "cookieID",
			Value: id.String(),
		}
		// Setting the session cookie
		http.SetCookie(res, userCookie)
		log.Println("Setting session now...")

		// Adding the cookie into the json
		sessInfo := pkg.SessionInfo{SessionID: userCookie.Value, Username: username}
		pkg.AddToSessionJson(sessInfo)

		if admin {
			http.Redirect(res, req, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(res, req, "/customer", http.StatusSeeOther)
		}
		return
	}
}

func Logout(res http.ResponseWriter, req *http.Request) {
	if !pkg.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}

	userCookie, _ := req.Cookie("cookieID")

	// Delete the session from the json
	pkg.DeleteSessionsfromJson(userCookie.Value)

	// Remove the cookie in the server
	userCookie = &http.Cookie{
		Name:   "cookieID",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, userCookie)

	// Redirect to the login page
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
