Bracket Pair Colorizer
Go To Method
Horizon Theme
Indent-Rainbow
vscode-icons
kite
wsl


CREATE TABLE public.dish
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    name character varying(10) COLLATE pg_catalog."default" NOT NULL,
    class character varying(20) COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    price money NOT NULL DEFAULT 20,
    imgurl text COLLATE pg_catalog."default" NOT NULL DEFAULT ''::text,
    CONSTRAINT stu_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.dish
    OWNER to postgres;