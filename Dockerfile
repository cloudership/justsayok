# syntax=docker/dockerfile:1

ARG ARCH=
FROM ${ARCH}golang:1.20-alpine AS build

WORKDIR /app

COPY *.go go.mod ./

RUN go build -o /justsayok

FROM ${ARCH}scratch

ENV PORT 3000

COPY --from=build /justsayok /justsayok

USER 1000:1000

ENTRYPOINT ["/justsayok"]
