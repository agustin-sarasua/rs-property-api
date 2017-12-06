FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

# Make port available to the world outside this container
EXPOSE 8080

# Define env variable
ENV NAME World

CMD ["go-wrapper", "run"] # ["app"]
