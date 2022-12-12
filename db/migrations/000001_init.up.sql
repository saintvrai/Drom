CREATE TABLE public.users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) not null,
    username VARCHAR(100) not null unique,
    password_hash VARCHAR(100) not null
);
CREATE TABLE public.clients(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL
);
CREATE TABLE public.сars(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    carbrand VARCHAR(100),
    free boolean not null default false,
    client_id UUID ,
    CONSTRAINT clients_fk FOREIGN KEY (client_id) REFERENCES public.clients(id) on delete cascade
);