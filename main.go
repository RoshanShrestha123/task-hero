package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	"github.com/RoshanShrestha123/task-hero/helper"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`create table if not exists todos(
	"id" integer not null primary key autoincrement,
	"task" text, 
	"isCompleted" integer,
	"createdAt" date, 
	"updatedAt" date
	)`)

	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("\nTask Hero at your service: 1(add) 2(list) 3(remove) 4(exit)")
		var option int
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			fmt.Println("Whats your plan for today?")
			reader := bufio.NewReader(os.Stdin)
			task, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			helper.Add(db, task)
			fmt.Printf("Task has been added, dnt forget to complete it by today!")

		case 2:

			fmt.Println("Id | Task | is Completed | created at | updated at")
			fmt.Println("___________________________________________________________")

			tasks := helper.List(db)

			for _, val := range tasks {
				fmt.Printf("%d | %s | %v | %s | %s\n", val.Id, val.Task, val.IsCompleted == 1, val.CreatedAt.Format("2006-01-02"), val.UpdatedAt.Format("2006-01-02"))
			}

		case 3:
			fmt.Println("Provide the id to delete it from database?")
			var id int
			fmt.Scanf("%d", &id)

			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			helper.Delete(db, id)
			tasks := helper.List(db)

			for _, val := range tasks {
				fmt.Printf("%d | %s | %v | %s | %s\n", val.Id, val.Task, val.IsCompleted == 1, val.CreatedAt.Format("2006-01-02"), val.UpdatedAt.Format("2006-01-02"))
			}

		case 4:
			fmt.Println("Bye..")
			os.Exit(0)

		default:
			fmt.Println("Invalid input!")

		}

	}
}
