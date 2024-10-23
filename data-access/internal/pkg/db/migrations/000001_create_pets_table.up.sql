CREATE TYPE diet AS ENUM ('unspecified', 'no_restrictions', 'grain_free', 'raw');
CREATE TYPE friendlyWith AS ENUM ('unspecified', 'littleDogs', 'bigDogs', 'cats', 'youngChildren', 'olderChildren');
CREATE TABLE IF NOT EXISTS pets(
   id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   birthdate DATE NOT NULL,
   description VARCHAR (300) NOT NULL,
   diet diet[],
   friendlyWith friendlyWith NOT NULL,
   pictureUrl VARCHAR (120) UNIQUE NOT NULL
);