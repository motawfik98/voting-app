package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "./handlers"
)

func main() {
  initDB()

  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Define HTTP routes
  e.File("/", "public/index.html")
  e.GET("/polls", handlers.GetPolls(db))
  e.PUT("/poll/:index", handlers.UpdatePoll(db))

  e.Logger.Fatal(e.Start(":9000"))
}
