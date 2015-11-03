FROM ubuntu:14.04

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y wget
RUN mkdir -p /srv/socketio/

ADD socket.io /srv/socketio/
ADD assets /srv/socketio/assets
ADD templates /srv/socketio/templates

WORKDIR /srv/socketio

CMD ["./socket.io"]

EXPOSE 3000
