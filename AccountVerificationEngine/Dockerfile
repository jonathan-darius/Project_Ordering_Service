FROM python:3.9

WORKDIR /AccountVerificationEngine

ADD Main.py .
ADD config .

RUN pip install pika==1.3.1 python-dotenv==0.21.0

COPY . .

CMD ["python","Main.py"]