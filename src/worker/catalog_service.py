import json

from exceptions import ImpossibleToConnect
from kafka import KafkaProducer


_DEFAULT_KAFKA_TOPIC = "products-integration"
_DEFAULT_KAFKA_SERVER = "kafka:9092"


def get_producer():
    print("getting producer")
    producer: KafkaProducer = None
    try:
        producer = KafkaProducer(
            bootstrap_servers=_DEFAULT_KAFKA_SERVER,
            api_version=(0, 10, 2),
            value_serializer=lambda v: json.dumps(v).encode("utf-8"),
            key_serializer=lambda v: json.dumps(v).encode("utf-8"),
        )
        print("producer ok")
        return producer
    except Exception as e:
        raise ImpossibleToConnect(e)


def send_message(producer: KafkaProducer, message: dict):
    # fire and forget message
    try:
        result = producer.send(
            _DEFAULT_KAFKA_TOPIC, key={"key": message.__hash__}, value=message
        )
    except Exception as e:
        raise ImpossibleToConnect(e)
