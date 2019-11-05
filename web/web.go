package web

import (
	"gopkg.in/macaron.v1"
	"tcf/routers"
	"strings"
	"html/template"
	"net/http"
)

func RunWeb() {
	m := macaron.Classic()
	m.Use(macaron.Renderer(
		macaron.RenderOptions{
			Directory:  "templates",
			Extensions: []string{".tmpl", ".html"},
			Funcs: []template.FuncMap{map[string]interface{}{
				"Replace": func(str *string) string {
					t := strings.Replace(*str, "\n", "<br>", -1)
					return t
				},
				"Split": func(str *string) []string {
					return strings.Split(*str, ",")
				},
				"unescaped": func(x string) interface{} { return template.HTML(x) },
			}},
			Delims:          macaron.Delims{"{{", "}}"},
			Charset:         "UTF-8",
			IndentJSON:      true,
			IndentXML:       true,
			PrefixJSON:      []byte("macaron"),
			PrefixXML:       []byte("macaron"),
			HTMLContentType: "text/html",
		}))

	m.Get("/tcf/index/", routers.GetIndex)
	m.Post("/tcf/evluate/", routers.GetTcfParas)

	m.Use(macaron.Renderer(
		macaron.RenderOptions{
			Directory:  "templates",
			Extensions: []string{".tmpl", ".html"},
			Funcs: []template.FuncMap{map[string]interface{}{
				"Replace": func(str *string) string {
					t := strings.Replace(*str, "\n", "<br>", -1)
					return t
				},
				"Split": func(str *string) []string {
					return strings.Split(*str, ",")
				},
				"unescaped": func(x string) interface{} { return template.HTML(x) },
			}},
			Delims:          macaron.Delims{"{{", "}}"},
			Charset:         "UTF-8",
			IndentJSON:      true,
			IndentXML:       true,
			PrefixJSON:      []byte("macaron"),
			PrefixXML:       []byte("macaron"),
			HTMLContentType: "text/html",
		}))

	http.ListenAndServe("0.0.0.0:4001", m)
}


