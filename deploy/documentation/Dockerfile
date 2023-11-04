FROM node:18
WORKDIR /usr/src/app
COPY ./documentation/ .
RUN npm ci
CMD ["npm", "run", "start"]