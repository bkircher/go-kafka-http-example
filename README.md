# go-kafka-http-example

Small microservice in Go that reads from a Kafka topic and speaks HTTP to
another service.

## Dependencies

- Confluent's Apache Kafka Go client:
  <https://github.com/confluentinc/confluent-kafka-go>
- Uber's fast, structured, leveled Go logger: <https://github.com/uber-go/zap>

## Steps

Make sure you have go installed:

    brew install go

    $ go version
    go version go1.22.2 darwin/arm64

Then create a new module:

    $ mkdir mkdir go-kafka-http-example
    $ cd go-kafka-http-example
    $ go mod init kafka-http-example
    go: creating new go.mod: module kafka-http-example

Then install some dependencies:

    go get -u github.com/confluentinc/confluent-kafka-go/v2/kafka
    go get -u go.uber.org/zap

With this and the code in `main.go` you already have set everything up that is
needed:

    $ ls -1
    go.mod
    go.sum
    main.go

Then you do

    go build

And dang, you get a binary.

Alright, lets add a Kafka broker to the story.

    docker-compose up

And produce message while reading them.
