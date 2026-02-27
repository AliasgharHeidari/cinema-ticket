# get-session-list

![session-list](./session-list.png)

## description

1. user call get-session-list endpoint.
2. service returns request's response (session list in JSON)

## api-contract

## session-list

```go
Name: get-session-list
Method: Get
Url: http://127.0.0.1:3000/session
Params:
Header:
Body:
Errors:
    - code: 500
      Name: Internal server error
      Body:
          {
            "error" : "failed to get session list, try again later",
          }
Responses:
    - code: 200
      Name: OK
      Body:
          {
            "session-list" : {
                "movie" : {
                    "Name" : "",
                    "Year" : "",
                    "Genre" : "",
                    "Director" : "",
                },
                "room" : {
                    "Name" : "",
                    "Number" : "",
                },
                "time" : {
                    "Start" : "",
                    "End" : "",
                    "Duration" : "",
                },
                "session-id" : (string),
            }
          }

```
