FROM node:latest

COPY . /var/app

WORKDIR /var/app

RUN npm install simple-json-replay-server

EXPOSE 8008

CMD ["node", "/var/app/node_modules/simple-json-replay-server/src/main.js"]
