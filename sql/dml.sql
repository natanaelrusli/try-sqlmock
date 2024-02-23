INSERT INTO books (title, description, cover, author_id, stock, updated_at, created_at)
SELECT 
    'Book Title ' || generate_series,
    'Description for Book ' || generate_series,
    '/path/to/cover' || generate_series || '.jpg',
    (random() * 10 + 1)::int,
    (random() * 100)::int,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
FROM generate_series(1, 100);
