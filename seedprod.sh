#!/bin/bash

# Change directory to the location of your SQL script files
cd foodloop-db-files/seeders

# Connect to the PostgreSQL database
PGPASSWORD=Mbk8LLupzXCgauwwu4tl2x1V5qzbA3wF psql -h dpg-co1s6tn79t8c73cfhs8g-a.singapore-postgres.render.com -U foodloopfp foodloopdb <<EOF

\i /Users/m.quek/capstone/fl-server/foodloop-db-files/01_schema.sql
\i /Users/m.quek/capstone/fl-server/foodloop-db-files/seeders/02_insertPeople.sql
\i /Users/m.quek/capstone/fl-server/foodloop-db-files/seeders/03_insertFoodlist.sql
\i /Users/m.quek/capstone/fl-server/foodloop-db-files/seeders/04_insertPeopleToFoodlist.sql
\i /Users/m.quek/capstone/fl-server/foodloop-db-files/seeders/05_insertFood.sql
\i /Users/m.quek/capstone/fl-server/foodloop-db-files/seeders/06_insertTag.sql
\i /Users/m.quek/capstone/fl-server/foodloop-db-files/seeders/07_insertFoodToTag.sql
\i /Users/m.quek/capstone/fl-server/foodloop-db-files/seeders/08_insertRestaurant.sql
\i /Users/m.quek/capstone/fl-server/foodloop-db-files/seeders/09_insertRestaurantToFood.sql
\i /Users/m.quek/capstone/fl-server/foodloop-db-files/seeders/10_insertFoodlistToFood.sql


