package main

import (
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

func main() {
  initDB()

  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Define HTTP routes
  e.GET("/polls", showPolls)
  e.PUT("/polls", addPoll)
  e.PUT("/polls/:id", updatePoll)

  e.Logger.Fatal(e.Start(":9000"))
}
