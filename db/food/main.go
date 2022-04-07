package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/moromin/PFC-balancer/db/food/config"
	"github.com/moromin/PFC-balancer/db/food/db"
	"github.com/moromin/PFC-balancer/services/food/models"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	c := config.LoadConfig()

	db := db.Init(c.DBUrl)

	f, err := os.Open("./foods.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	copyDataCSVToDB(db, r, ctx)
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

func copyDataCSVToDB(db *sql.DB, r *csv.Reader, ctx context.Context) {
	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.PrepareContext(ctx, createFood)
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

		food := convertDataToFood(record)
		_, err = stmt.ExecContext(ctx, food.Name, food.Protein, food.Fat, food.Carbohydrate, food.Category)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
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
func convertDataToFood(record []string) (food models.Food) {
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
