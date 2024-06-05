#!/bin/sh

for i in $(seq 1 10000)
do
    curl -X POST http://localhost:8082/api/v1/chatservers \
         -H "Content-Type: application/json" \
         -d "{\"name\": \"server $i\"}"
done
