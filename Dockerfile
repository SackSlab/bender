FROM golang:1.14.3-alpine AS build
WORKDIR /usr/src/app

# statically compile binary
ENV CGO_ENABLED=0

COPY . .
RUN apk update; \
    apk add --virtual build-dependencies build-base

RUN make build-prod

FROM scratch
WORKDIR /usr/src/app

ARG APP_PORT

COPY --from=build /usr/src/app/bin/prod/bender .

EXPOSE ${APP_PORT}
CMD [ "./bender" ]
