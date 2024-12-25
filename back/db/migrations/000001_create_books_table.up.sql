CREATE TABLE IF NOT EXISTS books(
    isbn VARCHAR(13) PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    author VARCHAR(50) NOT NULL,
    publisher VARCHAR(50) NOT NULL,
    yearPublished VARCHAR(4) NOT NULL,
    description TEXT,
    amount INTEGER,
    CONSTRAINT amount CHECK (amount >= 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    userFullName VARCHAR(50) NOT NULL,
    isbn VARCHAR(13) REFERENCES books(isbn) ON DELETE CASCADE, 
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO books (isbn, title, author, publisher, yearPublished, description, amount)
VALUES
    ('9780306406157', 'The Great Gatsby', 'F. Scott Fitzgerald', 'Random House', '2019', 'A story of wealth, ambition, and tragedy in 1920s America.', 3),
    ('9781400033871', 'To Kill a Mockingbird', 'Harper Lee', 'J.B. Lippincott & Co.', '1960', 'A tale of racial injustice in the American South.', 1),
    ('9780451160924', '1984', 'George Orwell', 'Penguin Books', '1949', 'A dystopian novel about totalitarianism and surveillance.', 0);
