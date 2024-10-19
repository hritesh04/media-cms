-- Enum
CREATE TYPE role AS ENUM ('author', 'user');

-- Create Categories table
CREATE TABLE Categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO Categories (name) VALUES ('General'),('Culture'),('Geography'),('HealthAndFitness'),('History'),('Psychology'),('Mathematics'),('Natural'),('LifeStyle'),('Philosophy'),('SocialScience'),('Technology')