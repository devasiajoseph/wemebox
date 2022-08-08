CREATE TABLE user_account(
       user_account_id serial PRIMARY KEY,
       email VARCHAR (355) UNIQUE NOT NULL,
       password VARCHAR (1024) NOT NULL,       
       full_name VARCHAR (250),
       address_line_1 VARCHAR (255),
       address_line_2 VARCHAR (255),
       city VARCHAR (255),
       state_province VARCHAR (255),
       country_code VARCHAR (3) NOT NULL default 'IND',
       postal_code VARCHAR(20),
       phone VARCHAR(20),
       user_role VARCHAR(10) NULL,
       active BOOLEAN NOT NULL default false,
       created_on TIMESTAMP NOT NULL,
       last_login TIMESTAMP
       );


CREATE INDEX index_user_account_email ON user_account(email);
CREATE INDEX index_user_account_phone ON user_account(phone);


CREATE TABLE user_keys(
       user_keys_id serial PRIMARY KEY,
       user_account_id integer NOT NULL,
       registration_key VARCHAR (255),
       forgot_password_key VARCHAR (255),
       CONSTRAINT user_registration_user_account_id_fkey FOREIGN KEY (user_account_id)
       REFERENCES user_account (user_account_id) MATCH SIMPLE 
       ON DELETE CASCADE
       );

CREATE TABLE user_session(
       user_session_id serial PRIMARY KEY,
       user_account_id integer NOT NULL,
       auth_token VARCHAR(64),
       session_expiry TIMESTAMP NOT NULL,
       CONSTRAINT user_session_user_account_id_fkey FOREIGN KEY (user_account_id)
       REFERENCES user_account (user_account_id) MATCH SIMPLE 
       ON DELETE CASCADE
       );
      

