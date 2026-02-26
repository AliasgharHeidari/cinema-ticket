# delete-session

![delete-session](./delete-session.png)

# description
1. admin put session-id in request body and call `delete-session` endpoint.
2. server checks the posibility and returns request response.

# api-contract

# delete-session

```go
Name: delete-session
Method: DELETE
Url: http://127.0.0.1:9898/session
Param:
Body:
    {
        "session-id" : string,
    }
Errors:
    - code: 500
      Name: Internal server error
      Body:
          {
            "error" : "failed to delete session, please try again later",
          }
    - code: 400
      Name: bad request
      Body:
          {
            "error" : "invalid request body",
          }
    - code: 404
      Name: not found
      Body:
          {
            "error" : "session does not exist",
          }
Responses:
    - code: 201
      Name: created
      Body:
          {
            "message" : "session deleted successfully",
          }

```