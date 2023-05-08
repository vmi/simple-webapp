package model

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
    var err error
    db, err = gorm.Open("sqlite3", "db/sample.db")
    if err != nil {
        panic(fmt.Sprintf("failed to connect database: %v", err))
    }
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Todo{})
}
