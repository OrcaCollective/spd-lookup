package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	officers := loadOfficersFromCSV(os.Getenv("SEED_FILE"))
	fmt.Printf("loaded %d officers\n", len(officers))

	db := newDBClient()
	err := db.clearOfficerTable()
	if err != nil {
		log.Panic("error clearing db", err)
	}
	fmt.Println("officers db cleared")

	c := make(chan error, len(officers))
	for _, ofc := range officers {
		go func(ofc *officer) {
			c <- db.insertOfficer(ofc)
		}(ofc)
	}

	successfulInserts := len(officers)
	for i := 0; i < cap(c); i++ {
		insertErr := <-c
		if insertErr != nil {
			successfulInserts--
			fmt.Printf("error inserting officer %s\n", insertErr)
		}
	}

	fmt.Printf("successfully inserted %d rows into the officers table\n", successfulInserts)
}

type officer struct {
	badgeNumber     string
	firstName       string
	middleName      string
	lastName        string
	title           string
	unit            string
	unitDescription string
}

func loadOfficersFromCSV(inputFile string) []*officer {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	officers := []*officer{}
	scanner.Scan()
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")

		if len(parts) == 9 {
			parts[1] += parts[2]
			for i := 2; i < 8; i++ {
				parts[i] = parts[i+1]
			}
		}

		officers = append(officers, &officer{
			badgeNumber:     parts[0],
			title:           parts[1],
			unit:            parts[2],
			unitDescription: parts[3],
			lastName:        parts[4],
			firstName:       parts[5],
			middleName:      parts[6],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return officers
}

type dbClient struct {
	pool *pgxpool.Pool
}

func newDBClient() *dbClient {
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

	pool, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Panic(fmt.Sprintf("Unable to create db connection: %v", err))
	}

	return &dbClient{
		pool: pool,
	}
}

func (db *dbClient) clearOfficerTable() error {
	_, err := db.pool.Exec(context.Background(), "delete from officers")
	return err
}

func (db *dbClient) insertOfficer(ofc *officer) error {
	_, err := db.pool.Exec(context.Background(), `
	insert into officers (badge_number, first_name, middle_name, last_name, title, unit, unit_description)
	values ($1, $2, $3, $4, $5, $6, $7)
	`, ofc.badgeNumber, ofc.firstName, ofc.middleName, ofc.lastName, ofc.title, ofc.unit, ofc.unitDescription)

	return err
}
