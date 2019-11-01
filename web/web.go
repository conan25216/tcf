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
	m.Get("/hello/*", routers.GetParameter)
	m.Get("/hello/index", routers.GetIndex)
	m.Post("/processor/go/", routers.GetTcfParas)

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

	http.ListenAndServe("0.0.0.0:9527", m)
}


