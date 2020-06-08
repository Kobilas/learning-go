DROP DATABASE IF EXISTS movies;
CREATE DATABASE movies;
SHOW DATABASES;
USE movies;

CREATE TABLE genre (
	genreId INT NOT NULL AUTO_INCREMENT,
    genreName VARCHAR(30) NOT NULL,
    CONSTRAINT pk_genre
		PRIMARY KEY (genreId)
);

CREATE TABLE director (
	directorId INT NOT NULL AUTO_INCREMENT,
    firstName VARCHAR(30),
    lastName VARCHAR(30),
    birthDate DATE,
    CONSTRAINT pk_director
		PRIMARY KEY (directorId)
);

CREATE TABLE rating (
	ratingId INT NOT NULL AUTO_INCREMENT,
    ratingName VARCHAR(5) NOT NULL,
    CONSTRAINT pk_rating
		PRIMARY KEY (ratingId)
);

CREATE TABLE movie (
	movieId INT NOT NULL AUTO_INCREMENT,
    genreId INT NOT NULL,
    directorId INT,
    ratingId INT,
    title VARCHAR(128) NOT NULL,
    releaseDate DATE,
    CONSTRAINT pk_movie
		PRIMARY KEY (movieId),
	CONSTRAINT fk_movie_genre
		FOREIGN KEY (genreId) REFERENCES genre(genreId),
	CONSTRAINT fk_movie_director
		FOREIGN KEY (directorId) REFERENCES director(directorId),
	CONSTRAINT fk_movie_rating
		FOREIGN KEY (ratingId) REFERENCES rating(ratingId)
);

CREATE TABLE actor (
	actorId INT NOT NULL AUTO_INCREMENT,
    firstName VARCHAR(30) NOT NULL,
    lastName VARCHAR(30) NOT NULL,
    birthDate DATE,
    CONSTRAINT pk_actor
		PRIMARY KEY (actorId)
);

CREATE TABLE castMember (
	castMemberId INT NOT NULL AUTO_INCREMENT,
    actorId INT NOT NULL,
    movieId INT NOT NULL,
    rle VARCHAR(50) NOT NULL,
    CONSTRAINT pk_castMember
		PRIMARY KEY (castMemberId),
	CONSTRAINT fk_castMember_actor
		FOREIGN KEY (actorId) REFERENCES actor(actorId),
	CONSTRAINT fk_castMember_movie
		FOREIGN KEY (movieId) REFERENCES movie(movieId)
);