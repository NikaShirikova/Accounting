package server

import (
	"database/sql"
	"fmt"
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

var Database *sql.DB

func ExecInsertTable() {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	_, errRes := Database.Exec(
		"INSERT INTO Accounting.income (sum, type, place, date) VALUES (54545.777, 'Перевод по СБП', 'Роман Александрович Т.', $1)", timeNow)
	if errRes != nil {
		fmt.Println(errRes)
		return
	}
	_, errExpenses := Database.Exec(
		"INSERT INTO Accounting.Expenses (sum, category, place, date) VALUES (54545.777, 'Перевод по СБП', 'Роман Александрович Т.', $1)", timeNow)
	if errExpenses != nil {
		fmt.Println(errExpenses)
		return
	}
}

func QueryInsertTable() {
	rows, err := Database.Query("SELECT sum, type FROM Accounting.income")
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
