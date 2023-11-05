FROM node:18
WORKDIR /usr/src/app
COPY ./documentation/ .
RUN npm ci
RUN npm run build
CMD ["npm", "run", "serve"]