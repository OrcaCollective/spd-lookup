package data

import (
	"context"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// RentonOfficer is the object model for LPD officers
type RentonOfficer struct {
	LastName           string `json:"last_name,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	Rank string `json:"rank,omitempty"`
	Department string `json:"department,omitempty"`
	Division string `json:"division,omitempty"`
	Shift string `json:"shift,omitempty"`
	AdditionalInfo string `json:"additional_info,omitempty"`
	Badge string `json:"badge,omitempty"`
}

// rentonOfficer is an internal intermediary between the returned SQL rows data
// and the actual JSON return itself.
type rentonOfficer struct {
	LastName           nulls.String
	FirstName       nulls.String
	MiddleName nulls.String
	Rank nulls.String
	Department nulls.String
	Division nulls.String
    Shift nulls.String
	AdditionalInfo nulls.String
	Badge nulls.String
}

// RentonOfficerMetadata retrieves metadata describing the RentonOfficer struct
func (c *Client) RentonOfficerMetadata() *DepartmentMetadata {
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
				"FieldName": "middle_name",
				"Label":     "Middle Name",
			},
			{
				"FieldName": "rank",
				"Label":     "Officer Rank",
			},
			{
				"FieldName": "department",
				"Label":     "Officer Department",
			},
			{
				"FieldName": "division",
				"Label":     "Officer Division",
			},
			{
				"FieldName": "shift",
				"Label":     "Shift",
			},
			{
				"FieldName": "additional_info",
				"Label":     "additional information (including retirement date)",
			},
			{
				"FieldName": "badge",
				"Label":     "Badge number",
			},
		},
		LastAvailableRosterDate: "2021-05-01",
		Name:                    "Renton PD",
		ID:                      "rpd",
		SearchRoutes: map[string]*SearchRouteMetadata{
			"exact": {
				Path:        "/renton/officer",
				QueryParams: []string{"first_name", "last_name"},
			},
			"fuzzy": {
				Path:        "/renton/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

// RentonSearchOfficerByName returns an officer by their first or last name.
func (c *Client) RentonSearchOfficerByName(firstName, lastName string) ([]*RentonOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.middle_name,
				o.rank,
				o.department,
				o.division,
				o.shift,
                o.additional_info,
                o.badge_number
			FROM renton_officers o
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

	return rentonMarshalOfficerRows(rows)
}

// RentonFuzzySearchByName returns an list of officers by their full name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) RentonFuzzySearchByName(name string) ([]*RentonOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.middle_name,
				o.rank,
				o.department,
				o.division,
				o.shift,
                o.additional_info,
                o.badge_number
			FROM renton_officers o
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

	return rentonMarshalOfficerRows(rows)
}

// RentonFuzzySearchByFirstName returns an list of officers by their first name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) RentonFuzzySearchByFirstName(firstName string) ([]*RentonOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.middle_name,
				o.rank,
				o.department,
				o.division,
				o.shift,
                o.additional_info,
                o.badge_number

			FROM renton_officers o
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

	return rentonMarshalOfficerRows(rows)
}

// RentonFuzzySearchByLastName returns an list of officers by their last name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) RentonFuzzySearchByLastName(lastName string) ([]*RentonOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.last_name,
				o.first_name,
				o.middle_name,
				o.rank,
				o.department,
				o.division,
				o.shift,
                o.additional_info,
                o.badge_number

			FROM renton_officers o
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

	return rentonMarshalOfficerRows(rows)
}

// rentonMarshalOfficerRows takes SQL return objects and marshals them onto the
// RentonOfficer object for return as JSON by the API.
func rentonMarshalOfficerRows(rows pgx.Rows) ([]*RentonOfficer, error) {
	officers := []*RentonOfficer{}
	for rows.Next() {
		ofc := rentonOfficer{}
		err := rows.Scan(
			&ofc.LastName,
			&ofc.FirstName,
			&ofc.MiddleName,
			&ofc.Rank,
			&ofc.Department,
			&ofc.Division,
			&ofc.Shift,
			&ofc.AdditionalInfo,
			&ofc.Badge,
		)

		if err != nil {
			return nil, err
		}

		returnOfficer := RentonOfficer{
			ofc.LastName.String,
			ofc.FirstName.String,
			ofc.MiddleName.String,
			ofc.Rank.String,
			ofc.Department.String,
			ofc.Division.String,
			ofc.Shift.String,
			ofc.AdditionalInfo.String,
			ofc.Badge.String,
		}

		officers = append(officers, &returnOfficer)
	}
	return officers, nil
}
