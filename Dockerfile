FROM golang:alpine

COPY . .

RUN go mod download
RUN go build

# Running server
CMD [ "./profanity" ]
