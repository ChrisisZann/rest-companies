# XM companies service

a small service to create, patch, delete and get companies

## Table of Contents

- Dependencies
- Configuration file
- Installation
- Run as docker container
- Usage
- Events
- Testing Examples


## Dependencies

```bash
go get github.com/golang-jwt/jwt/v5
go get github.com/google/uuid
go get github.com/gorilla/websocket
go get github.com/lib/pq
go get github.com/tsawler/toolbox
```

## Configuration file

optionally load configuration file with <-c> option: -c example_file.json

defualt file name is config.json

add in Dockerfile by modifiying

CMD ["/usr/local/bin/xmapi/api","-c","example_file.json"]

example:
```json
{
    "jwt_key": "dummy_key",
    "db_user": "dummy_username",
    "db_password": "dummy_password",
    "db_host": "dummy_hostname",
    "db_name": "dummy_db_name",
    "log_dir": "/var/log/xm-companies"
}
```

## Installation

Download project:

```bash
git clone https://github.com/chrisisZann/xm-companies.git
```

## Run as docker container

- Run as standalone service
```bash
docker build .
docker run -it --name xmapi -p <host_port>:8888 <container-image>
```

- Run with docker compose
```bash
docker compose up -d
```

## Usage

1. Create user by sending POST request on /user with parameters username and password
2. use the JWT in response to access protected endpoints(POST, PATCH, DELETE /company)

if token expires (after 24hours) login to receive new token

GET /company does not require authentication

## Events

Connect via websocket to listen to mutating events of the microservice
ws://127.0.0.1:port/ws

## Testing Examples

Create user
```bash
curl -X POST "http://127.0.0.1:8888/user?username=<USERNAME>&password=<PASSWORD>"
```

Login
```bash
curl -X POST "http://127.0.0.1:8888/login?username=<USERNAME>&password=<PASSWORD>"
```

Create company
```bash
curl -X POST -H "Authorization: Bearer <your_token>"\
"http://127.0.0.1:8888/auth-company?name=<value>&description=<value>&registered=<value>&type=<value>&amount_of_employees=<value>" 
```

Find company
```bash
curl -X GET "http://127.0.0.1:8888/company?name=<COMPANY_NAME>"
````

Patch company field
```bash
curl -X PATCH -H "Authorization: Bearer <your_token>"\
"http::/127.0.0.1:8888/auth-company?name=<COMPANY_NAME>&field=<key>&value=<value>" 
```

Delete company field
```bash
curl -X DELETE  -H "Authorization: Bearer <your_token>"\
"http://127.0.0.1:8888/auth-company?name=<COMPANY_NAME>"
```