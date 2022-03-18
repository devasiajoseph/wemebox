CREATE TABLE domain(
       domain_id serial PRIMARY KEY,
       domain_name varchar(255)
);


CREATE TABLE user_account(
       user_account_id serial PRIMARY KEY,
       phone VARCHAR(20) UNIQUE NOT NULL,
       email VARCHAR (355) UNIQUE NOT NULL default '',
       password VARCHAR (1024) NOT NULL,       
       full_name VARCHAR (250) NOT NULL default '',
       active BOOLEAN NOT NULL default false,
       created_on TIMESTAMP NOT NULL,
       last_login TIMESTAMP,
       
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
       role VARCHAR (20) not null default 'user'
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

create table page_bkp (
    page_id  serial primary key,
    page_type varchar(20),
    page_title varchar(1024),
    page_markdown text,
    page_html text,
    page_slug varchar(1024),
    last_edited  TIMESTAMP,
    domain_id integer,
    CONSTRAINT page_domain_id_fkey FOREIGN KEY (domain_id)
    REFERENCES domain (domain_id) MATCH SIMPLE 
    ON DELETE CASCADE
);


create table page (
    page_id  serial primary key,
    page_slug varchar(1024),
    domain_id integer,
    page_file varchar(255),
    base_page_file varchar(255),
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
    last_edited  TIMESTAMP
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