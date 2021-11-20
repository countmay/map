package main

import (
	
	"log"
	"net/http"
	"text/template"
	"encoding/json"
	"io/ioutil"

	
	"database/sql"
   "fmt"
    _ "github.com/lib/pq"

)



type Ticket struct{
// ticketid string
// //dateandtime timestamp
// Isopen bool
// body string
X float64 `json: "x"`
Y float64 `json: "y"`
}



var allpoints []Ticket
func main() {

	// 
	http.HandleFunc("/", mapHandler)
	log.Fatal(http.ListenAndServe(":3010", nil))

}

func  mapHandler(w http.ResponseWriter, r *http.Request){
allpoints=nil

// connection to database
connStr := "postgres://countmay:eebc2f64-3d74-4517-9d7f-fba1e58d35dc@134.122.78.80/taldykol?sslmode=disable"
db, err := sql.Open("postgres", connStr)

    if err != nil {
        // panic(err)
    } 


//error if db falls
	rows, err := db.Query("SELECT ST_X(coordinates), ST_Y(coordinates) FROM tickets")
	if err != nil {
		// panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		p := Ticket{}
		err := rows.Scan(&p.X,&p.Y)
		if err != nil {
			// panic(err)
		}
allpoints=append(allpoints,p)
	}

    defer db.Close()

	fmt.Println(allpoints)


	// j, _ := json.Marshal(allpoints)
	// fmt.Println(j)

	file, _ := json.MarshalIndent(allpoints, "", " ")
 
	_ = ioutil.WriteFile("test.json", file, 0644)

	tpl, _ := template.ParseFiles("page.html")

	// if r.URL.Path != "/" {
	// 	h.ErrorHandler(w, r, "404")
	// 	return
	// }
	//need Error handling


	//nil 
	tpl.Execute(w,nil)
}



