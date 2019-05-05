package models

import (
  "errors"
  "github.com/jinzhu/gorm"
)

type Poll struct {
  gorm.Model
  Name string `gorm:"not null;unique" json:"name"`
  Topic string `gorm:"not null" json:"topic"`
  Src string `gorm:"not null" json:"src"`
  Upvotes int `gorm:"not null" json:"upvotes"`
  Downvotes int `gorm:"not null" json:"downvotes"`
}

type PollCollection struct {
  Polls []Poll `json:"items"`
}

func GetPolls(db *gorm.DB) []Poll {
  var polls []Poll
  db.Find(&polls)
  return polls
}

func UpdatePoll(db *gorm.DB, index int, upvotes int, downvotes int) (int, error) {
  var poll Poll
  db.First(&poll, index)
  if poll.ID == 0 {
    return index, errors.New("Element not found")
  }
  poll.Upvotes = upvotes
  poll.Downvotes = downvotes
  db.Save(&poll)
  return int(poll.ID), nil
}

func (poll *Poll) AfterCreate(scope *gorm.Scope) error {
	ID := int(poll.ID)
	hash := GenerateHash(ID)
	scope.DB().Model(poll).Update("Hash", hash)
	return nil
}
