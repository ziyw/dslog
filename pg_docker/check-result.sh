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
  type VARCHAR NOT NULL,
  text VARCHAR NOT NULL
);

INSERT INTO logs(type, text) 
VALUES ('INFO', 'Hello world from client');
INSERT INTO logs(type, text) 
VALUES ('ERROR', 'server side exception');
INSERT INTO logs(type, text) 
VALUES ('WARN', 'This might now work');

