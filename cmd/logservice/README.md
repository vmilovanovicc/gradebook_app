# Standalone Log Service

## Build
`go build gradebook_app/cmd/logservice`

This will create the binary `logservice.exe`.

---

## Run 
`./logservice.exe`

![logservice](../../img/cmd_logservice.png "Run logservice")

---
## Test

Test your service in Postman. Send a POST request. Target `/log`. 

![postman](../../img/postman_logservice.png "Send request")

This should log the body's request in `app.log`.

![logs](../../img/app_log.png "Log")

---

To stop the service, press any key.
