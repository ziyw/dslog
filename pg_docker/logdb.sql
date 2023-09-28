--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9 (Homebrew)
-- Dumped by pg_dump version 14.9 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: dslog; Type: TABLE; Schema: public; Owner: ziyan
--

CREATE TABLE public.dslog (
    id integer NOT NULL,
    type character varying NOT NULL,
    text character varying NOT NULL
);


ALTER TABLE public.dslog OWNER TO ziyan;

--
-- Name: dslog_id_seq; Type: SEQUENCE; Schema: public; Owner: ziyan
--

ALTER TABLE public.dslog ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.dslog_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: dslog; Type: TABLE DATA; Schema: public; Owner: ziyan
--

COPY public.dslog (id, type, text) FROM stdin;
1	INFO	Hello world from client
2	ERROR	server side exception
3	WARN	This might now work
\.


--
-- Name: dslog_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ziyan
--

SELECT pg_catalog.setval('public.dslog_id_seq', 3, true);


--
-- PostgreSQL database dump complete
--

