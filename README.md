# docker-fake-response-server

A Docker image to always respond HTTP status 200.

This is intended to check container platform like ECS.

You can specify the server port with environment value `FAKE_SERVER_PORT`.


## Usage

```
$ docker run --rm --name fake-server -d -p 8888:8000 -e "FAKE_SERVER_PORT=8000" sudix/docker-fake-response-server
$ curl -i http://localhost:8888/somenotfoundpath
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 16 May 2017 07:53:03 GMT
Content-Length: 35

{"message":"Hello!","success":true}
```
