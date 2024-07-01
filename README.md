# To-do List API
REST API for To-do List Web App

## Prerequisite
- Go version 1.20+
- Postgresql

## Running
### Using Docker
#### Postgresql

```
docker run --name postgresql -e POSTGRES_PASSWORD=password -d -v ~/home/user/postgres-data:/var/lib/postgresql/data -p 5432:5432 postgres
```

#### Go
copy the **.env.example** to **.env** and fill the values then run:
```
./start.sh
```
