FROM golang:1.20rc3-alpine3.17 AS build
RUN apk update && apk add -q make
COPY . /airportdetector
WORKDIR /airportdetector

RUN make clean build

FROM scratch
WORKDIR /app
COPY --from=build /airportdetector/bin/airport_detector /app/airport_detector
COPY --from=build /airportdetector/api/* /app

CMD ["./airport_detector"]