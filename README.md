# GET
curl --location --request GET 'http://localhost:8080/api/v1/person/'

# POST
curl --location --request POST 'http://localhost:8080/api/v1/person/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 2,
    "firstname": "halose2",
    "lastname": "hai"
}'

# PUT
curl --location --request PUT 'http://localhost:8080/api/v1/person/:id' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 2,
    "firstname": "halose2",
    "lastname": "hai"
}'

# DELETE
curl --location --request DELETE 'http://localhost:8080/api/v1/person/:id'