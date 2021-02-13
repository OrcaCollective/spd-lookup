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

	expected := []byte(`[{"id":"spd","name":"Seattle PD","last_available_roster_date":"today","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/seattle/officer","query_params":["badge","first_name","last_name"]}}},{"id":"tpd","name":"Tacoma PD","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/tacoma/officer","query_params":["last_name"]}}}]` + "\n")

	if !bytes.Equal(expected, resp) {
		t.Errorf("Expected resp: %s, got %s", expected, resp)
	}
}
