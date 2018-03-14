--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: auth_token; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE auth_token (
    id integer NOT NULL,
    token character varying,
    user_id integer
);


ALTER TABLE auth_token OWNER TO postgres;

--
-- Name: auth_token_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE auth_token_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE auth_token_id_seq OWNER TO postgres;

--
-- Name: auth_token_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE auth_token_id_seq OWNED BY auth_token.id;


--
-- Name: events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE events (
    id integer NOT NULL,
    user_id integer,
    title character varying(255),
    text character varying,
    list_notifications json,
    notifications character varying,
    create_at timestamp without time zone,
    date timestamp without time zone,
    complete boolean DEFAULT false,
    important boolean DEFAULT false
);


ALTER TABLE events OWNER TO postgres;

--
-- Name: events_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE events_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE events_id_seq OWNER TO postgres;

--
-- Name: events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE events_id_seq OWNED BY events.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE users (
    id integer NOT NULL,
    firstname character varying(126),
    lastname character varying(126),
    email character varying(126),
    phone character varying(11),
    date_register timestamp without time zone,
    role integer DEFAULT 0,
    subscribe boolean DEFAULT false,
    date_subscribe timestamp without time zone
);


ALTER TABLE users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE users_id_seq OWNED BY users.id;


--
-- Name: auth_token id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY auth_token ALTER COLUMN id SET DEFAULT nextval('auth_token_id_seq'::regclass);


--
-- Name: events id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY events ALTER COLUMN id SET DEFAULT nextval('events_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);


--
-- Data for Name: auth_token; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY auth_token (id, token, user_id) FROM stdin;
1	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJJc3N1ZWRBdCI6MTUxNzMwODk5MCwiaWQiOjV9.aKSKo7POmxhowC-fPt9cUdjIOWPuTskw9Eibm3WOmtM	5
2	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJJc3N1ZWRBdCI6MTUxNzMwOTQ5MSwiaWQiOjV9.4q_dwqecr36fPJEcKfo8hTjnJ0usiYIu8brLq8-Bxf8	5
3	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJJc3N1ZWRBdCI6MTUxNzMwOTQ5MywiaWQiOjV9.CB-i_2aZj6QF4_JUV7lswA7XsNa8IzPnH5o8RmOpYRA	5
4	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJJc3N1ZWRBdCI6MTUxNzMxMDM2NiwiaWQiOjZ9.Y9949hZk0iH4UuD8fN64vf1iR11P1RN0EjxStFKreIM	6
5	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJJc3N1ZWRBdCI6MTUxNzMxMDU4NSwiaWQiOjd9.RUtLYFEg-hpJ1TGditbSTfbY4AQ685wIG_lyJlv8CAE	7
6	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJJc3N1ZWRBdCI6MTUxNzMxMDYyMSwiaWQiOjZ9.0lsJQR3JiEp05bdgS2gt-xGRbTprRzIj28w3sjwFejA	6
7	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJJc3N1ZWRBdCI6MTUxNzMxMDYzMCwiaWQiOjZ9.jFRPIGnDE1K20fnxSnXmfcDO8NkhpDL6lxTyXyKOJoM	6
8	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJJc3N1ZWRBdCI6MTUxNzMxMDYzNywiaWQiOjZ9.HLr1Rg3S5UtwUo7LvgzLKp70MZYxUehnIFQiYTwUiXg	6
9	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJJc3N1ZWRBdCI6MTUxNzMxMDgyMywiaWQiOjd9.h6-jnVwn9DeySk1nwPZ_NIsHgTdBXaxBXvFJqFJBKFA	7
10	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE1MDAwLCJJc3N1ZWRBdCI6MTUxNzMxMDg0NCwiaWQiOjd9.w5ktB0GqIl5XCUOfNd13D7i6I0pJx4hwQ3WPHqzvlMM	7
\.


--
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY events (id, user_id, title, text, list_notifications, notifications, create_at, date, complete, important) FROM stdin;
9	6	Test Event	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	2018-01-31 00:00:00	f	t
10	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	2018-02-01 00:00:00	f	t
11	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	0001-01-01 00:00:00	f	t
12	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	0001-01-01 00:00:00	f	t
13	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	0001-01-01 00:00:00	f	t
14	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	0001-01-01 00:00:00	f	t
15	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	2018-02-01 10:00:00	f	t
16	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	2018-02-01 13:00:00	f	t
17	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	2018-02-02 12:30:00	f	t
18	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	2018-02-02 15:30:00	f	t
19	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	2018-02-02 18:30:00	f	t
20	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	2018-02-05 18:30:00	f	t
21	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	2018-02-05 13:30:00	f	t
22	6	Test Event 1	Test Test Test Test	{\n  "vk": true,\n  "facebook": true,\n  "telegram": false,\n  "viber": true,\n  "mail": false\n}	111	2018-01-31 00:00:00	2018-02-05 09:30:00	f	t
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY users (id, firstname, lastname, email, phone, date_register, role, subscribe, date_subscribe) FROM stdin;
1	Даниил	Вилявин	89505105005@mail.ru	89161647949	2018-01-23 12:01:34.531855	1	t	2018-01-23 12:01:34.531855
2	Веселый	Молочник	fun@milk.ru	1111	2018-01-23 16:37:21.303649	1	t	2018-01-23 16:37:21.303649
4	Иван	Иванов	ivan@ivan.com	8888	2018-01-29 20:38:07.783093	0	f	\N
3	Сергей	Савтыра	admin@calday.org	89162870886	2018-01-29 20:31:18.516606	1	t	2018-01-29 20:31:18.516606
5	Иван	Иванов	ivan@ivan.com	88881	2018-01-30 12:38:20.740985	0	f	\N
7	Иван	Иванов	ivan@ivan.com	88883	2018-01-30 13:09:05.427843	0	f	\N
6	Иван	Иванов	ivan@ivan.com	88882	2018-01-30 12:54:58.735033	0	f	\N
\.


--
-- Name: auth_token_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('auth_token_id_seq', 10, true);


--
-- Name: events_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('events_id_seq', 22, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('users_id_seq', 7, true);


--
-- Name: auth_token auth_token_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY auth_token
    ADD CONSTRAINT auth_token_pkey PRIMARY KEY (id);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: users_id_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX users_id_uindex ON users USING btree (id);


--
-- PostgreSQL database dump complete
--

