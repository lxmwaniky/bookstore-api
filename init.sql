-- Create books table
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publication_year INTEGER,
    genre VARCHAR(100),
    description TEXT,
    cover_image_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Insert test data
INSERT INTO books (title, author, publication_year, genre, description) VALUES
('The Go Programming Language', 'Alan Donovan', 2015, 'Programming', 'A comprehensive guide to Go programming'),
('Clean Code', 'Robert Martin', 2008, 'Programming', 'A handbook of agile software craftsmanship'),
('The Pragmatic Programmer', 'David Thomas', 1999, 'Programming', 'From journeyman to master'),
('Design Patterns', 'Gang of Four', 1994, 'Programming', 'Elements of reusable object-oriented software');

-- Verify the data
SELECT * FROM books;