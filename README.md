# _No name =(_

API for a cool web application that will help you play tabletop role-playing games online

### Setup

- Initialize project (it will copy sample files to not sample files)

```
make init
```

- Run docker containers (Database, test database and adminer)

```
docker-compose up -d
```

- Run migrations for actual database and test database

```
~/migrate -path migrations -database "postgres://<username>:<password>@<host>:<port>/<db_name>?sslmode=disable" up
```

- Run tests to check if everything is ok

```
make test
```

- Build app

```
make build
```

- Run app

```
./apiserver
```

- Enjoy!

### Used Technologies

- Go v1.20
- PostgreSQL v14