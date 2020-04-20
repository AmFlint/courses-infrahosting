FROM node:12.10.0

WORKDIR /home/app

COPY package.* .

RUN npm i

COPY . .

CMD [ "node", "server.js" ]
