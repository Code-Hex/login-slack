## login-slack

This is a library to get login api-token for Slack.

## Synopsis

```go
token, err := loginslack.Login("email@email.com", "password", "workspace")
if err != nil {
    log.Fatal(err)
}
log.Println(token) // print api-token like xoxs-*
```