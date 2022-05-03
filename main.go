package main

import (
	config "Products/config"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	Products "github.com/hmartins98/Contracts/Products"
)

// var c cache.ICache
// var c database.IDatabase
var db *sql.DB

func main() {

	var err error
	db, err = config.GetPostgresDB()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//teste

	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatalf("Failed to listener %v", err)
	}

	s := grpc.NewServer()
	Products.RegisterProductsContractServer(s, &server{})

	log.Printf("listening on port 5000")

	if error := s.Serve(lis); error != nil {
		log.Fatalf("Failed to serve: %v", error)
	}
}
