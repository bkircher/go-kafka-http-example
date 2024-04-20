# go-kafka-http-example

Small microservice in Go that reads from a Kafka topic and speaks HTTP to
another service.

## Dependencies

- Confluent's Apache Kafka Go client:
  <https://github.com/confluentinc/confluent-kafka-go>
- Uber's fast, structured, leveled Go logger: <https://github.com/uber-go/zap>

## Steps

Make sure you have go installed.

    brew install go

or

    $ go version
    go version go1.22.2 darwin/arm64

Then create a new module

    $ mkdir mkdir go-kafka-http-example
    $ cd go-kafka-http-example
    $ go mod init kafka-http-example
    go: creating new go.mod: module kafka-http-example

Then we install some dependencies:

    go get -u github.com/confluentinc/confluent-kafka-go/v2/kafka
    go get -u go.uber.org/zap

With this and the code in `main.go` you already have set everything up we need:

    $ ls -1
    go.mod
    go.sum
    main.go

Then you do

    go build

And dang, you get a binary.

    $ ./kafka-http-example
    %3|1713627093.757|FAIL|rdkafka#consumer-1| [thrd:localhost:9092/bootstrap]: localhost:9092/bootstrap: Connect to ipv4#127.0.0.1:9092 failed: Connection refused (after 2ms in state CONNECT)
    panic: runtime error: invalid memory address or nil pointer dereference
    [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x1044fba1c]

    goroutine 1 [running]:
    main.main()
            /Users/KircherB/src/go-kafka-http-example/main.go:46 +0x42c

Alright, lets add a Kafka broker to the story.
