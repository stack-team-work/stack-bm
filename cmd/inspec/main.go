package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/stack_api?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	var result map[string]interface{}
	var results []map[string]interface{}
	
	// Show columns
	db.Raw("SHOW COLUMNS FROM game_cp").Scan(&results)
	fmt.Println("=== game_cp columns ===")
	for _, r := range results {
		fmt.Printf("  %v\n", r)
	}
	
	// Test create
	fmt.Println("\n=== Test create ===")
	r := db.Exec("INSERT INTO game_cp (name) VALUES ('test')")
	fmt.Printf("  RowsAffected: %d, Error: %v\n", r.RowsAffected, r.Error)
	
	// Test select
	db.Raw("SELECT * FROM game_cp LIMIT 1").Scan(&result)
	fmt.Printf("  Result: %v\n", result)
	
	// Cleanup
	db.Exec("DELETE FROM game_cp WHERE name = 'test'")
}
