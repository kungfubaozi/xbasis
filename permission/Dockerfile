FROM alpine:latest

RUN mkdir /app

WORKDIR /app

ADD conf/admin.json /app/admin.json
ADD conf/route.json /app/route.json
ADD conf/safe.json /app/safe.json
ADD conf/user.json /app/user.json

ADD permission /app/permission

CMD ["./permission"]