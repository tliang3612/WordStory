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

	filename := flag.String("file", "gopher.json", "the json file with the goword story")

	valid := false

	fmt.Println("Do you wish to start a console or a web story? Type c or w")
	var input string

	//do while valid is false. ok is true in first loop
	for ok := true; ok; ok = valid {
		fmt.Scan(&input)

		if input == "c" {
			RunConsoleStory(*filename)
		} else if input == "w" {
			RunWebStory(*filename)
		} else {
			fmt.Println("Enter a valid entry")
			valid = false
		}
	}

}

func RunWebStory(filename string) {
	//define flags
	port := flag.Int("port", 3000, "the port to start the word story on web")

	// parse defined flags
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", filename)

	story := createStory(filename)

	tpl := template.Must(template.New("web").Parse(webStoryTemplate))

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

func RunConsoleStory(filename string) {
	story := createStory(filename)

	tpl := template.Must(template.New("console").Parse(consoleStoryTemplate))

	chp := "intro"

	go_word.ExecuteDefaultChapter(story, chp, tpl)

	var valid bool

	// default values
	var quit = false
	var input = "0"

	for quit == false {
		valid = false

		fmt.Println("Select the chapter to go to as an integer")
		fmt.Scan(&input)

		if input == "q" {
			quit = true
			continue
		}

		for valid == false {
			valid = go_word.ExecuteInput(input, story, &chp, tpl)
			if valid == false {
				fmt.Println("Please enter valid input")
				fmt.Scan(&input)
				continue
			}
			valid = true
		}
	}

}

func createStory(filename string) go_word.Story {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	story, err := go_word.CreateStory(f)
	if err != nil {
		panic(err)
	}

	return story
}

func pathFunc(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "/story" || path == "/story/" {
		path = "/story/intro" //set default path
	}

	//trim "/story/" from returned path
	return path[len("/story/"):]
}

var webStoryTemplate = `
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
					<li>
						<a href= "/story/{{.Chapter }}">{{.Text }}</a>
					</li>
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

var consoleStoryTemplate = `
{{.Title}}

{{range .Paragraphs}}
{{.}}
{{end}}

{{range .Options}}
{{.Text}}
{{end}}



`
