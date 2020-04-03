/* IF EXISTS (
    SELECT * FROM pg_catalog.pg_proc JOIN pg_namespace ON pg_catalog.pg_proc.pronamespace = pg_namespace.oid WHERE proname = 'db_update') THEN
    DROP FUNCTION public.db_update();
 */
CREATE OR REPLACE FUNCTION public.db_update() RETURNS void LANGUAGE 'plpgsql'
AS $$
BEGIN
    ------------------------------------------------------------
	-- SEQUENCES
	------------------------------------------------------------
    /* IF NOT EXISTS (SELECT 1 FROM pg_class where relname = 'tprojects_id_seq' ) THEN
		CREATE SEQUENCE public.tprojects_id_seq;
		ALTER SEQUENCE public.tprojects_id_seq OWNER TO postgres;
	END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_class where relname = 'tcatalogs_id_seq' ) THEN
		CREATE SEQUENCE public.tcatalogs_id_seq;
		ALTER SEQUENCE public.tcatalogs_id_seq OWNER TO postgres;
	END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_class where relname = 'tfiles_id_seq' ) THEN
		CREATE SEQUENCE public.tfiles_id_seq;
		ALTER SEQUENCE public.tfiles_id_seq OWNER TO postgres;
	END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_class where relname = 'tprocedures_id_seq' ) THEN
		CREATE SEQUENCE public.tprocedures_id_seq;
		ALTER SEQUENCE public.tprocedures_id_seq OWNER TO postgres;
	END IF;    

    IF NOT EXISTS (SELECT 1 FROM pg_class where relname = 'trights_id_seq' ) THEN
		CREATE SEQUENCE public.trights_id_seq;
		ALTER SEQUENCE public.trights_id_seq OWNER TO postgres;
	END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_class where relname = 'troles_id_seq' ) THEN
		CREATE SEQUENCE public.troles_id_seq;
		ALTER SEQUENCE public.troles_id_seq	OWNER TO postgres;
	END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_class where relname = 'tusers_id_seq' ) THEN
		CREATE SEQUENCE public.tusers_id_seq;
		ALTER SEQUENCE public.tusers_id_seq	OWNER TO postgres;
	END IF; */
	---------------------------------------------------------------------
	-- TABLES
	---------------------------------------------------------------------
    -- AERO TRIP
    /* IF NOT EXISTS (SELECT 1 FROM pg_class where relname = 'passenger_seq' ) THEN
		CREATE SEQUENCE public.tusers_id_seq;
		ALTER SEQUENCE public.tusers_id_seq	OWNER TO postgres;
	END IF; */

	IF NOT EXISTS(SELECT 1 FROM information_schema.tables
		WHERE table_schema = 'public' AND table_name = 'Passenger') THEN
		CREATE TABLE public.Passenger
		(
			id integer NOT NULL,
			name character varying(50) COLLATE pg_catalog."default" NOT NULL
		)
		WITH (OIDS = FALSE)	TABLESPACE pg_default;
		ALTER TABLE public.Passenger OWNER to postgres;
	END IF;

	IF NOT EXISTS(SELECT 1 FROM information_schema.tables
		WHERE table_schema = 'public' AND table_name = 'Pass_in_trip') THEN
		CREATE TABLE public.Pass_in_trip
		(
			id integer NOT NULL,
			trip integer NOT NULL,
            passenger integer NOT NULL,
            place character varying(50) COLLATE pg_catalog."default" NOT NULL
		)
		WITH (OIDS = FALSE)	TABLESPACE pg_default;
		ALTER TABLE public.Pass_in_trip OWNER to postgres;
	END IF;

	IF NOT EXISTS(SELECT 1 FROM information_schema.tables
		WHERE table_schema = 'public' AND table_name = 'Company') THEN
		CREATE TABLE public.Company
		(
			id integer NOT NULL,
			name character varying(50) COLLATE pg_catalog."default" NOT NULL
		)
		WITH (OIDS = FALSE)	TABLESPACE pg_default;
		ALTER TABLE public.Company OWNER to postgres;
	END IF;

	IF NOT EXISTS(SELECT 1 FROM information_schema.tables
		WHERE table_schema = 'public' AND table_name = 'Trip') THEN
		CREATE TABLE public.Trip
		(
			id integer NOT NULL,
            company integer NOT NULL,
			plane character varying(50) COLLATE pg_catalog."default" NOT NULL,
            town_from character varying(50) COLLATE pg_catalog."default" NOT NULL,
            town_to character varying(50) COLLATE pg_catalog."default" NOT NULL,
            time_out timestamp with time zone NOT NULL,
            time_in timestamp with time zone NOT NULL
		)
		WITH (OIDS = FALSE)	TABLESPACE pg_default;
		ALTER TABLE public.Trip OWNER to postgres;
	END IF;



    -----------------------------------------------------------------------------------
    IF NOT EXISTS(SELECT * FROM public.Company WHERE ID = 1) THEN
        INSERT INTO public.Company(id, name) VALUES (1, 'Don_avia');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Company WHERE ID = 2) THEN
        INSERT INTO public.Company(id, name) VALUES (2, 'Aeroflot');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Company WHERE ID = 3) THEN
        INSERT INTO public.Company(id, name) VALUES (3, 'Dale_avia');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Company WHERE ID = 4) THEN
        INSERT INTO public.Company(id, name) VALUES (4, 'air_France');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Company WHERE ID = 5) THEN
        INSERT INTO public.Company(id, name) VALUES (5, 'British_AW');
    END IF;

    -----------------------------------------------------------------------------------

    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 1) THEN
        INSERT INTO public.Passenger(id, name) VALUES (1, 'Bruce Willis');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 2) THEN
        INSERT INTO public.Passenger(id, name) VALUES (2, 'George Clooney');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 3) THEN
        INSERT INTO public.Passenger(id, name) VALUES (3, 'Kevin Costner');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 4) THEN
        INSERT INTO public.Passenger(id, name) VALUES (4, 'Donald Sutherland');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 5) THEN
        INSERT INTO public.Passenger(id, name) VALUES (5, 'Jennifer Lopez');
    END IF;

    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 6) THEN
        INSERT INTO public.Passenger(id, name) VALUES (6, 'Ray Liotta');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 7) THEN
        INSERT INTO public.Passenger(id, name) VALUES (7, 'Samuel L. Jackson');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 8) THEN
        INSERT INTO public.Passenger(id, name) VALUES (8, 'Nikole Kidman');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 9) THEN
        INSERT INTO public.Passenger(id, name) VALUES (9, 'Alan Rickman');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 10) THEN
        INSERT INTO public.Passenger(id, name) VALUES (10, 'Kurt Russell');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 11) THEN
        INSERT INTO public.Passenger(id, name) VALUES (11, 'Harrison Ford');
    END IF;

    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 12) THEN
        INSERT INTO public.Passenger(id, name) VALUES (12, 'Russell Crowe');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 13) THEN
        INSERT INTO public.Passenger(id, name) VALUES (13, 'Steve Martin');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 14) THEN
        INSERT INTO public.Passenger(id, name) VALUES (14, 'Michael Caine');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 15) THEN
        INSERT INTO public.Passenger(id, name) VALUES (15, 'Angelina Jolie');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 16) THEN
        INSERT INTO public.Passenger(id, name) VALUES (16, 'Mel Gibson');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 17) THEN
        INSERT INTO public.Passenger(id, name) VALUES (17, 'Michael Douglas');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 18) THEN
        INSERT INTO public.Passenger(id, name) VALUES (18, 'John Travolta');
    END IF;                        

    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 19) THEN
        INSERT INTO public.Passenger(id, name) VALUES (19, 'Sylvester Stallone');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 20) THEN
        INSERT INTO public.Passenger(id, name) VALUES (20, 'Tommy Lee Jones');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 21) THEN
        INSERT INTO public.Passenger(id, name) VALUES (21, 'Catherine Zeta-Jones');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 22) THEN
        INSERT INTO public.Passenger(id, name) VALUES (22, 'Antonio Banderas');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 23) THEN
        INSERT INTO public.Passenger(id, name) VALUES (23, 'Kim Basinger');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 24) THEN
        INSERT INTO public.Passenger(id, name) VALUES (24, 'Sam Neill');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 25) THEN
        INSERT INTO public.Passenger(id, name) VALUES (25, 'Gary Oldman');
    END IF; 

    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 26) THEN
        INSERT INTO public.Passenger(id, name) VALUES (26, 'ClINT Eastwood');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 27) THEN
        INSERT INTO public.Passenger(id, name) VALUES (27, 'Brad Pitt');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 28) THEN
        INSERT INTO public.Passenger(id, name) VALUES (28, 'Johnny Depp');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 29) THEN
        INSERT INTO public.Passenger(id, name) VALUES (29, 'Pierce Brosnan');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 30) THEN
        INSERT INTO public.Passenger(id, name) VALUES (30, 'Sean Connery');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 31) THEN
        INSERT INTO public.Passenger(id, name) VALUES (31, 'Bruce Willis');
    END IF;
    IF NOT EXISTS(SELECT * FROM public.Passenger WHERE ID = 37) THEN
        INSERT INTO public.Passenger(id, name) VALUES (37, 'Mullah Omar');
    END IF; 
    
    -----------------------------------------------------------------------------------

    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 1) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (1, 1100, 1, '1a');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 2) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (2, 1123, 3, '2a');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 3) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (3, 1123, 1, '4c');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 4) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (4, 1123, 6, '4b');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 5) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (5, 1124, 2, '2d');
    END IF; 

    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 6) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (6, 1145, 3, '2c');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 7) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (7, 1181, 1, '1a');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 8) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (8, 1181, 6, '1b');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 9) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (9, 1181, 8, '3c');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 10) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (10, 1181, 5, '1b');
    END IF; 

    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 11) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (11, 1182, 5, '4b');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 12) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (12, 1187, 8, '3a');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 13) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (13, 1188, 8, '3a');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 14) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (14, 1182, 9, '6d');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 15) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (15, 1145, 5, '1d');
    END IF; 

    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 16) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (16, 1187, 10, '3d');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 17) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (17, 8882, 37, '1a');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 18) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (18, 7771, 37, '1c');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 19) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (19, 7772, 37, '1a');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 20) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (20, 8881, 37, '1d');
    END IF; 

    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 21) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (21, 7778, 10, '2a');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 22) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (22, 7772, 10, '3a');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 23) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (23, 7771, 11, '4a');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 24) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (24, 7771, 11, '1b');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 25) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (25, 7771, 11, '5a');
    END IF; 

    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 26) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (26, 7772, 12, '1d');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 27) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (27, 7773, 13, '2d');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 28) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (28, 7772, 13, '1b');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 29) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (29, 8882, 14, '3d');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 30) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (30, 7771, 14, '4d');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 31) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (31, 7771, 14, '5d');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Pass_in_trip WHERE ID = 32) THEN
        INSERT INTO public.Pass_in_trip(id, trip, passenger, place) VALUES (32, 7772, 14, '1c');
    END IF; 

    -----------------------------------------------------------------------------------

    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1100) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1100, 4, 'Boeing', 'Rostov', 'Paris', '1900-01-01T14:30:00.000Z', '1900-01-01T17:50:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1101) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1101, 4, 'Boeing', 'Paris', 'Rostov', '1900-01-01T08:12:00.000Z', '1900-01-01T11:45:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1123) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1123, 3, 'TU-154', 'Rostov', 'Vladivostok', '1900-01-01T16:20:00.000Z', '1900-01-01T03:40:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1124) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1124, 3, 'TU-154', 'Vladivostok', 'Rostov', '1900-01-01T09:00:00.000Z', '1900-01-01T19:50:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1145) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1145, 2, 'IL-86', 'Moscow', 'Rostov', '1900-01-01T09:35:00.000Z', '1900-01-01T11:23:00.000Z');
    END IF; 

    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1146) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1146, 2, 'IL-86', 'Rostov', 'Moscow', '1900-01-01T17:55:00.000Z', '1900-01-01T20:01:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1181) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1181, 1, 'TU-134', 'Rostov', 'Moscow', '1900-01-01T06:12:00.000Z', '1900-01-01T08:01:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1182) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1182, 1, 'TU-134', 'Moscow', 'Rostov', '1900-01-01T12:35:00.000Z', '1900-01-01T14:30:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1187) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1187, 1, 'TU-134', 'Rostov', 'Moscow', '1900-01-01T15:42:00.000Z', '1900-01-01T17:39:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1188) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1188, 1, 'TU-134', 'Moscow', 'Rostov', '1900-01-01T22:50:00.000Z', '1900-01-01T00:48:00.000Z');
    END IF; 

    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1195) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1195, 1, 'TU-154', 'Rostov', 'Moscow', '1900-01-01T23:30:00.000Z', '1900-01-01T01:11:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 1196) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (1196, 1, 'TU-154', 'Moscow', 'Rostov', '1900-01-01T04:00:00.000Z', '1900-01-01T05:45:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 7771) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (7771, 5, 'Boeing', 'London', 'Singapore', '1900-01-01T01:00:00.000Z', '1900-01-01T11:00:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 7772) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (7772, 5, 'Boeing', 'Singapore', 'London', '1900-01-01T12:00:00.000Z', '1900-01-01T02:00:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 7773) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (7773, 5, 'Boeing', 'London', 'Singapore', '1900-01-01T03:00:00.000Z', '1900-01-01T13:00:00.000Z');
    END IF; 


    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 7774) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (7774, 5, 'Boeing', 'Singapore', 'London', '1900-01-01T14:00:00.000Z', '1900-01-01T06:00:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 7775) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (7775, 5, 'Boeing', 'London', 'Singapore', '1900-01-01T09:00:00.000Z', '1900-01-01T20:00:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 7776) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (7776, 5, 'Boeing', 'Singapore', 'London', '1900-01-01T18:00:00.000Z', '1900-01-01T08:00:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 7777) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (7777, 5, 'Boeing', 'London', 'Singapore', '1900-01-01T18:00:00.000Z', '1900-01-01T06:00:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 7778) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (7778, 5, 'Boeing', 'Singapore', 'London', '1900-01-01T22:00:00.000Z', '1900-01-01T12:00:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 8881) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (8881, 5, 'Boeing', 'London', 'Paris', '1900-01-01T03:00:00.000Z', '1900-01-01T04:00:00.000Z');
    END IF; 
    IF NOT EXISTS(SELECT * FROM public.Trip WHERE ID = 8882) THEN
        INSERT INTO public.Trip(id, company, plane, town_from, town_to, time_out, time_in) VALUES (8882, 5, 'Boeing', 'Paris', 'London', '1900-01-01T22:00:00.000Z', '1900-01-01T23:00:00.000Z');
    END IF; 

    -----------------------------------------------------------------------------------------------
    -----------------------------------------------------------------------------------------------
    -----------------------------------------------------------------------------------------------

	IF NOT EXISTS(SELECT 1 FROM information_schema.tables
		WHERE table_schema = 'public' AND table_name = 'ttasks') THEN
		CREATE TABLE public.ttasks
		(
			id integer NOT NULL,
			text character varying(500) COLLATE pg_catalog."default" NOT NULL,
            stars integer NOT NULL,
            col1 character varying(50),
            col2 character varying(50),
            hard boolean

		)
		WITH (OIDS = FALSE)	TABLESPACE pg_default;
		ALTER TABLE public.ttasks OWNER to postgres;
	END IF;
/*
    IF NOT EXISTS(SELECT * FROM public.ttasks WHERE ID = 1) THEN
        INSERT INTO public.ttasks(id, text, stars, col1, col2, hard) VALUES (1, 'Вывести имена всех когда-либо обслуживаемых пассажиров авиакомпаний', 1, 'name', '', FALSE);
    END IF;
*/
	IF NOT EXISTS(SELECT 1 FROM information_schema.tables
		WHERE table_schema = 'public' AND table_name = 'tanswers') THEN
		CREATE TABLE public.tanswers
		(
			id integer NOT NULL,
            task_id integer NOT NULL,
            order_ integer NOT NULL,
            col1 character varying(50),
            col2 character varying(50)
		)
		WITH (OIDS = FALSE)	TABLESPACE pg_default;
		ALTER TABLE public.tanswers OWNER to postgres;
	END IF;

    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 1) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (1, 1, 0, 'Bruce Willis', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 2) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (2, 1, 0, 'George Clooney', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 3) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (3, 1, 0, 'Kevin Costner', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 4) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (4, 1, 0, 'Donald Sutherland', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 5) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (5, 1, 0, 'Jennifer Lopez', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 6) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (6, 1, 0, 'Ray Liotta', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 7) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (7, 1, 0, 'Samuel L. Jackson', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 8) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (8, 1, 0, 'Nikole Kidman', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 9) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (9, 1, 0, 'Alan Rickman', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 10) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (10, 1, 0, 'Kurt Russell', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 11) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (11, 1, 0, 'Harrison Ford', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 12) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (12, 1, 0, 'Russell Crowe', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 13) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (13, 1, 0, 'Steve Martin', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 14) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (14, 1, 0, 'Michael Caine', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 15) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (15, 1, 0, 'Angelina Jolie', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 16) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (16, 1, 0, 'Mel Gibson', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 17) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (17, 1, 0, 'Michael Douglas', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 18) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (18, 1, 0, 'John Travolta', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 19) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (19, 1, 0, 'Sylvester Stallone', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 20) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (20, 1, 0, 'Tommy Lee Jones', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 21) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (21, 1, 0, 'Catherine Zeta-Jones', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 22) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (22, 1, 0, 'Antonio Banderas', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 23) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (23, 1, 0, 'Kim Basinger', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 24) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (24, 1, 0, 'Sam Neill', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 25) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (25, 1, 0, 'Gary Oldman', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 26) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (26, 1, 0, 'ClINT Eastwood', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 27) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (27, 1, 0, 'Brad Pitt', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 28) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (28, 1, 0, 'Johnny Depp', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 29) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (29, 1, 0, 'Pierce Brosnan', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 30) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (30, 1, 0, 'Sean Connery', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 31) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (31, 1, 0, 'Bruce Willis', ''); END IF;
    IF NOT EXISTS(SELECT * FROM public.tanswers WHERE ID = 32) THEN INSERT INTO public.tanswers(id, task_id, order_, col1, col2) VALUES (32, 1, 0, 'Mullah Omar', ''); END IF;

END
$$;