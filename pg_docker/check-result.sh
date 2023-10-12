#!/bin/zsh
docker ps 
docker exec -it <conatiner_id> bash 
psql -U ziyan 

### Table Content ### 
 id |    title    |          content          |         created_at         
----+-------------+---------------------------+----------------------------
  1 | First Entry | The very first note entry | 2023-09-13 20:50:00.105541
  2 | Entry 2     | Second try                | 2023-09-13 20:50:11.487881
  3 | Entry three | Third time is a charm     | 2023-09-13 20:50:57.312694
(3 rows)

INSERT INTO note(title, content, created_at) 
VALUES ('DevNote', 'Should be very small', CURRENT_TIMESTAMP);

// db name: logdb 
psql logdb;

CREATE TABLE dslog (
  id INT GENERATED ALWAYS AS IDENTITY,
  created_at TIMESTAMP NOT NULL,
  logType VARCHAR NOT NULL, 
  logMsg VARCHAR NOT NULL
);

INSERT INTO dslog(created_at, logType, logMsg) 
VALUES (Now(), 'INFO', 'Hello world from client');
INSERT INTO dslog(created_at, logType, logMsg) 
VALUES (Now(), 'ERROR', 'server side exception');
INSERT INTO dslog(created_at, logType, logMsg) 
VALUES (Now(), 'WARN', 'This might now work');

DROP TABLE dslog; 

pg_dump logdb > ~/Desktop/logdb.sql // run directly, not in psql 


IMG_NAME=postgres-test-img
APP_NAME=pg-test-app

build:
	docker image build . -t $(IMG_NAME)
	docker volume create pgdata 
	docker container run -d --rm -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=ziyan --name $(APP_NAME) $(IMG_NAME) -v pgdata:/var/lib/postgresql/data

docker volume create pgdata
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
docker run -d \
  --name my-pg \
  -p 5432:5432 \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_USER=postgres \
  -v pgdata:/var/lib/postgresql/data \
  my-pg-img

docker run -d \
--name my-pg \
-p 5432:5432 \
-v pgdata:/var/lib/postgresql/data \
postgres 


docker run -d \
	--name my-postgres \
  -e POSTGRES_USER=ziyan \
	-e POSTGRES_PASSWORD=mysecretpassword \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v pgdata:/var/lib/postgresql/data \
  my-postgres-img


docker image build . -t my-postgres-img
docker volume create pgdata 


docker exec mycontainer pgdata: /path/to/test.sh
docker cp c:\myfolder\myfile.txt dummy:/root/myfile.txt

docker cp logdb.sql my-postgres:/logdb.sql
docker exec my-postgres psql -U ziyan < logdb.sql

psql -U ziyan -d logdb -f logdb.sql

https://1kevinson.com/how-to-create-a-postgres-database-in-docker/

