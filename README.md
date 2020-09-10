## REST API using Onion Architecture

A boilerplate project that designed and aimed for enterprise-level using Onion Architecture

### Prerequisite

- Go 1.13+
- SQLite3
- Gorm
- Gin
- Heimdall

### Features

- [x] Multiple server
- [x] A dynamic filter and search given by query
- [x] Cache an object
- [x] A dynamic limit and offset for retrieving products
- [x] Graceful shutdown
- [x] Struct level validation
- [ ] gRPC
- [ ] Testing
- [ ] Integrate with Sentry for logging
- [ ] Circuit Breaker
- [ ] API documentation

### Key Notes

- Circuit breaker is not fully implemented
- Remove `products.db` if you can't insert the database during/after migration

### Usage

Using go modules

```shell script
go install
go run main.go
```

Using docker

```shell script
docker build -t <name_your_app> .
docker run -p 8080:8081 -it <name_your_app>
```

### API Endpoints

| Methods | Endpoints | Description |
| --------- | ----------- | --------- |
| `GET` | /api/v1/product/search/?Search= | Retrieve a products given by `product_name` |
| `GET` | /api/v1/product/cache | Retrieve and cache a products, duration of caching set to 15 minutes |
| `GET` | /api/v1/product/limit/?Limit= | Retrieve a products with limited result |
| `GET` | /api/v1/product/all-products | Fetching and retrieve all products |
| `GET` | /api/v1/product/offset/?Offset= | Retrieve and skip the rows from beginning before returned a result |
| `GET` | /api/v1/product/detailed/:product_id | Fetching a detailed product given by `product_id` |
| `PUT` | /api/v1/product/update/:product_id | Fetching and updated a product given by `product_id` |
| `POST` | /api/v1/product/create | Create a new product |
| `DELETE` | /api/v1/product/delete/:product_id | Delete a product given by `product_id` |