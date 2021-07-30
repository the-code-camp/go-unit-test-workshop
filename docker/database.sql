CREATE TABLE public.customers (
  customer_id SERIAL,
  "name" varchar(100) NOT NULL,
  date_of_birth date NOT NULL,
  city varchar(100) NOT NULL,
  zipcode varchar(10) NOT NULL,
  status int4 NOT NULL DEFAULT 1,
  CONSTRAINT customers_pkey PRIMARY KEY (customer_id)
);


CREATE TABLE public.accounts (
  account_id SERIAL,
  customer_id int4 NOT NULL,
  opening_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  account_type varchar(10) NOT NULL,
  amount numeric(10,2) NOT NULL,
  status int4 NOT NULL DEFAULT 1,
  CONSTRAINT accounts_pkey PRIMARY KEY (account_id),
  CONSTRAINT accounts_fk FOREIGN KEY (customer_id) REFERENCES public.customers(customer_id)
);


CREATE TABLE public.transactions (
  transaction_id SERIAL,
  account_id int4 NOT NULL,
  amount numeric(10,2) NOT NULL,
  transaction_type varchar(10) NOT NULL,
  transaction_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id),
  CONSTRAINT transactions_fk FOREIGN KEY (account_id) REFERENCES public.accounts(account_id)
);


CREATE TABLE public.users (
  username varchar(20) NOT NULL,
  "password" varchar(20) NOT NULL,
  "role" varchar(20) NOT NULL,
  customer_id int4 NULL,
  created_on timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT users_pkey PRIMARY KEY (username)
);


INSERT INTO customers VALUES
  (2000,'Steve Postgres','1978-12-15','Delhi','110075',1),
  (2001,'Arian Postgres','1988-05-21','Newburgh, NY','12550',1),
  (2002,'Hadley Postgres','1988-04-30','Englewood, NJ','07631',1),
  (2003,'Ben Postgres','1988-01-04','Manchester, NH','03102',0),
  (2004,'Nina Postgres','1988-05-14','Clarkston, MI','48348',1),
  (2005,'Osman Postgres','1988-11-08','Hyattsville, MD','20782',0);

INSERT INTO accounts VALUES
  (95470,2000,'2020-08-22 10:20:06', 'saving', 6823.23, 1),
  (95471,2002,'2020-08-09 10:27:22', 'checking', 3342.96, 1),
  (95472,2001,'2020-08-09 10:35:22', 'saving', 7000, 1),
  (95473,2001,'2020-08-09 10:38:22', 'saving', 5861.86, 1);

INSERT INTO users VALUES
  ('admin','abc123','admin', NULL, '2020-08-09 10:27:22'),
  ('2001','abc123','user', 2001, '2020-08-09 10:27:22'),
  ('2000','abc123','user', 2000, '2020-08-09 10:27:22');
