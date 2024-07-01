# Transaction Core Service API

## Prerequisite
- Go version 1.20+
- MySQL

## Running
### Using Docker
#### MySQL

```
docker run --name mysql-temp -p 3306:3306 -d mysql:latest
```

#### Go
copy the **.env.example** to **.env** and fill the values then run:
```
./start.sh
```
