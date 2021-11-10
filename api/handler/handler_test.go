package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Handler_Ping(t *testing.T) {
	router := NewRouter(NewHandler(&MockDatabase{}))
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/ping")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status %d", res.StatusCode)
	}
	defer res.Body.Close()
	resp, _ := ioutil.ReadAll(res.Body)
	if !bytes.Equal([]byte("🏓 P O N G 🏓"), resp) {
		t.Errorf("Expected resp: 🏓 P O N G 🏓, got %s", resp)
	}
}

func Test_Handler_DescribeDepartments(t *testing.T) {
	router := NewRouter(NewHandler(&MockDatabase{}))
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/departments")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status %d", res.StatusCode)
	}
	defer res.Body.Close()
	resp, _ := ioutil.ReadAll(res.Body)

	expected := []byte(`[{"id":"spd","name":"Seattle PD","last_available_roster_date":"today","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/seattle/officer","query_params":["badge","first_name","last_name"]}}},{"id":"tpd","name":"Tacoma PD","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/tacoma/officer","query_params":["last_name"]}}},{"id":"ppb","name":"Portland PB","last_available_roster_date":"today","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/portland/officer","query_params":["badge","first_name","last_name","employee_id","helmet_id","helmet_id_three_digit"]},"fuzzy":{"path":"/portland/officer/search","query_params":["first_name","last_name"]}}},{"id":"apd","name":"Auburn PD","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/auburn/officer","query_params":["badge","first_name","last_name"]},"fuzzy":{"path":"/auburn/officer/search","query_params":["first_name","last_name"]}}},{"id":"lpd","name":"Lakewood PD","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/lakewood/officer","query_params":["first_name","last_name"]},"fuzzy":{"path":"/lakewood/officer/search","query_params":["first_name","last_name"]}}},{"id":"fakepd","name":"Renton","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"fake":{"path":"/fake","query_params":["fake","fake2"]}}},{"id":"fakepd","name":"Thurston County","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"fake":{"path":"/fake","query_params":["fake","fake2"]}}},{"id":"fakepd","name":"Bellevue","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"fake":{"path":"/fake","query_params":["fake","fake2"]}}},{"id":"fakepd","name":"Port of Seattle","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"fake":{"path":"/fake","query_params":["fake","fake2"]}}},{"id":"opd","name":"Olympia PD","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/olympia/officer","query_params":["badge","first_name","last_name"]},"fuzzy":{"path":"/olympia/officer/search","query_params":["first_name","last_name"]}}}]` + "\n")

	if !bytes.Equal(expected, resp) {
		t.Errorf("Expected resp: %s, got %s", expected, resp)
	}
}
