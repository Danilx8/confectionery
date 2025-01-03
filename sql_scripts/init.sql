CREATE DATABASE IF NOT EXISTS cakes;
USE cakes;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    login varchar(255) COLLATE utf8_bin NOT NULL,
    password varchar(25) NOT NULL,
    role varchar(255) NOT NULL,
    full_name varchar(255),
    photo_url varchar(510),
    PRIMARY KEY (login, password),
    UNIQUE (login)
);

DROP TABLE IF EXISTS `items`;
CREATE TABLE `items` (
    name varchar(255) NOT NULL,
    size varchar(100) NOT NULL,
    PRIMARY KEY (name)
);

DROP TABLE IF EXISTS `tooling_types`;
CREATE TABLE `tooling_types` (
    name varchar(255) NOT NULL,
    PRIMARY KEY (name)
);

DROP TABLE IF EXISTS `suppliers`;
CREATE TABLE `suppliers` (
    name varchar(255) NOT NULL,
    address varchar(255),
    delivery_time int NOT NULL,
    PRIMARY KEY (name)
);

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
    id int(11) NOT NULL AUTO_INCREMENT,
    date DATE NOT NULL,
    name varchar(255) NOT NULL,
    item varchar(255) NOT NULL,
    orderer varchar(255) NOT NULL COLLATE utf8_bin,
    assigned_manager varchar(255) COLLATE utf8_bin,
    price decimal(10, 2),
    expected_fulfilment_date date,
    examples varchar(1000),
    PRIMARY KEY (id, date),
    FOREIGN KEY (item) REFERENCES `items`(name),
    FOREIGN KEY (orderer) REFERENCES `users`(login),
    FOREIGN KEY (assigned_manager) REFERENCES `users`(login)
);

DROP TABLE IF EXISTS `toolings`;
CREATE TABLE `toolings` (
    marking varchar(255) NOT NULL,
    type varchar(255) NOT NULL,
    properties json,
    PRIMARY KEY (marking),
    FOREIGN KEY (type) REFERENCES `tooling_types`(name)
);

DROP TABLE IF EXISTS `operation_specifications`;
CREATE TABLE `operation_specifications` (
    item varchar(255) NOT NULL,
    operation varchar(255) NOT NULL,
    sequence_number int(11) NOT NULL,
    tooling_type varchar(255),
    required_time int(3) NOT NULL,
    PRIMARY KEY (item, operation, sequence_number),
    FOREIGN KEY (item) REFERENCES `items`(name)
);

DROP TABLE IF EXISTS `premade_specifications`;
CREATE TABLE `premade_specifications` (
    item varchar(255) NOT NULL,
    premade varchar(255) NOT NULL,
    amount int(5),
    PRIMARY KEY (item, premade),
    FOREIGN KEY (item) REFERENCES `items`(name),
    FOREIGN KEY (premade) REFERENCES `items`(name)
);

DROP TABLE IF EXISTS `cake_decorations`;
CREATE TABLE `cake_decorations` (
    article varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    unit varchar(50) NOT NULL,
    amount int(5) NOT NULL,
    supplier varchar(255),
    image varchar(1000),
    type varchar(255) NOT NULL,
    cost_price decimal(10, 2) NOT NULL,
    weight decimal(10, 2) NOT NULL,
    PRIMARY KEY (article),
    FOREIGN KEY (supplier) REFERENCES `suppliers` (name)
);

DROP TABLE IF EXISTS `cake_decoration_specifications`;
CREATE TABLE `cake_decoration_specifications` (
    item varchar(255) NOT NULL,
    cake_decoration varchar(255) NOT NULL,
    amount int(4) NOT NULL,
    PRIMARY KEY (item, cake_decoration),
    FOREIGN KEY (item) REFERENCES `items`(name),
    FOREIGN KEY (cake_decoration) REFERENCES `cake_decorations`(article)
);

DROP TABLE IF EXISTS `ingredients`;
CREATE TABLE `ingredients` (
    article varchar(255) NOT NULL,
    name varchar(100) NOT NULL,
    unit varchar(50) NOT NULL,
    amount int(4) NOT NULL,
    supplier varchar(255),
    image varchar(1000),
    type varchar(255) NOT NULL,
    cost_price decimal(10, 2) NOT NULL,
    gost varchar(255),
    packing varchar(100),
    specs varchar(1000),
    PRIMARY KEY (article),
    FOREIGN KEY (supplier) REFERENCES `suppliers`(name)
);

DROP TABLE IF EXISTS `ingredient_specifications`;
CREATE TABLE `ingredient_specifications` (
    item varchar(255) NOT NULL,
    ingredient varchar(255) NOT NULL,
    amount int(5),
    PRIMARY KEY (item, ingredient),
    FOREIGN KEY (item) REFERENCES `items`(name),
    FOREIGN KEY (ingredient) REFERENCES `ingredients`(article)
)