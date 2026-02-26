# get-session-seat

![session-seat](./session-seat.png)

## description

1. user put session-id in request url and call get-session-seat endpoint.
2. service returns request's response (avalible session seats in JSON)

## api-contract

## session-seat

```go
Name: get-session-seat
Method: Get
Url: http://127.0.0.1:3000/sessions/seat/:session-id
Params: 
    - Name: session-id
      Type: string
Header:
Body:
Errors:
    - code: 500
      Name: Internal server error
      Body:
          {
            "error" : "failed to get session seats, try again later",
          }
Responses:
    - code: 200
      Name: Ok
      Body:
        {
            "available seats" : {
                "Row" : "1",
                "seat": [1,2,3,4,5],
                "Row" : "2",
                "seat": [1,2,3,4,5],

                ...
                
            },
        }