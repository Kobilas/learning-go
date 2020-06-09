DROP DATABASE IF EXISTS hotel;
CREATE DATABASE hotel;
USE hotel;

CREATE TABLE amenity (
	amenityId INT NOT NULL AUTO_INCREMENT,
    amenityName VARCHAR(20) NOT NULL,
	amenityType VARCHAR(20),
    CONSTRAINT pk_amenity
		PRIMARY KEY (amenityId)
);

CREATE TABLE bed (
	bedId INT NOT NULL AUTO_INCREMENT,
    size VARCHAR(10) NOT NULL,
    CONSTRAINT pk_bed
		PRIMARY KEY (bedId)
);

CREATE TABLE room (
	roomId INT NOT NULL AUTO_INCREMENT,
    bedId INT NOT NULL,
    amenityId INT,
    CONSTRAINT pk_room
		PRIMARY KEY (roomId),
	CONSTRAINT fk_room_bed
		FOREIGN KEY (bedId) REFERENCES bed(bedId),
	CONSTRAINT fk_room_amenity
		FOREIGN KEY (amenityId) REFERENCES amenity(amenityId)
);

CREATE TABLE guest (
	guestId INT NOT NULL AUTO_INCREMENT,
    firstName VARCHAR(20) NOT NULL,
    lastName VARCHAR(30) NOT NULL,
    address VARCHAR(50) NOT NULL,
    phone VARCHAR(10) NOT NULL,
    email VARCHAR(50) NOT NULL,
    CONSTRAINT pk_guest
		PRIMARY KEY (guestId)
);

CREATE TABLE reservation (
	reservationId INT NOT NULL AUTO_INCREMENT,
    startDate DATE NOT NULL,
    endDate DATE,
    guestId INT NOT NULL,
    roomId INT NOT NULL,
	CONSTRAINT pk_reservation
		PRIMARY KEY (reservationId),
	CONSTRAINT fk_reservation_guest
		FOREIGN KEY (guestId) REFERENCES guest(guestId),
	CONSTRAINT fk_reservation_room
		FOREIGN KEY (roomId) REFERENCES room(roomId)
);