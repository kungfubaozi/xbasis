FROM alpine:latest

RUN mkdir /app

WORKDIR /app

ADD conf/manage.json /app/init.json
ADD permission /app/permission

CMD ["./permission"]