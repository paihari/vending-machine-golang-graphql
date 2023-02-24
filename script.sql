CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table clients
CREATE TABLE clients (
    id SERIAL,
    universal_id uuid DEFAULT uuid_generate_v4 (),
    name TEXT,
    full_name TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
);

INSERT INTO clients(name, full_name) VALUES ('AEVO', 'Avaloq Evolution');
INSERT INTO clients(name, full_name) VALUES ('ASSL', 'Avaloq Sourcing CH');
INSERT INTO clients(name, full_name) VALUES ('ASAP', 'Avaloq Sourcing AP');
INSERT INTO clients(name, full_name) VALUES ('APLH', 'Aplha Bank AG');





CREATE TABLE cloud_providers (
    id SERIAL,
    universal_id uuid DEFAULT uuid_generate_v4 (),
    name TEXT,
    full_name TEXT,
    root_cid TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
);

INSERT INTO cloud_providers(name, full_name, root_cid) VALUES ('AWS', 'Amazon Web Services', 'ROOT AWS ID');
INSERT INTO cloud_providers(name, full_name, root_cid) VALUES ('OCI', 'Oracle Cloud Infrastructure', 'ROOT OCI ID');
INSERT INTO cloud_providers(name, full_name, root_cid) VALUES ('GCP', 'Google Cloud', 'ROOT GCP ID' );
INSERT INTO cloud_providers(name, full_name, root_cid) VALUES ('AZURE', 'Microsoft Azure', 'ROOT AZURE ID');



CREATE TABLE classes (
    id SERIAL,
    universal_id uuid DEFAULT uuid_generate_v4 (),
    name TEXT,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
)

INSERT INTO classes(name, description) VALUES ('free', 'The service is deployed into a free-tier account.');
INSERT INTO classes(name, description) VALUES ('trial', 'The service represents a sponsored proof-of-concept and has to abide to budget restrictions.');
INSERT INTO classes(name, description) VALUES ('payg', 'Provisioning resources immedeately increases the cloud bill.');
INSERT INTO classes(name, description) VALUES ('commit', 'The service budget is planned ahead and resources are accounted against a credit commitment.');




CREATE TABLE stages (
    id SERIAL,
    universal_id uuid DEFAULT uuid_generate_v4 (),
    name TEXT,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
)

INSERT INTO stages(name, description) VALUES ('development', 'The service is still under development.');
INSERT INTO stages(name, description) VALUES ('uat', 'The service is in user acceptance testing, data is not persisted.');
INSERT INTO stages(name, description) VALUES ('production', 'The service is in production and has to comply with the securty and compliance guidelines.');


CREATE TABLE residents (
    id SERIAL,
    universal_id uuid DEFAULT uuid_generate_v4 (),
    name VARCHAR(32) UNIQUE,
    description TEXT,
    purchase_order_id TEXT,
    client TEXT,
    cloud_provider TEXT,
    resident_cid TEXT,
    root_cid TEXT,
    class TEXT,
    stage TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
);
