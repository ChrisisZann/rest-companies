# XM companies service

a small service to create, patch, delete and get companies

## Table of Contents

- Dependencies
- Configuration file
- Installation
- Run as docker container
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

load configuration file with <-c> option: -c example_file.json

defualt file is config.json

add in Dockerfile by modifiying

CMD ["/usr/local/bin/xmapi/api","-c","example_file.json"]

## Installation

Download project:

```bash
git clone https://github.com/chrisisZann/xm-companies.git
```

## Run as docker container

- Run as standalone service
```bash
docker build .
docker run -it --name xmapi -p 8080:8888 <container-image>
```

- Run with docker compose
```bash
docker compose up -d
```
## Testing Examples

Create user
```bash
curl -X POST http://192.168.1.11:8888/user?username=<USERNAME>&password=<PASSWORD>
```

Login
```bash
curl -X POST http://192.168.1.11:8888/login?username=<USERNAME>&password=<PASSWORD>
```

Create company
```bash
curl -X POST http://192.168.1.11:8888/auth-company?name=<value>&description=<value>&registered=<value>&type=<value>&amount_of_employees=<value> -H "Authorization: Bearer <your_token>"
```

Find company
```bash
curl -X GET http://192.168.1.11:8888/company?name=<COMPANY_NAME>
````

Patch company field
```bash
curl -X PATCH http://192.168.1.11:8888/auth-company?name=<COMPANY_NAME>&field=<key>&value=<value> -H "Authorization: Bearer <your_token>"
```

Delete company field
```bash
curl -X DELETE http://192.168.1.11:8888/auth-company?name=<COMPANY_NAME> -H "Authorization: Bearer <your_token>"
```