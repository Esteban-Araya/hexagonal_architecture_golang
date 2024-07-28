package database

import (
    "database/sql"
    _ "github.com/lib/pq"
    "fmt"
	"log"
	"net"
)

func getIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	adress := conn.LocalAddr().(*net.UDPAddr)

	return adress.IP
}

func InitDB() (DB *sql.DB, err error){
	
    server := "192.168.0.21"
	if server == "" {
		server = getIP().String()
	} 

	connURL := fmt.Sprintf("postgres://postgres:postgres@%s/postgres?sslmode=disable", server)
	db, err := sql.Open("postgres", connURL)
	if err != nil {
        return db, err
    }

	if err = db.Ping(); err != nil {
        return db, err
    }
	fmt.Println("Connected to DB")

	return db, nil

   
}