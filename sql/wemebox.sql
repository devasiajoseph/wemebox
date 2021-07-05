
CREATE TABLE user_account(
       user_account_id serial PRIMARY KEY,
       phone VARCHAR(20) UNIQUE NOT NULL,
       email VARCHAR (355) UNIQUE NOT NULL default '',
       password VARCHAR (1024) NOT NULL,       
       full_name VARCHAR (250) NOT NULL default '',
       active BOOLEAN NOT NULL default false,
       created_on TIMESTAMP NOT NULL,
       last_login TIMESTAMP,
       role VARCHAR (20) not null default 'user'
);



create table tag (
       tag_id serial primary key,
       tag_name varchar(256) not null unique
);


create table page (
    page_id  serial primary key,
    page_type varchar(20),
    page_title varchar(1024),
    page_markdown text,
    page_html text,
    page_slug varchar(1024),
    last_edited  TIMESTAMP
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