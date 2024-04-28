package talker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

/* Struct representing a response from the listener. */
type Response struct {
	Msg string `json:"Msg"` // The status message
	Name string `json:"Name,omitempty"` // The name of the challenge
	Link string `json:"Link,omitempty"` // The submission link if the challenge is completed
}

/* Sends a request to the submission website */
func SubmitFlag(flag string, resTemp chan *Response) {
	url := "http://bwsi-lab-submission.vercel.app/api/listener/submit"

	rawBody := map[string]string{"flag": flag}
	body, err := json.Marshal(rawBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	response := &Response{}
	derr := json.NewDecoder(res.Body).Decode(response)

	if derr != nil {
		fmt.Println(err)
		return
	}

	resTemp <- response

}
