package main

import (
	"AccountingService/server"
	"AccountingService/server/api"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)

const CountArray = 5

func main() {
	conn := "user=root password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Println(err)
		return
	}

	server.Database = db
	defer db.Close()

	//var incomes [CountArray]server.Income
	//var expenses [CountArray]server.Expenses

	//fmt.Println("Доходы")
	//AddRowsArrayIncome(incomes)
	//AddRowsArray(expenses, CountArray)
	//timeNow := time.Now().Format("2006-01-02 15:04:05")
	//income := &server.Income{0, 434, "ЗП", "Криптософт", timeNow}
	//editDB(income)
	//
	//expenses := &server.Expenses{0, 434, "Продукты", "Магнит", timeNow}
	//editDB(expenses)

	go RunServer(":8081")
	RunServer(":8080")
}

func RunServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/income", api.PrintIncomes)
	mux.HandleFunc("/expenses", api.PrintExpenses)
	mux.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Addr: ", addr, "URL: ", r.URL.String())
		})

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	fmt.Println("Start server at", addr)
	server.ListenAndServe()
}

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

func AddRowsArrayIncome(inc [CountArray]server.Income) {
	for i := 0; i < CountArray; i++ {
		randSum := 0 + rand.Intn(100000)
		inc[i] = server.Income{
			i, float64(randSum), "Перевод", "Ширикова Г.В.", time.Now().Format("2006-01-02 15:04:05")}
		go PrintIncome(inc[i])
	}
}

func PrintIncome(inc server.Income) {
	fmt.Printf("ID: %d Сумма: %g, Тип: %s, От кого: %s, Дата и время: %s \n",
		inc.Id, inc.Sum, inc.Type, inc.Place, inc.Date)
	runtime.Gosched()
}
