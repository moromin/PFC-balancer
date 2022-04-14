CREATE TABLE "foods" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "protein" real NOT NULL,
  "fat" real NOT NULL,
  "carbohydrate" real NOT NULL,
  "category" integer NOT NULL
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL
);

CREATE TABLE "recipes" (
  "id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "user_id" bigint NOT NULL
);

CREATE TABLE "procedures" (
  "id" bigserial PRIMARY KEY,
  "text" varchar NOT NULL,
  "recipe_id" bigint NOT NULL
);

CREATE TABLE "recipe_food" (
  "id" bigserial PRIMARY KEY,
  "recipe_id" bigint NOT NULL,
  "food_id" bigint NOT NULL,
  "amount" real NOT NULL
);

ALTER TABLE "recipes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "procedures" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");

ALTER TABLE "recipe_food" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("id");

ALTER TABLE "recipe_food" ADD FOREIGN KEY ("food_id") REFERENCES "foods" ("id");
