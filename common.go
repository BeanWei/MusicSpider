package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
)

func reqHandler(reqmethod, url, data string) map[string]interface{} {
	client := &http.Client{}
	dataByte := bytes.NewBuffer([]byte(data))
	req, _ := http.NewRequest(reqmethod, url, dataByte)
	resp, _ := client.Do(req)
	status := strconv.Itoa(resp.StatusCode)
	if status == "200" {
		body, _ := ioutil.ReadAll(resp.Body)
		feedback := map[string]interface{}{"status": status, "result": body}
		return feedback
	}
	feedback := map[string]interface{}{"status": status, "result": ""}
	return feedback
}

func checkDicKey(dict map[string]int, thisKey ...string) bool {
	for _, k := range thisKey {
		_, ok := dict[k]
		if !ok {
			return ok
		}
	}
	return true

}
