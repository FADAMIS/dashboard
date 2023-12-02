config:
	scripts/config.sh

docker:
	git pull
	docker compose build
	docker compose up -d