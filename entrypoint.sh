#!/bin/sh
chmod 600 /usr/mongo-secret
/usr/bin/mongod --keyFile /usr/mongo-secret --bind_ip_all --replSet rs0
