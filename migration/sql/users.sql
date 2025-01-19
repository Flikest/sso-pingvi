CREATE TABLE users(
    id UUID
    name VARCHAR(255) UNIQUE NOT NULL
    pass VARCHAR(150) NOT NULL
    avatar VARCHAR(1000) NOT NULL
    about_me TEXT
)