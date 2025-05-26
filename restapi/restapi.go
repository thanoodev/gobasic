package restapi

import (
	"encoding/json"
	"net/http"
)

var profileList = []Profile{
	{Id: 1, Name: "John Doe"},
	{Id: 2, Name: "Jane Smith"},
}

func Init() {
	getProfile()
	addProfile()
	http.ListenAndServe(":8080", nil)
}

func getProfile() {
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := json.Marshal(profileList)
		if err != nil {
			http.Error(w, "Failed to marshal profile", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	})
}

func addProfile() {
	http.HandleFunc("/addprofile", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var profile Profile
		if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		// Add profile to the list
		profileList = append(profileList, profile)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(PostProfileResponse{
			Status:  "Profile created successfully",
			Profile: profile,
		})
	})
}

type PostProfileResponse struct {
	Status  string
	Profile Profile
}

type Profile struct {
	Id   int
	Name string
}
