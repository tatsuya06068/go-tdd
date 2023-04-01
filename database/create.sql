CREATE DATABASE IF NOT EXISTS go_db;

USE go_db;

CREATE TABLE IF NOT EXISTS tasks(
    task_id int NOT NULL AUTO_INCREMENT,
    task_name varchar(255) NOT NULL,
    PRIMARY KEY (task_id)
);

INSERT INTO tasks (task_name) VALUES('task_1'),('task_2'),('task_3'),('task_4'),('task_5'),('task_6'),('task_7'),('task_8'),('task_9'),('task_10');




