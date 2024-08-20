package common

import "gorm.io/gorm"

var db *gorm.DB

func SetDB(database *gorm.DB) { // TODO remove this
	db = database
}
