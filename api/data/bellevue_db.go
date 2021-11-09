package data

import (
	"context"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// BellevueOfficer is the object model for BPD officers
type BellevueOfficer struct {
	LastName           string `json:"last_name,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	Title string `json:"title,omitempty"`
	Unit string `json:"unit,omitempty"`
	Notes string `json:"notes,omitempty"`
	Badge string `json:"badge,omitempty"`
}

// bellevueOfficer is an internal intermediary between the returned SQL rows data
// and the actual JSON return itself.
type bellevueOfficer struct {
	LastName           nulls.String
	FirstName       nulls.String
	Title nulls.String
	Unit nulls.String
	Notes nulls.String
	Badge nulls.String
}

// BellevueOfficerMetadata retrieves metadata describing the BellevueOfficer struct
func (c *Client) BellevueOfficerMetadata() *DepartmentMetadata {
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
				"FieldName": "unit",
				"Label":     "Officer unit",
			},
			{
				"FieldName": "notes",
				"Label":     "additional information (including retirement date)",
			},
			{
				"FieldName": "badge",
				"Label":     "Badge number",
			},
		},
		LastAvailableRosterDate: "2021-05-01",
		Name:                    "Bellevue PD",
		ID:                      "bpd",
		SearchRoutes: map[string]*SearchRouteMetadata{
			"exact": {
				Path:        "/bellevue/officer",
				QueryParams: []string{"first_name", "last_name"},
			},
			"fuzzy": {
				Path:        "/bellevue/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

// BellevueSearchOfficerByName returns an officer by their first or last name.
func (c *Client) BellevueSearchOfficerByName(firstName, lastName string) ([]*BellevueOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.title,
				o.unit,
				o.notes,
                o.badge
			FROM bellevue_officers o
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

	return bellevueMarshalOfficerRows(rows)
}

// BellevueFuzzySearchByName returns an list of officers by their full name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) BellevueFuzzySearchByName(name string) ([]*BellevueOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.title,
				o.unit,
				o.notes,
                o.badge
			FROM bellevue_officers o
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

	return bellevueMarshalOfficerRows(rows)
}

// BellevueFuzzySearchByFirstName returns an list of officers by their first name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) BellevueFuzzySearchByFirstName(firstName string) ([]*BellevueOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.title,
				o.unit,
				o.notes,
                o.badge

				FROM bellevue_officers o
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

	return bellevueMarshalOfficerRows(rows)
}

// BellevueFuzzySearchByLastName returns an list of officers by their last name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) BellevueFuzzySearchByLastName(lastName string) ([]*BellevueOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.title,
				o.unit,
				o.notes,
                o.badge

			FROM bellevue_officers o
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

	return bellevueMarshalOfficerRows(rows)
}

// bellevueMarshalOfficerRows takes SQL return objects and marshals them onto the
// BellevueOfficer object for return as JSON by the API.
func bellevueMarshalOfficerRows(rows pgx.Rows) ([]*BellevueOfficer, error) {
	officers := []*BellevueOfficer{}
	for rows.Next() {
		ofc := bellevueOfficer{}
		err := rows.Scan(
			&ofc.LastName,
			&ofc.FirstName,
			&ofc.Title,
			&ofc.Unit,
			&ofc.Notes,
			&ofc.Badge,
		)

		if err != nil {
			return nil, err
		}

		returnOfficer := BellevueOfficer{
			ofc.LastName.String,
			ofc.FirstName.String,
			ofc.Title.String,
			ofc.Unit.String,
			ofc.Notes.String,
			ofc.Badge.String,
		}

		officers = append(officers, &returnOfficer)
	}
	return officers, nil
}
