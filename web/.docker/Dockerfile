FROM node:21-alpine3.17 as builder

WORKDIR /app

COPY ./web/package.json ./
COPY ./web/package-lock.json ./

RUN npm install

#RUN npm install --save nuxt

RUN npm install nuxt

COPY ./web .

RUN npm run build

EXPOSE 3000

CMD node ./.output/server/index.mjs