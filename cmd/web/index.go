package handlers

import (
	"DentalAppointmentSystem/pkg"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {
	if pkg.AlreadyLoggedIn(req) {
		user := pkg.GetUser(res, req)
		if user.IsAdmin {
			http.Redirect(res, req, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(res, req, "/customer", http.StatusSeeOther)
		}
	}
	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
