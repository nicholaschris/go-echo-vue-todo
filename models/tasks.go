// models/tasks.go

package models

import (
  "database/sql"

  _ "github.com/mattn/go-sqlite3"
)

type Task struct {
  ID int `json.id`
  Name string `json:"name"`
}

type TaskCollection struct {
  Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
  sql := "SELECT * FROM tasks"
  rows, err := db.Query(sql)
  // Exit if the SQL doesn't work for some reason
  if err != nil {
    panic(err)
  }
  // make sure to cleanup when the program exists
  defer rows.Close()

  result := TaskCollection{}
  for rows.Next() {
    task := Task{}
    err2 := rows.Scan(&task.ID, &task.Name)
    // Exit if we get an error
    if err2 != nil {
      panic(err2)
    }
    result.Tasks = append(result.Tasks, task)
    }
    return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
  sql := "INSERT INTO tasks(name) VALUES(?)"
  // Create a prepared sql statement
  stmt, err := db.Prepare(sql)
  // Exit if we get an error
  if err != nil {
    panic(err)
  }
  // make sure to cleanup after the program exits
  defer stmt.Close()

  // Replace the '?' in our prepared statement with 'name'
  result, err2 := stmt.Exec(name)
  // Exit if we get an error
  if err2 != nil {
    panic(err2)
  }

  return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
  sql := "DELETE FROM tasks WHERE id = ?"

  // Create prepared SQL statement
  stmt, err := db.Prepare(sql)
  // Exit if we get an error
  if err != nil {
    panic(err)
  }

  // Replace the '?' in the prepared statement with 'id'
  result, err2 := stmt.Exec(id)
  // Exit if we get an error
  if err2 != nil {
    panic(err2)
  }

  return result.RowsAffected()
}
