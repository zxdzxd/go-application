package spotifyclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/mux"
)

var (
	ClientId      string
	ClientSecret  string
	token_api_url string = "https://accounts.spotify.com/api/token"
	grantType     string = "client_credentials"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func SfClient() {
	router := mux.NewRouter()

	router.HandleFunc("/getArtist", getArtistById)

	http.ListenAndServe(":8001", router)
}

func getArtistById(w http.ResponseWriter, r *http.Request) {
	access_token := getAccessToken()
	fmt.Printf("Received access token - %s", access_token)

}

func getAccessToken() string {

	ClientId = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")

	formdata := url.Values{}
	formdata.Set("grant_type", grantType)
	formdata.Set("client_id", ClientId)
	formdata.Set("client_secret", ClientSecret)

	encoded_formdata := formdata.Encode()

	req, err := http.NewRequest(http.MethodPost, token_api_url, bytes.NewBufferString(encoded_formdata))
	if err != nil {
		log.Printf("Unable to form request, error - %s", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Unable to get response, error - %s", err)
	}

	defer resp.Body.Close()

	// respBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Printf("unable to read response body, error - %s", err)
	// }

	// respBodyString := string(respBody)

	// fmt.Println(respBodyString)
	var tokenResponse TokenResponse

	respBodydecoder := json.NewDecoder(resp.Body)
	// since Body is a ReadCloser the second time it will find it already read.
	// so during second read it will throw EOF error
	if err := respBodydecoder.Decode(&tokenResponse); err != nil {
		log.Printf("unable to decode Json Response, error - %s", err)
	}
	fmt.Println(tokenResponse)

	return tokenResponse.AccessToken

}
