package data

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DepartmentMetadata is the structure of the metadata returned by the SQL DB
type DepartmentMetadata struct {
	ID                      string                          `json:"id"`
	Name                    string                          `json:"name"`
	LastAvailableRosterDate string                          `json:"last_available_roster_date"`
	Fields                  []map[string]string             `json:"fields"`
	SearchRoutes            map[string]*SearchRouteMetadata `json:"search_routes"`
}

// SearchRouteMetadata describes the search routes of a deparment
type SearchRouteMetadata struct {
	Path        string   `json:"path"`
	QueryParams []string `json:"query_params"`
}

// DatabaseInterface describes database functions
type DatabaseInterface interface {
	SeattleOfficerMetadata() *DepartmentMetadata
	SeattleGetOfficerByBadge(badge string) ([]*SeattleOfficer, error)
	SeattleSearchOfficerByName(firstName, lastName string) ([]*SeattleOfficer, error)
	SeattleFuzzySearchByName(name string) ([]*SeattleOfficer, error)
	SeattleFuzzySearchByFirstName(firstName string) ([]*SeattleOfficer, error)
	SeattleFuzzySearchByLastName(lastName string) ([]*SeattleOfficer, error)

	SeattleGetOfficerByBadgeHistorical(badge string) ([]*SeattleOfficer, error)

	TacomaOfficerMetadata() *DepartmentMetadata
	TacomaSearchOfficerByName(firstName, lastName string) ([]*TacomaOfficer, error)
	TacomaFuzzySearchByName(name string) ([]*TacomaOfficer, error)
	TacomaFuzzySearchByFirstName(firstName string) ([]*TacomaOfficer, error)
	TacomaFuzzySearchByLastName(lastName string) ([]*TacomaOfficer, error)

	PortlandOfficerMetadata() *DepartmentMetadata
	PortlandSearchOfficersByBadge(badge string) ([]*PortlandOfficer, error)
	PortlandSearchOfficersByEmployeeId(employee_id string) ([]*PortlandOfficer, error)
	PortlandSearchOfficersByHelmetId(helmet_id string) ([]*PortlandOfficer, error)
	PortlandSearchOfficersByHelmetIdThreeDigit(helmet_id_three_digit string) ([]*PortlandOfficer, error)
	PortlandSearchOfficersByName(firstName, lastName string) ([]*PortlandOfficer, error)
	PortlandFuzzySearchByName(name string) ([]*PortlandOfficer, error)
	PortlandFuzzySearchByFirstName(firstName string) ([]*PortlandOfficer, error)
	PortlandFuzzySearchByLastName(lastName string) ([]*PortlandOfficer, error)

	AuburnOfficerMetadata() *DepartmentMetadata
	AuburnGetOfficerByBadge(badge string) ([]*AuburnOfficer, error)
	AuburnSearchOfficerByName(firstName, lastName string) ([]*AuburnOfficer, error)
	AuburnFuzzySearchByName(name string) ([]*AuburnOfficer, error)
	AuburnFuzzySearchByFirstName(firstName string) ([]*AuburnOfficer, error)
	AuburnFuzzySearchByLastName(lastName string) ([]*AuburnOfficer, error)

	LakewoodOfficerMetadata() *DepartmentMetadata
	LakewoodSearchOfficerByName(firstName, lastName string) ([]*LakewoodOfficer, error)
	LakewoodFuzzySearchByName(name string) ([]*LakewoodOfficer, error)
	LakewoodFuzzySearchByFirstName(firstName string) ([]*LakewoodOfficer, error)
	LakewoodFuzzySearchByLastName(lastName string) ([]*LakewoodOfficer, error)
}

// Client is the client used to connect to the db
type Client struct {
	pool *pgxpool.Pool
}

// NewClient is the constructor for Client
func NewClient(username, password, host, dbName string) *Client {
	var pool *pgxpool.Pool

	pgxConfig, err := pgxpool.ParseConfig(fmt.Sprintf("postgresql://%s", host))
	if err != nil {
		log.Panic(fmt.Sprintf("Unable to create postgres connection config: %v", err))
	}

	pgxConfig.ConnConfig.Port = 5432
	pgxConfig.ConnConfig.Database = dbName
	pgxConfig.ConnConfig.User = username
	pgxConfig.ConnConfig.Password = password

	pool, err = pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Panic(fmt.Sprintf("Unable to create db connection: %v", err))
	}

	return &Client{pool: pool}
}
