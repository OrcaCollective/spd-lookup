package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

type officer struct {
	BadgeNumber     string `json:"badge_number,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	MiddleName      string `json:"middle_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Title           string `json:"title,omitempty"`
	Unit            string `json:"unit,omitempty"`
	UnitDescription string `json:"unit_description,omitempty"`
}

type databaseInterface interface {
	getOfficerByBadge(badge string) (*officer, error)
	searchOfficerByName(firstName, lastName string) ([]*officer, error)
}

type dbClient struct {
	pool *pgxpool.Pool
}

func newDBClient() *dbClient {
	var pool *pgxpool.Pool
	if os.Getenv("DATABASE_URL") != "" {
		pgxConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Panic(fmt.Sprintf("Unable to create postgres connection config: %v", err))
		}
		pool, err = pgxpool.ConnectConfig(context.Background(), pgxConfig)
		if err != nil {
			log.Panic(fmt.Sprintf("Unable to create db connection: %v", err))
		}
	} else {
		// local set up
		username := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")
		pgxConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgresql://%s", host))
		if err != nil {
			log.Panic(fmt.Sprintf("Unable to create postgres connection config: %v", err))
		}

		pgxConfig.ConnConfig.Port = 5432
		pgxConfig.ConnConfig.Database = dbName
		pgxConfig.ConnConfig.User = username
		pgxConfig.ConnConfig.Password = password

		pool, err = pgxpool.ConnectConfig(context.Background(), pgxConfig)
		if err != nil {
			log.Panic(fmt.Sprintf("Unable to create db connection: %v", err))
		}
	}

	return &dbClient{pool: pool}
}

func (db *dbClient) getOfficerByBadge(badge string) (*officer, error) {
	ofc := officer{}
	err := db.pool.QueryRow(context.Background(),
		`
			SELECT
				badge_number,
				first_name,
				middle_name,
				last_name,
				title,
				unit,
				unit_description
			FROM get_officer_by_badge_p (badge_number := $1);
		`, badge,
	).Scan(
		&ofc.BadgeNumber,
		&ofc.FirstName,
		&ofc.MiddleName,
		&ofc.LastName,
		&ofc.Title,
		&ofc.Unit,
		&ofc.UnitDescription,
	)

	return &ofc, err
}

func (db *dbClient) searchOfficerByName(firstName, lastName string) ([]*officer, error) {
	rows, err := db.pool.Query(context.Background(), `
	SELECT
		badge_number,
		first_name,
		middle_name,
		last_name,
		title,
		unit,
		unit_description
	FROM search_officer_by_name_p(first_name := $1, last_name := $2);`, firstName, lastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	officers := []*officer{}
	for rows.Next() {
		ofc := officer{}
		err := rows.Scan(
			&ofc.BadgeNumber,
			&ofc.FirstName,
			&ofc.MiddleName,
			&ofc.LastName,
			&ofc.Title,
			&ofc.Unit,
			&ofc.UnitDescription,
		)

		if err != nil {
			return nil, err
		}
		officers = append(officers, &ofc)
	}
	return officers, nil
}
