package main

import (
	"time"

	"github.com/hiroki-Fukumoto/matching-app/api/route"
)

const location = "Asia/Tokyo"

// @title Matching app
// @version 1.0
// @description  Matching app
// @host localhost:8080
func main() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	r := route.SetupRouter()

	r.Run(":8080")
}
