# library-api
This microservice serves APIs for managing books

Build Project Steps
Goto library-app directory and excute the following commands

#docker-compose build
#docker-compose up


API-1
curl --location 'http://localhost:8080/books' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Second Animal Book",
    "author": "Kripa Sharma",
    "isbn": "false"
}'
Response: 
{
    "id": 2,
    "title": "Second Animal Book",
    "author": "Kripa Sharma",
    "isbn": "false"
}

API-2
curl --location --request PUT 'http://localhost:8080/books/2' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Second Fiction Book",
    "author": "Kripa Sharma",
    "isbn": "false"
}'
Response:
{
    "id": 2,
    "title": "Second Fiction Book",
    "author": "Kripa Sharma",
    "isbn": "false"
}

API-3
curl --location 'http://localhost:8080/books'
Response:
[
    {
        "id": 1,
        "title": "First Animal Book",
        "author": "Kripa Shankar Sharma",
        "isbn": "true"
    },
    {
        "id": 2,
        "title": "Second Fiction Book",
        "author": "Kripa Sharma",
        "isbn": "false"
    }
]

API-4
curl --location --request DELETE 'http://localhost:8080/books/1'
Response:
http.status - 204
