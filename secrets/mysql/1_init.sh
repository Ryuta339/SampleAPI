#!/bin/bash

COMMAND='
CREATE DATABASE books;
USE books;
CREATE TABLE book (
	id INT AUTO_INCREMENT, 
	title VARCHAR(40), 
	content VARCHAR(40),
	primary key (id)
);'

mysql --user=root \
      --password=password123 \
	  --execute="$COMMAND"
