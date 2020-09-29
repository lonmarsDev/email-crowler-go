package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/realpamisa/gmail-extractor/decode"
	"github.com/realpamisa/gmail-extractor/models"
)

func GetAllMails(token string, userID string) ([]byte, error) { //getting all users emails
	response, err := http.Get("https://gmail.googleapis.com/gmail/v1/users/" + userID + "/messages?access_token=" + token)
	if err != nil {
		return nil, fmt.Errorf("failed getting message : %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	contents = []byte(contents)
	var newContent models.Message_content
	err = json.Unmarshal(contents, &newContent)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range newContent.Messages {
		EmailExtractor(token, userID, v.Id)
	}

	return nil, nil
}

func EmailExtractor(token string, userID string, messageID string) (string, error) {
	response, err := http.Get("https://gmail.googleapis.com/gmail/v1/users/" + userID + "/messages/" + messageID + "?access_token=" + token)
	if err != nil {
		fmt.Errorf("failed getting message : %s", err.Error())
		return "", nil
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("failed reading response body: %s", err.Error())
		return "", nil
	}

	contents = []byte(contents)
	var newContents models.Emails
	err = json.Unmarshal(contents, &newContents)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i, v := range newContents.Payload.Parts {
		if i == 0 {
			// var newContent = []byte(v.Body.Data_string)
			decode.HandleEmailDecode(v.Body.Data_string)

		}

	}

	return "", nil
}
