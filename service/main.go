package main

import (
	"fmt"

	"log"

	"strconv"

	"database/sql"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "unicorn"
	password = "rainbow"
	dbname   = "alltheway"
)

var db = initDb()

func initDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	log.SetPrefix("all-the-way:")
	router := gin.Default()
	router.POST("/run/:runner_id", postRun)
	router.GET("/runs/:runner_id", getRuns)
	router.GET("/run/:run_id", getRun)
	log.Output(1, "Starting servier in port 8080")
	defer db.Close()
	router.Run("localhost:8080")
}

func getRuns(c *gin.Context) {
	runner_id, err := strconv.Atoi(c.Param("runner_id"))
	if err != nil {
		print("error here")
		log.Fatal(err)
	}
	log.Output(1, fmt.Sprintf("Fetching runs for runner %d", runner_id))
	var run_id int
	var distance float32

	rows, err := db.Query("select run_id, distance from run where runner_id = $1;", runner_id)
	if err != nil {
		print("error here")
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&run_id, &distance)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(run_id, distance)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}

func getRun(c *gin.Context) {
	print("nil")
}
