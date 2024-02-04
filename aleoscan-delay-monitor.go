package main

import (
	"aleoscan-delay-monitor/prometh"
	"aleoscan-delay-monitor/sqlexec"
	"flag"
	"log"
	"time"
)

var pgDSN = flag.String("pgdsn", "", "Postgres connection string,example:postgres://postgres:123456@localhost:5432/postgres?sslmode=disable")
var Interval = flag.Int("interval", 1, "Interval in minutes to run the monitor")
var pushAddr = flag.String("push-address", "", "Address of the Pushgateway to send metrics,example:http://127.0.0.1:9091")

func main() {
	flag.Parse()

	if *pushAddr == "" || *pgDSN == "" {
		log.Fatalln("push-address and pgdsn must be provided")
	}

	db, err := sqlexec.InitDB(*pgDSN)
	if err != nil {
		log.Fatalln(err)
	}

	for {

		height, ts, err := sqlexec.GetHeightAndDelay(db)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Duration(*Interval) * time.Minute)
		}
		delay:= time.Now().Unix() - ts
		log.Printf("height:%d, ts:%d,delay:%d", height, ts, delay)
		prometh.Push(*pushAddr, "aleoscan_cur_height", height)
		prometh.Push(*pushAddr, "aleoscan_cur_timestamp_delay_second",delay)

		time.Sleep(time.Duration(*Interval) * time.Minute)
	}
}
