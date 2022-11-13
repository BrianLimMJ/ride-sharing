package main

//netstat -a -o to list all ports

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//"database/sql"
	//"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Passenger Object
type Passenger struct {
	firstName    string `json:"firstName"`
	lastName     string `json:"lastName"`
	mobileNumber int    `json:"mobileNumber"`
	emailAddress string `json:"emailAddress"`
}

// Array of Passengers
type AllPassengers struct {
	Passengers map[string]Passenger `json:"Passengers"`
}

// Add data to Array of Passengers
var passengers map[string]Passenger = map[string]Passenger{
	"p1": Passenger{"Kajus", "Bullock", 1234578, "KajusB@gmail.com"},
	"p2": Passenger{"Eileen", "Hubbard", 2345781, "EileenH@gmail.com"},
	"p3": Passenger{"Rome", "Hawes", 3457812, "RomeH@gmail.com"},
}

// Array of Drivers
type Driver struct {
	firstName        string `json:"firstName"`
	lastName         string `json:"lastName"`
	mobileNumber     int    `json:"mobileNumber"`
	emailAddress     string `json:"emailAddress"`
	identificationNo string `json:"identificationNo"`
	carLicense       string `json:"carLicense"`
	isAvailable      bool   `json:"isAvailable"`
}

type AllDrivers struct {
	Drivers map[string]Driver `json:"Drivers"`
}

// Add data to Array of Drivers
var drivers map[string]Driver = map[string]Driver{
	"d1": Driver{"Esmay", "Cope", 87654321, "EsmayC@gmail.com", "S2271158Z", "ABC123", true},
	"d2": Driver{"Aaliyah", "Cooley", 76543218, "AaliyaC@gmail.com", "S2271158Z", "BC123A", false},
	"d3": Driver{"Beatrice", "Conroy", 65432187, "BeatriceC@gmail.com", "S2271158Z", "C123AB", true},
}

type AllTrip struct {
	Trips map[string]Trip `json:"Trips"`
}

type Trip struct {
	startPostal  int    `json:"startPostal"`
	endPostal    int    `json:"endPostal"`
	dateTime     string `json:"dateTime"`
	driver_id    string `json:"driver_id"`
	passenger_id string `json:"passenger_id"`
}

// func main() {
// 	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
// 		n, err := fmt.Fprintf(w, "Hellow World")
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
// 	})
// 	_ = http.ListenAndServe(":3000")
// }

func main() {
	router := mux.NewRouter()
	//router.HandleFunc("api/v1/courses", coursesFilter)
	router.HandleFunc("api/v1/Drivers/{driver_id}", alldrivers)
	//fmt.Println(time.Now().Date())
	fmt.Println("Listening at port 5000")
	fmt.Println(drivers)
	log.Fatal(http.ListenAndServe(":5000", router))

}

func alldrivers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	/*if r.Mthod == "GET"
	{
		if v, ok := params["course_id"]; ok{
			json.NewEncoder(w).Encode(v)
		}
	} else {
		if v, ok := params["course_id"]; ok{
			delete(courses, params["course_id"])
			fmt.Fprintf(w, params["Course_id"]+" Deleted")
		}
	}
	*/
	if v, ok := drivers[params["driver_id"]]; ok {
		if r.Method == "GET" {
			json.NewEncoder(w).Encode(v)
		} else {
			delete(drivers, params["driver_id"])
			fmt.Fprintf(w, params["Driver_id"]+" Deleted ")
		}
	} else {
		fmt.Fprintf(w, "Invalid Driver ID")
	}
}
