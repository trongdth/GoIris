package main

import (
	"database/sql"
	"github.com/kataras/iris"
	_ "github.com/go-sql-driver/mysql"
)

const (
	USER_ID = "user_id"
	ACTION  = "action"
	MYSQL_USER = MYSQL_USERNAME
	MYSQL_PWD = MYSQL_MYPWD
)


func main() {
	db, err := sql.Open("mysql", MYSQL_USER + ":" + MYSQL_PWD + "@/autonomous_log")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	iris.Post("/log", func(ctx *iris.Context) {
		user_id := ctx.FormValue(USER_ID)
		action := ctx.FormValue(ACTION)

		ctx.Write("user_id %s, action %s", user_id, action)

		// Prepare statement for inserting data
		stmtIns, err := db.Prepare("INSERT INTO log(user_id, action) VALUES(?, ?)") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

		_, err = stmtIns.Exec(user_id, action)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	})
	iris.Listen(":8080")
}