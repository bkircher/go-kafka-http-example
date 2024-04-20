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

or

    $ go version
    go version go1.22.2 darwin/arm64

Then create a new module

    $ go mod init kafka-http-example
    go: creating new go.mod: module kafka-http-example

Then we install some dependencies:

    go get -u github.com/confluentinc/confluent-kafka-go/v2/kafka
    go get -u go.uber.org/zap

Then you do

    go build

And dang, you get:
