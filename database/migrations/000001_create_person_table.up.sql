CREATE TABLE person (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

INSERT INTO person (name, email)
VALUES 
('John Doe', 'john.doe@example.com'),
('Jane Doe', 'jane.doe@example.com');
