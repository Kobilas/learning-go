DROP DATABASE IF EXISTS hotel;
CREATE DATABASE hotel;
SHOW DATABASES;
USE hotel;


CREATE TABLE room (
    roomId INT NOT NULL AUTO_INCREMENT,
    CONSTRAINT pk_room
        PRIMARY KEY (roomId)
);

CREATE TABLE amenities (
    amenityId INT NOT NULL AUTO_INCREMENT,
    amenityName VARCHAR(25) NOT NULL,
    roomID INT NOT NULL,
    CONSTRAINT pk_amenity
        PRIMARY KEY (amenityId),
    CONSTRAINT fk_room_id
        FOREIGN KEY (roomID) REFERENCES room(roomID)
);

CREATE TABLE bed (
    bedID INT NOT NULL AUTO_INCREMENT,
    size VARCHAR(10) NOT NULL,
    roomID INT,
    CONSTRAINT pk_bed
        PRIMARY KEY (bedID),
    CONSTRAINT fk_bedroom_id
        FOREIGN KEY (roomID) REFERENCES room(roomID)
);

CREATE TABLE reservation (
    reservationID INT NOT NULL AUTO_INCREMENT,
    startDate DATE NOT NULL,
    endDate DATE,
    ownerID INT NOT NULL,
    roomId INT NOT NULL,
    numGuests INT,
    CONSTRAINT pk_reservation
        PRIMARY KEY (reservationID),
    CONSTRAINT fk_reservation_room
        FOREIGN KEY (roomId) REFERENCES room(roomId)
);

CREATE TABLE guest (
    guestID INT NOT NULL AUTO_INCREMENT,
    firstName VARCHAR(25) NOT NULL,
    lastName VARCHAR(25) NOT NULL,
    address VARCHAR(50) NOT NULL,
    phone VARCHAR(10) NOT NULL,
    email VARCHAR(50) NOT NULL,
    reservationID INT NOT NULL,
    CONSTRAINT pk_guest
        PRIMARY KEY (guestID),
    CONSTRAINT fk_reservation_id
        FOREIGN KEY (reservationID) REFERENCES reservation(reservationID)
);

ALTER TABLE reservation 
	ADD CONSTRAINT fk_reservation_owner
		FOREIGN KEY (ownerId) REFERENCES guest(guestId);