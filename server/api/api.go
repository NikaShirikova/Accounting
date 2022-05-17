package api

import (
	"AccountingService/server"
	"fmt"
	"html/template"
	"net/http"
)

//Содержит маршруты

func PrintIncomes(w http.ResponseWriter, r *http.Request) {
	rows, err := server.Database.Query("SELECT * FROM Accounting.income")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()
	incomes := []server.Income{}

	for rows.Next() {
		income := server.Income{}
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

func PrintExpenses(w http.ResponseWriter, r *http.Request) {
	rows, err := server.Database.Query("SELECT * FROM Accounting.expenses")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer rows.Close()
	expenses := []server.Expenses{}

	for rows.Next() {
		cost := server.Expenses{}
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
