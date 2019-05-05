package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mssql"
  "./models"

)
var db *gorm.DB
func initDB() {
  var err error
  db, err = gorm.Open("mssql", "sqlserver://remote:mohamed@localhost:1433?database=poll_voting")
  if err == nil {
    db.AutoMigrate(&models.Poll{})
  }
}
