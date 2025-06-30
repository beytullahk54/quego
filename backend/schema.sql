CREATE DATABASE IF NOT EXISTS bookstore;
USE bookstore;

CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
);

-- Örnek veriler
INSERT INTO books (title, author) VALUES 
    ('Savaş ve Barış', 'Tolstoy'),
    ('1984', 'George Orwell'); 