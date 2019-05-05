package main

import (
  "github.com/jinzhu/gorm"
)

type Poll struct {
  gorm.Model
  Name string `gorm:"not null;unique"`
  Topic string `gorm:"not null"`
  Src string `gorm:"not null"`
  Upvotes int `gorm:"not null"`
  Downvotes int `gorm:"not null"`
}


func (poll *Poll) AfterCreate(scope *gorm.Scope) error {
	ID := int(poll.ID)
	hash := generateHash(ID)
	scope.DB().Model(poll).Update("Hash", hash)
	return nil
}
