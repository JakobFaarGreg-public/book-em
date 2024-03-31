# book-em: A Library Management System

This is the repository containing the Library Management System created by me.

## The Idea

- A user can borrow available books and reserve already lent out books.
- book-em lends a book to a user, or reserves it and informs the user of their placement in the queue of reservations.

## Idea Overview

- Frontend
  - CLI
    - It returns the book(s), if it is available.
    - Alternatively it informs the user of their placement in the queue of reservations.
  - Website
    - Search through books and see their availability
- Backend

  - A monolithic server that receives a request from the frontend and processes it.

- Database
  - A relational database that contains all the books.

## Project Overview

- DevOps:

  - Static Analysis:
    - Linting
    - Formatting
    - Style checking
    - Regression testing
  - Conventions:
    - Monorepo
    - [Conventional-Commits](https://www.conventionalcommits.org/en/v1.0.0/) (Controls the versioning)
    - Continuous Integration (No branches, no pull requests)
    - Test Driven Development
  - Containerization
    - A devcontainer is provided

- Issue Tracking
  - [Github Project](https://github.com/orgs/JakobFaarGreg-public/projects/1/views/1)

## Running the Application

Run the server from `./backend` with `go run main.go` _remember to insert connection string_

Run the database from `./db` with `docker compose up`

Run the curl script from `./frontend/curl` with `bash get_book.sh <id_of_book>` _might need to `chmod +x get_book.sh`_
