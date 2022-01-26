package data

import (
	"context"
	"fmt"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// SeattleOfficer is the object model for SPD officers
type SeattleOfficer struct {
	Date            string `json:"date,omitempty"`
	Badge           string `json:"badge,omitempty"`
	FullName        string `json:"full_name,omitempty"`
	Title           string `json:"title,omitempty"`
	Unit            string `json:"unit,omitempty"`
	UnitDescription string `json:"unit_description,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	MiddleName      string `json:"middle_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Current         bool   `json:"is_current"`
}

// seattleOfficer is an internal intermediary between the returned SQL rows data
// and the actual JSON return itself.
type seattleOfficer struct {
	Date            time.Time
	Badge           nulls.String
	FullName        nulls.String
	Title           nulls.String
	Unit            nulls.String
	UnitDescription nulls.String
	FirstName       nulls.String
	MiddleName      nulls.String
	LastName        nulls.String
	Current         bool
}

// SeattleOfficerMetadata retrieves metadata describing the SeattleOfficer struct
func (c *Client) SeattleOfficerMetadata() *DepartmentMetadata {
	var date time.Time
	err := c.pool.QueryRow(context.Background(),
		`
			SELECT max(date) as date
			FROM seattle_officers;
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
			{
				"FieldName": "full_name",
				"Label":     "Full Name",
			},
			{
				"FieldName": "is_current",
				"Label":     "On Current Roster",
			},
		},
		LastAvailableRosterDate: date.Format("2006-01-02"),
		Name:                    "Seattle PD",
		ID:                      "seattle-wa",
		SearchRoutes: map[string]*SearchRouteMetadata{
			"exact": {
				Path:        "/seattle-wa/officer",
				QueryParams: []string{"badge", "first_name", "last_name"},
			},
			"fuzzy": {
				Path:        "/seattle-wa/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
			"historical-exact": {
				Path:        "/seattle-wa/officer/historical",
				QueryParams: []string{"badge"},
			},
		},
	}
}

// SeattleGetOfficerByBadge returns an officer by their badge. It searches the full historical
// roster list but only returns the most recent entry.
func (c *Client) SeattleGetOfficerByBadge(badge string) ([]*SeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			WITH max_roster AS (SELECT MAX(date) max_date FROM seattle_officers)
			SELECT
				o.date,
				o.badge,
				o.full_name,
				o.first_name,
				o.middle_name,
				o.last_name,
				o.title,
				o.unit,
				o.unit_description,
				CASE WHEN o.date = m.max_date THEN TRUE ELSE FALSE END is_current
			FROM seattle_officers o
			CROSS JOIN max_roster m
			WHERE o.badge = $1
			ORDER BY o.date DESC
			LIMIT 1;
		`,
		badge,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return seattleMarshalOfficerRows(rows)
}

// SeattleGetOfficerByBadgeHistorical returns an officer by their badge. It searches
// the full historical roster list and returns all entries in descending date order.
func (c *Client) SeattleGetOfficerByBadgeHistorical(badge string) ([]*SeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			WITH max_roster AS (SELECT MAX(date) max_date FROM seattle_officers)
			SELECT
				o.date,
				o.badge,
				o.full_name,
				o.first_name,
				o.middle_name,
				o.last_name,
				o.title,
				o.unit,
				o.unit_description,
				CASE WHEN o.date = m.max_date THEN TRUE ELSE FALSE END is_current
			FROM seattle_officers o
			CROSS JOIN max_roster m
			WHERE o.badge = $1
			ORDER BY o.date DESC;
		`,
		badge,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return seattleMarshalOfficerRows(rows)
}

// SeattleSearchOfficerByName returns an officer by their first or last name. It searches the full historical
// roster list but only returns the most recent entry.
func (c *Client) SeattleSearchOfficerByName(firstName, lastName string) ([]*SeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			WITH o AS (
				SELECT
					*,
					row_number() over (partition by o.badge order by o.date desc) seqnum
				FROM seattle_officers o
				WHERE LOWER(o.first_name) LIKE LOWER($1)
				AND LOWER(o.last_name) LIKE LOWER($2)
			),
			max_roster AS (SELECT MAX(date) max_date FROM seattle_officers)
			SELECT
				o.date,
				o.badge,
				o.full_name,
				o.first_name,
				o.middle_name,
				o.last_name,
				o.title,
				o.unit,
				o.unit_description,
				CASE WHEN o.date = max_roster.max_date THEN TRUE ELSE FALSE END is_current
			FROM o, max_roster
			WHERE o.seqnum = 1
			ORDER BY 
				o.date DESC,
				o.full_name;
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

// SeattleFuzzySearchByName returns an list of officers by their full name using fuzzy matching.
// It searches the full historical roster list but only returns the last availabe roster entry. Entries
// are sorted by date in descending order.
func (c *Client) SeattleFuzzySearchByName(name string) ([]*SeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			WITH o AS (
				SELECT
					*,
					row_number() over (partition by o.badge order by o.date desc) seqnum
				FROM seattle_officers o
				WHERE LOWER(first_name || ' ' || last_name) % LOWER($1)
			),
			max_roster AS (SELECT MAX(date) max_date FROM seattle_officers)
			SELECT
				o.date,
				o.badge,
				o.full_name,
				o.first_name,
				o.middle_name,
				o.last_name,
				o.title,
				o.unit,
				o.unit_description,
				CASE WHEN o.date = max_roster.max_date THEN TRUE ELSE FALSE END is_current
			FROM o, max_roster
			WHERE o.seqnum = 1
			ORDER BY
				o.date DESC,
				SIMILARITY(LOWER(o.first_name || ' ' || o.last_name), LOWER($1)) DESC;
		`,
		name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return seattleMarshalOfficerRows(rows)
}

// SeattleFuzzySearchByFirstName returns an list of officers by their first name using fuzzy matching.
// It searches the full historical roster list but only returns the last availabe roster entry. Entries
// are sorted by date in descending order.
func (c *Client) SeattleFuzzySearchByFirstName(firstName string) ([]*SeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			WITH o AS (
				SELECT
					*,
					row_number() over (partition by o.badge order by o.date desc) seqnum
				FROM seattle_officers o
				WHERE LOWER(first_name) % LOWER($1)
			),
			max_roster AS (SELECT MAX(date) max_date FROM seattle_officers)
			SELECT
				o.date,
				o.badge,
				o.full_name,
				o.first_name,
				o.middle_name,
				o.last_name,
				o.title,
				o.unit,
				o.unit_description,
				CASE WHEN o.date = max_roster.max_date THEN TRUE ELSE FALSE END is_current
			FROM o, max_roster
			WHERE o.seqnum = 1
			ORDER BY
				o.date DESC,
				SIMILARITY(LOWER(o.first_name), LOWER($1)) DESC;
		`,
		firstName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return seattleMarshalOfficerRows(rows)
}

// SeattleFuzzySearchByLastName returns an list of officers by their last name using fuzzy matching.
// It searches the full historical roster list but only returns the last availabe roster entry. Entries
// are sorted by date in descending order.
func (c *Client) SeattleFuzzySearchByLastName(lastName string) ([]*SeattleOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			WITH o AS (
				SELECT
					*,
					row_number() over (partition by o.badge order by o.date desc) seqnum
				FROM seattle_officers o
				WHERE LOWER(last_name) % LOWER($1)
			),
			max_roster AS (SELECT MAX(date) max_date FROM seattle_officers)
			SELECT
				o.date,
				o.badge,
				o.full_name,
				o.first_name,
				o.middle_name,
				o.last_name,
				o.title,
				o.unit,
				o.unit_description,
				CASE WHEN o.date = max_roster.max_date THEN TRUE ELSE FALSE END is_current
			FROM o, max_roster
			WHERE o.seqnum = 1
			ORDER BY
				o.date DESC,
				SIMILARITY(LOWER(o.last_name), LOWER($1)) DESC;
		`,
		lastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return seattleMarshalOfficerRows(rows)
}

// seattleMarshalOfficerRows takes SQL return objects and marshals them onto the
// SeattleOfficer object for return as JSON by the API.
func seattleMarshalOfficerRows(rows pgx.Rows) ([]*SeattleOfficer, error) {
	officers := []*SeattleOfficer{}
	for rows.Next() {
		ofc := seattleOfficer{}
		err := rows.Scan(
			&ofc.Date,
			&ofc.Badge,
			&ofc.FullName,
			&ofc.FirstName,
			&ofc.MiddleName,
			&ofc.LastName,
			&ofc.Title,
			&ofc.Unit,
			&ofc.UnitDescription,
			&ofc.Current,
		)

		if err != nil {
			return nil, err
		}

		returnOfficer := SeattleOfficer{
			ofc.Date.Format("2006-01-02"),
			ofc.Badge.String,
			ofc.FullName.String,
			ofc.Title.String,
			ofc.Unit.String,
			ofc.UnitDescription.String,
			ofc.FirstName.String,
			ofc.MiddleName.String,
			ofc.LastName.String,
			ofc.Current,
		}

		officers = append(officers, &returnOfficer)
	}
	return officers, nil
}
