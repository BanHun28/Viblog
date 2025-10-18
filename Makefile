.PHONY: help dev build up down clean

help:
	@echo "Available commands:"
	@echo "  make dev          - Start development environment"
	@echo "  make build        - Build all services"
	@echo "  make up           - Start all services"
	@echo "  make down         - Stop all services"
	@echo "  make clean        - Remove all containers and volumes"

dev:
	docker-compose up -d

build:
	docker-compose build

up:
	docker-compose -f docker-compose.prod.yml up -d

down:
	docker-compose down

clean:
	docker-compose down -v --rmi local
	docker-compose -f docker-compose.prod.yml down -v --rmi local
