package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Handler_PortlandOfficerMetadata(t *testing.T) {
	router := NewRouter(NewHandler(&MockDatabase{}))
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/portland/metadata")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status %d", res.StatusCode)
	}
	defer res.Body.Close()
	resp, _ := ioutil.ReadAll(res.Body)

	expected := []byte(`{"id":"ppb","name":"Portland PB","last_available_roster_date":"today","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/portland/officer","query_params":["badge","first_name","last_name","employee_id","helmet_id","helmet_id_three_digit"]},"fuzzy":{"path":"/portland/officer/search","query_params":["first_name","last_name"]}}}` + "\n")

	if !bytes.Equal(expected, resp) {
		t.Errorf("Expected resp: %s, got %s", expected, resp)
	}
}

func Test_Handler_PortlandStrictMatch(t *testing.T) {
	t.Parallel()
	router := NewRouter(NewHandler(&MockDatabase{}))
	ts := httptest.NewServer(router)
	defer ts.Close()

	for _, tt := range [...]struct {
		name               string
		firstName          string
		lastName           string
		badge              string
		employeeId         string
		helmetId           string
		helmetIdThreeDigit string
		expectedStatus     int
		expectedBody       []byte
	}{
		{
			name:           "no params",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte("at least one of the following parameters must be provided: badge, first_name, last_name"),
		},
		{
			name:           "badge not found",
			badge:          "badge_not_found",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte("[]\n"),
		},
		{
			name:           "badge search, db error",
			badge:          "db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: get officer by badge db error"),
		},
		{
			name:           "badge search, officer found",
			badge:          "1",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"first_name":"first","last_name":"ppb","gender":null,"officer_rank":null,"employee_id":"1","helmet_id":"1","helmet_id_three_digit":"111","salary":null,"badge":"1","cops_photo_profile_link":null,"cops_photo_has_photo":null,"employed_3_12_21":null,"employed_12_28_20":null,"employed_10_01_20":null,"retired_6_1_20":null,"retired_or_cert_revoked":null,"retired_or_cert_revoked_date":null,"hire_year":null,"hire_date":null,"state_cert_date":null,"state_cert_level":null,"rrt":null,"rrt_2016":null,"rrt_2018_niiya_email":null,"rrt_2018":null,"rrt_2019":null,"rrt_2020":null,"sound_truck_training_2020":null,"instructed_for_dpsst":null,"instructed_for_less_lethal":null,"involved_in_ois_uof":null,"notes":null}]` + "\n"),
		},
		{
			name:           "employee id not found",
			employeeId:     "employee_id_not_found",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte("[]\n"),
		},
		{
			name:           "employee id search, db error",
			employeeId:     "db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: get officer by employee id db error"),
		},
		{
			name:           "employee id search, officer found",
			employeeId:     "1",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"first_name":"first","last_name":"ppb","gender":null,"officer_rank":null,"employee_id":"1","helmet_id":"1","helmet_id_three_digit":"111","salary":null,"badge":"1","cops_photo_profile_link":null,"cops_photo_has_photo":null,"employed_3_12_21":null,"employed_12_28_20":null,"employed_10_01_20":null,"retired_6_1_20":null,"retired_or_cert_revoked":null,"retired_or_cert_revoked_date":null,"hire_year":null,"hire_date":null,"state_cert_date":null,"state_cert_level":null,"rrt":null,"rrt_2016":null,"rrt_2018_niiya_email":null,"rrt_2018":null,"rrt_2019":null,"rrt_2020":null,"sound_truck_training_2020":null,"instructed_for_dpsst":null,"instructed_for_less_lethal":null,"involved_in_ois_uof":null,"notes":null}]` + "\n"),
		},
		{
			name:           "helmet id not found",
			helmetId:       "helmet_id_not_found",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte("[]\n"),
		},
		{
			name:           "helmet id search, db error",
			helmetId:       "db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: get officer by helmet id db error"),
		},
		{
			name:           "helmet id search, officer found",
			helmetId:       "1",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"first_name":"first","last_name":"ppb","gender":null,"officer_rank":null,"employee_id":"1","helmet_id":"1","helmet_id_three_digit":"111","salary":null,"badge":"1","cops_photo_profile_link":null,"cops_photo_has_photo":null,"employed_3_12_21":null,"employed_12_28_20":null,"employed_10_01_20":null,"retired_6_1_20":null,"retired_or_cert_revoked":null,"retired_or_cert_revoked_date":null,"hire_year":null,"hire_date":null,"state_cert_date":null,"state_cert_level":null,"rrt":null,"rrt_2016":null,"rrt_2018_niiya_email":null,"rrt_2018":null,"rrt_2019":null,"rrt_2020":null,"sound_truck_training_2020":null,"instructed_for_dpsst":null,"instructed_for_less_lethal":null,"involved_in_ois_uof":null,"notes":null}]` + "\n"),
		},
		{
			name:               "helmet id three digits not found",
			helmetIdThreeDigit: "helmet_id_three_digit_not_found",
			expectedStatus:     http.StatusOK,
			expectedBody:       []byte("[]\n"),
		},
		{
			name:               "helmet id three digits search, db error",
			helmetIdThreeDigit: "db_error",
			expectedStatus:     http.StatusInternalServerError,
			expectedBody:       []byte("error getting officer: get officer by helmet id three digits db error"),
		},
		{
			name:               "helmet id three digits search, officer found",
			helmetIdThreeDigit: "111",
			expectedStatus:     http.StatusOK,
			expectedBody:       []byte(`[{"first_name":"first","last_name":"ppb","gender":null,"officer_rank":null,"employee_id":"1","helmet_id":"1","helmet_id_three_digit":"111","salary":null,"badge":"1","cops_photo_profile_link":null,"cops_photo_has_photo":null,"employed_3_12_21":null,"employed_12_28_20":null,"employed_10_01_20":null,"retired_6_1_20":null,"retired_or_cert_revoked":null,"retired_or_cert_revoked_date":null,"hire_year":null,"hire_date":null,"state_cert_date":null,"state_cert_level":null,"rrt":null,"rrt_2016":null,"rrt_2018_niiya_email":null,"rrt_2018":null,"rrt_2019":null,"rrt_2020":null,"sound_truck_training_2020":null,"instructed_for_dpsst":null,"instructed_for_less_lethal":null,"involved_in_ois_uof":null,"notes":null}]` + "\n"),
		},
		{
			name:           "name search, db error",
			firstName:      "db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: get officer by name db error"),
		},
		{
			name:           "name search, officers found",
			lastName:       "test",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"first_name":"first","last_name":"poo","gender":null,"officer_rank":null,"employee_id":"2","helmet_id":"1","helmet_id_three_digit":"222","salary":null,"badge":"2","cops_photo_profile_link":null,"cops_photo_has_photo":null,"employed_3_12_21":null,"employed_12_28_20":null,"employed_10_01_20":null,"retired_6_1_20":null,"retired_or_cert_revoked":null,"retired_or_cert_revoked_date":null,"hire_year":null,"hire_date":null,"state_cert_date":null,"state_cert_level":null,"rrt":null,"rrt_2016":null,"rrt_2018_niiya_email":null,"rrt_2018":null,"rrt_2019":null,"rrt_2020":null,"sound_truck_training_2020":null,"instructed_for_dpsst":null,"instructed_for_less_lethal":null,"involved_in_ois_uof":null,"notes":null},{"first_name":"test","last_name":"poo","gender":null,"officer_rank":null,"employee_id":"3","helmet_id":"1","helmet_id_three_digit":"333","salary":null,"badge":"3","cops_photo_profile_link":null,"cops_photo_has_photo":null,"employed_3_12_21":null,"employed_12_28_20":null,"employed_10_01_20":null,"retired_6_1_20":null,"retired_or_cert_revoked":null,"retired_or_cert_revoked_date":null,"hire_year":null,"hire_date":null,"state_cert_date":null,"state_cert_level":null,"rrt":null,"rrt_2016":null,"rrt_2018_niiya_email":null,"rrt_2018":null,"rrt_2019":null,"rrt_2020":null,"sound_truck_training_2020":null,"instructed_for_dpsst":null,"instructed_for_less_lethal":null,"involved_in_ois_uof":null,"notes":null},{"first_name":"first","last_name":"ppb","gender":null,"officer_rank":null,"employee_id":"1","helmet_id":"1","helmet_id_three_digit":"111","salary":null,"badge":"1","cops_photo_profile_link":null,"cops_photo_has_photo":null,"employed_3_12_21":null,"employed_12_28_20":null,"employed_10_01_20":null,"retired_6_1_20":null,"retired_or_cert_revoked":null,"retired_or_cert_revoked_date":null,"hire_year":null,"hire_date":null,"state_cert_date":null,"state_cert_level":null,"rrt":null,"rrt_2016":null,"rrt_2018_niiya_email":null,"rrt_2018":null,"rrt_2019":null,"rrt_2020":null,"sound_truck_training_2020":null,"instructed_for_dpsst":null,"instructed_for_less_lethal":null,"involved_in_ois_uof":null,"notes":null}]` + "\n"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/portland/officer?badge=%s&first_name=%s&last_name=%s&employee_id=%s&helmet_id=%s&helmet_id_three_digit=%s", ts.URL, tt.badge, tt.firstName, tt.lastName, tt.employeeId, tt.helmetId, tt.helmetIdThreeDigit))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)
			if !bytes.Equal(tt.expectedBody, resp) {
				t.Errorf("Expected resp %s; got %s", tt.expectedBody, resp)
			}
		})
	}
}

