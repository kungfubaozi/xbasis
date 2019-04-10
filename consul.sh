#!/usr/bin/env bash

#docker run -d  -p 8400:8400 -p 8500:8500 -p 8600:53/udp \
#    -e 'CONSUL_LOCAL_CONFIG={"skip_leave_on_interrupt": true}' \
#    consul agent -server -bootstrap-expect=1 -ui -node=node11 -client=0.0.0.0

consul agent -dev -server -data-dir /tmp/consul -client=192.168.80.67 -node=node2 -ui