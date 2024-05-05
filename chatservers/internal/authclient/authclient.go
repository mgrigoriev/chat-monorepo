package authclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const authURL = "http://users:8080/api/v1/users/auth"

type currentUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Authenticate(token string) (userID int, err error) {
	data := map[string]string{"token": token}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	httpClient := http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Post(authURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("authentication failed: %d", resp.StatusCode)
	}

	var user currentUser

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}
