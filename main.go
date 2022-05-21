package main

import (
	"AccountingService/server"
	"AccountingService/server/api"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"time"
)

func editDB(work server.WorkWithDB) {
	switch work.(type) {
	case *server.Income:
		fmt.Println("Работа с таблицей accounting.income:")
	case *server.Expenses:
		fmt.Println("Работа с таблицей accounting.expenses:")
	default:
		fmt.Println("Такой таблицы нет!")
	}

	fmt.Println("Добавим строку в таблицу")
	errInsert := work.ExecInsertTable()
	if errInsert != nil {
		fmt.Println(errInsert)
	}
	fmt.Println("Выведем строки из таблицы")
	errSelect := work.QuerySelectTable()
	if errSelect != nil {
		fmt.Println(errSelect)
	}
	fmt.Println("Удалим строку из таблицы")
	var id int
	fmt.Println("Введите ID строки, которую хотите удалить")
	fmt.Fscan(os.Stdin, &id)
	errDelete := work.QueryDeleteRow(id)
	if errDelete != nil {
		fmt.Println(errDelete)
	}
	fmt.Println("Работа с БД успешно проведена")
}

func main() {
	conn := "user=root password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Println(err)
		return
	}

	server.Database = db
	defer db.Close()

	timeNow := time.Now().Format("2006-01-02 15:04:05")
	income := &server.Income{0, 434, "ЗП", "Криптософт", timeNow}
	editDB(income)

	expenses := &server.Expenses{0, 434, "Продукты", "Магнит", timeNow}
	editDB(expenses)

	http.HandleFunc("/income", api.PrintIncomes)
	http.HandleFunc("/expenses", api.PrintExpenses)

	fmt.Println("Date...")
	http.ListenAndServe(":80", nil)
}
