# hadar
This is a simple application in go using redis.

## How to run

### Run the redis + redis-commander + hadar

You can use the docker compose file to run:
- [redis](https://hub.docker.com/_/redis)
- [redis-commander](https://hub.docker.com/r/rediscommander/redis-commander)
- hadar

### Run just Redis
```
docker-compose up
```
You can use docker to run redis
```
docker run --name redis -p 6379:6379 -d redis
```

### Run the application
```
go run main.go
```
