#!/usr/bin/env bash
#(1)
#$ curl -i http://localhost:9999/
#HTTP/1.1 200 OK
#Date: Mon, 12 Aug 2019 16:52:52 GMT
#Content-Length: 18
#Content-Type: text/html; charset=utf-8
#<h1>Hello Gee</h1>
#
#(2)
#$ curl "http://localhost:9999/hello?name=geektutu"
#hello geektutu, you're at /hello
#
#(3)
curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
#{"password":"1234","username":"geektutu"}
#
#(4)
#$ curl "http://localhost:9999/xxx"
#404 NOT FOUND: /xxx