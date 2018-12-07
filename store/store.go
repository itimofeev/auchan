package store

import (
	"bitbucket.org/Axxonsoft/axxoncloudgo/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/itimofeev/auchan/util"
	"time"
)

func NewStore(connectURL string) *Store {
	opts, err := pg.ParseURL(connectURL)
	util.CheckErr(err, "pg.ParseURL")
	db := pg.Connect(opts)
	util.CheckErr(createSchema(db))

	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}
		util.Log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	return &Store{}
}

type Store struct {
	db pg.DB
}

func createSchema(db *pg.DB) error {
	err := db.CreateTable((*model.FileInfo)(nil), &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}
