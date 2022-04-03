import pandas as pd

from catalog_service import (
    send_message,
    get_producer,
)


def main():
    products_df = pd.read_csv("csv/products.csv")

    total_messages = len(products_df.index)
    counter = 0

    producer = get_producer()

    for item in products_df.to_dict(orient="records"):
        send_message(producer, item)
        counter = counter + 1
        print(f"message sent {counter} / {total_messages}")

    producer.flush(30)


if __name__ == "__main__":
    main()
    print("finished sending messages")
