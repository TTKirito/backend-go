
CREATE TYPE "positions" as ENUM (
    'Design',
    'Develop'
);
CREATE TYPE "genders" as ENUM (
    'Man',
    'Women'
);
create TYPE "status" as ENUM (
    'Active',
    'Inactive'
);

CREATE TABLE accounts (
    "id" bigserial PRIMARY KEY,
    "owner" varchar NOT NULL,
    "position" positions NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "gender" genders NOT NULL,
    "dob" timestamptz NOT NULL,
    "status" status NOT NULL
);

CREATE INDEX ON accounts ("owner");


CREATE TYPE "event_types" as ENUM (
    'Event',
    'Meeting'
);

CREATE TYPE "visit_types" as ENUM (
    'Office',
    'Online'
);

CREATE TABLE events (
    "id" bigserial PRIMARY KEY,
    "title" varchar,
    "start_time" timestamptz NOT NULL,
    "end_time" timestamptz NOT NULL,
    "is_emegency" BIT NOT NULL DEFAULT '0',
    "owner" bigserial NOT NULL,
    "note" varchar,
    "type" event_types NOT NULL,
    "visit_type" visit_types NOT NULL,
    "meeting" varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON events ("owner", "start_time", "end_time");

COMMENT ON COLUMN "events"."start_time" IS 'required';

ALTER TABLE "events" ADD FOREIGN KEY ("owner") REFERENCES "accounts" ("id");


CREATE TABLE locations (
    "id" bigserial PRIMARY KEY,
    "event" bigserial NOT NULL,
    "lat" varchar NOT NULL,
    "long" varchar NOT NULL,
    "block_no" integer,
    "apartment_name" varchar,
    "apartment_number" integer,
    "street" varchar NOT NULL,
    "city" varchar NOT NULL,
    "country" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())

);

CREATE INDEX ON locations ("event");

ALTER TABLE "locations" ADD FOREIGN KEY ("event") REFERENCES "events" ("id");


CREATE TABLE participants(
    "id" bigserial PRIMARY KEY,
    "account" bigserial NOT NULL,
    "event" bigserial NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON participants ("event", "account");

ALTER TABLE "participants" ADD FOREIGN KEY ("event") REFERENCES "events" ("id");
ALTER TABLE "participants" ADD FOREIGN KEY ("account") REFERENCES "accounts" ("id");

