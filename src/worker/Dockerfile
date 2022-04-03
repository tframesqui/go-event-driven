FROM python:3.9.6
ENV PYTHONUNBUFFERED=1
WORKDIR /worker
COPY main.py /worker/main.py
COPY requirements.txt /worker/
RUN pip install -r requirements.txt
COPY . /worker/