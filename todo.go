// todo.go
package main

import (
  	// "net/http"
    "database/sql"
    "go-echo-vue/handlers"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    _ "github.com/mattn/go-sqlite3"
    // "github.com/labstack/echo/engine/standard"
)

// todo.go
func main() {

  db := initDB("storage.db")
  migrate(db)

  e := echo.New()

  // CORS default
  //Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
  // e.Use(middleware.CORS())

  // CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://0.0.0.0:8000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

  // e.File("/", "public/index.html")
  e.GET("/tasks", handlers.GetTasks(db))
  e.PUT("/tasks", handlers.PutTask(db))
  e.DELETE("/tasks/:id", handlers.DeleteTask(db))

  e.Logger.Fatal(e.Start(":1323"))
}

func initDB(filepath string) *sql.DB {
  db, err := sql.Open("sqlite3", filepath)

  if err != nil {
    panic(err)
  }

  if db == nil {
    panic("db nil")
  }

  return db
}

func migrate(db *sql.DB) {
  sql := `
    CREATE TABLE IF NOT EXISTS tasks(
      id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
      name VARCHAR NOT NULL
    )
  `

  _, err := db.Exec(sql)
  if err != nil {
    panic(err)
  }
}
