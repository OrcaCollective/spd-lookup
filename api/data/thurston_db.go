package data

import (
	"context"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// ThurstonCountyOfficer is the object model for BPD officers
type ThurstonCountyOfficer struct {
	LastName  string `json:"last_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	Title     string `json:"title,omitempty"`
	CallSign  string `json:"call_sign,omitempty"`
}

// thurstonCountyOfficer is an internal intermediary between the returned SQL rows data
// and the actual JSON return itself.
type thurstonCountyOfficer struct {
	LastName  nulls.String
	FirstName nulls.String
	Title     nulls.String
	CallSign  nulls.String
}

// ThurstonCountyOfficerMetadata retrieves metadata describing the ThurstonCountyOfficer struct
func (c *Client) ThurstonCountyOfficerMetadata() *DepartmentMetadata {
	return &DepartmentMetadata{
		Fields: []map[string]string{
			{
				"FieldName": "last_name",
				"Label":     "Last Name",
			},
			{
				"FieldName": "first_name",
				"Label":     "First Name",
			},
			{
				"FieldName": "title",
				"Label":     "Officer Title",
			},
			{
				"FieldName": "call_sign",
				"Label":     "Call Sign",
			},
		},
		LastAvailableRosterDate: "2021-05-01",
		Name:                    "Thurston County Sheriff's Department",
		ID:                      "tcsd",
		SearchRoutes: map[string]*SearchRouteMetadata{
			"exact": {
				Path:        "/thurston_county/officer",
				QueryParams: []string{"first_name", "last_name"},
			},
			"fuzzy": {
				Path:        "/thurston_county/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

// ThurstonCountySearchOfficerByName returns an officer by their first or last name.
func (c *Client) ThurstonCountySearchOfficerByName(firstName, lastName string) ([]*ThurstonCountyOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.title,
                o.call_sign
			FROM thurston_officers o
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

	return thurstonCountyMarshalOfficerRows(rows)
}

// ThurstonCountyFuzzySearchByName returns an list of officers by their full name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) ThurstonCountyFuzzySearchByName(name string) ([]*ThurstonCountyOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.title,
                o.call_sign

			FROM thurston_officers o
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

	return thurstonCountyMarshalOfficerRows(rows)
}

// ThurstonCountyFuzzySearchByFirstName returns an list of officers by their first name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) ThurstonCountyFuzzySearchByFirstName(firstName string) ([]*ThurstonCountyOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.title,
                o.call_sign

				FROM thurston_officers o
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

	return thurstonCountyMarshalOfficerRows(rows)
}

// ThurstonCountyFuzzySearchByLastName returns an list of officers by their last name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) ThurstonCountyFuzzySearchByLastName(lastName string) ([]*ThurstonCountyOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.title,
                o.call_sign

			FROM thurston_officers o
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

	return thurstonCountyMarshalOfficerRows(rows)
}

// thurstonCountyMarshalOfficerRows takes SQL return objects and marshals them onto the
// ThurstonCountyOfficer object for return as JSON by the API.
func thurstonCountyMarshalOfficerRows(rows pgx.Rows) ([]*ThurstonCountyOfficer, error) {
	officers := []*ThurstonCountyOfficer{}
	for rows.Next() {
		ofc := thurstonCountyOfficer{}
		err := rows.Scan(
			&ofc.LastName,
			&ofc.FirstName,
			&ofc.Title,
			&ofc.CallSign,
		)

		if err != nil {
			return nil, err
		}

		returnOfficer := ThurstonCountyOfficer{
			ofc.LastName.String,
			ofc.FirstName.String,
			ofc.Title.String,
			ofc.CallSign.String,
		}

		officers = append(officers, &returnOfficer)
	}
	return officers, nil
}
