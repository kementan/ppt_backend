CREATE TABLE ppt_roles (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE ppt_services (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE ppt_service_endpoints (
    id BIGSERIAL PRIMARY KEY,
    service_id BIGSERIAL NOT NULL,
    endpoint VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX ON "ppt_service_endpoints" ("service_id");
ALTER TABLE "ppt_service_endpoints" ADD FOREIGN KEY ("service_id") REFERENCES "ppt_services" ("id");

CREATE TABLE ppt_service_accesses (
    id BIGSERIAL PRIMARY KEY,
    role_id BIGSERIAL NOT NULL,
    service_id BIGSERIAL NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX ON "ppt_service_accesses" ("role_id");
ALTER TABLE "ppt_service_accesses" ADD FOREIGN KEY ("role_id") REFERENCES "ppt_roles" ("id");

CREATE TABLE ppt_users (
    id BIGINT PRIMARY KEY,
    role_id BIGSERIAL NOT NULL,
    name VARCHAR(50) NOT NULL,
    nik VARCHAR(20) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX ON "ppt_users" ("role_id");
ALTER TABLE "ppt_users" ADD FOREIGN KEY ("role_id") REFERENCES "ppt_roles" ("id");