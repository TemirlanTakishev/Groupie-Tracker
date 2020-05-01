package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

var artists []Artist

func mainPage(w http.ResponseWriter, r *http.Request) {
	// будет располагаться главная странитца с группами
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, artists)
	if err != nil {
		panic(err)
	}

}

func Artisti(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[8:]
	p, _ := strconv.Atoi(id)
	t, err := template.ParseFiles("artist.html")
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, artists[p-1])
	if err != nil {
		panic(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/artist/", Artisti)
	fmt.Print("live and serve")

	apiUrl := "https://groupietrackers.herokuapp.com/api/artists"
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil) //обращаемся к нашей ссылке
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", nil)
}
