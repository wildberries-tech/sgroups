FROM node:18 as builder
WORKDIR /usr/src/app
COPY ./documentation/ .
RUN npm ci
RUN npm run build

FROM node:18
WORKDIR /usr/src/swarm_doc
COPY --from=builder /usr/src/app/build .
COPY --from=builder /usr/src/app/node_modules ./node_modules
COPY --from=builder /usr/src/app/package.json ./package.json
CMD ["npm", "run", "serve"]