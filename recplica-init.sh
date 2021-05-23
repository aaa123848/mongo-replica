#!/bin/sh
mongo --eval "rs.initiate({_id:'rs0', members: [{_id: 0, host: 'golang-mongo-1:27017'}, { _id: 1, host: 'golang-mongo-2:27017'}, { _id: 2, host: 'golang-mongo-3:27017'}]})"
# echo "rs.initiate({_id:'rs0', members: [{_id: 0, host: 'golang-mongo-1:27017'}, { _id: 1, host: 'golang-mongo-2:27017'}, { _id: 2, host: 'golang-mongo-3:27017'}]})" | mongo

mongo admin --eval "db.createUser({user: 'root', pwd: '1234', roles: ['root']})"
# echo "db.createUser({user: 'root', pwd: '1234', roles: ['root']})" | mongo admin