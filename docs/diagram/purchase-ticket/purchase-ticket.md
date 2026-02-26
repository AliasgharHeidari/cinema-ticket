# purchase-ticket

![purchase-ticket](./purchase-ticket.png)

## description

1. user put phone-number in request body and call purchase-ticket endpoint.
2. service returns request's response

## api-contract

## purchase-ticket

```go
Name: purchase-ticket
Method: POST
Url: http://127.0.0.1:3000/payment/example-generated-link
Params:
Header:
Body:
    {
        "mobile-number" : (string),
    }
Errors:
    - code: 500
      Name: Internal server error
      Body:
          {
            "error" : "purchase failed, please try again",
          }
Responses:
    - code: 200
      Name: Ok
      Body:
        {
          "messsage" : "purchase compeleted successfully",
          "ticket" : {
            "movie" : "example",
            "room"  : 2,
            "row"   : 3,
            "seat"  : 4
          }
        }
```
