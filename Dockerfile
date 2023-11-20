FROM golang:1.21.3 AS dashboard
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o dashboard .

FROM node:21-bookworm

COPY --from=dashboard /app/dashboard /bin

WORKDIR /app/frontend
COPY frontend .

RUN npm install
RUN npm run build

RUN apt update -y
RUN apt install -y debian-keyring debian-archive-keyring apt-transport-https
RUN curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg
RUN curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | tee /etc/apt/sources.list.d/caddy-stable.list
RUN apt update -y
RUN apt install -y caddy systemctl
RUN systemctl enable --now caddy

COPY services/Caddyfile /etc/caddy/Caddyfile
COPY services/dashboard.service /etc/systemd/system/dashboard.service

RUN systemctl enable --now dashboard

EXPOSE 80
EXPOSE 443

CMD systemctl start dashboard && systemctl start caddy && node build/index.js