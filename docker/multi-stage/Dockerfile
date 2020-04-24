FROM golang:1.13-alpine as build
WORKDIR /work
ADD . .
RUN go build -o app

FROM scratch
WORKDIR /
COPY --from=build /work/app .
CMD [ "/app" ]
