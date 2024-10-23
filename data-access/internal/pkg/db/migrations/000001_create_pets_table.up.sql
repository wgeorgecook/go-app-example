CREATE TYPE diet AS ENUM ('unspecified', 'no_restrictions', 'grain_free', 'raw');
CREATE TYPE friendly_with AS ENUM ('unspecified', 'little_dogs', 'big_dogs', 'cats', 'young_children', 'older_children');
CREATE TABLE IF NOT EXISTS pets(
   id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   birthdate DATE NOT NULL,
   description VARCHAR (300) NOT NULL,
   diet diet[],
   friendly_with friendly_with[] NOT NULL,
   picture_url VARCHAR (256) UNIQUE NOT NULL
);

INSERT INTO pets 
   (name, birthdate, description, diet, friendly_with, picture_url) 
VALUES
   ('Silas', 
   '2015-06-1', 
   'Medium-sized cutie with brown splotches that look like a heart and a peanut', 
   '{"no_restrictions"}',
    '{"little_dogs", "big_dogs", "cats", "young_children", "older_children"}', 
    'https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fstatic.businessinsider.com%2Fimage%2F5484d9d1eab8ea3017b17e29%2Fimage.jpg&f=1&nofb=1&ipt=3f4de2a11914ecac5af0ac35457a5e6c191c520cc2b0141c2f09f597f40cb27a&ipo=images'
    );