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

	body := []byte("{" + `"flag": ` + `"` + flag + `"` + "}")

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	response := &Response{}
	derr := json.NewDecoder(res.Body).Decode(response)

	fmt.Println(response)
	if derr != nil {
		fmt.Println(err)
		return
	}

	resTemp <- response

}
