# GET
curl --location --request GET 'http://localhost:8080/'

# POST
curl --location --request POST 'http://localhost:8080/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 2,
    "firstname": "halose2",
    "lastname": "hai"
}'

# PUT
curl --location --request PUT 'http://localhost:8080/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 2,
    "firstname": "halose2",
    "lastname": "hai"
}'

# DELETE
curl --location --request DELETE 'http://localhost:8080/2'