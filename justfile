set dotenv-load := false
IS_PROD := env_var_or_default("IS_PROD", "")
COMPOSE_FILE := "--file=docker-compose.yml" + (
    if IS_PROD != "true" {" --file=docker-compose.override.yml"} else {""}
)
DC := "docker-compose " + COMPOSE_FILE
INTEGRATION_TEST_PATH := "./integration"


# Show all available recipes
default:
    @just --list --unsorted

# Create the .env file from the template
dotenv:
    @([ ! -f .env ] && cp .env.example .env) || true

# Build the containers
build: dotenv
	{{ DC }} build

# Spin up all (or one) service
up service="": dotenv
	{{ DC }} up -d {{ service }}

# Tear down containers
down:
	{{ DC }} down

# Pull all docker images
pull:
    {{ DC }} pull

@test-int test_path="":
    echo "Starting up spd-lookup services for integration testing"
    {{ DC }} up -d &> /dev/null
    ( while [ "$(docker logs -n 1 $(docker ps | grep spd_lookup_api | awk -F ' ' '{print $1}'))" != "starting server on port 1312" ] ; do sleep 1s; done ) &> /dev/null
    echo "spd-lookup services started, beginning integration testing"
    echo "***************************************"
    go test -v -tags=integrations {{ INTEGRATION_TEST_PATH }} -count=1 -run={{ test_path }}
    echo "spd-lookup integration testing completed, removing services"
    echo "***************************************"
    {{ DC }} down &> /dev/null
