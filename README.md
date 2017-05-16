# docker-fake-response-server

A Docker image to always respond HTTP status 200.

This is intended to check container platform like ECS.

You can specify the server port like below.

```
$ docker run --rm --name fake-server -d -p 8888:8888 -e "FAKE_SERVER_PORT=8888" sudix/docker-fake-response-server
```
