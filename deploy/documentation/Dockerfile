FROM node:20
WORKDIR /usr/src/app
COPY ./documentation/ .
RUN npm i
RUN npm run build
CMD ["npm", "run", "serve"]