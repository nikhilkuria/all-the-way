```
-- public.run definition

-- Drop table

-- DROP TABLE public.run;

CREATE TABLE public.run (
	run_id int4 NOT NULL,
	distance numeric NOT NULL,
	duration numeric NOT NULL,
	place varchar NOT NULL,
	runner_id int4 NOT NULL
);
```