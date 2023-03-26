# gochive
'go + archive' is 'gochive!', which means that this project is for string golang posts in reddit and medium. because i have no time to read so many post of golnag in reddit or medium. so i decided to store the post that uploaded in the reddit or medium. And i will provide the post like a newspaper form.   

## TODO
1. database - on going
2. reddit worker
3. REST API
4. Automatically store

## .env
```
DB_DSN=<DB_DSN>
```

## MYSQL
- mysql container install
```
docker pull mysql
```

- run container of mysql
```
export MYSQL_ROOT_PASSWORD=<PASSWORD>
docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD -d -p 3306:3306 mysql:latest
```

- execute mysql root server
```
docker exec -it mysql-container bash

mysql -u root -p
```

- create database
```sql
CREATE DATABASE gochive CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
use gochive;
```

- create mysql table
```sql
CREATE TABLE post
(
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    author VARCHAR(100) NOT NULL,
    origin VARCHAR(30) NOT NULL,
    title VARCHAR(200) NOT NULL,
    content VARCHAR(200) NOT NULL,
    link VARCHAR(2083) NOT NULL,
    published_at VARCHAR(30) NOT NULL,
    created_at DATETIME NOT NULL
);
```

- insert mysql entity
```sql
INSERT INTO post (author,origin, title, content, link, published_at, created_at) VALUES("gyu","reddit", "gyu", "test", "test", "2012-12-12", "2012-12-12") 
ON DUPLICATE KEY UPDATE link="test";
```