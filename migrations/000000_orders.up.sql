Create table if not exists orders(
     id int GENERATED ALWAYS AS IDENTITY,
     products varchar(200) NOT NULL,
     total float NOT NULL
);

INSERT INTO orders (products,  total) values ('produto1',  10.0);
INSERT INTO orders (products,  total) values ('produto2',  20.0);