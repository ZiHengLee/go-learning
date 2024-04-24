package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/lib/pq"
)

func TestNotifyListen(t *testing.T) {
	connStr := "user=postgres dbname=postgres password=postgres*5432 host=10.14.80.53 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("LISTEN table_changes")
	if err != nil {
		log.Fatal(err)
	}

	listener := pq.NewListener(connStr, 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Fatal(err)
		}
	})

	listener.Listen("table_changes")
	go func() {
		for {
			select {
			case n := <-listener.Notify:
				var notification map[string]interface{}
				err := json.Unmarshal([]byte(n.Extra), &notification)
				if err != nil {
					log.Println("Error parsing notification:", err)
					continue
				}

				operation, ok := notification["operation"].(string)
				if !ok {
					log.Println("Invalid operation type")
					continue
				}

				switch operation {
				case "INSERT":
					fmt.Println("Received INSERT notification:")
					fmt.Println("Table:", notification["table"])
					fmt.Println("Data:", notification["data"])
				case "DELETE":
					fmt.Println("Received DELETE notification:")
					fmt.Println("Table:", notification["table"])
					fmt.Println("Data:", notification["data"])
				default:
					fmt.Println("Unsupported operation:", operation)
				}
			}
		}
	}()

	select {}
}
