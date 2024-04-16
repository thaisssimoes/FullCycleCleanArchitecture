Create table if not exists products(
     id int primary key,
     name varchar(50),
     quantity int,
     price float
);




INSERT INTO products (name, quantity, price) values ("produto1", 1, 10.0);
INSERT INTO products (name, quantity, price) values ("produto2", 2, 20.0);