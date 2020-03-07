package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
    db *gorm.DB
    sqlConnection = "root:password@(127.0.0.1:10039)/demo?charset=utf8&parseTime=True&loc=Local"
)

func init() {
    //打开数据库连接
    var err error
    db, err = gorm.Open("mysql", sqlConnection)
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&todoModel{})
}


