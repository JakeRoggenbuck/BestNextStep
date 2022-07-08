# Daft

## Use
```sh
curl -v -X GET "http://localhost:1357"

curl -v -X GET "http://localhost:1357/api/v1" -u Admin:adminpassword

curl -v -X GET "http://localhost:1357/api/v1/all" -u Admin:adminpassword

curl -v -X POST "http://localhost:1357/api/v1/new-user" -u Admin:adminpassword -F "name=Jake" -F "password=123"
```
