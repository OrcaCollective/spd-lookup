package main

import (
	"context"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

type tacomaOfficer struct {
	FirstName  string       `json:"first_name,omitempty"`
	LastName   string       `json:"last_name,omitempty"`
	Title      string       `json:"title,omitempty"`
	Department string       `json:"department,omitempty"`
	Salary     nulls.String `json:"salary,omitempty"`
}

func (db *dbClient) tacomaSearchOfficerByName(firstName, lastName string) ([]*tacomaOfficer, error) {
	rows, err := db.pool.Query(context.Background(), `
	SELECT
		first_name,
		last_name,
		title,
		department,
		salary
	FROM tacoma_search_officer_by_name_p(first_name := $1, last_name := $2);`, firstName, lastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return marshalTacomaOfficerRows(rows)
}

func (db *dbClient) tacomaFuzzySearchByName(name string) ([]*tacomaOfficer, error) {
	rows, err := db.pool.Query(context.Background(), `
	SELECT
		first_name,
		last_name,
		title,
		department,
		salary
	FROM tacoma_fuzzy_search_officer_by_name_p(full_name := $1);`, name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return marshalTacomaOfficerRows(rows)
}

func (db *dbClient) tacomaFuzzySearchByFirstName(firstName string) ([]*tacomaOfficer, error) {
	rows, err := db.pool.Query(context.Background(), `
	SELECT
		first_name,
		last_name,
		title,
		department,
		salary
	FROM tacoma_fuzzy_search_officer_by_first_name_p(first_name := $1);`, firstName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return marshalTacomaOfficerRows(rows)
}

func (db *dbClient) tacomaFuzzySearchByLastName(lastName string) ([]*tacomaOfficer, error) {
	rows, err := db.pool.Query(context.Background(), `
	SELECT
		first_name,
		last_name,
		title,
		department,
		salary
	FROM tacoma_fuzzy_search_officer_by_last_name_p(last_name := $1);`, lastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return marshalTacomaOfficerRows(rows)
}

func marshalTacomaOfficerRows(rows pgx.Rows) ([]*tacomaOfficer, error) {
	officers := []*tacomaOfficer{}
	for rows.Next() {
		ofc := tacomaOfficer{}
		err := rows.Scan(
			&ofc.FirstName,
			&ofc.LastName,
			&ofc.Title,
			&ofc.Department,
			&ofc.Salary,
		)

		if err != nil {
			return nil, err
		}
		officers = append(officers, &ofc)
	}
	return officers, nil
}
