package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/realpamisa/gmail-extractor/models"
)

func GetUserID(token string) (string, error) {

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		fmt.Errorf("failed getting user info: %s", err.Error())
		return "", nil
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("failed reading response body: %s", err.Error())
		return "", nil
	}
	contents = []byte(contents)
	var newContent models.Content
	json.Unmarshal(contents, &newContent)

	return newContent.Id, nil
}
