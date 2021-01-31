package data

import (
	"context"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// SeattleOfficer is the object model for SPD officers
type SeattleOfficer struct {
	BadgeNumber     string       `json:"badge_number,omitempty"`
	FirstName       string       `json:"first_name,omitempty"`
	MiddleName      nulls.String `json:"middle_name,omitempty"`
	LastName        string       `json:"last_name,omitempty"`
	Title           string       `json:"title,omitempty"`
	Unit            string       `json:"unit,omitempty"`
	UnitDescription nulls.String `json:"unit_description,omitempty"`
}

// SeattleOfficerMetadata retrieves metadata describing the SeattleOfficer struct
func (c *Client) SeattleOfficerMetadata() []map[string]string {
	return []map[string]string{
		{
			"FieldName": "badge_number",
			"Label":     "Badge Number",
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
			"FieldName": "unit_description",
			"Label":     "Unit Description",
		},
	}
}

// SeattleGetOfficerByBadge invokes seattle_get_officer_by_badge_p
func (c *Client) SeattleGetOfficerByBadge(badge string) (*SeattleOfficer, error) {
	ofc := SeattleOfficer{}
	err := c.pool.QueryRow(context.Background(),
		`
			SELECT
				badge_number,
				first_name,
				middle_name,
				last_name,
				title,
				unit,
				unit_description
			FROM seattle_get_officer_by_badge_p (badge_number := $1);
		`,
		badge,
	).Scan(
		&ofc.BadgeNumber,
		&ofc.FirstName,
		&ofc.MiddleName,
		&ofc.LastName,
		&ofc.Title,
		&ofc.Unit,
		&ofc.UnitDescription,
	)

	return &ofc, err
}

// SeattleSearchOfficerByName invokes seattle_search_officer_by_name_p
func (c *Client) SeattleSearchOfficerByName(firstName, lastName string) ([]*SeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				badge_number,
				first_name,
				middle_name,
				last_name,
				title,
				unit,
				unit_description
			FROM seattle_search_officer_by_name_p(first_name := $1, last_name := $2);
		`,
		firstName,
		lastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return seattleMarshalOfficerRows(rows)
}

// SeattleFuzzySearchByName invokes seattle_fuzzy_search_officer_by_name_p
func (c *Client) SeattleFuzzySearchByName(name string) ([]*SeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				badge_number,
				first_name,
				middle_name,
				last_name,
				title,
				unit,
				unit_description
			FROM seattle_fuzzy_search_officer_by_name_p(full_name := $1);
		`,
		name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return seattleMarshalOfficerRows(rows)
}

// SeattleFuzzySearchByFirstName invokes seattle_fuzzy_search_officer_by_first_name_p
func (c *Client) SeattleFuzzySearchByFirstName(firstName string) ([]*SeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				badge_number,
				first_name,
				middle_name,
				last_name,
				title,
				unit,
				unit_description
			FROM seattle_fuzzy_search_officer_by_first_name_p(first_name := $1);
		`,
		firstName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return seattleMarshalOfficerRows(rows)
}

// SeattleFuzzySearchByLastName invokes seattle_fuzzy_search_officer_by_last_name_p
func (c *Client) SeattleFuzzySearchByLastName(lastName string) ([]*SeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
				badge_number,
				first_name,
				middle_name,
				last_name,
				title,
				unit,
				unit_description
			FROM seattle_fuzzy_search_officer_by_last_name_p(last_name := $1);
		`,
		lastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return seattleMarshalOfficerRows(rows)
}

func seattleMarshalOfficerRows(rows pgx.Rows) ([]*SeattleOfficer, error) {
	officers := []*SeattleOfficer{}
	for rows.Next() {
		ofc := SeattleOfficer{}
		err := rows.Scan(
			&ofc.BadgeNumber,
			&ofc.FirstName,
			&ofc.MiddleName,
			&ofc.LastName,
			&ofc.Title,
			&ofc.Unit,
			&ofc.UnitDescription,
		)

		if err != nil {
			return nil, err
		}
		officers = append(officers, &ofc)
	}
	return officers, nil
}
