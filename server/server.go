package server

import (
	"database/sql"
	"fmt"
)

var Database *sql.DB

type WorkWithDB interface {
	ExecInsertTable() error
	QuerySelectTable() error
	QueryDeleteRow(id int) error
}

type Income struct {
	Id    int
	Sum   float64
	Type  string
	Place string
	Date  string
}

type Expenses struct {
	Id       int
	Sum      float64
	Category string
	Place    string
	Date     string
}

func (inc *Income) ExecInsertTable() error {
	_, errRes := Database.Exec(
		"INSERT INTO Accounting.income (sum, type, place, date) VALUES ($1, $2, $3, $4)",
		inc.Sum, inc.Type, inc.Place, inc.Date)
	if errRes != nil {
		return errRes
	}
	fmt.Println("Строка успешно вставлена!")
	return nil
}

func (exp *Expenses) ExecInsertTable() error {
	_, errExpenses := Database.Exec(
		"INSERT INTO Accounting.expenses (sum, category, place, date) VALUES ($1, $2, $3, $4)",
		exp.Sum, exp.Category, exp.Place, exp.Date)
	if errExpenses != nil {
		return errExpenses
	}
	fmt.Println("Строка успешно вставлена!")
	return nil
}

func (inc *Income) QuerySelectTable() error {
	rows, err := Database.Query("SELECT sum, type FROM Accounting.income")
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var sum float64
		var types string
		errScan := rows.Scan(&sum, &types)
		if errScan != nil {
			return errScan
		}
		fmt.Println(sum, types)
	}
	fmt.Println("Строки успешно выведены!")
	return nil
}

func (exp *Expenses) QuerySelectTable() error {
	rows, err := Database.Query("SELECT sum, category FROM Accounting.expenses")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var sum float64
		var types string
		errScan := rows.Scan(&sum, &types)
		if errScan != nil {
			fmt.Println(errScan)
			return errScan
		}
		fmt.Println(sum, types)
	}
	fmt.Println("Строки успешно выведены!")
	return nil
}

func (inc *Income) QueryDeleteRow(id int) error {
	_, err := Database.Query("Delete from accounting.income where id=$1", id)
	if err != nil {
		return err
	}
	fmt.Println("Строка успешно удалена!")
	return nil
}

func (exp *Expenses) QueryDeleteRow(id int) error {
	_, err := Database.Query("Delete from accounting.expenses where id=$1", id)
	if err != nil {
		return err
	}
	fmt.Println("Строка успешно удалена!")
	return nil
}
