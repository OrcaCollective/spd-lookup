package data

import (
	"context"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// TacomaOfficer is the object model for Tacoma PD officers
type TacomaOfficer struct {
	FirstName  string       `json:"first_name,omitempty"`
	LastName   string       `json:"last_name,omitempty"`
	Title      string       `json:"title,omitempty"`
	Department string       `json:"department,omitempty"`
	Salary     nulls.String `json:"salary,omitempty"`
}

// TacomaOfficerMetadata retreives metadata describing the TacomaOfficer struct
func (c *Client) TacomaOfficerMetadata() []map[string]string {
	return []map[string]string{
		{
			"FieldName": "first_name",
			"Label":     "First Name",
		},
		{
			"FieldName": "last_name",
			"Label":     "Last Name",
		},
		{
			"FieldName": "title",
			"Label":     "Title",
		},
		{
			"FieldName": "department",
			"Label":     "Department",
		},
		{
			"FieldName": "salary",
			"Label":     "Salary",
		},
	}
}

// TacomaSearchOfficerByName invokes tacoma_search_officer_by_name_p
func (c *Client) TacomaSearchOfficerByName(firstName, lastName string) ([]*TacomaOfficer, error) {
	rows, err := c.pool.Query(context.Background(), `
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

// TacomaFuzzySearchByName invokes tacoma_fuzzy_search_officer_by_name_p
func (c *Client) TacomaFuzzySearchByName(name string) ([]*TacomaOfficer, error) {
	rows, err := c.pool.Query(context.Background(), `
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

// TacomaFuzzySearchByFirstName invokes tacoma_fuzzy_search_officer_by_first_name_p
func (c *Client) TacomaFuzzySearchByFirstName(firstName string) ([]*TacomaOfficer, error) {
	rows, err := c.pool.Query(context.Background(), `
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

// TacomaFuzzySearchByLastName invokes tacoma_fuzzy_search_officer_by_last_name_p
func (c *Client) TacomaFuzzySearchByLastName(lastName string) ([]*TacomaOfficer, error) {
	rows, err := c.pool.Query(context.Background(), `
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

func marshalTacomaOfficerRows(rows pgx.Rows) ([]*TacomaOfficer, error) {
	officers := []*TacomaOfficer{}
	for rows.Next() {
		ofc := TacomaOfficer{}
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
