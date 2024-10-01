package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type OAuthTokenManager struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

func NewOAuthTokenManager() OAuthTokenManager {
	tm := OAuthTokenManager{}
	tm.LoadSavedToken()
	tm.RequestOAuthToken()
	tm.SaveToken()
	return tm
}

func (tm *OAuthTokenManager) LoadSavedToken() {
	file, err := os.Open(filepath.Join("configs", "auth.json"))
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(tm); err != nil {
		log.Fatalf("error unmarshalling JSON: %s", err)
	}

}

func (tm OAuthTokenManager) SaveToken() {
	file, err := os.OpenFile(filepath.Join("configs", "auth.json"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("failed to open file for writing: %s", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err = encoder.Encode(&tm); err != nil {
		log.Fatalf("error encoding JSON: %s", err)
	}
}

func (tm *OAuthTokenManager) RequestOAuthToken() error {
	data := fmt.Sprintf("grant_type=refresh_token&refresh_token=%s&client_id=%s&client_secret=%s", tm.RefreshToken, tm.ClientID, tm.ClientSecret)

	req, err := http.NewRequest("POST", "https://id.twitch.tv/oauth2/token", bytes.NewBufferString(data))
	if err != nil {
		return fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %s", body)
	}

	err = json.Unmarshal(body, tm)
	if err != nil {
		return fmt.Errorf("could not unmarshal response: %v", err)
	}

	return nil
}
