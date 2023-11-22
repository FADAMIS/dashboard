FROM node:21-bookworm

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build

EXPOSE 3000

CMD node build/index.js