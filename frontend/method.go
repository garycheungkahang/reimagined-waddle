package main

import(
	// "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	// "github.com/gin-gonic/gin"
)

func create(name string) {

   db, err := sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/")
   if err != nil {
       panic(err)
   }
   defer db.Close()

   _,err = db.Exec("CREATE DATABASE "+name)
   if err != nil {
       panic(err)
   }

   _,err = db.Exec("USE "+name)
   if err != nil {
       panic(err)
   }

   _,err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
   if err != nil {
       panic(err)
   }}