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

//const mysqlUrl = "tyFeng:J0]nt4D_3-NbO>8|GgryV-ry.?G{@tcp(arxivinfo.cvheva0xliby.us-east-1.rds.amazonaws.com:3306)/arxivInfo?parseTime=true"
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

	err = sqlc.PutAllTopics()
	if err != nil {
		log.Fatal(err)
		return
	}

	// var tick <-chan time.Time = time.Tick(6 * time.Hour)
	// //loop to ...
	// for range tick {
	// 	go func() {
	// 		CreateLogs()
	// 		err = sqlc.PutAllTopics()
	// 		if err != nil {
	// 			log.Fatal(err)
	// 			return
	// 		}
	// 	}()

	// }

}
