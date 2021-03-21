package data

import (
	"context"

	"github.com/gobuffalo/nulls"
	"github.com/jackc/pgx/v4"
)

// PortlandOfficer is the object model for PPB officers
type PortlandOfficer struct {
	FirstName                   string          `json:"first_name,omitempty"`
	LastName                    string          `json:"last_name,omitempty"`
	Gender                      nulls.String    `json:"gender,omitempty"`
	OfficerRank                 nulls.String    `json:"officer_rank,omitempty"`
	EmployeeID                  nulls.String    `json:"employee_id,omitempty"`
	HelmetID                    nulls.String    `json:"helmet_id,omitempty"`
	HelmetIDThreeDigit          nulls.String    `json:"helmet_id_three_digit,omitempty"`
	Salary                      nulls.String    `json:"salary,omitempty"`
	Badge                       nulls.String    `json:"badge,omitempty"`
	CopsPhotoProfileLink        nulls.String    `json:"cops_photo_profile_link,omitempty"`
	CopsPhotoHasPhoto           nulls.String    `json:"cops_photo_has_photo,omitempty"`
	Employed_3_12_21            nulls.String    `json:"employed_3_12_21,omitempty"`
	Employed_12_28_20           nulls.String    `json:"employed_12_28_20,omitempty"`
	Employed_10_01_20           nulls.String    `json:"employed_10_01_20,omitempty"`
	Retired_6_1_20              nulls.String    `json:"retired_6_1_20,omitempty"`
	RetiredOrCertRevoked        nulls.String    `json:"retired_or_cert_revoked,omitempty"`
	RetiredOrCertRevokedDate    nulls.String    `json:"retired_or_cert_revoked_date,omitempty"`
	HireYear                    nulls.String    `json:"hire_year,omitempty"`
	HireDate                    nulls.String    `json:"hire_date,omitempty"`
	StateCertDate               nulls.String    `json:"state_cert_date,omitempty"`
	StateCertLevel              nulls.String    `json:"state_cert_level,omitempty"`
	RRT                         nulls.String    `json:"rrt,omitempty"`
	RRT2016                     nulls.String    `json:"rrt_2016,omitempty"`
	RRT2018NiiyaEmail           nulls.String    `json:"rrt_2018_niiya_email,omitempty"`
	RRT2018                     nulls.String    `json:"rrt_2018,omitempty"`
	RRT2019                     nulls.String    `json:"rrt_2019,omitempty"`
	RRT2020                     nulls.String    `json:"rrt_2020,omitempty"`
	SoundTruckTraining          nulls.String    `json:"sound_truck_training_2020,omitempty"`
	InstructedForDpsst          nulls.String    `json:"instructed_for_dpsst,omitempty"`
	InstructedForLessLethal     nulls.String    `json:"instructed_for_less_lethal,omitempty"`
	InvolvedInOisUof            nulls.String    `json:"involved_in_ois_uof,omitempty"`
	Notes                       nulls.String    `json:"notes,omitempty"`
}

