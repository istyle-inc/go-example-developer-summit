# example / Apache Kafka golang producer for developer summit2018

producer_example - basic producer

required [dep](https://github.com/golang/dep)

## Installing librdkafka

[confluentinc/confluent-kafka-go installing-librdkafka](https://github.com/confluentinc/confluent-kafka-go#installing-librdkafka)

## Started

```bash
$ make build
```

```bash
$ goproducer produce --file data/data.log
```

### For MacOS

```bash
export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
```
