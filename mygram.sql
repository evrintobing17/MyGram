-- -------------------------------------------------------------
-- TablePlus 4.6.2(410)
--
-- https://tableplus.com/
--
-- Database: mygram
-- Generation Time: 2022-03-28 04:41:16.7740
-- -------------------------------------------------------------

CREATE DATABASE IF NOT EXISTS mygram;

DROP TABLE IF EXISTS "public"."comments";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS comments_id_seq;

-- Table Definition
CREATE TABLE "public"."comments" (
    "id" int4 NOT NULL DEFAULT nextval('comments_id_seq'::regclass),
    "user_id" int4,
    "photos_id" int4,
    "message" text NOT NULL,
    "created_at" date,
    "updated_at" date,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."photos";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS photos_id_seq;

-- Table Definition
CREATE TABLE "public"."photos" (
    "id" int4 NOT NULL DEFAULT nextval('photos_id_seq'::regclass),
    "title" varchar(255) NOT NULL,
    "caption" varchar(255),
    "photo_url" varchar(255),
    "user_id" int4,
    "created_at" date,
    "updated_at" date,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."social_medias";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS social_medias_id_seq;

-- Table Definition
CREATE TABLE "public"."social_medias" (
    "id" int4 NOT NULL DEFAULT nextval('social_medias_id_seq'::regclass),
    "name" varchar(50) NOT NULL,
    "social_media_url" text NOT NULL,
    "user_id" int4,
    "created_at" date,
    "updated_at" date,
    PRIMARY KEY ("id")
);

DROP TABLE IF EXISTS "public"."users";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "public"."users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "username" varchar(50) NOT NULL,
    "email" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    "age" int4 NOT NULL,
    "created_at" date,
    "updated_at" date,
    PRIMARY KEY ("id")
);

ALTER TABLE "public"."comments" ADD FOREIGN KEY ("photos_id") REFERENCES "public"."photos"("id") ON DELETE CASCADE;
ALTER TABLE "public"."comments" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE;
ALTER TABLE "public"."photos" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE;
ALTER TABLE "public"."social_medias" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("id") ON DELETE CASCADE;
