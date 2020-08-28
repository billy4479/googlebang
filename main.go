package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

var (
	addr       = flag.String("addr", "localhost:4479", "Specify the host and port where to run")
	redirectTo = flag.String("to", "https://www.google.com/search?q=%s&ie=UTF-8", "The default redirect with %%s where there will be the search string")

	bangMatch = regexp.MustCompile("!.* .*$")
)

func main() {
	flag.Parse()
	http.HandleFunc("/search", search)
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Println("Listening on", *addr)
	http.ListenAndServe(*addr, nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	input := r.Form.Get("search")
	if input == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	hasBang := bangMatch.MatchString(input)
	if !hasBang {
		http.Redirect(w, r, fmt.Sprintf(*redirectTo, input), http.StatusSeeOther)
		return
	}

	acronym := ""
	search := ""
	endAcronym := false
	for _, c := range input {
		if !endAcronym {
			if c != ' ' {
				if c == '!' {
					continue
				}
				acronym += string(c)
			} else {
				endAcronym = true
			}
		} else {
			search += string(c)
		}
	}

	for _, bang := range bangs {
		if bang.acronym == acronym {
			http.Redirect(w, r, fmt.Sprintf(bang.searchString, url.QueryEscape(search)), http.StatusFound)
			return
		}
	}
	http.Redirect(w, r, fmt.Sprintf(*redirectTo, input), http.StatusSeeOther)
}
