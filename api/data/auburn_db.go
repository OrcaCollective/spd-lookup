package data

import (
	"context"
	"fmt"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// AuburnOfficer is the object model for LPD officers
type AuburnOfficer struct {
	Date            string `json:"date,omitempty"`
	Badge           string `json:"badge,omitempty"`
	Title           string `json:"title,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
}

// auburnOfficer is an internal intermediary between the returned SQL rows data
// and the actual JSON return itself.
type auburnOfficer struct {
	Date            time.Time
	Badge           nulls.String
	Title           nulls.String
	FirstName       nulls.String
	LastName        nulls.String
}

// AuburnOfficerMetadata retrieves metadata describing the AuburnOfficer struct
func (c *Client) AuburnOfficerMetadata() *DepartmentMetadata {
	var date time.Time
	err := c.pool.QueryRow(context.Background(),
		`
			SELECT max(date) as date
			FROM auburn_officers;
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
				"FieldName": "badge",
				"Label":     "Badge",
			},
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
		},
		LastAvailableRosterDate: date.Format("2006-01-02"),
		Name:                    "Auburn PD",
		ID:                      "apd",
		SearchRoutes: map[string]*SearchRouteMetadata{
			"exact": {
				Path:        "/auburn/officer",
				QueryParams: []string{"badge", "first_name", "last_name"},
			},
			"fuzzy": {
				Path:        "/auburn/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

// AuburnGetOfficerByBadge returns an officer by their badge.
func (c *Client) AuburnGetOfficerByBadge(badge string) ([]*AuburnOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.badge,
				o.first_name,
				o.last_name,
				o.title
			FROM auburn_officers o
			WHERE o.badge = $1;
		`,
		badge,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return auburnMarshalOfficerRows(rows)
}

// AuburnSearchOfficerByName returns an officer by their first or last name.
func (c *Client) AuburnSearchOfficerByName(firstName, lastName string) ([]*AuburnOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.badge,
				o.first_name,
				o.last_name,
				o.title
			FROM auburn_officers o
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

	return auburnMarshalOfficerRows(rows)
}

// AuburnFuzzySearchByName returns an list of officers by their full name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) AuburnFuzzySearchByName(name string) ([]*AuburnOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.badge,
				o.first_name,
				o.last_name,
				o.title
			FROM auburn_officers o
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

	return auburnMarshalOfficerRows(rows)
}

// AuburnFuzzySearchByFirstName returns an list of officers by their first name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) AuburnFuzzySearchByFirstName(firstName string) ([]*AuburnOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.badge,
				o.first_name,
				o.last_name,
				o.title
			FROM auburn_officers o
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

	return auburnMarshalOfficerRows(rows)
}

// AuburnFuzzySearchByLastName returns an list of officers by their last name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) AuburnFuzzySearchByLastName(lastName string) ([]*AuburnOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.badge,
				o.first_name,
				o.last_name,
				o.title
			FROM auburn_officers o
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

	return auburnMarshalOfficerRows(rows)
}

// auburnMarshalOfficerRows takes SQL return objects and marshals them onto the
// AuburnOfficer object for return as JSON by the API.
func auburnMarshalOfficerRows(rows pgx.Rows) ([]*AuburnOfficer, error) {
	officers := []*AuburnOfficer{}
	for rows.Next() {
		ofc := auburnOfficer{}
		err := rows.Scan(
			&ofc.Date,
			&ofc.Badge,
			&ofc.FirstName,
			&ofc.LastName,
			&ofc.Title,
		)

		if err != nil {
			return nil, err
		}

		returnOfficer := AuburnOfficer{
			ofc.Date.Format("2006-01-02"),
			ofc.Badge.String,
			ofc.Title.String,
			ofc.FirstName.String,
			ofc.LastName.String,
		}

		officers = append(officers, &returnOfficer)
	}
	return officers, nil
}
