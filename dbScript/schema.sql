-- Create necessary PG extensions
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- # ======================== tables schema =======================
-- Author table
CREATE TABLE IF NOT EXISTS author (
    uid UUID DEFAULT gen_random_uuid() NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_digest VARCHAR(511),

    first_name VARCHAR(63) NOT NULL,
    last_name VARCHAR(63) NOT NULL,
    avatar_url VARCHAR(255) NOT NULL,
    about TEXT,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT author_id_pkey PRIMARY KEY (uid),
    CONSTRAINT author_email_uniq UNIQUE (email)

);

CREATE TABLE IF NOT EXISTS blog (
    uid UUID DEFAULT gen_random_uuid() NOT NULL,

    status VARCHAR(5) NOT NULL DEFAULT 'draft',

    title TEXT,
    image_url VARCHAR(255),
    authored_by_id UUID NOT NULL,
    authored_on TIMESTAMPTZ NOT NULL,
    published_on timestamptz,

    body TEXT,
    categories VARCHAR[] default '{}',

    tags VARCHAR[] default '{}',

    claps int4 NOT NULL DEFAULT 0,
    read_time INT NOT NULL DEFAULT 0,
    access_count int4 NOT NULL DEFAULT 0,

    CONSTRAINT blog_id_pkey PRIMARY KEY (uid),
    CONSTRAINT blog_author_fkey FOREIGN KEY (authored_by_id) REFERENCES author
);


CREATE TABLE IF NOT EXISTS comment (
  uid UUID DEFAULT gen_random_uuid() NOT NULL,
  blog_id UUID REFERENCES blog NOT NULL,
  commented_by_id UUID REFERENCES author NOT NULL,
  created_at TIMESTAMPTZ NOT NULL
);
