/* Matthew Kobilas
06/08/20 */
DROP DATABASE IF EXISTS vinylrecordshop;
CREATE DATABASE vinylrecordshop;
SHOW DATABASES;
USE vinylrecordshop;

CREATE TABLE album (
	albumId INT NOT NULL AUTO_INCREMENT,
    albumTitle VARCHAR(100) NOT NULL,
    label VARCHAR(50),
    releaseDate DATE,
    price DECIMAL(5,2),
    CONSTRAINT pk_album
		PRIMARY KEY (albumId)
);
DESCRIBE album;

CREATE TABLE artist (
	artistId INT NOT NULL AUTO_INCREMENT,
    artistFirstName VARCHAR(25),
    artistLastName VARCHAR(50) NOT NULL,
    CONSTRAINT pk_artist
		PRIMARY KEY (artistId)
);
DESCRIBE artist;

CREATE TABLE band (
	bandId INT NOT NULL AUTO_INCREMENT,
    bandName VARCHAR(50) NOT NULL,
    CONSTRAINT pk_band
		PRIMARY KEY (bandId)
);
DESCRIBE band;

CREATE TABLE song (
	songId INT NOT NULL AUTO_INCREMENT,
    songTitle VARCHAR(100) NOT NULL,
    videoUrl VARCHAR(100),
    bandId INT,
    CONSTRAINT pk_song
		PRIMARY KEY (songId),
	CONSTRAINT fk_song
		FOREIGN KEY (bandId) REFERENCES band(bandId)
);
DESCRIBE song;

CREATE TABLE songAlbum (
	songId INT,
    albumId INT,
    CONSTRAINT pk_songAlbum
		PRIMARY KEY (songId, albumId),
	CONSTRAINT fk_songAlbum_song
		FOREIGN KEY (songId) REFERENCES song(songId),
	CONSTRAINT fk_songAlbum_album
		FOREIGN KEY (albumId) REFERENCES album(albumId)
);
DESCRIBE songAlbum;

CREATE TABLE bandArtist (
	bandId INT,
    artistId INT,
    CONSTRAINT pk_bandArtist
		PRIMARY KEY (bandId, artistId),
	CONSTRAINT fk_bandArtist_band
		FOREIGN KEY (bandId) REFERENCES band(bandId),
	CONSTRAINT fk_bandArtist_artist
		FOREIGN KEY (artistId) REFERENCES artist(artistId)
);
DESCRIBE bandArtist;