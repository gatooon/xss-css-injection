package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	IPAddr  string `yaml:"IP_Address"`
	DN      string `yaml:"DN"`
	Port    string `yaml:"Port"`
	CrtPath string `yaml:"Crt_path"`
	KeyPath string `yaml:"Key_path"`
}

func Parse_conf(yml_path string) Conf {
	conf := Conf{}

	if yml_path == "" {
		fmt.Println("No conf file provided")
		os.Exit(1)
	}

	source, err := ioutil.ReadFile(yml_path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = yaml.Unmarshal([]byte(source), &conf)
	if err != nil {
		log.Println(err)
	}

	CheckField(conf)

	return conf
}

func CheckField(conf Conf) {
	is_empty := false
	if conf.IPAddr == "" {
		fmt.Println("IP_Address field is empty")
		is_empty = true
	}
	if conf.Port == "" {
		fmt.Println("Port field is empty")
		is_empty = true
	}
	if conf.CrtPath == "" {
		fmt.Println("Crt_path field is empty")
		is_empty = true
	}
	if conf.KeyPath == "" {
		fmt.Println("Key_path field is empty")
		is_empty = true
	}
	if conf.DN == "" {
		fmt.Println("DN field is empty")
		is_empty = true
	}
	if is_empty == true {
		os.Exit(1)
	}
}
