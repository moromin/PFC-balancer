CREATE TABLE foods (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "protein" real NOT NULL,
  "fat" real NOT NULL,
  "carbohydrate" real NOT NULL,
  "category" integer NOT NULL
);
