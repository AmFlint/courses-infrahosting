# First Step - Build project FROM JSX to HTML and JS With Webpack
FROM node:10 as build

WORKDIR /home/node

COPY package*.json ./

RUN npm install

COPY . ./

ARG REACT_APP_BACKEND_URL

RUN npm run build

# Second Step - Server website

FROM nginx:alpine

COPY --from=build /home/node/build /usr/share/nginx/html
