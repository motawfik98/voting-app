package main

import (
  "net/http"
  "github.com/labstack/echo"
)

func showPolls(c echo.Context) error {
  return c.JSON(http.StatusOK, "GET Polls")
}

func addPoll(c echo.Context) error {
  return c.JSON(http.StatusOK, "PUT Poll")
}

func updatePoll(c echo.Context) error {
  return c.JSON(http.StatusOK, "UPDATE POLL " + c.Param("id"))
}
