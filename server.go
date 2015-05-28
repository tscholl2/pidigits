package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"./digits"
)

type requestMessage struct {
	Start  int `json:"start"`
	Length int `json:"length"`
}

type responseMessage struct {
	Digits string `json:"digits"`
	Error  string `json:"err"`
}

func onRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// check for errors and read request
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{err:%s}", err.Error())))
		return
	}
	start, err := strconv.Atoi(r.FormValue("start"))
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{err:\"%s\"}", err.Error())))
		return
	}
	length, err := strconv.Atoi(r.FormValue("length"))
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{err:\"%s\"}", err.Error())))
		return
	}
	reqMsg := requestMessage{Start: start, Length: length}
	if reqMsg.Start < 0 || reqMsg.Start > 10000000 {
		w.Write([]byte("{err:\"Invalid starting place.\"}"))
		return
	}
	if reqMsg.Length <= 0 || reqMsg.Length > 100 {
		w.Write([]byte("{err:\"Invalid length.\"}"))
		return
	}
	// reply
	var resMsg responseMessage
	chars := digits.Get(reqMsg.Start, reqMsg.Length)
	resMsg.Digits = string(*chars)
	bytes, err := json.Marshal(resMsg)
	if err != nil {
		bytes = []byte(fmt.Sprintf("{err:%s}", err.Error()))
	}
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/", onRequest)
	http.ListenAndServe(":8888", nil)
}
