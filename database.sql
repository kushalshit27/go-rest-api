BEGIN;


CREATE TABLE public.blogs
(
    title character varying(255) NOT NULL,
    description text NOT NULL,
    status boolean NOT NULL,
    created_at timestamp(0) without time zone NOT NULL,
    updated_at timestamp(0) without time zone NOT NULL,
    deleted_at timestamp(0) without time zone,
    created_by bigint NOT NULL,
    id integer NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public.passwords
(
    id bigint NOT NULL,
    password character varying(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE public.users
(
    name varying(255) NOT NULL,
    email varying(255) NOT NULL,
    age integer NOT NULL,
    role character varying NOT NULL,
    created_at timestamp(0) without time zone NOT NULL,
    updated_at timestamp(0) without time zone NOT NULL,
    deleted_at timestamp(0) without time zone,
    status boolean NOT NULL,
    id integer NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE public.blogs
    ADD FOREIGN KEY (created_by)
    REFERENCES public.users (id)
    NOT VALID;

END;