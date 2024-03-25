\c default_database;

CREATE TABLE author (
    id int NOT NULL PRIMARY KEY,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL
);

CREATE TABLE book (
    id int NOT NULL PRIMARY KEY,
    author_id int NOT NULL,
    title varchar(255) NOT NULL,
    FOREIGN KEY (author_id) REFERENCES author(id)
);


INSERT INTO author (id, first_name, last_name) VALUES
(1, 'Stephen', 'King'),
(2, 'J.K.', 'Rowling'),
(3, 'George', 'Orwell'),
(4, 'Harper', 'Lee'),
(5, 'Agatha', 'Christie');

INSERT INTO book (id, author_id, title) VALUES
(1, 1, 'The Shining'),
(2, 1, 'Misery'),
(3, 2, 'Harry Potter and the Philosopher''s Stone'),
(4, 3, '1984'),
(5, 4, 'To Kill a Mockingbird'),
(6, 5, 'Murder on the Orient Express'),
(7, 3, 'Animal Farm'),
(8, 1, 'It'),
(9, 2, 'Harry Potter and the Chamber of Secrets'),
(10, 4, 'Go Set a Watchman');
