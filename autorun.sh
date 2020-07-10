#!bin/sh

go run main.go &
sleep 1
curl http://localhost:1234
