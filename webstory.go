package go_word

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("default").Parse(defaultHandlerTemplate))
}

var defaultHandlerTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>{{ .Title }}}</title>	
	</head>

	<body>
		<div id="mid">
			<h1>{{ .Title }}</h1>	
				{{range .Paragraphs}}
					<p>{{.}}</p>
				{{end}}
			<ul>
				{{range .Options}}
					<span id="choice"></span>
					<li href= /story/"{{.Chapter }}">{{.Text }}</li>
				{{ end }}
			</ul>
		</div>

		<style>
			body {
				background-color: #DAC0A3;
			}

			#mid {
				max-width: 600px;
				margin: auto;
				left: 30px;
				padding: 20px;
				border: 5px solid #0F2C59;
				background-color: #EADBC8;
			}

			h1 {
				text-align: center;
			}

			a {
				padding: 10px;
				text-decoration: none;
			}

			#choice {
				font-size: 48px;
			}
		</style>
	</body>
</html>`

type handler struct {
	s        Story
	t        *template.Template
	pathFunc func(r *http.Request) string
}

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	if t == nil {
		t = tpl
	}

	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFunc = fn
	}
}

func defaultPathFunc(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "intro"
	}

	return path
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	//set default hanlder properties
	h := handler{s, tpl, defaultPathFunc}

	for _, opt := range opts {
		if opt != nil {
			opt(&h)
		}
	}

	return h
}

// Called everytime an http request is made
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := h.pathFunc(r)

	//Check if path is a valid key for story
	if chapter, ok := h.s[path]; ok {
		//write to http server the chapter contents if no err
		if err := h.t.Execute(w, chapter); err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)

		}
		return
	}

	fmt.Printf("Error connecting to path %s \n", path)

}
