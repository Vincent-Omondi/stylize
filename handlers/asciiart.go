// handlers/asciiart.go
package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/Vincent-Omondi/stylize/asciiart"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	var errors []string

	// check if the method is not POST
	if r.Method != http.MethodPost {
		errors = append(errors, "Method not allowed")
		ErrorHandler(w, r, http.StatusMethodNotAllowed, errors)
		return
	}
	// receive tesxt and banner from text field and banner option
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	asciiChars, err := asciiart.LoadAsciiChars(banner + ".txt")
	if err != nil {
		errors = append(errors, "Error 500: Internal Server Error")
		ErrorHandler(w, r, http.StatusInternalServerError, errors)
		return
	}

	// check if the text field is empty
	if len(text) == 0 {
		errors = append(errors, "Error 400: Bad Request - Text field is empty")
		ErrorHandler(w, r, http.StatusBadRequest, errors)
		return
	}
	if err != nil {
		log.Println(err)
		errors = append(errors, err.Error())
		ErrorHandler(w, r, http.StatusBadRequest, errors)
		return
	}

	// split the text using  newline and carriage return and creates the art and appends it to asciiArt
	str := strings.Split(text, "\r\n")
	var asciiArt string
	for _, wrd := range str {
		asciiArt += asciiart.PrintAsciiArt(wrd, asciiChars)
	}

	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(asciiArt))
		return
	}

	data := AsciiData{
		Text:   text,
		Banner: banner,
		Result: asciiArt,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		errors = append(errors, err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError, errors)
		return
	}

	// Execute the data and send it to tmpl
	err = tmpl.Execute(w, data)
	if err != nil {
		errors = append(errors, err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError, errors)
	}
}
