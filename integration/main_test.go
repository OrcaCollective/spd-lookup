package integration

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// Default test server
var testServer = "http://localhost:1312"

// validateFunc are for subtests that share a single setup
type validateFunc func(context.Context, *testing.T, string)

// Timeout value for integration tests
const ctxTimeout = 30 * time.Second

// Main integration test runner
func TestIntegrations(t *testing.T) {
	profile := "integration"
	ctx, _ := context.WithTimeout(context.Background(), ctxTimeout)

	// Parallel tests
	t.Run("parallel", func(t *testing.T) {
		tests := []struct {
			name      string
			validator validateFunc
		}{
			{"TestPing", testPing},
			{"TestDepartments", testDepartments},
			{"TestSeattleStrict", testSeattleStrict},
			{"TestSeattleFuzzy", testSeattleFuzzy},
			{"TestSeattleHistorical", testSeattleHistorical},
		}
		for _, tc := range tests {
			tc := tc
			if ctx.Err() == context.DeadlineExceeded {
				t.Fatalf("Unable to run more tests (deadline exceeded)")
			}
			t.Run(tc.name, func(t *testing.T) {
				tc.validator(ctx, t, profile)
			})
		}
	})
}

// Test ping endpoint
func testPing(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	expectedResponse := []byte(`🏓 P O N G 🏓`)
	resp, err := http.Get(testServer + "/ping")
	if err != nil {
		t.Errorf("Unspecified error with request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Unspecified error with reading response: %v", err)
	} else if !bytes.Equal(expectedResponse, body) {
		t.Errorf("Ping failed, got:%s, want:%s", body, expectedResponse)
	}
}

