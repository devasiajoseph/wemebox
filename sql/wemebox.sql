CREATE TABLE domain(
       domain_id serial PRIMARY KEY,
       domain_name varchar(255) not null unique,
       domain_dir varchar(255) not null unique
);


/*user tables*/
CREATE TABLE user_account(
       user_account_id serial PRIMARY KEY,
       full_name VARCHAR (250),   
       email VARCHAR (355) UNIQUE NOT NULL,
       password VARCHAR (1024) NOT NULL,
       phone VARCHAR(20),
       role VARCHAR(10) NOT NULL default 'user',
       active BOOLEAN NOT NULL default false,
       created_on TIMESTAMP NOT NULL,
       last_login TIMESTAMP
       );


CREATE INDEX index_user_account_email ON user_account(email);
CREATE INDEX index_user_account_phone ON user_account(phone);

CREATE TABLE user_profile(
       user_profile_id serial PRIMARY KEY,
       address_line_1 VARCHAR (255),
       address_line_2 VARCHAR (255),
       city VARCHAR (255),
       state_province VARCHAR (255),
       country_id VARCHAR (3) NOT NULL default 'US',
       postal_code VARCHAR(20),
       user_account_id integer NOT NULL,
       CONSTRAINT user_profile_user_account_id_fkey FOREIGN KEY (user_account_id)
       REFERENCES user_account (user_account_id) MATCH SIMPLE 
       ON DELETE CASCADE
);

CREATE TABLE user_keys(
       user_keys_id serial PRIMARY KEY,
       user_account_id integer NOT NULL,
       key_name VARCHAR (255),
       key_value VARCHAR (255),
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

CREATE TABLE domain_user_role (
       role_id serial PRIMARY KEY,
       role VARCHAR (20) not null default 'user',
       user_account_id integer,
       domain_id integer,
       CONSTRAINT user_account_id_fkey FOREIGN KEY (user_account_id)
       REFERENCES user_account (user_account_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       CONSTRAINT domain_id_fkey FOREIGN KEY (domain_id)
       REFERENCES domain (domain_id) MATCH SIMPLE 
       ON DELETE CASCADE
);


create table tag (
       tag_id serial primary key,
       tag_name varchar(256) not null unique
);



create table page (
    page_id  serial primary key,
    page_slug varchar(1024),
    page_file varchar(255),
    base_page_file varchar(255) not null default 'base.html',
    domain_id integer,
    CONSTRAINT page_domain_id_fkey FOREIGN KEY (domain_id)
    REFERENCES domain (domain_id) MATCH SIMPLE 
    ON DELETE CASCADE
);


create table blog (
    blog_id  serial primary key,
    blog_type varchar(20),
    blog_title varchar(1024),
    blog_markdown text,
    blog_html text,
    blog_slug varchar(1024),
    created TIMESTAMP,
    last_edited  TIMESTAMP,
    domain_id integer,
    CONSTRAINT blog_domain_id_fkey FOREIGN KEY (domain_id)
    REFERENCES domain (domain_id) MATCH SIMPLE 
    ON DELETE CASCADE
);



create table blog_tag(
       blog_tag_id serial primary key,
       blog_id integer,
       tag_id integer,
       unique (blog_id, tag_id),
       CONSTRAINT tag_id_fkey FOREIGN KEY (tag_id)
       REFERENCES tag (tag_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       CONSTRAINT blog_id_fkey FOREIGN KEY (blog_id)
       REFERENCES blog (blog_id) MATCH SIMPLE 
       ON DELETE CASCADE
);


create table account(
       subs_id serial primary key,
       domain_id integer unique,
       balance integer,
       CONSTRAINT account_domain_id_fkey FOREIGN KEY (domain_id)
       REFERENCES domain (domain_id) MATCH SIMPLE 
       ON DELETE CASCADE

);

create table trx(
       trx_id serial primary key,
       domain_id integer,
       credit integer not null default 0,
       debit integer not null default 0,
       trx_time TIMESTAMP not null default now(),
       notes text not null default '',
       CONSTRAINT trx_domain_id_fkey FOREIGN KEY (domain_id)
       REFERENCES domain (domain_id) MATCH SIMPLE 
       ON DELETE CASCADE
);