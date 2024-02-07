CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    phone varchar(20) NOT NULL,
    gender varchar(20) NOT NULL,
    age int NOT NULL,
    email varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    vdays int DEFAULT NULL
);

INSERT INTO employees(name, phone, gender, age, email, address, vdays) VALUES ('Alex Balder','89171356392','male',31,'abalder@gmail.com','New York, Main Street, 91',NULL);
INSERT INTO employees(name, phone, gender, age, email, address, vdays) VALUES ('Mikhail Voronin','89168352563','male',31,'vor@gmail.com','Thief Guild',30);
INSERT INTO employees(name, phone, gender, age, email, address, vdays) VALUES ('Boris Petrov','89172212233','male',23,'bpetrov@gmail.com','Moscow 42 Rybalko Street',NULL);
INSERT INTO employees(name, phone, gender, age, email, address, vdays) VALUES ('Mikhail Vorobiev','89174412233','male',35,'bpetrov@gmail.com','Moscow 14 Lermontov Street',NULL);
INSERT INTO employees(name, phone, gender, age, email, address, vdays) VALUES ('Vasily Danilov','89178323232','male',38,'vdanilov@gmail.com','St. Petersburg, Stalin Street, 11',NULL);
INSERT INTO employees(name, phone, gender, age, email, address, vdays) VALUES ('Mikhail Eremeev','89174512733','male',34,'mkeremeev@gmail.com','Moscow 15 First Street',NULL);
