#!/bin/bash

#fail on the first error
set -e

#create db
echo "creating simpleform db"
createdb simpleform

echo "setting up simpleform db"
psql --echo-queries --dbname simpleform --file db.sql

#create config file
echo "creating default config file"
cp config.json.sample config.json
