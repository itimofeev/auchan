package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/itimofeev/auchan/util"
	"os"
	"strconv"
)

type Prod struct {
	ID           int64
	CategoryName string
	EAM          int64
	Title        string
	Price        float64
	Picture      string
	Link         string
}

func main() {
	//r := csv.NewReader(strings.NewReader(in))

	f, err := os.Create("cmd.sql")
	util.CheckErr(err, "os.create")
	w := bufio.NewWriter(f)

	csvFile, _ := os.Open("/Users/ilyatimofee/Downloads/itemsAuchan - Лист1.csv")
	r := csv.NewReader(bufio.NewReader(csvFile))

	records, err := r.ReadAll()
	util.CheckErr(err, "read all")

	for i := 1; i < len(records); i++ {
		r := records[i]
		id, err := strconv.Atoi(r[0])
		if err != nil {
			continue
		}
		eam, err := strconv.Atoi(r[2])
		if err != nil {
			continue
		}

		price, err := strconv.ParseFloat(r[6], 64)
		if err != nil {
			continue
		}

		p := Prod{
			ID:           int64(id),
			CategoryName: r[1],
			EAM:          int64(eam),
			Title:        r[4],
			Price:        price,
			Picture:      r[8],
			Link:         r[11],
		}

		_, err = w.WriteString(fmt.Sprintf("INSERT INTO product("+
			"name, auchan_id, image_url, category_name, link, price) VALUE ("+
			"%s, %d, %s, %s, %s, %f)/n", p.Title, p.ID, "https://images3.alphacoders.com/258/thumb-1920-258059.jpg", p.CategoryName, p.Link, p.Price))

		util.CheckErr(err, "write")
	}

	w.Flush()
}
