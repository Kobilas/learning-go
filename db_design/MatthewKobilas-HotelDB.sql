DROP DATABASE IF EXISTS hotel_software_guild;
CREATE DATABASE hotel_software_guild;
USE hotel_software_guild;

CREATE TABLE room (
	room_id INT NOT NULL AUTO_INCREMENT,
		CONSTRAINT pk_room
		PRIMARY KEY (room_id),
    beds ENUM("single", "double", "suite") NOT NULL,
    ada_access BOOL NOT NULL,
    standard_occupancy INT NOT NULL,
    max_occupancy INT NOT NULL,
    base_price DECIMAL(3, 2) NOT NULL,
    price_per_added_guest DECIMAL(2, 2) NOT NULL
);

CREATE TABLE amenity (
	amenity_id INT NOT NULL AUTO_INCREMENT,
		CONSTRAINT pk_amenity
		PRIMARY KEY (amenity_id),
    name VARCHAR(20) NOT NULL
);

CREATE TABLE room_amenity (
	room_id INT NOT NULL,
		CONSTRAINT fk_room_amenity_room
		FOREIGN KEY (room_id) REFERENCES room(room_id),
    amenity_id INT NOT NULL,
		CONSTRAINT fk_room_amenity_amenity
		FOREIGN KEY (amenity_id) REFERENCES amenity(amenity_id),
    CONSTRAINT pk_room_amenity
	PRIMARY KEY (room_id, amenity_id)
);

CREATE TABLE guest (
	guest_id INT NOT NULL AUTO_INCREMENT,
		CONSTRAINT pk_guest
		PRIMARY KEY (guest_id),
    first_name VARCHAR(25) NOT NULL,
    last_name VARCHAR(25) NOT NULL,
    address VARCHAR(25) NOT NULL,
    city VARCHAR(25) NOT NULL,
    /* state has 51 variants for the enum to account for Washington, DC */
    state ENUM("AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL",
			   "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME",
               "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH",
               "OK", "OR", "MD", "MA", "MI", "MN", "MS", "MO", "PA", "RI",
               "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY"),
	zip VARCHAR(5) NOT NULL,
    phone VARCHAR(10) NOT NULL
);

CREATE TABLE reservation (
	guest_id INT NOT NULL,
		CONSTRAINT fk_reservation_guest
		FOREIGN KEY (guest_id) REFERENCES guest(guest_id),
	room_id INT NOT NULL,
		CONSTRAINT fk_reservation_room
		FOREIGN KEY (room_id) REFERENCES room(room_id),
	CONSTRAINT pk_reservation
	PRIMARY KEY (guest_id, room_id),
    adults INT NOT NULL,
    children INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    total_cost DECIMAL(4, 2)
);