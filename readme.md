# RESTful GO

A tiny project to demonstrate principles of simple REST API and packages in GO

## Usage
- ```GET /``` - shows a form for sending a request to ```POST /posting```
- ```GET /articles``` - returns JSON of all objects in database table

## Structure
- ```/main.go``` - entry point. Handles requests and works as a controller.
- ```/testdb/main.go``` - database controller with methods to manipulate data of database.
- ```/databasetest/main.go``` - database tests (just to demonstrate simple features in a simple way)
- ```/article/article.go``` - article object

