package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	go_word "github.com/tliang3612/wordstory"
)

func main() {
	//define flags
	port := flag.Int("port", 3000, "the port to start the word story on web")
	filename := flag.String("file", "gopher.json", "the json file with the goword story")

	// parse defined flags
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := go_word.DecodeJsonToStruct(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyTemplate))

	h := go_word.NewHandler(
		story,
		go_word.WithTemplate(tpl),
		go_word.WithPathFunc(pathFunc),
	)

	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting the server on port %d \n", *port)
	//Create http server with port and handler
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

}

func pathFunc(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "/story" || path == "/story/" {
		path = "/stories/"
	}

	//trim "/story/" from returned path
	return path[len("/story/"):]
}

var storyTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>{{.Title}}}</title>	
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
					<li href= "/story/{{.Chapter }}">{{.Text }}</li>
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
