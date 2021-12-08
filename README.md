# reperio-backend-assessment

This is Reperio Health's GoLang backend assessment. 

## Prerequisite

- install + setup golang 1.17
- install sqlite3 
- install ide with linting tools for golang

## Environment Variables

- PORT
- ENVIRONMENT

## To Run 

`PORT=3000 ENVIRONMENT=tory go run main.go`

## To Build

`go build -o $GOPATH/bin/reperio-backend-assessment main.go`

## Objectives

- There is a bug hidden that deals with formatting of string being sent to the weather api, please find and fix 
- Add a v2 route for current-weather that implements saving to the database
- There is a bug with the query builder please find and fix
- Sort the forecast data by coldest or hottest day using a query param
- Implement a cast somewhere from interface{} to an actual type
- Add a piece of middleware that checks if the current location has a record in the database and saves true or false in the request context (bonus)
