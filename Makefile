run:
	air -c .air.toml

up:
	docker-compose up --force-recreate --remove-orphans
