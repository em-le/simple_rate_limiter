PROJECT:=rate-limiter

export GO111MODULE=on

HAS_DOCKER_COMPOSE_WITH_DASH:=$(shell which docker-compose)
ifdef HAS_DOCKER_COMPOSE_WITH_DASH
	DOCKER_COMPOSE=docker-compose
else
	DOCKER_COMPOSE=docker compose
endif

export LIMITER=LEAKY_BUCKET_2

run:
	 go run main.go
run_locust:
	locust
start_local:
	$(DOCKER_COMPOSE) -f docker-compose.yaml up -d --wait

