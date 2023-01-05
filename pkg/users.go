package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type User struct {
	Username string `json:"Username"`
	Password []byte `json:"Password"`
	IsAdmin  bool   `json:"Is Admin`
	Userinfo UserInfo
}

type UserInfo struct {
	Firstname string `json:"First Name"`
	Lastname  string `json:"Last Name"`
}

type UserAppointment struct {
	User User
	App  []Appointment
}

var UsersList []User
var AdminsList []User

func CheckExistingUser(username string) bool {
	// Loop through all users and check if user exist
	for _, userData := range UsersList {
		if userData.Username == username {
			return true
		}
	}

	// Loop through all admin and check if admin exist
	for _, userData := range AdminsList {
		if userData.Username == username {
			return true
		}
	}

	// User does not exist
	log.Printf("%s does not exist!", username)
	return false
}

// Function to get the user
func GetUser(res http.ResponseWriter, req *http.Request) User {
	myCookie, err := req.Cookie("cookieID")
	if err != nil {
		log.Println("User cookie not found!")
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	var myUser User
	if username, ok := checkSessionId(myCookie.Value); ok {
		myUser = GetUserInfo(username)
	} else {
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	return myUser
}

func GetUserInfo(username string) (userinfo User) {
	// Loop through all users and check if user exist
	for i, userData := range UsersList {
		if userData.Username == username {
			return UsersList[i]
		}
	}

	// Loop through all admins and check if admin exist
	for i, userData := range AdminsList {
		if userData.Username == username {
			return AdminsList[i]
		}
	}

	// User does not exist
	log.Printf("%s does not exist!", username)
	return
}

func AddToUsersJson(user User) {
	// Add into the list
	UsersList = append(UsersList, user)

	result, err := json.MarshalIndent(UsersList, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	// Write into json file
	err = os.WriteFile("cmd/json/users.json", result, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully added into users.json")
}

func AddToAdminsJson(user User) {
	// Add into the list
	AdminsList = append(AdminsList, user)

	result, err := json.MarshalIndent(AdminsList, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	// Write into json file
	err = os.WriteFile("cmd/json/admin.json", result, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully added into admin.json")
}

func UpdateUserFromJson(userlist []User) {
	result, err := json.MarshalIndent(userlist, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	// Write to file with the new list
	err = os.WriteFile("cmd/json/users.json", result, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully updated users.json")
}

func UpdateUserUsername(user User, newUsername string) []User {
	for i, u := range UsersList {
		if u.Username == user.Username {
			UsersList[i].Username = newUsername
		}
	}

	return UsersList
}

func UpdateUserPassword(user User, newPassword []byte) []User {
	for i, u := range UsersList {
		if u.Username == user.Username {
			UsersList[i].Password = newPassword
		}
	}

	return UsersList
}

func GetUserPassword(username string) ([]byte, bool) {
	// Loop through all the user to get the password of the respective user
	for _, userData := range UsersList {
		if userData.Username == username {
			return userData.Password, userData.IsAdmin
		}
	}

	// Loop through all the admin to get the password of the respective user
	for _, userData := range AdminsList {
		if userData.Username == username {
			return userData.Password, userData.IsAdmin
		}
	}

	return nil, false
}

func RetrieveUsersFromJson(userlist *[]User, path string, filename string) {
	data, err := OpenFile(path, filename)
	if err != nil {
		log.Fatal(err)
	}

	// Store into the list using the data from the json file
	json.Unmarshal(data, &userlist)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUserFromJson(username string) {
	// Open file to pull the most updated from users.json
	data, err := OpenFile("cmd/json/", "users.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &UsersList)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the find the user that need to be deleted
	for i, user := range UsersList {
		if user.Username == username {
			UsersList[i] = UsersList[len(UsersList)-1]
			UsersList = UsersList[:len(UsersList)-1]
		}
	}

	// Overwrite the json file
	result, err := json.MarshalIndent(UsersList, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("cmd/json/users.json", result, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Successfully updated users.json
	log.Println("Successfully updated users.json")

	// Loop through the find the session that need to be deleted
	for i, ses := range RecordedSession {
		if ses.Username == username {
			RecordedSession[i] = RecordedSession[len(RecordedSession)-1]
			RecordedSession = RecordedSession[:len(RecordedSession)-1]
		}
	}

	// Overwrite the json file
	result, err = json.MarshalIndent(&RecordedSession, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("cmd/json/sessions.json", result, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Successfully updated sessions.json
	log.Println("Sucessfully updated sessions.json")

	// Loop through the find the appointment that need to be deleted
	for i, app := range AppointmentsList {
		if app.Name == username {
			AppointmentsList[i] = AppointmentsList[len(AppointmentsList)-1]
			AppointmentsList = AppointmentsList[:len(AppointmentsList)-1]
		}
	}

	// Overwrite the json file
	result, err = json.MarshalIndent(&AppointmentsList, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("cmd/json/appointments.json", result, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Successfully updated appointments.json
	log.Println("Sucessfully updated appointments.json")
}
