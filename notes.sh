git clone https://github.com/ChrisisZann/xm-companies.git
git pull origin main

docker exec -it 7806a52a26e8 sh


http://192.168.1.11:8888/user?username=curl_user&password=1234

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImN1cmxfdXNlciIsImlzcyI6InlpcHBlZSIsInN1YiI6ImNvbXBhbmllcyIsImV4cCI6MTcyNzI2ODQzNSwibmJmIjoxNzI3MTgyMDM1LCJpYXQiOjE3MjcxODIwMzV9.j2Opux233PUNLVN8nzDUv4c3B0quy8KpKqMucZjdsEk

curl -X POST "http://192.168.1.11:8888/user?username=docker&password=1234"
curl -X POST "http://127.0.0.1:8888/user?username=docker&password=1234"


curl -X POST http://192.168.1.11:8888/login?username=<USERNAME>&password=<1234>

curl -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImN1cmxfdXNlciIsImlzcyI6InlpcHBlZSIsInN1YiI6ImNvbXBhbmllcyIsImV4cCI6MTcyNzI2ODQzNSwibmJmIjoxNzI3MTgyMDM1LCJpYXQiOjE3MjcxODIwMzV9.j2Opux233PUNLVN8nzDUv4c3B0quy8KpKqMucZjdsEk" \
"http://192.168.1.11:8888/auth-company?name=new_company&description=desc&registered=true&type=NonProfit&amount_of_employees=666"


curl -X GET "http://192.168.1.11:8888/company?name=<COMPANY_NAME>"

curl -X PATCH -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImN1cmxfdXNlciIsImlzcyI6InlpcHBlZSIsInN1YiI6ImNvbXBhbmllcyIsImV4cCI6MTcyNzI2ODQzNSwibmJmIjoxNzI3MTgyMDM1LCJpYXQiOjE3MjcxODIwMzV9.j2Opux233PUNLVN8nzDUv4c3B0quy8KpKqMucZjdsEk" \
"http://192.168.1.11:8888/auth-company?name=<COMPANY_NAME>&field=<key>&value=<value>" 

curl -X DELETE  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImN1cmxfdXNlciIsImlzcyI6InlpcHBlZSIsInN1YiI6ImNvbXBhbmllcyIsImV4cCI6MTcyNzI2ODQzNSwibmJmIjoxNzI3MTgyMDM1LCJpYXQiOjE3MjcxODIwMzV9.j2Opux233PUNLVN8nzDUv4c3B0quy8KpKqMucZjdsEk" \
"http://192.168.1.11:8888/auth-company?name=new_company"

