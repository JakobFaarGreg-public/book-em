## How to run the database

Step 1: In the `./db` folder run `docker compose up`

Step 2: `docker run --name pgadmin-database -p 5051:80 -e "PGADMIN_DEFAULT_EMAIL=user@bookem.com" -e "PGADMIN_DEFAULT_PASSWORD=bookem" -d dpage/pgadmin4`
