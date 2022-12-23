import sys
import pika

from config.settings import settings
from Send import SendingMail
import os

ROOT_DIR = os.path.abspath(os.curdir)


def main():
    credentials = pika.PlainCredentials(
        settings.RABBIT_USERNAME,
        settings.RABBIT_PASSWORD,
    )
    conn = pika.ConnectionParameters(
        settings.RABBIT_URL,
        settings.RABBIT_PORT,
        "/",
        credentials,
    )
    Connection = pika.BlockingConnection(conn).channel()
    Connection.queue_declare(
        queue=settings.RABBIT_QUEUE
    )

    Connection.basic_consume(
        queue=settings.RABBIT_QUEUE,
        auto_ack=True,
        on_message_callback=SendingMail
    )
    print('[*] Engine Start Waiting for messages. To exit press CTRL+C')
    print(f"Time \t\t\t\t\t\tMessage")
    Connection.start_consuming()


if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        sys.exit(0)
