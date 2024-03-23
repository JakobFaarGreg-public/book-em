#! bin/bash

curl -X GET "http://127.0.0.1:3333/v1/book?book="$1
curl -X GET "http://127.0.0.1:3333/v1/book"
