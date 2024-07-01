# Transaction Core Service API

## Prerequisite
- Go version 1.20+
- MySQL

## Running
### Using Docker Compose
```
docker-compose run --build
```

## Endpoint List
### [GET] Get Limit List
```
URL: http://localhost:5000/api/v1/limits
```

### [POST] Checkout
```
URL: http://localhost:5000/api/v1/transactions/checkout
BODY: {
    "consumer_id": 1,
    "limit_id": 1,
    "amount": 150000,
    "asset_name": "Soft Case"
}
```
