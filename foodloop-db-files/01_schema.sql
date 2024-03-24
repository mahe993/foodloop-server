\c postgres;
DROP DATABASE IF EXISTS foodloopdb;
CREATE DATABASE foodloopdb;
\c foodloopdb;

CREATE SCHEMA foodloop;

ALTER DATABASE foodloopdb SET search_path TO foodloop;

CREATE TABLE foodloop.people (
    peopleID SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE foodloop.foodlist (
    foodlistID SERIAL PRIMARY KEY,
    foodlistName VARCHAR(100),
    foodlistTime VARCHAR(20),
    foodlistDay VARCHAR(20),
    foodlistCurrIdx INTEGER
);

CREATE TABLE foodloop.restaurant (
    restaurantID SERIAL PRIMARY KEY,
    restaurantName VARCHAR(50),
    descriptions VARCHAR(250)
);

CREATE TABLE foodloop.food (
    foodID SERIAL PRIMARY KEY,
    foodName VARCHAR(50),
    descriptions VARCHAR(250)
);

CREATE TABLE foodloop.tag (
    tagID SERIAL PRIMARY KEY,
    tagName VARCHAR(20)
);

CREATE TABLE foodloop.peopleToFoodlist (
    peopleID INT,
    foodlistID INT,
    PRIMARY KEY (peopleID, foodlistID),
    CONSTRAINT fk_people
        FOREIGN KEY (peopleID)
            REFERENCES foodloop.people(peopleID),
    CONSTRAINT fk_foodlist
        FOREIGN KEY (foodlistID)
            REFERENCES foodloop.foodlist(foodlistID)
);

CREATE TABLE foodloop.foodlistToFood (
    foodlistID INT,
    foodID INT,
    foodIndex INT,
    PRIMARY KEY (foodlistID, foodIndex),
    CONSTRAINT fk_foodlist
        FOREIGN KEY (foodlistID)
            REFERENCES foodloop.foodlist(foodlistID),
    CONSTRAINT fk_food
        FOREIGN KEY (foodID)
            REFERENCES foodloop.food(foodID)
);

CREATE TABLE foodloop.restaurantToFood (
    restaurantID INT,
    foodID INT,
    PRIMARY KEY (restaurantID, foodID),
    CONSTRAINT fk_restaurant
        FOREIGN KEY (restaurantID)
            REFERENCES foodloop.restaurant(restaurantID),
    CONSTRAINT fk_food
        FOREIGN KEY (foodID)
            REFERENCES foodloop.food(foodID)
);

CREATE TABLE foodloop.restaurantToTag (
    restaurantID INT,
    tagID INT,
    PRIMARY KEY (restaurantID, tagID),
    CONSTRAINT fk_restaurant
        FOREIGN KEY (restaurantID)
            REFERENCES foodloop.restaurant(restaurantID),
    CONSTRAINT fk_tag
        FOREIGN KEY (tagID)
            REFERENCES foodloop.tag(tagID)
);

CREATE TABLE foodloop.foodToTag (
    foodID INT,
    tagID INT,
    PRIMARY KEY (foodID, tagID),
    CONSTRAINT fk_food
        FOREIGN KEY (foodID)
            REFERENCES foodloop.food(foodID),
    CONSTRAINT fk_tag
        FOREIGN KEY (tagID)
            REFERENCES foodloop.tag(tagID)
);