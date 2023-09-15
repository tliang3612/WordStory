package go_word

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
)

// return next chapter and false if invalid input. true if valid
func ExecuteInput(input string, s Story, chp *string, t *template.Template) bool {

	//convert input to int
	index, err := strconv.Atoi(input)

	if err != nil {
		panic(err)
	}

	if chapter, ok := s[*chp]; ok {

		n := len(s[*chp].Options)

		if index >= n {
			fmt.Println("Index out of bounds")
			return false
		}

		gotoChp := chapter.Options[index].Chapter

		//write to console the chapter contents if no err
		if err := t.Execute(os.Stdout, s[gotoChp]); err != nil {
			log.Printf("%v", err)
			panic(err)
		}
		*chp = gotoChp
		return true
	}

	return false
}

func ExecuteDefaultChapter(s Story, chp string, t *template.Template) {
	if err := t.Execute(os.Stdout, s[chp]); err != nil {
		log.Printf("%v", err)
		panic(err)
	}
}
