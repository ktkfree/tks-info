CREATE DATABASE tks;
\c tks;
CREATE TABLE application_groups
(
    name character varying(50) COLLATE pg_catalog."default",
    id uuid primary key,
    type bigint,
    workflow_id character varying(100) COLLATE pg_catalog."default",
    status integer,
    status_desc character varying(10000) COLLATE pg_catalog."default",
    cluster_id uuid,
    external_label character varying(50) COLLATE pg_catalog."default",
    updated_at timestamp with time zone,
    created_at timestamp with time zone
);
CREATE TABLE applications
(
    id uuid primary key,
    type bigint,
    app_group_id uuid,
    endpoint character varying(200) COLLATE pg_catalog."default",
    metadata json,
    updated_at timestamp with time zone,
    created_at timestamp with time zone
);