func Test_Handler_PortlandFuzzySearch(t *testing.T) {
	t.Parallel()
	router := NewRouter(NewHandler(&MockDatabase{}))
	ts := httptest.NewServer(router)
	defer ts.Close()

	for _, tt := range [...]struct {
		name           string
		firstName      string
		lastName       string
		expectedStatus int
		expectedBody   []byte
	}{
		{
			name:           "no params",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte("at least one of the following parameters must be provided: first_name, last_name"),
		},
		{
			name:           "name fuzzy search, db error",
			firstName:      "db",
			lastName:       "error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: fuzzy search by name db error"),
		},
		{
			name:           "name fuzzy search, officers found",
			firstName:      "test",
			lastName:       "test",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"first_name":"first","last_name":"ppb","gender":null,"officer_rank":null,"employee_id":"1","helmet_id":"1","helmet_id_three_digit":"111","salary":null,"badge":"1","cops_photo_profile_link":null,"cops_photo_has_photo":null,"employed_3_12_21":null,"employed_12_28_20":null,"employed_10_01_20":null,"retired_6_1_20":null,"retired_or_cert_revoked":null,"retired_or_cert_revoked_date":null,"hire_year":null,"hire_date":null,"state_cert_date":null,"state_cert_level":null,"rrt":null,"rrt_2016":null,"rrt_2018_niiya_email":null,"rrt_2018":null,"rrt_2019":null,"rrt_2020":null,"sound_truck_training_2020":null,"instructed_for_dpsst":null,"instructed_for_less_lethal":null,"involved_in_ois_uof":null,"notes":null}]` + "\n"),
		},
		{
			name:           "first name fuzzy search, db error",
			firstName:      "db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: fuzzy search by first name db error"),
		},
		{
			name:           "first name fuzzy search, officers found",
			firstName:      "test",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"first_name":"first","last_name":"ppb","gender":null,"officer_rank":null,"employee_id":"1","helmet_id":"1","helmet_id_three_digit":"111","salary":null,"badge":"1","cops_photo_profile_link":null,"cops_photo_has_photo":null,"employed_3_12_21":null,"employed_12_28_20":null,"employed_10_01_20":null,"retired_6_1_20":null,"retired_or_cert_revoked":null,"retired_or_cert_revoked_date":null,"hire_year":null,"hire_date":null,"state_cert_date":null,"state_cert_level":null,"rrt":null,"rrt_2016":null,"rrt_2018_niiya_email":null,"rrt_2018":null,"rrt_2019":null,"rrt_2020":null,"sound_truck_training_2020":null,"instructed_for_dpsst":null,"instructed_for_less_lethal":null,"involved_in_ois_uof":null,"notes":null}]` + "\n"),
		},
		{
			name:           "last name fuzzy search, db error",
			lastName:       "db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: fuzzy search by last name db error"),
		},
		{
			name:           "last name fuzzy search, officers found",
			lastName:       "test",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"first_name":"first","last_name":"ppb","gender":null,"officer_rank":null,"employee_id":"1","helmet_id":"1","helmet_id_three_digit":"111","salary":null,"badge":"1","cops_photo_profile_link":null,"cops_photo_has_photo":null,"employed_3_12_21":null,"employed_12_28_20":null,"employed_10_01_20":null,"retired_6_1_20":null,"retired_or_cert_revoked":null,"retired_or_cert_revoked_date":null,"hire_year":null,"hire_date":null,"state_cert_date":null,"state_cert_level":null,"rrt":null,"rrt_2016":null,"rrt_2018_niiya_email":null,"rrt_2018":null,"rrt_2019":null,"rrt_2020":null,"sound_truck_training_2020":null,"instructed_for_dpsst":null,"instructed_for_less_lethal":null,"involved_in_ois_uof":null,"notes":null}]` + "\n"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/portland/officer/search?first_name=%s&last_name=%s", ts.URL, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)
			if !bytes.Equal(tt.expectedBody, resp) {
				t.Errorf("Expected resp %s; got %s", tt.expectedBody, resp)
			}
		})
	}
}
