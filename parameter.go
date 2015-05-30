package httpserver

import (
	"fmt"
)

type Params struct {
	name          string
	port          int
	root_path     string
	template_path string
	config_path   string
}

func SetServerName(parameter *Params, name string) {
	parameter.name = name
}

func SetServerPort(parameter *Params, port int) {
	parameter.port = port
}

func SetServerRoot(parameter *Params, path string) {
	parameter.root_path = path
}

func SetServerTemplate(parameter *Params, path string) {
	parameter.template_path = path
}

func ListenPort(parameter *Params) string {
	listen_port := fmt.Sprintf(":%d", parameter.port)
	return listen_port
}

func RootPath(parameter *Params) string {
	return parameter.root_path
}

func TemplatePath(parameter *Params) string {
	return parameter.template_path
}

func ConfigPath(parameter *Params) string {
	return parameter.config_path
}

func PrintParameter(parameter *Params) {
	fmt.Printf("[server name]\t%s\n", parameter.name)
	fmt.Printf("[address]\tlocalhost:%d\n", parameter.port)
	fmt.Printf("[root]\t\t%s\n", parameter.root_path)
	fmt.Printf("[tempalte]\t%s\n", parameter.template_path)
	fmt.Printf("[config]\t%s\n", parameter.config_path)
}
