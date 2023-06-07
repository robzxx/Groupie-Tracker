package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

//STRUCTURE

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Creationdate int      `json:"creationDate"`
	Firstalbum   string   `json:"firstAlbum"`
	Loc          string   `json:"locations"` //récupération du lien
	Locations    struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	}
	Dates     string `json:"dates"` //récupération du lien
	Dates_loc struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	}
}

/*
type Relations struct {
	Id             int      `json:"id"`
	DatesLocations []string `json:"datesLocations"`
}
*/

func main() {

	var Api_artist []Artists //Tableau de struct
	/*var Api_dateslocations []Relations*/

	//LIENS API//

	response_artists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists") //Permet d'aller chercher les artists dans l'API
	if err != nil {
		fmt.Println("ERROR")
	}

	/*
		response_dateslocations, err := http.Get("https://groupietrackers.herokuapp.com/api/relation") //Permet d'aller chercher tous les artistes au lieu d'en avoir qu'un seul (API)
		if err != nil {
			fmt.Println("ERROR")
		}


			response_locations, err := http.Get("https://groupietrackers.herokuapp.com/api/Locations") //Permet d'aller chercher tous les artistes au lieu d'en avoir qu'un seul (API)
			if err != nil {
				fmt.Println("ERROR")
			}

			response_dates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates") //Permet d'aller chercher tous les artistes au lieu d'en avoir qu'un seul (API)
			if err != nil {
				fmt.Println("ERROR")
			}
	*/

	//DATA_STRUCTS//

	responseData_artists, err := ioutil.ReadAll(response_artists.Body)
	if err != nil {
		log.Fatal(err)
	}

	/*
		responseData_dateslocations, err := ioutil.ReadAll(response_dateslocations.Body)
		if err != nil {
			log.Fatal(err)
		}


			responseData_locations, err := ioutil.ReadAll(response_locations.Body)
			if err != nil {
				log.Fatal(err)
			}

			responseData_dates, err := ioutil.ReadAll(response_dates.Body)
			if err != nil {
				log.Fatal(err)
			}
	*/

	//UNMARSHAL//

	json.Unmarshal(responseData_artists, &Api_artist)
	/*
		json.Unmarshal(responseData_dateslocations, &Api_artist)
		json.Unmarshal(responseData_locations, &Api_artist)
		json.Unmarshal(responseData_dates, &Api_artist)
	*/

	//BARRE DE RECHERCHE

	/*Scan_input := r.FormValue("Scan_input")*/

	//TEMPLATES//

	tpl := template.Must(template.ParseGlob("html/*"))

	//PAGE INDEX

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index.html", Api_artist) // Faire afficher la page avec toutes les pochettes d'ablums
	})

	//PAGE MENU

	http.HandleFunc("/menu.html", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "menu.html", Api_artist) // Faire afficher la page avec toutes les pochettes d'ablums
	})

	//PAGE ABOUT

	http.HandleFunc("/about.html", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "about.html", Api_artist)
	})

	//PAGE ARTISTES//

	http.HandleFunc("/artists/", func(w http.ResponseWriter, r *http.Request) {
		ArtistsURL := r.URL.RequestURI()[9:]         //pour "scan" le chiffre après le artists/
		ArtistsAtoi, err := strconv.Atoi(ArtistsURL) //pour convretir le string en int et récupérer un chiffre
		if err != nil {
			log.Fatal(err)
		}
		tpl.ExecuteTemplate(w, "artists.html", Api_artist[ArtistsAtoi-1]) //permet de faire afficher la page pour chaque artist
	})

	/*http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./assets/*")))) //link tous le css (non fonctionnel actuellement)*/
	log.Fatal(http.ListenAndServe(":8080", nil)) //Port du server
}
