package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
	"time"
)

type Income struct {
	Id    int
	Sum   float64
	Type  string
	Place string
	Date  time.Time
}

type Expenses struct {
	Id       int
	Sum      float64
	Category string
	Place    string
	Date     time.Time
}

var database *sql.DB

func main() {
	conn := "user=root password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Println(err)
		return
	}

	database = db
	defer db.Close()

	//execInsertTable()
	//queryInsertTable()

	http.HandleFunc("/incomes", printIncomes)
	http.HandleFunc("/expenses", printExpenses)

	fmt.Println("Date...")
	http.ListenAndServe(":8181", nil)
}

func printIncomes(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("SELECT * FROM Accounting.income")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()
	incomes := []Income{}

	for rows.Next() {
		income := Income{}
		errScan := rows.Scan(&income.Id, &income.Sum, &income.Type, &income.Place, &income.Date)
		if errScan != nil {
			fmt.Println(errScan)
			continue
		}
		incomes = append(incomes, income)
	}

	tempIncomes, errTemp := template.ParseFiles("templates/indexInc.html")
	if errTemp != nil {
		fmt.Println(errTemp)
		return
	}
	tempIncomes.Execute(w, incomes)
	fmt.Println("OK")
}

func printExpenses(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("SELECT * FROM Accounting.expenses")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()
	expenses := []Expenses{}

	for rows.Next() {
		cost := Expenses{}
		errScan := rows.Scan(&cost.Id, &cost.Sum, &cost.Category, &cost.Place, &cost.Date)
		if errScan != nil {
			fmt.Println(errScan)
			continue
		}
		expenses = append(expenses, cost)
	}

	tempIncomes, errTemp := template.ParseFiles("templates/indexExp.html")
	if errTemp != nil {
		fmt.Println(errTemp)
		return
	}
	tempIncomes.Execute(w, expenses)
	fmt.Println("OK")
}

func execInsertTable() {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	_, errRes := database.Exec(
		"INSERT INTO Accounting.income (sum, type, place, date) VALUES (54545.777, 'Перевод по СБП', 'Роман Александрович Т.', $1)", timeNow)
	if errRes != nil {
		fmt.Println(errRes)
		return
	}
	_, errExpenses := database.Exec(
		"INSERT INTO Accounting.Expenses (sum, category, place, date) VALUES (54545.777, 'Перевод по СБП', 'Роман Александрович Т.', $1)", timeNow)
	if errExpenses != nil {
		fmt.Println(errExpenses)
		return
	}
	fmt.Println("OK")
}

func queryInsertTable() {
	rows, err := database.Query("SELECT sum, type FROM Accounting.income")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var sum float64
		var types string
		errScan := rows.Scan(&sum, &types)
		if errScan != nil {
			fmt.Println(errScan)
			return
		}
		fmt.Println(sum, types)
	}
}
