package repositories

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	host     = "ec2-52-202-146-43.compute-1.amazonaws.com"
	database = "d7s7hkambeivdc"
	username = "pzosaqcfhlswjt"
	port     = 5432
	password = "8df6747ba4aea4396e5ffb13cd0614c51caf7082f856dd1f09b20ae5c8e44eb1"
)

func ReturnDangerLocations() *sql.DB {

	db, err := sql.Open("postgres", "postgres://pzosaqcfhlswjt:8df6747ba4aea4396e5ffb13cd0614c51caf7082f856dd1f09b20ae5c8e44eb1@ec2-52-202-146-43.compute-1.amazonaws.com:5432/d7s7hkambeivdc")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
