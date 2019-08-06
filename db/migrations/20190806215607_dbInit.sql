
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- Create Customer Gender Table

CREATE TABLE IF NOT EXISTS miu.genders(
   id serial PRIMARY KEY,
   name VARCHAR (10)
);

-- Create Customer Account table

CREATE TABLE IF NOT EXISTS miu.accounts(
   id serial PRIMARY KEY,
   first_name VARCHAR (50) NOT NULL,
   last_name VARCHAR (50),
   password VARCHAR (50) NOT NULL,
   email_address VARCHAR (355) UNIQUE NOT NULL,
   phone_number VARCHAR (50),
   date_of_birth DATE,
   gender_id INT, 
   created_at TIMESTAMP NOT NULL,
   verified_at TIMESTAMP,
   last_login_at TIMESTAMP NOT NULL,
   last_modified_at TIMESTAMP NOT NULL
);


ALTER TABLE miu.accounts ADD CONSTRAINT fkey_customer_account_gender_id FOREIGN KEY (gender_id)
      REFERENCES miu.genders (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION;

-- Create Customer Device table

CREATE TABLE miu.devices(
   id SERIAL PRIMARY KEY,
   serial_number VARCHAR (50),
   name VARCHAR (50),
   operating_system VARCHAR (50)
);

-- Create Account Device table

CREATE TABLE miu.account_devices(
   account_id INTEGER NOT NULL,
   device_id INTEGER NOT NULL,
  PRIMARY KEY (account_id, device_id),
  CONSTRAINT fkey_account_device_account_id FOREIGN KEY (account_id)
      REFERENCES miu.accounts (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT fkey_account_device_device_id FOREIGN KEY (device_id)
      REFERENCES miu.devices (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create Customer Location table

CREATE TABLE miu.locations(
   id SERIAL PRIMARY KEY,
   account_id INTEGER NOT NULL,
   city VARCHAR (50),
   region VARCHAR (50),
   country VARCHAR (355) NOT NULL,
   ip_address VARCHAR (50),
   recorded_at TIMESTAMP NOT NULL,
   CONSTRAINT fkey_customer_loc_account_id FOREIGN KEY (account_id)
      REFERENCES miu.accounts (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

-- Create Generic Config Table

CREATE TABLE miu.configs(
   id SERIAL PRIMARY KEY,
   name VARCHAR(100) NOT NULL,
   type VARCHAR(50),
   value VARCHAR(5000)
);

-- Create Pins Table

CREATE TABLE miu.pins(
   email_address VARCHAR(355) PRIMARY KEY,
   pin INTEGER NOT NULL,
   purpose VARCHAR(50),
   expired_at TIMESTAMP 
);

-- Create Revoked Tokens Table

CREATE TABLE miu.revoked_tokens(
   jti uuid PRIMARY KEY,
   token VARCHAR(700) NOT NULL,
   account_id INTEGER,
   expired_at TIMESTAMP,
   revoked_at TIMESTAMP
);



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE miu.account_devices;
DROP TABLE miu.configs;
DROP TABLE miu.devices;
DROP TABLE miu.locations;
DROP TABLE miu.revoked_tokens;
DROP TABLE miu.pins;
DROP TABLE miu.accounts;
DROP TABLE miu.genders;