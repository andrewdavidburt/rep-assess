# reperio-backend-assessment

This is Reperio Health's GoLang backend assessment. 

## Notes from Candidate (Andrew Burt)
- I've uploaded all of this also to a github repo at https://github.com/andrewdavidburt/rep-assess
- I've included full notes to my resolutions to the Objectives below the Objectives section (Notes on Completion of Objectives)

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

## Notes on Completion of Objectives
1. Fixed. Ordering and types were off. (packages/weather.go)
2. I added a v2 route that implements saving to the database. That said, after fixing the bugs relating to the db query, the code in the v1 route which saved to the database started to work, so I wasn't sure whether to also remove this from the v1 route. I decided to leave it, so now both the v1 route saves to the database, and the v2 route, which I created and implemented, also saves to the database. (functions/v2/weather.go, handlers/v2/currentWeather.go, routes/currentWeather.go, routes/versions.go)
3. Fixed. Condition and types were off. (database/query.go)
4. Added a query param to the forecast route "sort", which has options "asc" and "desc" (ascending and descending), which implements sort and orders by the forecast:forecastday:day:avgtemp_c json fields of each returned day's forecast in ascending or descending order. (handlers/v1/forecast.go)
5. Added 2 open interface to type casts. One is fully integrated into the code (functions/v2/weather.go). The other is sort of arbitrary, since it didn't result in something that was useable by the following step, but demonstrates another way to cast to type: (packages/weather.go).
6. (Bonus Objective) This is unfinished. I stubbed the middleware (middleware/common.go, routes/versions.go), and mostly implemented the database query wrapper (I think I successfully implemented Query in types/database.go, but the select query wrapper in database/query.go is not quite right), but didn't complete troubleshooting the actual db query, or saving to the context.

7. Additionally: Possible issue: The forecast route only returns 1-3 days of data from the forecast api "days" parameter. 4+ day requests only return 3. Possible external API limitation? 