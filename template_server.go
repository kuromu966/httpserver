package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/kuromu966/httpserver/tpl_error"
	"github.com/kuromu966/httpserver/tpl_sample"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"text/template"
)

var ServerParameter = Params{
	name:          "template server",
	port:          8080,
	root_path:     "/",
	template_path: ".",
	config_path:   ".",
}

type IndexConfig struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type ContentConfig struct {
	Name     string `json:"name"`
	Uri      string `json:"uri"`
	Template string `json:"template"`
	Content  string `json:"content"`
}

type ServerConfig struct {
	IndexList   []IndexConfig   `json:"index"`
	ContentList []ContentConfig `json:"contents"`
}

var configuration ServerConfig

func get_content_config(uri string) *ContentConfig {
	var result *ContentConfig
	var notfound *ContentConfig

	for _, value := range configuration.ContentList {
		if value.Name == "sys_notfound" {
			notfound = &value
		}

		if value.Uri != uri {
			continue
		}
		result = &value
		break
	}

	if result == nil && notfound == nil {
		message := "Failed to found 'sys_notfound' template."
		panic(message)
	}

	if result == nil {
		result = notfound
	}

	return result
}

func get_path(filename string) *string {
	base_path := ServerParameter.template_path
	path := filepath.Join(base_path, filename)
	return &path
}

func get_template(tpl_name string) *template.Template {

	var path *string

	for _, value := range configuration.IndexList {
		if value.Name != tpl_name {
			continue
		}
		path = get_path(value.Path)
		break
	}

	if path == nil {
		message := fmt.Sprintf("Failed to found '%s' template.", tpl_name)
		panic(message)
	}

	tpl, err := template.ParseFiles(*path)
	if err != nil {
		panic(err)
	}

	return tpl
}

func get_template_from_config(c ContentConfig) *template.Template {
	if c.Template == "" {
		message := fmt.Sprintf("Failed to find 'tempalte' attribtes on '%s' content.", c.Name)
		panic(message)
	}
	return get_template(c.Template)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	uri := r.URL.Path[1:]
	content_config := get_content_config(uri)

	switch content_config.Template {
	case "sample":
		tpl := get_template_from_config(*content_config)
		content_parameter_file_path := get_path(content_config.Content)
		page := tpl_sample.LoadContent(*content_parameter_file_path)
		err := tpl.Execute(w, page)
		if err != nil {
			panic(err)
		}

	case "sys_notfound":
		tpl := get_template_from_config(*content_config)
		content_parameter_file_path := get_path(content_config.Content)
		page := tpl_error.LoadContent(*content_parameter_file_path)
		err := tpl.Execute(w, page)
		if err != nil {
			panic(err)
		}

	default:
		tpl := get_template_from_config(*content_config)
		content_parameter_file_path := get_path(content_config.Content)
		page := tpl_error.LoadContent(*content_parameter_file_path)
		err := tpl.Execute(w, page)
		if err != nil {
			panic(err)
		}

	}

}

func SetServerParameter(name string, port int, root string, tpl string, config string) {
	ServerParameter = Params{
		name:          name,
		port:          port,
		root_path:     root,
		template_path: tpl,
		config_path:   config,
	}
}

func load_configuration(p *Params) {
	path := ConfigPath(p)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &configuration)
}

func Run() {
	PrintParameter(&ServerParameter)
	load_configuration(&ServerParameter)
	http.HandleFunc(RootPath(&ServerParameter), viewHandler)
	http.ListenAndServe(ListenPort(&ServerParameter), nil)
}
