CREATE TABLE users(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) not null,
    username VARCHAR(100) not null unique,
    password_hash VARCHAR(100) not null
);
CREATE TABLE clients(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    money INT DEFAULT 0
);
CREATE TABLE cars(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    carbrand VARCHAR(100),
    free boolean not null default true,
    clientid UUID,
    CONSTRAINT clients_fk FOREIGN KEY (clientid) REFERENCES clients(id)
);