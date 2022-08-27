--
-- PostgreSQL database dump
--

-- Dumped from database version 10.21 (Ubuntu 10.21-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.21 (Ubuntu 10.21-0ubuntu0.18.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: blog; Type: TABLE; Schema: public; Owner: devasia
--

CREATE TABLE public.blog (
    blog_id integer NOT NULL,
    blog_type character varying(20),
    blog_title character varying(1024),
    blog_markdown text,
    blog_html text,
    blog_slug character varying(1024),
    created timestamp without time zone,
    last_edited timestamp without time zone,
    domain_id integer
);


ALTER TABLE public.blog OWNER TO devasia;

--
-- Name: blog_blog_id_seq; Type: SEQUENCE; Schema: public; Owner: devasia
--

CREATE SEQUENCE public.blog_blog_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.blog_blog_id_seq OWNER TO devasia;

--
-- Name: blog_blog_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: devasia
--

ALTER SEQUENCE public.blog_blog_id_seq OWNED BY public.blog.blog_id;


--
-- Name: blog_tag; Type: TABLE; Schema: public; Owner: devasia
--

CREATE TABLE public.blog_tag (
    blog_tag_id integer NOT NULL,
    blog_id integer,
    tag_id integer
);


ALTER TABLE public.blog_tag OWNER TO devasia;

--
-- Name: blog_tag_blog_tag_id_seq; Type: SEQUENCE; Schema: public; Owner: devasia
--

CREATE SEQUENCE public.blog_tag_blog_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.blog_tag_blog_tag_id_seq OWNER TO devasia;

--
-- Name: blog_tag_blog_tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: devasia
--

ALTER SEQUENCE public.blog_tag_blog_tag_id_seq OWNED BY public.blog_tag.blog_tag_id;


--
-- Name: domain; Type: TABLE; Schema: public; Owner: devasia
--

CREATE TABLE public.domain (
    domain_id integer NOT NULL,
    domain_name character varying(255),
    domain_dir character varying(255)
);


ALTER TABLE public.domain OWNER TO devasia;

--
-- Name: domain_domain_id_seq; Type: SEQUENCE; Schema: public; Owner: devasia
--

CREATE SEQUENCE public.domain_domain_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.domain_domain_id_seq OWNER TO devasia;

--
-- Name: domain_domain_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: devasia
--

ALTER SEQUENCE public.domain_domain_id_seq OWNED BY public.domain.domain_id;


--
-- Name: domain_user_role; Type: TABLE; Schema: public; Owner: devasia
--

CREATE TABLE public.domain_user_role (
    role_id integer NOT NULL,
    role character varying(20) DEFAULT 'user'::character varying NOT NULL,
    user_account_id integer,
    domain_id integer
);


ALTER TABLE public.domain_user_role OWNER TO devasia;

--
-- Name: domain_user_role_role_id_seq; Type: SEQUENCE; Schema: public; Owner: devasia
--

CREATE SEQUENCE public.domain_user_role_role_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.domain_user_role_role_id_seq OWNER TO devasia;

--
-- Name: domain_user_role_role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: devasia
--

ALTER SEQUENCE public.domain_user_role_role_id_seq OWNED BY public.domain_user_role.role_id;


--
-- Name: page; Type: TABLE; Schema: public; Owner: devasia
--

CREATE TABLE public.page (
    page_id integer NOT NULL,
    page_slug character varying(1024),
    page_file character varying(255),
    base_page_file character varying(255) DEFAULT 'base.html'::character varying NOT NULL,
    domain_id integer
);


ALTER TABLE public.page OWNER TO devasia;

--
-- Name: page_page_id_seq; Type: SEQUENCE; Schema: public; Owner: devasia
--

CREATE SEQUENCE public.page_page_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.page_page_id_seq OWNER TO devasia;

--
-- Name: page_page_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: devasia
--

ALTER SEQUENCE public.page_page_id_seq OWNED BY public.page.page_id;


--
-- Name: tag; Type: TABLE; Schema: public; Owner: devasia
--

CREATE TABLE public.tag (
    tag_id integer NOT NULL,
    tag_name character varying(256) NOT NULL
);


ALTER TABLE public.tag OWNER TO devasia;

--
-- Name: tag_tag_id_seq; Type: SEQUENCE; Schema: public; Owner: devasia
--

CREATE SEQUENCE public.tag_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tag_tag_id_seq OWNER TO devasia;

--
-- Name: tag_tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: devasia
--

ALTER SEQUENCE public.tag_tag_id_seq OWNED BY public.tag.tag_id;


--
-- Name: user_account; Type: TABLE; Schema: public; Owner: devasia
--

CREATE TABLE public.user_account (
    user_account_id integer NOT NULL,
    phone character varying(20) NOT NULL,
    email character varying(355) DEFAULT ''::character varying NOT NULL,
    password character varying(1024) NOT NULL,
    full_name character varying(250) DEFAULT ''::character varying NOT NULL,
    active boolean DEFAULT false NOT NULL,
    created_on timestamp without time zone NOT NULL,
    last_login timestamp without time zone
);


ALTER TABLE public.user_account OWNER TO devasia;

--
-- Name: user_account_user_account_id_seq; Type: SEQUENCE; Schema: public; Owner: devasia
--

CREATE SEQUENCE public.user_account_user_account_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_account_user_account_id_seq OWNER TO devasia;

--
-- Name: user_account_user_account_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: devasia
--

ALTER SEQUENCE public.user_account_user_account_id_seq OWNED BY public.user_account.user_account_id;


--
-- Name: user_session; Type: TABLE; Schema: public; Owner: devasia
--

CREATE TABLE public.user_session (
    user_session_id integer NOT NULL,
    user_account_id integer NOT NULL,
    auth_token character varying(64),
    session_expiry timestamp without time zone NOT NULL
);


ALTER TABLE public.user_session OWNER TO devasia;

--
-- Name: user_session_user_session_id_seq; Type: SEQUENCE; Schema: public; Owner: devasia
--

CREATE SEQUENCE public.user_session_user_session_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_session_user_session_id_seq OWNER TO devasia;

--
-- Name: user_session_user_session_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: devasia
--

ALTER SEQUENCE public.user_session_user_session_id_seq OWNED BY public.user_session.user_session_id;


--
-- Name: blog blog_id; Type: DEFAULT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.blog ALTER COLUMN blog_id SET DEFAULT nextval('public.blog_blog_id_seq'::regclass);


--
-- Name: blog_tag blog_tag_id; Type: DEFAULT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.blog_tag ALTER COLUMN blog_tag_id SET DEFAULT nextval('public.blog_tag_blog_tag_id_seq'::regclass);


--
-- Name: domain domain_id; Type: DEFAULT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.domain ALTER COLUMN domain_id SET DEFAULT nextval('public.domain_domain_id_seq'::regclass);


--
-- Name: domain_user_role role_id; Type: DEFAULT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.domain_user_role ALTER COLUMN role_id SET DEFAULT nextval('public.domain_user_role_role_id_seq'::regclass);


--
-- Name: page page_id; Type: DEFAULT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.page ALTER COLUMN page_id SET DEFAULT nextval('public.page_page_id_seq'::regclass);


--
-- Name: tag tag_id; Type: DEFAULT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.tag ALTER COLUMN tag_id SET DEFAULT nextval('public.tag_tag_id_seq'::regclass);


--
-- Name: user_account user_account_id; Type: DEFAULT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.user_account ALTER COLUMN user_account_id SET DEFAULT nextval('public.user_account_user_account_id_seq'::regclass);


--
-- Name: user_session user_session_id; Type: DEFAULT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.user_session ALTER COLUMN user_session_id SET DEFAULT nextval('public.user_session_user_session_id_seq'::regclass);


--
-- Data for Name: blog; Type: TABLE DATA; Schema: public; Owner: devasia
--

COPY public.blog (blog_id, blog_type, blog_title, blog_markdown, blog_html, blog_slug, created, last_edited, domain_id) FROM stdin;
\.


--
-- Data for Name: blog_tag; Type: TABLE DATA; Schema: public; Owner: devasia
--

COPY public.blog_tag (blog_tag_id, blog_id, tag_id) FROM stdin;
\.


--
-- Data for Name: domain; Type: TABLE DATA; Schema: public; Owner: devasia
--

COPY public.domain (domain_id, domain_name, domain_dir) FROM stdin;
1	localhost	localhost
2	palathrainfra.com	palathrainfra.com
3	wemebox.com	wemebox.com
4	centipair.com	centipair.com
\.


--
-- Data for Name: domain_user_role; Type: TABLE DATA; Schema: public; Owner: devasia
--

COPY public.domain_user_role (role_id, role, user_account_id, domain_id) FROM stdin;
\.


--
-- Data for Name: page; Type: TABLE DATA; Schema: public; Owner: devasia
--

COPY public.page (page_id, page_slug, page_file, base_page_file, domain_id) FROM stdin;
1	home	home.html	base.html	1
2	home	home.html	base.html	2
3	about	about.html	base.html	2
4	contact	contact.html	base.html	2
5	our-business	our-business.html	base.html	2
6	services	services.html	base.html	2
7	projects	projects.html	base.html	2
8	home	home.html	base.html	3
9	home	home.html	base.html	4
\.


--
-- Data for Name: tag; Type: TABLE DATA; Schema: public; Owner: devasia
--

COPY public.tag (tag_id, tag_name) FROM stdin;
\.


--
-- Data for Name: user_account; Type: TABLE DATA; Schema: public; Owner: devasia
--

COPY public.user_account (user_account_id, phone, email, password, full_name, active, created_on, last_login) FROM stdin;
\.


--
-- Data for Name: user_session; Type: TABLE DATA; Schema: public; Owner: devasia
--

COPY public.user_session (user_session_id, user_account_id, auth_token, session_expiry) FROM stdin;
\.


--
-- Name: blog_blog_id_seq; Type: SEQUENCE SET; Schema: public; Owner: devasia
--

SELECT pg_catalog.setval('public.blog_blog_id_seq', 1, false);


--
-- Name: blog_tag_blog_tag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: devasia
--

SELECT pg_catalog.setval('public.blog_tag_blog_tag_id_seq', 1, false);


--
-- Name: domain_domain_id_seq; Type: SEQUENCE SET; Schema: public; Owner: devasia
--

SELECT pg_catalog.setval('public.domain_domain_id_seq', 4, true);


--
-- Name: domain_user_role_role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: devasia
--

SELECT pg_catalog.setval('public.domain_user_role_role_id_seq', 1, false);


--
-- Name: page_page_id_seq; Type: SEQUENCE SET; Schema: public; Owner: devasia
--

SELECT pg_catalog.setval('public.page_page_id_seq', 9, true);


--
-- Name: tag_tag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: devasia
--

SELECT pg_catalog.setval('public.tag_tag_id_seq', 1, false);


--
-- Name: user_account_user_account_id_seq; Type: SEQUENCE SET; Schema: public; Owner: devasia
--

SELECT pg_catalog.setval('public.user_account_user_account_id_seq', 1, false);


--
-- Name: user_session_user_session_id_seq; Type: SEQUENCE SET; Schema: public; Owner: devasia
--

SELECT pg_catalog.setval('public.user_session_user_session_id_seq', 1, false);


--
-- Name: blog blog_pkey; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.blog
    ADD CONSTRAINT blog_pkey PRIMARY KEY (blog_id);


--
-- Name: blog_tag blog_tag_blog_id_tag_id_key; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.blog_tag
    ADD CONSTRAINT blog_tag_blog_id_tag_id_key UNIQUE (blog_id, tag_id);


--
-- Name: blog_tag blog_tag_pkey; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.blog_tag
    ADD CONSTRAINT blog_tag_pkey PRIMARY KEY (blog_tag_id);


--
-- Name: domain domain_pkey; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.domain
    ADD CONSTRAINT domain_pkey PRIMARY KEY (domain_id);


--
-- Name: domain_user_role domain_user_role_pkey; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.domain_user_role
    ADD CONSTRAINT domain_user_role_pkey PRIMARY KEY (role_id);


--
-- Name: page page_pkey; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.page
    ADD CONSTRAINT page_pkey PRIMARY KEY (page_id);


--
-- Name: tag tag_pkey; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (tag_id);


--
-- Name: tag tag_tag_name_key; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_tag_name_key UNIQUE (tag_name);


--
-- Name: user_account user_account_email_key; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.user_account
    ADD CONSTRAINT user_account_email_key UNIQUE (email);


--
-- Name: user_account user_account_phone_key; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.user_account
    ADD CONSTRAINT user_account_phone_key UNIQUE (phone);


--
-- Name: user_account user_account_pkey; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.user_account
    ADD CONSTRAINT user_account_pkey PRIMARY KEY (user_account_id);


--
-- Name: user_session user_session_pkey; Type: CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.user_session
    ADD CONSTRAINT user_session_pkey PRIMARY KEY (user_session_id);


--
-- Name: blog blog_domain_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.blog
    ADD CONSTRAINT blog_domain_id_fkey FOREIGN KEY (domain_id) REFERENCES public.domain(domain_id) ON DELETE CASCADE;


--
-- Name: blog_tag blog_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.blog_tag
    ADD CONSTRAINT blog_id_fkey FOREIGN KEY (blog_id) REFERENCES public.blog(blog_id) ON DELETE CASCADE;


--
-- Name: domain_user_role domain_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.domain_user_role
    ADD CONSTRAINT domain_id_fkey FOREIGN KEY (domain_id) REFERENCES public.domain(domain_id) ON DELETE CASCADE;


--
-- Name: page page_domain_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.page
    ADD CONSTRAINT page_domain_id_fkey FOREIGN KEY (domain_id) REFERENCES public.domain(domain_id) ON DELETE CASCADE;


--
-- Name: blog_tag tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.blog_tag
    ADD CONSTRAINT tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.tag(tag_id) ON DELETE CASCADE;


--
-- Name: domain_user_role user_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.domain_user_role
    ADD CONSTRAINT user_account_id_fkey FOREIGN KEY (user_account_id) REFERENCES public.user_account(user_account_id) ON DELETE CASCADE;


--
-- Name: user_session user_session_user_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: devasia
--

ALTER TABLE ONLY public.user_session
    ADD CONSTRAINT user_session_user_account_id_fkey FOREIGN KEY (user_account_id) REFERENCES public.user_account(user_account_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

