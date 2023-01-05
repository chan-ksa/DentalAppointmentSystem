package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type SessionInfo struct {
	SessionID string `json:"Session ID"`
	Username  string `json:"Username"`
}

var RecordedSession []SessionInfo

// Function to check for session id
func checkSessionId(id string) (string, bool) {
	// Loop through all the session
	for _, data := range RecordedSession {
		if data.SessionID == id {
			return data.Username, true
		}
	}

	log.Println("Session Id does not have a username")
	return "", false
}

// Function to check if already logged in
func AlreadyLoggedIn(req *http.Request) bool {
	myCookie, err := req.Cookie("cookieID")
	if err != nil {
		log.Println(err)
		return false
	}

	username, exist := checkSessionId(myCookie.Value)
	if exist {
		ok := CheckExistingUser(username)
		return ok
	} else {
		return false
	}
}

func GetCurrentSessionID(res http.ResponseWriter, req *http.Request) string {
	myCookie, err := req.Cookie("cookieID")

	if err != nil {
		log.Println("User Cookie not found!")
		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

	return myCookie.Value
}

// Function to add into the session json
func AddToSessionJson(si SessionInfo) {
	// When the RecordedSession is empty
	length := len(RecordedSession)

	if length == 0 {
		RecordedSession = append(RecordedSession, si)
	}

	// Check if the username already exist in the session id
	for i, sess := range RecordedSession {
		if sess.Username == si.Username {
			RecordedSession[i].SessionID = si.SessionID
			break
		} else if i == (length - 1) {
			RecordedSession = append(RecordedSession, si)
		}
	}

	result, err := json.MarshalIndent(RecordedSession, "", "")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("cmd/json/sessions.json", result, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Successfully added to json file
	log.Println("Successfully added into sessions.json")
}

func RetrieveSessionsFromJson(sessionlist *[]SessionInfo, path string, filename string) {
	data, err := OpenFile(path, filename)
	if err != nil {
		log.Fatal(err)
	}

	// Store into the list using the data from the json file
	json.Unmarshal(data, &sessionlist)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteSessionsfromJson(sessId string) {
	// Open file to pull the most updated from sessions.json
	data, err := OpenFile("cmd/json/", "sessions.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &RecordedSession)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the find the session that need to be deleted
	for i, sess := range RecordedSession {
		if sess.SessionID == sessId {
			RecordedSession[i] = RecordedSession[len(RecordedSession)-1]
			RecordedSession = RecordedSession[:len(RecordedSession)-1]
		}
	}

	// Overwrite the json file
	result, err := json.MarshalIndent(&RecordedSession, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("cmd/json/sessions.json", result, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Successfully updated
	log.Println("Sucessfully updated sessions.json")
}

func UpdateSessionUsername(cookieValue string, newUsername string) []SessionInfo {
	for i, sess := range RecordedSession {
		if sess.SessionID == cookieValue {
			RecordedSession[i].Username = newUsername
		}
	}

	return RecordedSession
}

func UpdateSessionsFromJson(sessionlist []SessionInfo) {
	result, err := json.MarshalIndent(sessionlist, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	// Write to file with the new list
	err = os.WriteFile("cmd/json/sessions.json", result, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully updated sessions.json")
}
