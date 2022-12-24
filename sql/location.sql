CREATE EXTENSION Postgis;

CREATE TABLE country(
       country_id serial PRIMARY KEY,
       country_name varchar(255) UNIQUE NOT NULL,
       country_phone_code varchar(10) NOT NULL
);

INSERT INTO country (country_name,country_phone_code) VALUES ('India','91');
CREATE TABLE state(
       state_id serial PRIMARY KEY,
       state_name varchar(255) NOT NULL,
       country_id integer,
       CONSTRAINT state_country_id_fkey FOREIGN KEY (country_id)
       REFERENCES country (country_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       unique (state_name, country_id)
);
INSERT INTO state (state_name, country_id) VALUES
('Kerala', 1);

CREATE TABLE district(
       district_id serial PRIMARY KEY,
       district_name varchar(255) NOT NULL,
       state_id integer,
       CONSTRAINT dsitrict_state_id_fkey FOREIGN KEY (state_id)
       REFERENCES state (state_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       unique (district_name, state_id)
);

INSERT INTO district (district_id, district_name, state_id) VALUES
(1,'Alappuzha',1),
(2,'Ernakulam',1),
(3,'Idukki',1),
(4,'Kannur',1),
(5,'Kasaragod',1),
(6,'Kollam',1),
(7,'Kottayam',1),
(8,'Kozhikode',1),
(9,'Malappuram',1),
(10,'Palakkad',1),
(11,'Pathanamthitta',1),
(12,'Thiruvananthapuram',1),
(13,'Thrissur',1),
(14,'Wayanad',1),
(15,'Lakshadweep',1),
(16,'Coimbatore',1),
(17,'Tenkasi',1),
(18,'Kanyakumari',1),
(19, 'Tirunelveli',1);


CREATE TABLE city_type (
       city_type_id serial PRIMARY KEY,
       city_type_name varchar(100)
);



CREATE TABLE city(
       city_id serial PRIMARY KEY,
       city_name varchar(255) NOT NULL,
       city_type varchar(32),
       location geography,
       latitude varchar(32) not null default '',
       longitude varchar(32) not null default '',
       district_id integer,
       district_centre boolean not null default false,
       CONSTRAINT city_district_id_fkey FOREIGN KEY (district_id)
       REFERENCES district (district_id) MATCH SIMPLE 
       ON DELETE CASCADE,
       unique (city_name, district_id)
);