package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	DB DbItem `yaml:"db"`
}

type DbItem struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

func main() {
	testCfg := DbItem{
		Driver: "testing",
		Name:   "testingname",
	}

	var dbcfg DbConfig

	yf, err := ioutil.ReadFile("./basic/postgres/config.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(yf, &dbcfg); err != nil {
		panic(err)
	}

	fmt.Println("yaml config: ", dbcfg)

	sql, err := sql.Open(dbcfg.DB.Driver, dbcfg.DB.DSN)
	if err != nil {
		panic(err)
	}

	// ctx := context.Background()

	if _, err := sql.Exec(fmt.Sprintf(`CREATE MATERIALIZED VIEW %s`, pq.QuoteIdentifier("haloview"))); err != nil {
		fmt.Println("error exec: ", err)
	}

	// ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	// insertEvery1Seconds(ctx, sql)
}

func insertEvery1Seconds(ctx context.Context, sql *sql.DB) {
	fmt.Println("fire this func")
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker done")
			return
		case t := <-ticker.C:
			fmt.Println("ticket: ", t)
			fmt.Println("random: ", rand.Intn(1000000))
			tx, err := sql.BeginTx(ctx, nil)
			if err != nil {
				panic(err)
			}

			tx.Query("create table testingaja")

			tx.Commit()
		}
	}
}
