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

Wait until the broker has finished starting. Then create a topic:

    $ docker exec -it broker bash -c '/bin/kafka-topics --bootstrap-server broker:29092 --create --topic my-topic --partitions 3 --replication-factor 1 --config retention.ms=3600000'
    Created topic my-topic.

Run our little consumer program:

    ./kafka-http-example

And start produce some messages in a second terminal:

    $ docker exec -it broker bash -c '/bin/kafka-console-producer --broker-list broker:29092 --topic my-topic'
    >Hello
    >Say something nice

Which gives us:

    $ ./kafka-http-example
    {"level":"info","ts":1714407627.183806,"caller":"go-kafka-http-example/main.go:40","msg":"Received message","topic":"my-topic","key":"","value":"Hello","partition":1,"offset":0}
    {"level":"info","ts":1714407631.0890791,"caller":"go-kafka-http-example/main.go:40","msg":"Received message","topic":"my-topic","key":"","value":"Say something nice","partition":1,"offset":1}

## Links

- Kafka Go Client library documentation:
  [https://docs.confluent.io/kafka-clients/go/current/overview.html](https://docs.confluent.io/kafka-clients/go/current/overview.html)
