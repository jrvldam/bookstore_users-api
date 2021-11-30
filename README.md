# bookstore_users-api
Users API

### Set database for the project
Install mysql client:
```shell
brew install mysql-client
```

Create a file `docker-compose.yml` with the content:
```yaml
version: '3'
services:
  mysql:
    image: mysql:latest
    ports:
      - 8083:3306
    volumes:
      - ./test-sql-2:/docker-entrypoint-initdb.d
    environment:
      MYSQL_ROOT_PASSWORD: BATMAN
      MYSQL_DATABASE: mysql
```

In the same directory where docker-compose.yml file:
```shell
docker-compose up
```

Check the expossed port by running:
```shell
docker ps
```

Open a mysql session:
```shell
mysql -P 49384 --protocol=tcp -u root -p
```

Create the schema
```shell
CREATE SCHEMA `users_db` DEFAULT CHARACTER SET utf8 COLLATE utf8_spanish2_ci;
```

Set the environment variables:
```shell
export mysql_users_username=root
export mysql_users_password=BATMAN
export mysql_users_host="127.0.0.1:49384"
export mysql_users_schema=users_db
```
