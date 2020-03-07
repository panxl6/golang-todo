package main

import "github.com/jinzhu/gorm"

type (
    // entity类
    todoModel struct {
        gorm.Model
        Title     string `json:"title"`
        Completed int    `json:"completed"`
    }

    // response entity
    transformedTodo struct {
        ID        uint   `json:"id"`
        Title     string `json:"title"`
        Completed bool   `json:"completed"`
    }
)