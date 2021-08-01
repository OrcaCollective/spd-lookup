package data

import (
	"context"
	"fmt"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// LakewoodOfficer is the object model for LPD officers
type LakewoodOfficer struct {
	Date            string `json:"date,omitempty"`
	Title           string `json:"title,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	Unit        string `json:"unit,omitempty"`
	UnitDescription string `json:"unit_description"`
}

// lakeOfficer is an internal intermediary between the returned SQL rows data
// and the actual JSON return itself.
type lakewoodOfficer struct {
	Date            time.Time
	Title           nulls.String
	LastName           nulls.String
	FirstName       nulls.String
	Unit        nulls.String
	UnitDescription nulls.String
}

// LakewoodOfficerMetadata retrieves metadata describing the LakewoodOfficer struct
func (c *Client) LakewoodOfficerMetadata() *DepartmentMetadata {
	var date time.Time
	err := c.pool.QueryRow(context.Background(),
		`
			SELECT max(date) as date
			FROM lakewood_officers;
		`).Scan(&date)

	if err != nil {
		fmt.Printf("DB Client Error: %s", err)
		return &DepartmentMetadata{}
	}

	return &DepartmentMetadata{
		Fields: []map[string]string{
			{
				"FieldName": "date",
				"Label":     "Roster Date",
			},
			{
				"FieldName": "title",
				"Label":     "Title",
			},
			{
				"FieldName": "last_name",
				"Label":     "Last Name",
			},
			{
				"FieldName": "first_name",
				"Label":     "First Name",
			},
			{
				"FieldName": "unit",
				"Label":     "Unit",
			},
			{
				"FieldName": "unit_descritpion",
				"Label": "Unit Description",
			},
		},
		LastAvailableRosterDate: date.Format("2006-01-02"),
		Name:                    "Lakewood PD",
		ID:                      "lpd",
		SearchRoutes: map[string]*SearchRouteMetadata{
			"exact": {
				Path:        "/lakewood/officer",
				QueryParams: []string{"first_name", "last_name"},
			},
			"fuzzy": {
				Path:        "/lakewood/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

// LakewoodSearchOfficerByName returns an officer by their first or last name.
func (c *Client) LakewoodSearchOfficerByName(firstName, lastName string) ([]*LakewoodOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.title,
				o.last_name,
				o.first_name,
				o.unit,
				o.unit_description
			FROM lakewood_officers o
			WHERE LOWER(o.first_name) LIKE LOWER($1) AND LOWER(o.last_name) LIKE LOWER($2)
			ORDER BY 
				o.last_name,
				o.first_name;
		`,
		firstName,
		lastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return lakewoodMarshalOfficerRows(rows)
}

// LakewoodFuzzySearchByName returns an list of officers by their full name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) LakewoodFuzzySearchByName(name string) ([]*LakewoodOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.title,
				o.last_name,
				o.first_name,
				o.unit,
				o.unit_description
			FROM lakewood_officers o
			WHERE LOWER(first_name || ' ' || last_name) % LOWER($1)
			ORDER BY
				SIMILARITY(LOWER(o.first_name || ' ' || o.last_name), LOWER($1)) DESC;
		`,
		name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return lakewoodMarshalOfficerRows(rows)
}

// LakewoodFuzzySearchByFirstName returns an list of officers by their first name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) LakewoodFuzzySearchByFirstName(firstName string) ([]*LakewoodOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.title,
				o.last_name,
				o.first_name,
				o.unit,
				o.unit_description
			FROM lakewood_officers o
			WHERE LOWER(first_name) % LOWER($1)
			ORDER BY
				SIMILARITY(LOWER(o.first_name), LOWER($1)) DESC;
		`,
		firstName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return lakewoodMarshalOfficerRows(rows)
}

// LakewoodFuzzySearchByLastName returns an list of officers by their last name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) LakewoodFuzzySearchByLastName(lastName string) ([]*LakewoodOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.title,
				o.last_name,
				o.first_name,
				o.unit,
				o.unit_description
			FROM lakewood_officers o
			WHERE LOWER(last_name) % LOWER($1)
			ORDER BY
				SIMILARITY(LOWER(o.last_name), LOWER($1)) DESC;
		`,
		lastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return lakewoodMarshalOfficerRows(rows)
}

// lakewoodMarshalOfficerRows takes SQL return objects and marshals them onto the
// AuburnOfficer object for return as JSON by the API.
func lakewoodMarshalOfficerRows(rows pgx.Rows) ([]*LakewoodOfficer, error) {
	officers := []*LakewoodOfficer{}
	for rows.Next() {
		ofc := lakewoodOfficer{}
		err := rows.Scan(
			&ofc.Date,
			&ofc.Title,
			&ofc.LastName,
			&ofc.FirstName,
			&ofc.Unit,
			&ofc.UnitDescription,
		)

		if err != nil {
			return nil, err
		}

		returnOfficer := LakewoodOfficer{
			ofc.Date.Format("2006-01-02"),
			ofc.Title.String,
			ofc.LastName.String,
			ofc.FirstName.String,
			ofc.Unit.String,
			ofc.UnitDescription.String,
		}

		officers = append(officers, &returnOfficer)
	}
	return officers, nil
}
