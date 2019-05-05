package handlers

import (
    "github.com/jinzhu/gorm"
    "net/http"
    "strconv"
    "github.com/labstack/echo"
    "../models"
)

type H map[string]interface{}

func GetPolls(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, models.GetPolls(db))
    }
}

func UpdatePoll(db *gorm.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var poll models.Poll
        c.Bind(&poll)
        index, _ := strconv.Atoi(c.Param("index"))
        id, err := models.UpdatePoll(db, index, poll.Upvotes, poll.Downvotes)
        if err == nil {
            return c.JSON(http.StatusCreated, H{
                "affected": id,
            })
        }
        return err
    }
}
