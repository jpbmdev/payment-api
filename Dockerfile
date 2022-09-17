FROM golang:alpine

WORKDIR /go/src/github.com/jpbmdev/payment-api

COPY . .

RUN go build .

EXPOSE 8080

ENV DB_CONNECTION_STRING=mongodb://testUser:testUser@mongodb:27017    
ENV DB_NAME=paymentDB    
ENV PORT=:8080    

ENTRYPOINT ["/go/src/github.com/jpbmdev/payment-api/payment-api"]