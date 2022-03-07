CREATE TABLE "users" (
  "user_id" bigserial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipes" (
  "recipe_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "content" varchar,
  "image_link" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "recipe_menu" (
  "id" bigserial PRIMARY KEY,
  "recipe_id" bigint NOT NULL,
  "menu_id" bigint NOT NULL
);

CREATE TABLE "menus" (
  "menu_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "menu_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "recipes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "menus" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "recipe_menu" ADD FOREIGN KEY ("recipe_id") REFERENCES "recipes" ("recipe_id");

ALTER TABLE "recipe_menu" ADD FOREIGN KEY ("menu_id") REFERENCES "menus" ("menu_id");
