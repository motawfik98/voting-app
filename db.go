package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mssql"
  "github.com/speps/go-hashids"

)
var db *gorm.DB
func initDB() {
  var err error
  db, err = gorm.Open("mssql", "sqlserver://remote:mohamed@localhost:1433?database=poll_voting")
  if err == nil {
    db.AutoMigrate(&Poll{})
  }
}

func generateHash(ID int) string {
  hd := hashids.NewData()
  hd.Salt = "xOBtdmJZxRcz^jkkyHfkrkT1*02bJUn+YQts0*xCeka%cGHCN1fjaC*faFtY"
  hd.MinLength = 8
  h, _ := hashids.NewWithData(hd)
  e, _ := h.Encode([]int{ID})
  return e
}
