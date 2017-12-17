package handlers

import (
  "database/sql"
  "net/http"
  "strconv"

  "go-echo-vue/models"

  "github.com/labstack/echo"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    return c.JSON(http.StatusOK, models.GetTasks(db))
  }
}

func PutTask(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    // Instantiate a new task
    var task models.Task
    // Map incoming JSON body to new Task
    c.Bind(&task)
    // Add a task to our new model
    id, err := models.PutTask(db, task.Name)
    // return a JSON response if succesful
    if err == nil {
      return c.JSON(http.StatusCreated, H{
        "created": id,
      })
    } else {
      return err
    }

  }
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    // Use model to delete a task
    _, err := models.DeleteTask(db, id)
    // Return a JSON response on success
    if err == nil {
      return c.JSON(http.StatusOK, H{
        "deleted": id,
      })
    } else {
      return err
    }
  }
}
