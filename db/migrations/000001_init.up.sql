CREATE TABLE public.user(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) not null,
    username VARCHAR(100) not null unique,
    password_hash VARCHAR(100) not null
);
CREATE TABLE public.client(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);
CREATE TABLE public.сar(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    carbrand VARCHAR(100),
    free boolean not null default false,
    client_id UUID ,
    CONSTRAINT client_fk FOREIGN KEY (client_id) REFERENCES public.client(id) on delete cascade
);