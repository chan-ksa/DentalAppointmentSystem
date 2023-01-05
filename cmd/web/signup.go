package handlers

import (
	"DentalAppointmentSystem/pkg"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Signup(res http.ResponseWriter, req *http.Request) {
	if pkg.AlreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		admin := req.FormValue("admin")
		if strings.TrimSpace(username) != "" && strings.TrimSpace(password) != "" {
			// check if username exist/ taken
			if pkg.CheckExistingUser(username) {
				http.Error(res, "Username already taken", http.StatusForbidden)
				// to do redirect
				return
			}

			bPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
			if err != nil {
				http.Error(res, "Internal server error", http.StatusInternalServerError)
				return
			}

			// Check if user / admin and add into json file
			if admin == "" {
				user := pkg.User{Username: username, Password: bPassword, IsAdmin: false, Userinfo: pkg.UserInfo{Firstname: firstname, Lastname: lastname}}
				pkg.AddToUsersJson(user)
			} else {
				admin := pkg.User{Username: username, Password: bPassword, IsAdmin: true, Userinfo: pkg.UserInfo{Firstname: firstname, Lastname: lastname}}
				pkg.AddToAdminsJson(admin)
			}
		} else {
			http.Error(res, "Username and Password cannot be empty", http.StatusForbidden)
			return
		}
		// redirect to main index
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return

	}
	tpl.ExecuteTemplate(res, "signup.gohtml", nil)
}
