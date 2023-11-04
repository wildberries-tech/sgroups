FROM node:18 as builder
WORKDIR /usr/src/app
COPY ./documentation/ .
RUN npm ci
CMD ["npm", "run", "start"]

# RUN npm run build

# FROM node:18
# WORKDIR /usr/src/swarm_doc
# COPY --from=builder /usr/src/app/build .
# CMD ["sleep", "infinity"]