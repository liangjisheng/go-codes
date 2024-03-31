### Group routes

This example shows how to group different routes in their own files and group them together in a orderly manner like this:

```go
func getRoutes() {
	v1 := router.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	v2 := router.Group("/v2")
	addPingRoutes(v2)
}
```

```shell
curl "http://127.0.0.1:8080/v1/users/"
curl "http://127.0.0.1:8080/v1/users/comments"
curl "http://127.0.0.1:8080/v1/users/pictures"
curl "http://127.0.0.1:8080/v1/ping/"
curl "http://127.0.0.1:8080/v2/ping/"
```
