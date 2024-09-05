FROM node:lts-alpine

RUN yarn global add serve

WORKDIR /app

COPY ./frontend/package.json ./
COPY ./frontend/yarn.lock ./

RUN yarn install

COPY ./frontend/ .

RUN yarn build

CMD ["serve", "-s", "dist", "-l", "tcp://0.0.0.0:10034" ]
