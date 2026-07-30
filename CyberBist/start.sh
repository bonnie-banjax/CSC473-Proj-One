#!/bin/sh
mongod --fork --logpath /var/log/mongodb.log --dbpath /data/db
node seed.js    # <--- Seeds initial menu data if DB is empty
node server.js