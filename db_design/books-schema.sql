DROP DATABASE IF EXISTS books;
CREATE DATABASE books;
SHOW DATABASES;
USE books;

CREATE TABLE book (
	bookId INT NOT NULL,
    title VARCHAR(100) NOT NULL,
    publicationDate DATETIME,
    CONSTRAINT pk_book
		PRIMARY KEY (bookId)
);

CREATE TABLE author (
	authorId INT NOT NULL AUTO_INCREMENT,
    firstName VARCHAR(25) NOT NULL,
    middleName VARCHAR(25),
    lastName VARCHAR(50) NOT NULL,
    gender VARCHAR(1),
    dateOfBirth DATETIME NOT NULL,
    dateOfDeath DATETIME,
    CONSTRAINT pk_author
		PRIMARY KEY (authorId)
);

CREATE TABLE authorBook (
	authorId INT NOT NULL,
    bookId INT NOT NULL,
    CONSTRAINT pk_authorBook
		PRIMARY KEY (authorId, bookId),
	CONSTRAINT fk_authorBook_author
		FOREIGN KEY (authorId) REFERENCES author(authorId),
	CONSTRAINT fk_authorBook_book
		FOREIGN KEY (bookId) REFERENCES book(bookId)
);

CREATE TABLE frmat (
	frmatId INT NOT NULL AUTO_INCREMENT,
    frmatName VARCHAR(12) NOT NULL,
    CONSTRAINT pk_frmat
		PRIMARY KEY (frmatId)
);

CREATE TABLE bookFrmat (
	bookId INT NOT NULL,
    frmatId INT NOT NULL,
    price DOUBLE,
    quantityOnHand INT,
    CONSTRAINT pk_bookFrmat
		PRIMARY KEY (bookId, frmatId),
	CONSTRAINT fk_bookFrmat_book
		FOREIGN KEY (bookId) REFERENCES book(bookId),
	CONSTRAINT fk_bookFrmat_frmat
		FOREIGN KEY (frmatId) REFERENCES frmat(frmatId)
);

CREATE TABLE genre (
	genreId INT NOT NULL AUTO_INCREMENT,
    genreName VARCHAR(25) NOT NULL,
    CONSTRAINT pk_genre
		PRIMARY KEY (genreId)
);

CREATE TABLE bookGenre (
	bookId INT NOT NULL,
    genreId INT NOT NULL,
    CONSTRAINT pk_bookGenre
		PRIMARY KEY (bookId, genreId),
	CONSTRAINT fk_bookGenre_book
		FOREIGN KEY (bookId) REFERENCES book(bookId),
	CONSTRAINT fk_bookGenre_genre
		FOREIGN KEY (genreId) REFERENCES genre(genreId)
);