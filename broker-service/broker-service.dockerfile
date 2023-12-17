# Comment this lines until the FROM alpine:latest
# # base go image
# FROM golang:1.21.5-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# RUN chmod +x /app/brokerApp

# # build a small docker image
FROM alpine:latest

RUN mkdir /app

# COPY --from=builder /app/brokerApp /app 
# FOR WINDOWS

COPY brokerApp /app

CMD [ "/app/brokerApp" ]