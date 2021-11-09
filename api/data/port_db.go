package data

import (
	"context"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// PortOfSeattleOfficer is the object model for BPD officers
type PortOfSeattleOfficer struct {
	Name string `json:"name,omitempty"`
	Rank string `json:"rank,omitempty"`
	Unit string `json:"unit,omitempty"`
	Badge string `json:"badge,omitempty"`
}

// portOfSeattleOfficer is an internal intermediary between the returned SQL rows data
// and the actual JSON return itself.
type portOfSeattleOfficer struct {
	Name nulls.String
	Rank nulls.String
	Unit nulls.String
	Badge nulls.String
}

// PortOfSeattleOfficerMetadata retrieves metadata describing the PortOfSeattleOfficer struct
func (c *Client) PortOfSeattleOfficerMetadata() *DepartmentMetadata {
	return &DepartmentMetadata{
		Fields: []map[string]string{
			{
				"FieldName": "name",
				"Label":     "Full Name",
			},
			{
				"FieldName": "rank",
				"Label":     "Officer Title",
			},
			{
				"FieldName": "unit",
				"Label":     "Officer unit",
			},
			{
				"FieldName": "badge",
				"Label":     "Badge number",
			},
		},
		LastAvailableRosterDate: "2021-05-01",
		Name:                    "Port Of Seattle PD",
		ID:                      "pospd",
		SearchRoutes: map[string]*SearchRouteMetadata{
			"exact": {
				Path:        "/port_of_seattle/officer",
				QueryParams: []string{"name"},
			},
			"fuzzy": {
				Path:        "/port_of_seattle/officer/search",
				QueryParams: []string{"name"},
			},
		},
	}
}

// PortOfSeattleSearchOfficerByName returns an officer by their name.
func (c *Client) PortOfSeattleSearchOfficerByName(name string) ([]*PortOfSeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.name,
				o.rank,
				o.unit,
                o.badge_number
			FROM port_of_seattle_officers o
			WHERE LOWER(o.name) LIKE LOWER($1)
			ORDER BY 
				o.name;
		`,
		name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return portOfSeattleMarshalOfficerRows(rows)
}

// PortOfSeattleFuzzySearchByName returns an list of officers by their full name using fuzzy matching.
// Entries are sorted by similarity to the full name in descending order.
func (c *Client) PortOfSeattleFuzzySearchByName(name string) ([]*PortOfSeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				o.name,
				o.rank,
				o.unit,
                o.badge_number

			FROM port_of_seattle_officers o
			WHERE LOWER(name || ' ' ) % LOWER($1)
			ORDER BY
				SIMILARITY(LOWER(' ' || o.name), LOWER($1)) DESC;
		`,
		name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return portOfSeattleMarshalOfficerRows(rows)
}

// portOfSeattleMarshalOfficerRows takes SQL return objects and marshals them onto the
// PortOfSeattleOfficer object for return as JSON by the API.
func portOfSeattleMarshalOfficerRows(rows pgx.Rows) ([]*PortOfSeattleOfficer, error) {
	officers := []*PortOfSeattleOfficer{}
	for rows.Next() {
		ofc := portOfSeattleOfficer{}
		err := rows.Scan(
			&ofc.Name,
			&ofc.Rank,
			&ofc.Unit,
			&ofc.Badge,
		)

		if err != nil {
			return nil, err
		}

		returnOfficer := PortOfSeattleOfficer{
			ofc.Name.String,
			ofc.Rank.String,
			ofc.Unit.String,
			ofc.Badge.String,
		}

		officers = append(officers, &returnOfficer)
	}
	return officers, nil
}
