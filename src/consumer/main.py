import json
from kafka import KafkaConsumer
import logging

from exceptions import ImpossibleToConsume

_DEFAULT_KAFKA_TOPIC = "products-integration"
_DEFAULT_KAFKA_SERVER = "kafka:9092"

# logging.basicConfig(level=logging.DEBUG)


def get_consumer():
    print("get consumer")
    try:
        consumer: KafkaConsumer = KafkaConsumer(
            bootstrap_servers=_DEFAULT_KAFKA_SERVER,
            api_version=(0, 10, 2),
            value_deserializer=lambda v: json.loads(v.decode("utf-8")),
            key_deserializer=lambda v: json.loads(v.decode("utf-8")),
        )
        consumer.subscribe([_DEFAULT_KAFKA_TOPIC])
        return consumer
    except Exception as e:
        raise ImpossibleToConsume(e)


def read_messages(consumer: KafkaConsumer):
    for message in consumer:
        print(f"{message.key}:, {message.timestamp}, {message.value}")

    print("End of messages")


if __name__ == "__main__":
    print("start reading from kafka...")
    consumer = get_consumer()
    print("consumer ok")

    read_messages(consumer)
