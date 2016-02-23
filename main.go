package main

import (
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/amaxwellblair/crud_app/app/models"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, t)
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/robots/":
		if r.Method == "GET" {
			getIndexRobots(w, r)
		} else {
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		}
	case "/robots/new":
		if r.Method == "GET" {
			getNewRobots(w, r)
		} else if r.Method == "POST" {
			postNewRobots(w, r)
		} else {
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		}
	case robotsShowID(r.URL.Path):
		if r.Method == "GET" {
			getShowRobot(w, r, path.Base(r.URL.Path))
		}
	default:
		http.NotFound(w, r)
	}
}

func robotsShowID(p string) string {
	b, err := path.Match("/robots/id/*", p)
	if err != nil {
		return ""
	}
	if b == true {
		return "/robots/id/" + path.Base(p)
	}
	return ""
}

func getIndexRobots(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/robots/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	s := mustOpenStore()
	defer s.Close()
	robots, err := s.All()
	t.Execute(w, robots)
}

func getNewRobots(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/robots/new.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, t)
}

func postNewRobots(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	function := r.FormValue("function")
	s := mustOpenStore()
	defer s.Close()

	s.CreateRobot(name, function)
	http.Redirect(w, r, "", http.StatusFound)
}

func getShowRobot(w http.ResponseWriter, r *http.Request, id string) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t, err := template.ParseFiles("app/views/robots/show.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	s := mustOpenStore()
	defer s.Close()

	rbt, err := s.Robot(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	t.Execute(w, rbt)
}

func mustOpenStore() *robots.Store {
	s := robots.NewStore("db/dev.db")
	err := s.Open()
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/robots/", robotsHandler)
	http.ListenAndServe(":8080", nil)
}
