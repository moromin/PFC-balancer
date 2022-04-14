package db

import (
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/moromin/PFC-balancer/services/food/models"
)

func loadFoodData(db *sql.DB) {
	if checkFoodData(db) {
		return
	}

	f, err := os.Open("./db/foods.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	copyDataCSVToDB(db, r)
}

func checkFoodData(db *sql.DB) bool {
	const q = `SELECT 1 FROM foods LIMIT 1`
	var res int
	if err := db.QueryRow(q).Scan(&res); err != nil {
		return false
	}
	return true
}

const createFood = `
INSERT INTO foods (
	name,
	protein,
	fat,
	carbohydrate,
	category
) VALUES (
	$1, $2, $3, $4, $5
)
`

func copyDataCSVToDB(db *sql.DB, r *csv.Reader) {
	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		food := convertRecordToFood(record)
		_, err = db.Exec(createFood, food.Name, food.Protein, food.Fat, food.Carbohydrate, food.Category)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

// data <-> record
// 0: name :1
// 1: protein :2
// 2: fat :3
// 3: carbohydrate :4
// 4: category :0
func convertRecordToFood(record []string) (food models.Food) {
	var err error

	food.Category, err = strconv.ParseInt(record[0], 10, 64)
	if err != nil {
		food.Category = 0
	}
	food.Name = record[1]
	food.Protein, err = strconv.ParseFloat(record[2], 64)
	if err != nil {
		food.Protein = 0
	}
	food.Fat, err = strconv.ParseFloat(record[3], 64)
	if err != nil {
		food.Fat = 0
	}
	food.Carbohydrate, err = strconv.ParseFloat(record[4], 64)
	if err != nil {
		food.Carbohydrate = 0
	}

	return
}
