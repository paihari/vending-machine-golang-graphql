CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
----------------------------------  v 0.1 ----------------------------------

-- FEDERALS
-- Having or relating to a system of government in which several states form a unity but remain independent in internal affairs.
-- AEVO, ASSL, ASAP ...

drop table federals;
CREATE TABLE federals (
    id SERIAL,
    uuid uuid DEFAULT uuid_generate_v4 (),
    name TEXT UNIQUE,
    full_name TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
);

INSERT INTO federals(name, full_name) VALUES ('AEVO', 'Avaloq Evolution');
INSERT INTO federals(name, full_name) VALUES ('ASSL', 'Avaloq Sourcing CH');
INSERT INTO federals(name, full_name) VALUES ('ASAP', 'Avaloq Sourcing AP');

-- CLOUD_PROVIDERS
-- List of Cloud Providers
-- AWS, GCP, OCI, AZURE etc ...


drop table cloud_providers;
CREATE TABLE cloud_providers (
    id SERIAL,
    uuid uuid DEFAULT uuid_generate_v4 (),
    name TEXT  UNIQUE,
    full_name TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
);


INSERT INTO cloud_providers(name, full_name) VALUES ('AWS', 'Amazon Web Services');
INSERT INTO cloud_providers(name, full_name) VALUES ('OCI', 'Oracle Cloud Infrastructure');
INSERT INTO cloud_providers(name, full_name) VALUES ('GCP', 'Google Cloud');
INSERT INTO cloud_providers(name, full_name) VALUES ('AZURE', 'Microsoft Azure');

-- CLOUD ESTATES
-- This are relation of the Federals at the heighest Level
-- Management Account to Accomodate Resident
-- Equivalent to AWS Organization/OCI Parent Tenency
-- Federals can have multiple Cloud Estates with Cloud Providers
-- Though it is understood, it may be single or couple of estates for each Federal entries


drop table cloud_estates;

CREATE TABLE cloud_estates (
    id SERIAL,
    uuid uuid DEFAULT uuid_generate_v4 (),
    name TEXT UNIQUE,
    description TEXT,
    federal TEXT REFERENCES federals(name),
    cloud_provider TEXT REFERENCES cloud_providers(name),
    federal_email_address TEXT,
    -- CLOUD ESTATE ID IS THE TOP LEVEL ID EG ORGANIZATION ID IN WHICH CHILD ACCOUNTS/RESIDENTS ARE BREWED
    cloud_estate_cid TEXT UNIQUE, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
);

INSERT INTO cloud_estates(name, description, federal, cloud_provider, cloud_estate_cid) VALUES ('AEVO-AWS-ESTATE', 'Cloud Estate for AEVO', 'AEVO', 'AWS', '-- ORG ID CREATED BY MANAGEMENT USER DURING INITIAL SIGIN IN--');



-- CLOUD ESTATE POLICY
-- THE POLICIES DECTATED BY FEDERAL LEVEL
-- NEED TO BE APPLIED TO ALL RESIDENTS/CHILD ACCOUNTS
-- EXAMPLE TAG POLICY

drop table cloud_estate_policies;
CREATE TABLE cloud_estate_policies (
    id SERIAL,
    uuid uuid DEFAULT uuid_generate_v4 (),
    name TEXT,
    description TEXT,
    cloud_estate TEXT REFERENCES cloud_estates(name),
    policy_type TEXT,
    -- ARN of the POLICY CREATED by the Management Account
    policy_cid TEXT,
    policy_json JSON,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
);

INSERT INTO cloud_estate_policies(name, description, cloud_estate, policy_type, policy_cid, policy_json) VALUES ('AEVO-AWS-ESTATE-TAG', 'TAG POLICY FROM FEDERAL LEVEL', 'AEVO-AWS_ESTATE', 'TAG', '--ARN--', 'JSON');

-- EXAMPLE OF TAG JSON
/* {
  "tags": {
    "CostCenter": {
      "tag_key": {
        "@@assign": "CostCenter",
        "@@operators_allowed_for_child_policies": [
          "@@none"
        ]
      }
    }
  }
}
 */



drop table clients;
CREATE TABLE clients (
    id SERIAL,
    uuid uuid DEFAULT uuid_generate_v4 (),
    name TEXT UNIQUE,
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


drop table classes;
CREATE TABLE classes (
    id SERIAL,
    uuid uuid DEFAULT uuid_generate_v4 (),
    name TEXT UNIQUE,
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


drop table stages;

CREATE TABLE stages (
    id SERIAL,
    uuid uuid DEFAULT uuid_generate_v4 (),
    name TEXT UNIQUE,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
);

INSERT INTO stages(name, description) VALUES ('development', 'The service is still under development.');
INSERT INTO stages(name, description) VALUES ('uat', 'The service is in user acceptance testing, data is not persisted.');
INSERT INTO stages(name, description) VALUES ('production', 'The service is in production and has to comply with the securty and compliance guidelines.');


-- RESIDENTS TABLE HOLDS MULTIPLE RESIDENT IN THE SPECIFIED CLOUD ESTATE
-- BUSINESS WISE IT HOLDS AWS CHILD ACCOUNTS/ OCI COMPARTMENTS WHERE THE WHOLE ECO-SYSTEM CAN BE BUILT
-- THE CLOUD ESTATE WHERE THE RESIDENT LIVES IS THE UNDERSTANDING BETWEEN THE FEDERAL AND CLOUD PROVIDER
-- THE POLICIES DECTATED BY FEDERAL LEVEL
-- NEED TO BE APPLIED TO ALL RESIDENTS/CHILD ACCOUNTS
-- EXAMPLE TAG POLICY


drop table residents;
CREATE TABLE residents (
    id SERIAL,
    uuid uuid DEFAULT uuid_generate_v4 (),
    name VARCHAR(32) UNIQUE,
    description TEXT,
    purchase_order TEXT,
    email_address TEXT,
    client TEXT REFERENCES clients(name),
    cloud_provider TEXT REFERENCES cloud_providers(name),
    -- CHILD ACCOUNT ID EG AWS
    resident_cid TEXT,
    -- ORGANIZATION ID/CLOUD ESTATE CID
    cloud_estate TEXT REFERENCES cloud_estates(name),
    cloud_estate_cid TEXT REFERENCES cloud_estates(cloud_estate_cid), 
    class TEXT REFERENCES classes(name),
    stage TEXT REFERENCES stages(name),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by TEXT,
    PRIMARY KEY(id)
);


-- RESIDENT ARE INSERTED BY VENDING MACHINES

INSERT INTO residents(name, description, purchase_order_id, email_address, client, cloud_provider, resident_cid, cloud_estate_cid, class, stage) 
VALUES ('AEVO_VEND1', 'The First Vending Machine', 'PO-123456', 'project@vending.avq', 'AEVO', 'AWS', '--ACCOUNT-ID--', '--ORG ESTATE ID', 'free', 'development');
