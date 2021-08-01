package data

import (
	"context"
	"fmt"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// OlympiaOfficer is the object model for LPD officers
type OlympiaOfficer struct {
	Date      string `json:"date,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Title     string `json:"title,omitempty"`
	Unit      string `json:"unit,omitempty"`
	Badge     string `json:"badge,omitempty"`
}

// olympiaOfficerOfficer is an internal intermediary between the returned SQL rows data
// and the actual JSON return itself.
type olympiaOfficer struct {
	Date      time.Time
	FirstName nulls.String
	LastName  nulls.String
	Title     nulls.String
	Unit      nulls.String
	Badge     nulls.String
}

// OlympiaOfficerMetadata retrieves metadata describing the OlympiaOfficer struct
func (c *Client) OlympiaOfficerMetadata() *DepartmentMetadata {
	var date time.Time
	err := c.pool.QueryRow(context.Background(),
		`
			SELECT max(date) as date
			FROM olympia_officers;
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
				"FieldName": "unit",
				"Label":     "Unit",
			},
			{
				"FieldName": "badge",
				"Label":     "Badge",
			},
		},
		LastAvailableRosterDate: date.Format("2006-01-02"),
		Name:                    "Olympia PD",
		ID:                      "opd",
		SearchRoutes: map[string]*SearchRouteMetadata{
			"exact": {
				Path:        "/olympia/officer",
				QueryParams: []string{"badge", "first_name", "last_name"},
			},
			"fuzzy": {
				Path:        "/olympia/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

// OlympiaGetOfficerByBadge returns an officer by their badge.
func (c *Client) OlympiaGetOfficerByBadge(badge string) ([]*OlympiaOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.first_name,
				o.last_name,
				o.title,
				o.unit,
				o.badge
			FROM olympia_officers o
			WHERE o.badge = $1;
		`,
		badge,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return olympiaMarshalOfficerRows(rows)
}

// OlympiaSearchOfficerByName returns an officer by their first or last name.
func (c *Client) OlympiaSearchOfficerByName(firstName, lastName string) ([]*OlympiaOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.first_name,
				o.last_name,
				o.title,
				o.unit,
				o.badge
			FROM olympia_officers o
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

	return olympiaMarshalOfficerRows(rows)
}

// OlympiaFuzzySearchByName returns an list of officers by their full name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) OlympiaFuzzySearchByName(name string) ([]*OlympiaOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.first_name,
				o.last_name,
				o.title,
				o.unit,
				o.badge
			FROM olympia_officers o
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

	return olympiaMarshalOfficerRows(rows)
}

// OlympiaFuzzySearchByFirstName returns an list of officers by their first name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) OlympiaFuzzySearchByFirstName(firstName string) ([]*OlympiaOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.first_name,
				o.last_name,
				o.title,
				o.unit,
				o.badge
			FROM olympia_officers o
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

	return olympiaMarshalOfficerRows(rows)
}

// OlympiaFuzzySearchByLastName returns an list of officers by their last name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) OlympiaFuzzySearchByLastName(lastName string) ([]*OlympiaOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.date,
				o.first_name,
				o.last_name,
				o.title,
				o.unit,
				o.badge
			FROM olympia_officers o
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

	return olympiaMarshalOfficerRows(rows)
}

// olympiaMarshalOfficerRows takes SQL return objects and marshals them onto the
// OlympiaOfficer object for return as JSON by the API.
func olympiaMarshalOfficerRows(rows pgx.Rows) ([]*OlympiaOfficer, error) {
	officers := []*OlympiaOfficer{}
	for rows.Next() {
		ofc := olympiaOfficer{}
		err := rows.Scan(
			&ofc.Date,
			&ofc.FirstName,
			&ofc.LastName,
			&ofc.Title,
			&ofc.Unit,
			&ofc.Badge,
		)

		if err != nil {
			return nil, err
		}

		returnOfficer := OlympiaOfficer{
			ofc.Date.Format("2006-01-02"),
			ofc.FirstName.String,
			ofc.LastName.String,
			ofc.Title.String,
			ofc.Unit.String,
			ofc.Badge.String,
		}

		officers = append(officers, &returnOfficer)
	}
	return officers, nil
}
