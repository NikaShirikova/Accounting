package api

import (
	"AccountingService/server"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
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

func AddIncomes(w http.ResponseWriter, r *http.Request) {
	sum := r.URL.Query().Get("sum")
	mType := r.URL.Query().Get("type")
	place := r.URL.Query().Get("place")
	if sum == "" || mType == "" || place == "" {
		fmt.Fprintln(w, "Не хватает параметров для добавления строки в таблицу income")
		return
	}
	floatSum, _ := strconv.Atoi(sum)
	inc := server.Income{0, float64(floatSum), mType, place, time.Now().Format("2006-01-02 15:04:05")}
	inc.ExecInsertTable()
	fmt.Println("OK ADD")
	fmt.Fprintln(w, "Строка успешно вставлена")
}

func AddExpenses(w http.ResponseWriter, r *http.Request) {
	sum := r.FormValue("sum")
	category := r.FormValue("category")
	place := r.FormValue("place")
	if sum == "" || category == "" || place == "" {
		fmt.Fprintln(w, "Не хватает параметров для добавления строки в таблицу expenses")
		return
	}
	floatSum, _ := strconv.Atoi(sum)
	exp := server.Expenses{0, float64(floatSum), category, place, time.Now().Format("2006-01-02 15:04:05")}
	exp.ExecInsertTable()
	fmt.Println("OK ADD")
	fmt.Fprintln(w, "Строка успешно вставлена")
}
