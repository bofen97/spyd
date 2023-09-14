package main

import (
	"log"
	"os"
	"time"

	SQLConn "github.com/bofen97/sqlc"
)

func CreateLogs() {
	logtime := time.Now().Format("3:04:5")
	logDir, err := os.Create("/tmp/topic_logs_" + logtime)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.SetOutput(logDir)

}

//const mysqlUrl = "root:@(127.0.0.1:3306)/arxivInfo?parseTime=true"

func main() {

	mysqlUrl := os.Getenv("sqlurl")
	if mysqlUrl == "" {
		log.Fatal("sqlurl is none")
		return
	}
	var sqlc = new(SQLConn.SQLConn)
	err := sqlc.Connect(mysqlUrl)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err = sqlc.CreateTable(); err != nil {
		log.Fatal(err)
		return
	}
	//start
	log.Print("task start \n")
	var tick <-chan time.Time = time.Tick(7 * time.Hour)
	//loop to ...
	for range tick {

		log.Print("run task download arxiv .. \n")

		go func() {
			//CreateLogs()
			err = sqlc.PutAllTopics()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

}
