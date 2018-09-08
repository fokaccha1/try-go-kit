```
$ docker-compose up
$ curl -d'{"id":2}' localhost:5000/get-user
{"id":2,"name":"George","age":25}
$ curl -d'{"name":"Jimmy", "age":45}' localhost:5000/create-user
{"id":7}
```

# Grafana
http://localhost:3000

| user | pass |
|:-----:|:-----:|
| try-go-kit | pass |
