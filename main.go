package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 1. เชื่อมต่อฐานข้อมูล (ดึงฟังก์ชันจาก database.go)
	_ = godotenv.Load()
	fmt.Println("เช็คค่า DB_URL:", os.Getenv("DB_URL"))

	initDB()
	defer db.Close()

	// 2. ตั้งค่า Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// 3. เริ่มทำงาน WebSocket (ดึงฟังก์ชันจาก websocket.go)
	http.HandleFunc("/", handleConnections)

	fmt.Println("🚀 WebSocket Server Started on port", port, "...")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
