# gobase layer
FROM golang:alpine AS go-base
ADD . /src
RUN cd /src && go build main.go

# final layer
FROM alpine
COPY --from=go-base /src/main /app/
EXPOSE 8090
ENTRYPOINT ["./app/main"]