// Test departments endpoint
func testDepartments(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	expectedResponse := []byte(`[{"id":"spd","name":"Seattle PD","last_available_roster_date":"2021-10-26","fields":[{"FieldName":"date","Label":"Roster Date"},{"FieldName":"badge","Label":"Badge"},{"FieldName":"first_name","Label":"First Name"},{"FieldName":"middle_name","Label":"Middle Name"},{"FieldName":"last_name","Label":"Last Name"},{"FieldName":"title","Label":"Title"},{"FieldName":"unit","Label":"Unit"},{"FieldName":"unit_description","Label":"Unit Description"},{"FieldName":"full_name","Label":"Full Name"},{"FieldName":"is_current","Label":"On Current Roster"}],"search_routes":{"exact":{"path":"/seattle/officer","query_params":["badge","first_name","last_name"]},"fuzzy":{"path":"/seattle/officer/search","query_params":["first_name","last_name"]},"historical-exact":{"path":"/seattle/officer/historical","query_params":["badge"]}}},{"id":"tpd","name":"Tacoma PD","last_available_roster_date":"2019","fields":[{"FieldName":"first_name","Label":"First Name"},{"FieldName":"last_name","Label":"Last Name"},{"FieldName":"title","Label":"Title"},{"FieldName":"department","Label":"Department"},{"FieldName":"salary","Label":"Salary 2019"}],"search_routes":{"exact":{"path":"/tacoma/officer","query_params":["first_name","last_name"]},"fuzzy":{"path":"/tacoma/officer/search","query_params":["first_name","last_name"]}}},{"id":"ppb","name":"Portland PB","last_available_roster_date":"2021-03-12","fields":[{"FieldName":"first_name","Label":"First Name"},{"FieldName":"last_name","Label":"Last Name"},{"FieldName":"gender","Label":"Gender"},{"FieldName":"officer_rank","Label":"Rank"},{"FieldName":"employee_id","Label":"Employee (Chest) ID"},{"FieldName":"helmet_id","Label":"Helmet #"},{"FieldName":"helmet_id_three_digit","Label":"3-Digit Helmet #"},{"FieldName":"salary","Label":"Fiscal Earnings 2019"},{"FieldName":"badge","Label":"Badge/DPSST Number"},{"FieldName":"cops_photo_profile_link","Label":"Cops.Photo Profile Link"},{"FieldName":"cops_photo_has_photo","Label":"Pic on Cops.photo (y/n)"},{"FieldName":"employed_3_12_21","Label":"Employed as of 3/12/21"},{"FieldName":"employed_12_28_20","Label":"Employed as of 12/28/20"},{"FieldName":"employed_10_01_20","Label":"Employed as of 10/01/20"},{"FieldName":"retired_6_1_20","Label":"Retired/Resigned as of 6/1/20"},{"FieldName":"retired_or_cert_revoked","Label":"Retired/Resigned as of 6/1/20 OR Cert Revoked (ever)"},{"FieldName":"retired_or_cert_revoked_date","Label":"Date of Cert Revoke"},{"FieldName":"hire_year","Label":"Hire Year"},{"FieldName":"hire_date","Label":"Hire Date"},{"FieldName":"state_cert_date","Label":"State Certification Date"},{"FieldName":"state_cert_level","Label":"State Certification Level"},{"FieldName":"rrt","Label":"RRT (Rapid Response Team) Member"},{"FieldName":"rrt_2016","Label":"RRT member as of 2016 via 2017 PPB AR"},{"FieldName":"rrt_2018_niiya_email","Label":"RRT member as of 2018 via Niiya Email"},{"FieldName":"rrt_2018","Label":"RRT Specific Training 2018"},{"FieldName":"rrt_2019","Label":"RRT Specific Training 2019"},{"FieldName":"rrt_2020","Label":"RRT Specific Training 2020"},{"FieldName":"sound_truck_training_2020","Label":"Sound Truck Training 2020"},{"FieldName":"instructed_for_dpsst","Label":"Has Instructed Course for DPSST 2017+"},{"FieldName":"instructed_for_less_lethal","Label":"Instructor for Less Lethal/Chemical Weapons Courses"},{"FieldName":"involved_in_ois_uof","Label":"Has Been Involved in OIS/Significant UoF Incident"},{"FieldName":"notes","Label":"Notes"}],"search_routes":{"exact":{"path":"/portland/officer","query_params":["badge","first_name","last_name","employee_id","helmet_id","helmet_id_three_digit"]},"fuzzy":{"path":"/portland/officer/search","query_params":["first_name","last_name"]}}},{"id":"apd","name":"Auburn PD","last_available_roster_date":"2021-06-07","fields":[{"FieldName":"date","Label":"Roster Date"},{"FieldName":"badge","Label":"Badge"},{"FieldName":"first_name","Label":"First Name"},{"FieldName":"last_name","Label":"Last Name"},{"FieldName":"title","Label":"Title"}],"search_routes":{"exact":{"path":"/auburn/officer","query_params":["badge","first_name","last_name"]},"fuzzy":{"path":"/auburn/officer/search","query_params":["first_name","last_name"]}}},{"id":"lpd","name":"Lakewood PD","last_available_roster_date":"2021-05-01","fields":[{"FieldName":"date","Label":"Roster Date"},{"FieldName":"title","Label":"Title"},{"FieldName":"last_name","Label":"Last Name"},{"FieldName":"first_name","Label":"First Name"},{"FieldName":"unit","Label":"Unit"},{"FieldName":"unit_descritpion","Label":"Unit Description"}],"search_routes":{"exact":{"path":"/lakewood/officer","query_params":["first_name","last_name"]},"fuzzy":{"path":"/lakewood/officer/search","query_params":["first_name","last_name"]}}},{"id":"rpd","name":"Renton PD","last_available_roster_date":"2021-05-01","fields":[{"FieldName":"last_name","Label":"Last Name"},{"FieldName":"first_name","Label":"First Name"},{"FieldName":"middle_name","Label":"Middle Name"},{"FieldName":"rank","Label":"Officer Rank"},{"FieldName":"department","Label":"Officer Department"},{"FieldName":"division","Label":"Officer Division"},{"FieldName":"shift","Label":"Shift"},{"FieldName":"additional_info","Label":"additional information (including retirement date)"},{"FieldName":"badge","Label":"Badge number"}],"search_routes":{"exact":{"path":"/renton/officer","query_params":["first_name","last_name"]},"fuzzy":{"path":"/renton/officer/search","query_params":["first_name","last_name"]}}},{"id":"tcsd","name":"Thurston County Sheriff's Department","last_available_roster_date":"2021-05-01","fields":[{"FieldName":"last_name","Label":"Last Name"},{"FieldName":"first_name","Label":"First Name"},{"FieldName":"title","Label":"Officer Title"},{"FieldName":"call_sign","Label":"Call Sign"}],"search_routes":{"exact":{"path":"/thurston_county/officer","query_params":["first_name","last_name"]},"fuzzy":{"path":"/thurston_county/officer/search","query_params":["first_name","last_name"]}}},{"id":"bpd","name":"Bellevue PD","last_available_roster_date":"2021-05-01","fields":[{"FieldName":"last_name","Label":"Last Name"},{"FieldName":"first_name","Label":"First Name"},{"FieldName":"title","Label":"Officer Title"},{"FieldName":"unit","Label":"Officer unit"},{"FieldName":"notes","Label":"additional information (including retirement date)"},{"FieldName":"badge","Label":"Badge number"}],"search_routes":{"exact":{"path":"/bellevue/officer","query_params":["badge","first_name","last_name"]},"fuzzy":{"path":"/bellevue/officer/search","query_params":["first_name","last_name"]}}},{"id":"pospd","name":"Port Of Seattle PD","last_available_roster_date":"2021-05-01","fields":[{"FieldName":"name","Label":"Full Name"},{"FieldName":"rank","Label":"Officer Title"},{"FieldName":"unit","Label":"Officer unit"},{"FieldName":"badge","Label":"Badge number"}],"search_routes":{"exact":{"path":"/port_of_seattle/officer","query_params":["badge","name"]},"fuzzy":{"path":"/port_of_seattle/officer/search","query_params":["name"]}}},{"id":"opd","name":"Olympia PD","last_available_roster_date":"2021-05-01","fields":[{"FieldName":"date","Label":"Roster Date"},{"FieldName":"first_name","Label":"First Name"},{"FieldName":"last_name","Label":"Last Name"},{"FieldName":"title","Label":"Title"},{"FieldName":"unit","Label":"Unit"},{"FieldName":"badge","Label":"Badge"}],"search_routes":{"exact":{"path":"/olympia/officer","query_params":["badge","first_name","last_name"]},"fuzzy":{"path":"/olympia/officer/search","query_params":["first_name","last_name"]}}}]` + "\n")
	resp, err := http.Get(testServer + "/departments")
	if err != nil {
		t.Errorf("Unspecified error with request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Unspecified error with reading response: %v", err)
	} else if !bytes.Equal(expectedResponse, body) {
		t.Logf("Length expected: %d", len(expectedResponse))
		t.Logf("Length received: %d", len(body))
		t.Errorf("Departments failed, got:%s, want:%s", body, expectedResponse)
	}
}
