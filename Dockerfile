FROM golang:1.15-alpine AS build

COPY /pkg/ /src/
RUN apk add --no-cache git
WORKDIR /src/
RUN ls -lah
RUN CGO_ENABLED=0 go build -o /bin/subgraph_monitoring

FROM ubuntu:latest
RUN apt-get update && apt-get install ca-certificates -y
COPY --from=build /bin/subgraph_monitoring /bin/subgraph_monitoring
ENTRYPOINT ["/bin/subgraph_monitoring"]
EXPOSE 2112

