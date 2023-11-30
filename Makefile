dashboard:
	cp scripts/dashboard.service /etc/systemd/system
	systemctl enable dashboard
	systemctl start dashboard

docker:
	git pull
	docker compose build
	docker compose up