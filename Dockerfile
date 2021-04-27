# FROM golang:1.12 as builder
# WORKDIR /module
# COPY go.mod /module/go.mod
# COPY go.sum /module/go.sum
# COPY .env /.env
# RUN go mod download

# COPY . /module
# RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./main.go

# FROM alpine
# RUN apk add tzdata && \
#     cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
#     echo "Asia/Bangkok" >  /etc/timezone && \
#     apk del tzdata
# WORKDIR /root/
# COPY --from=builder /module/bin .
# ENV GIN_MODE release
# EXPOSE 3000
# CMD ./app

FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run" , "main.go"]


#nodemon --exec go run main.go --signal SIGTERM