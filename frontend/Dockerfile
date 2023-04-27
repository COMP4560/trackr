# syntax=docker/dockerfile:1

FROM node:19-alpine

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
COPY .env.production ./
RUN npm install
RUN npm install -g serve

COPY . ./

RUN npm run build

EXPOSE 3000

CMD ["serve", "-s", "build"]
