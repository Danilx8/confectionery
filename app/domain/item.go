package domain

import (
	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	Name string `gorm:"primaryKey;column:name"`
	Size string `gorm:"column:size"`
}

type ItemRepository interface {
	Fetch() ([]Item, error)
}
