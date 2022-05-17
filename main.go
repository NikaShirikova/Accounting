package main

import (
	"AccountingService/server"
	"AccountingService/server/api"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	conn := "user=root password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Println(err)
		return
	}

	server.Database = db
	defer db.Close()

	//ExecInsertTable()
	//QueryInsertTable()

	http.HandleFunc("/incomes", api.PrintIncomes)
	http.HandleFunc("/expenses", api.PrintExpenses)

	fmt.Println("Date...")

	//managerCert := autocert.Manager{
	//	Cache:      autocert.DirCache("certs"),
	//	Prompt:     autocert.AcceptTOS,
	//	HostPolicy: autocert.HostWhitelist("DESKTOP-GM0R8G3"),
	//}
	//
	//server := &http.Server{
	//	Addr:      ":8181",
	//	TLSConfig: &tls.Config{GetCertificate: managerCert.GetCertificate},
	//}
	//
	//errTLS := server.ListenAndServeTLS("", "")
	//if errTLS != nil {
	//	fmt.Println(errTLS)
	//	return
	//}
	http.ListenAndServe(":8181", nil)
}
