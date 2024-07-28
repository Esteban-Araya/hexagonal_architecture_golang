package main 

import(
	"Project/pkg/settings/database"
	"Project/pkg/settings/rout"
	"net/http"
	"log"
)

func main(){
	db, err := database.InitDB() 

	if err != nil{
		log.Fatal(err.Error())
	}
	defer db.Close()
	
	server := http.Server{Addr:":8080", Handler: nil,}
	log.Println("START")
	err = rout.StartListener(&server, db) 
	if err != nil{ 
		log.Fatal(err.Error())
	}
	
}