// PortlandOfficerMetadata retrieves metadata describing the PortlandOfficer struct
func (c *Client) PortlandOfficerMetadata() *DepartmentMetadata {
	return &DepartmentMetadata{
		Fields: []map[string]string{
            {
                "FieldName": "first_name",
                "Label": "First Name",
            },
            {
                "FieldName": "last_name",
                "Label": "Last Name",
            },
            {
                "FieldName": "gender",
                "Label": "Gender",
            },
            {
                "FieldName": "officer_rank",
                "Label": "Rank",
            },
            {
                "FieldName": "employee_id",
                "Label": "Employee (Chest) ID",
            },
            {
                "FieldName": "helmet_id",
                "Label": "Helmet #",
            },
            {
                "FieldName": "helmet_id_three_digit",
                "Label": "3-Digit Helmet #",
            },
            {
                "FieldName": "salary",
                "Label": "Fiscal Earnings 2019",
            },
            {
                "FieldName": "badge",
                "Label": "Badge/DPSST Number",
            },
            {
                "FieldName": "cops_photo_profile_link",
                "Label": "Cops.Photo Profile Link",
            },
            {
                "FieldName": "cops_photo_has_photo",
                "Label": "Pic on Cops.photo (y/n)",
            },
            {
                "FieldName": "employed_3_12_21",
                "Label": "Employed as of 3/12/21",
            },
            {
                "FieldName": "employed_12_28_20",
                "Label": "Employed as of 12/28/20",
            },
            {
                "FieldName": "employed_10_01_20",
                "Label": "Employed as of 10/01/20",
            },
            {
                "FieldName": "retired_6_1_20",
                "Label": "Retired/Resigned as of 6/1/20",
            },
            {
                "FieldName": "retired_or_cert_revoked",
                "Label": "Retired/Resigned as of 6/1/20 OR Cert Revoked (ever)",
            },
            {
                "FieldName": "retired_or_cert_revoked_date",
                "Label": "Date of Cert Revoke",
            },
            {
                "FieldName": "hire_year",
                "Label": "Hire Year",
            },
            {
                "FieldName": "hire_date",
                "Label": "Hire Date",
            },
            {
                "FieldName": "state_cert_date",
                "Label": "State Certification Date",
            },
            {
                "FieldName": "state_cert_level",
                "Label": "State Certification Level",
            },
            {
                "FieldName": "rrt",
                "Label": "RRT (Rapid Response Team) Member",
            },
            {
                "FieldName": "rrt_2016",
                "Label": "RRT member as of 2016 via 2017 PPB AR",
            },
            {
                "FieldName": "rrt_2018_niiya_email",
                "Label": "RRT member as of 2018 via Niiya Email",
            },
            {
                "FieldName": "rrt_2018",
                "Label": "RRT Specific Training 2018",
            },
            {
                "FieldName": "rrt_2019",
                "Label": "RRT Specific Training 2019",
            },
            {
                "FieldName": "rrt_2020",
                "Label": "RRT Specific Training 2020",
            },
            {
                "FieldName": "sound_truck_training_2020",
                "Label": "Sound Truck Training 2020",
            },
            {
                "FieldName": "instructed_for_dpsst",
                "Label": "Has Instructed Course for DPSST 2017+",
            },
            {
                "FieldName": "instructed_for_less_lethal",
                "Label": "Instructor for Less Lethal/Chemical Weapons Courses",
            },
            {
                "FieldName": "involved_in_ois_uof",
                "Label": "Has Been Involved in OIS/Significant UoF Incident",
            },
            {
                "FieldName": "notes",
                "Label": "Notes",
            },
		},
		LastAvailableRosterDate: "2021-03-12",
		Name:                    "Portland PB",
		ID:                      "ppb",
		SearchRoutes: map[string]*SearchRouteMetadata{
			"exact": {
				Path:        "/portland/officer",
				QueryParams: []string{"badge", "first_name", "last_name", "employee_id", "helmet_id", "helmet_id_three_digit"},
			},
			"fuzzy": {
				Path:        "/portland/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

// PortlandSearchOfficersByBadge invokes portland_search_officer_by_badge_p
func (c *Client) PortlandSearchOfficersByBadge(badge string) ([]*PortlandOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
                first_name,
                last_name,
                gender,
                officer_rank,
                employee_id,
                helmet_id,
                helmet_id_three_digit,
                salary,
                badge,
                cops_photo_profile_link,
                cops_photo_has_photo,
                employed_3_12_21,
                employed_12_28_20,
                employed_10_01_20,
                retired_6_1_20,
                retired_or_cert_revoked,
                retired_or_cert_revoked_date,
                hire_year,
                hire_date,
                state_cert_date,
                state_cert_level,
                rrt,
                rrt_2016,
                rrt_2018_niiya_email,
                rrt_2018,
                rrt_2019,
                rrt_2020,
                sound_truck_training_2020,
                instructed_for_dpsst,
                instructed_for_less_lethal,
                involved_in_ois_uof,
                notes
			FROM portland_search_officer_by_badge_p (badge := $1);
		`,
		badge,
    )

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return portlandMarshalOfficerRows(rows)
}

// PortlandGetOfficerByEmployeeId invokes portland_search_officer_by_employee_p
func (c *Client) PortlandSearchOfficersByEmployeeId(employee_id string) ([]*PortlandOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
                first_name,
                last_name,
                gender,
                officer_rank,
                employee_id,
                helmet_id,
                helmet_id_three_digit,
                salary,
                badge,
                cops_photo_profile_link,
                cops_photo_has_photo,
                employed_3_12_21,
                employed_12_28_20,
                employed_10_01_20,
                retired_6_1_20,
                retired_or_cert_revoked,
                retired_or_cert_revoked_date,
                hire_year,
                hire_date,
                state_cert_date,
                state_cert_level,
                rrt,
                rrt_2016,
                rrt_2018_niiya_email,
                rrt_2018,
                rrt_2019,
                rrt_2020,
                sound_truck_training_2020,
                instructed_for_dpsst,
                instructed_for_less_lethal,
                involved_in_ois_uof,
                notes
			FROM portland_search_officer_by_employee_p (employee_id := $1);
		`,
		employee_id,
    )

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return portlandMarshalOfficerRows(rows)
}

// PortlandGetOfficerByHelmetId invokes portland_search_officer_by_helmet_p
func (c *Client) PortlandSearchOfficersByHelmetId(helmet_id string) ([]*PortlandOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
                first_name,
                last_name,
                gender,
                officer_rank,
                employee_id,
                helmet_id,
                helmet_id_three_digit,
                salary,
                badge,
                cops_photo_profile_link,
                cops_photo_has_photo,
                employed_3_12_21,
                employed_12_28_20,
                employed_10_01_20,
                retired_6_1_20,
                retired_or_cert_revoked,
                retired_or_cert_revoked_date,
                hire_year,
                hire_date,
                state_cert_date,
                state_cert_level,
                rrt,
                rrt_2016,
                rrt_2018_niiya_email,
                rrt_2018,
                rrt_2019,
                rrt_2020,
                sound_truck_training_2020,
                instructed_for_dpsst,
                instructed_for_less_lethal,
                involved_in_ois_uof,
                notes
			FROM portland_search_officer_by_helmet_p (helmet_id := $1);
		`,
		helmet_id,
    )

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return portlandMarshalOfficerRows(rows)
}

// PortlandGetOfficerByHelmetIdThreeDigit invokes portland_search_officer_by_helmet_p
func (c *Client) PortlandSearchOfficersByHelmetIdThreeDigit(helmet_id_three_digit string) ([]*PortlandOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
                first_name,
                last_name,
                gender,
                officer_rank,
                employee_id,
                helmet_id,
                helmet_id_three_digit,
                salary,
                badge,
                cops_photo_profile_link,
                cops_photo_has_photo,
                employed_3_12_21,
                employed_12_28_20,
                employed_10_01_20,
                retired_6_1_20,
                retired_or_cert_revoked,
                retired_or_cert_revoked_date,
                hire_year,
                hire_date,
                state_cert_date,
                state_cert_level,
                rrt,
                rrt_2016,
                rrt_2018_niiya_email,
                rrt_2018,
                rrt_2019,
                rrt_2020,
                sound_truck_training_2020,
                instructed_for_dpsst,
                instructed_for_less_lethal,
                involved_in_ois_uof,
                notes
			FROM portland_search_officer_by_helmet_three_digit_p (helmet_id_three_digit := $1);
		`,
		helmet_id_three_digit,
    )

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return portlandMarshalOfficerRows(rows)
}

// PortlandSearchOfficerByName invokes portland_search_officer_by_name_p
func (c *Client) PortlandSearchOfficersByName(firstName, lastName string) ([]*PortlandOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
                first_name,
                last_name,
                gender,
                officer_rank,
                employee_id,
                helmet_id,
                helmet_id_three_digit,
                salary,
                badge,
                cops_photo_profile_link,
                cops_photo_has_photo,
                employed_3_12_21,
                employed_12_28_20,
                employed_10_01_20,
                retired_6_1_20,
                retired_or_cert_revoked,
                retired_or_cert_revoked_date,
                hire_year,
                hire_date,
                state_cert_date,
                state_cert_level,
                rrt,
                rrt_2016,
                rrt_2018_niiya_email,
                rrt_2018,
                rrt_2019,
                rrt_2020,
                sound_truck_training_2020,
                instructed_for_dpsst,
                instructed_for_less_lethal,
                involved_in_ois_uof,
                notes
			FROM portland_search_officer_by_name_p(first_name := $1, last_name := $2);
		`,
		firstName,
		lastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return portlandMarshalOfficerRows(rows)
}

// PortlandFuzzySearchByName invokes portland_fuzzy_search_officer_by_name_p
func (c *Client) PortlandFuzzySearchByName(name string) ([]*PortlandOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
                first_name,
                last_name,
                gender,
                officer_rank,
                employee_id,
                helmet_id,
                helmet_id_three_digit,
                salary,
                badge,
                cops_photo_profile_link,
                cops_photo_has_photo,
                employed_3_12_21,
                employed_12_28_20,
                employed_10_01_20,
                retired_6_1_20,
                retired_or_cert_revoked,
                retired_or_cert_revoked_date,
                hire_year,
                hire_date,
                state_cert_date,
                state_cert_level,
                rrt,
                rrt_2016,
                rrt_2018_niiya_email,
                rrt_2018,
                rrt_2019,
                rrt_2020,
                sound_truck_training_2020,
                instructed_for_dpsst,
                instructed_for_less_lethal,
                involved_in_ois_uof,
                notes
			FROM portland_fuzzy_search_officer_by_name_p(full_name_v := $1);
		`,
		name,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return portlandMarshalOfficerRows(rows)
}

// PortlandFuzzySearchByFirstName invokes portland_fuzzy_search_officer_by_first_name_p
func (c *Client) PortlandFuzzySearchByFirstName(firstName string) ([]*PortlandOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
                first_name,
                last_name,
                gender,
                officer_rank,
                employee_id,
                helmet_id,
                helmet_id_three_digit,
                salary,
                badge,
                cops_photo_profile_link,
                cops_photo_has_photo,
                employed_3_12_21,
                employed_12_28_20,
                employed_10_01_20,
                retired_6_1_20,
                retired_or_cert_revoked,
                retired_or_cert_revoked_date,
                hire_year,
                hire_date,
                state_cert_date,
                state_cert_level,
                rrt,
                rrt_2016,
                rrt_2018_niiya_email,
                rrt_2018,
                rrt_2019,
                rrt_2020,
                sound_truck_training_2020,
                instructed_for_dpsst,
                instructed_for_less_lethal,
                involved_in_ois_uof,
                notes
			FROM portland_fuzzy_search_officer_by_first_name_p(first_name := $1);
		`,
		firstName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return portlandMarshalOfficerRows(rows)
}

// PortlandFuzzySearchByLastName invokes portland_fuzzy_search_officer_by_last_name_p
func (c *Client) PortlandFuzzySearchByLastName(lastName string) ([]*PortlandOfficer, error) {
	rows, err := c.pool.Query(context.Background(),
		`
			SELECT
                first_name,
                last_name,
                gender,
                officer_rank,
                employee_id,
                helmet_id,
                helmet_id_three_digit,
                salary,
                badge,
                cops_photo_profile_link,
                cops_photo_has_photo,
                employed_3_12_21,
                employed_12_28_20,
                employed_10_01_20,
                retired_6_1_20,
                retired_or_cert_revoked,
                retired_or_cert_revoked_date,
                hire_year,
                hire_date,
                state_cert_date,
                state_cert_level,
                rrt,
                rrt_2016,
                rrt_2018_niiya_email,
                rrt_2018,
                rrt_2019,
                rrt_2020,
                sound_truck_training_2020,
                instructed_for_dpsst,
                instructed_for_less_lethal,
                involved_in_ois_uof,
                notes
			FROM portland_fuzzy_search_officer_by_last_name_p(last_name := $1);
		`,
		lastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return portlandMarshalOfficerRows(rows)
}

func portlandMarshalOfficerRows(rows pgx.Rows) ([]*PortlandOfficer, error) {
	officers := []*PortlandOfficer{}
	for rows.Next() {
		ofc := PortlandOfficer{}
		err := rows.Scan(
            &ofc.FirstName,
            &ofc.LastName,
            &ofc.Gender,
            &ofc.OfficerRank,
            &ofc.EmployeeID,
            &ofc.HelmetID,
            &ofc.HelmetIDThreeDigit,
            &ofc.Salary,
            &ofc.Badge,
            &ofc.CopsPhotoProfileLink,
            &ofc.CopsPhotoHasPhoto,
            &ofc.Employed_3_12_21,
            &ofc.Employed_12_28_20,
            &ofc.Employed_10_01_20,
            &ofc.Retired_6_1_20,
            &ofc.RetiredOrCertRevoked,
            &ofc.RetiredOrCertRevokedDate,
            &ofc.HireYear,
            &ofc.HireDate,
            &ofc.StateCertDate,
            &ofc.StateCertLevel,
            &ofc.RRT,
            &ofc.RRT2016,
            &ofc.RRT2018NiiyaEmail,
            &ofc.RRT2018,
            &ofc.RRT2019,
            &ofc.RRT2020,
            &ofc.SoundTruckTraining,
            &ofc.InstructedForDpsst,
            &ofc.InstructedForLessLethal,
            &ofc.InvolvedInOisUof,
            &ofc.Notes,
		)

		if err != nil {
			return nil, err
		}
		officers = append(officers, &ofc)
	}
	return officers, nil
}
