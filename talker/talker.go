package talker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Msg string `json:"Msg"`
}

func SubmitFlag(flag string, resTemp chan *Response) {
	url := "http://localhost:3000/api/listener/submit"

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
