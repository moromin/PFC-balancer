CREATE TABLE food (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "protein" real NOT NULL,
  "fat" real NOT NULL,
  "carbohydrate" real NOT NULL,
  "category" integer NOT NULL
);

CREATE TABLE users (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);
