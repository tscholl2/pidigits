package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tscholl2/pihash/digits"
)

type requestMessage struct {
	Place  int `json:"place"`
	Length int `json:"length"`
}

type responseMessage struct {
	Digits string `json:"digits"`
	Error  string `json:"err"`
}

func onRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// check for errors and read request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{err:%s}", err.Error())))
		return
	}
	var reqMsg requestMessage
	json.Unmarshal(body, &reqMsg)
	if reqMsg.Place < 0 || reqMsg.Place > 1000000000 {
		w.Write([]byte("{err:\"Invalid starting place.\"}"))
		return
	}
	if reqMsg.Length <= 0 || reqMsg.Length > 100 {
		w.Write([]byte("{err:\"Invalid length.\"}"))
		return
	}
	// reply
	var resMsg responseMessage
	chars := digits.Get(reqMsg.Place, reqMsg.Length)
	resMsg.Digits = string(*chars)
	bytes, err := json.Marshal(resMsg)
	if err != nil {
		bytes = []byte(fmt.Sprintf("{err:%s}", err.Error()))
	}
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/", onRequest)
	http.ListenAndServe(":8899", nil)
}
