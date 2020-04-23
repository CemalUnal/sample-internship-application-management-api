#!/bin/bash

mongorestore --host mongodb:27017 --gzip --archive=/mongo-seed/dump.gz --drop --db Company
