package main

import "bestlease.deals/db"

func main() {
	conn := db.DBConnect()
	print(conn)
}
