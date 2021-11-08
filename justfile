set dotenv-load := false
IS_PROD := env_var_or_default("IS_PROD", "")
COMPOSE_FILE := "--file=docker-compose.yml" + (
    if IS_PROD != "true" {" --file=docker-compose.override.yml"} else {""}
)
DC := "docker-compose " + COMPOSE_FILE


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
