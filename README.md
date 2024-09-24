# XM companies service

a small service to create, patch, delete and get companies

## Table of Contents

- Dependencies
- Configuration file
- Installation

## Dependencies

github.com/golang-jwt/jwt/v5
github.com/google/uuid
github.com/gorilla/websocket
github.com/lib/pq
github.com/tsawler/toolbox

## Configuration file

localy load Configuration file:
-c example_file.json

default is config.json


defualt file is config.json

change in Dockerfile by modifiying

CMD ["/usr/local/bin/xmapi/api","-c","example_file.json"]

## Installation

Instructions on how to install and set up the project. For example:

```bash
git clone https://github.com/chrisisZann/xm-companies.git
cd yourproject

docker build .

docker run -it --name xmapi -p 8080:8888 <container-image>
docker run -it --name xmapi -p 8080:8888 sha256:cdfda30d132bd9b1ee408275ef4b6a75ef018c62c4c95d921def3352c8ceb5b2

or

docker compose up -d

