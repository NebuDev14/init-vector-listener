package talker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	msg	string	`json:"msg"`
}

func SubmitFlag(flag string) {
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

	

}
