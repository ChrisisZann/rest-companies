

CREATE DATABASE chrisis_home
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_GB.UTF-8'
    LC_CTYPE = 'en_GB.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;


CREATE SCHEMA IF NOT EXISTS xm
    AUTHORIZATION postgres;

-- Table: xm.companies

-- DROP TABLE IF EXISTS xm.companies;

CREATE TABLE IF NOT EXISTS xm.companies
(
    company_uuid uuid NOT NULL,
    company_name character varying(15) COLLATE pg_catalog."default" NOT NULL,
    description character varying(3000) COLLATE pg_catalog."default",
    amount_of_employees integer NOT NULL,
    registered boolean NOT NULL,
    company_type character varying(20) COLLATE pg_catalog."default" NOT NULL,
    sys_creation_date timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sys_update_date timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT companies_pkey PRIMARY KEY (company_uuid),
    CONSTRAINT companies_company_name_key UNIQUE (company_name)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS xm.companies
    OWNER to postgres;

-- Table: xm.users

-- DROP TABLE IF EXISTS xm.users;

CREATE TABLE IF NOT EXISTS xm.users
(
    username character varying(16) COLLATE pg_catalog."default" NOT NULL,
    password character varying(16) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (username, password)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS xm.users
    OWNER to postgres;