USE hotel_software_guild;

/* 1 - Write a query that returns a list of reservation that end in July 2023, including the name of the guest, the room number(s),
	   and the reservation dates. */
SELECT
	guest.full_name,
    reservation.room_id,
    reservation.start_date,
    reservation.end_date
FROM reservation
INNER JOIN guest
	ON reservation.guest_id = guest.guest_id
WHERE reservation.end_date BETWEEN "2023-07-01" AND "2023-07-31";
/*
Matthew Kobilas	205	2023-06-28	2023-07-02
Bettyann Seery	303	2023-07-28	2023-07-29
Walter Holaway	204	2023-07-13	2023-07-14
Wilfred Vise	401	2023-07-18	2023-07-21
*/

/* 2 - Write a query that returns a list of all reservations for rooms with a jacuzzi, displaying the guest's name,
	   the room number, and the dates of the reservation. */
SELECT
	guest.full_name,
    reservation.room_id,
    reservation.start_date,
    reservation.end_date
FROM reservation
INNER JOIN guest
	ON reservation.guest_id = guest.guest_id
INNER JOIN room_amenity
	ON reservation.room_id = room_amenity.room_id
INNER JOIN amenity
	ON room_amenity.amenity_id = amenity.amenity_id
WHERE amenity.amenity_id = 2;
/*
Karie Yang	201	2023-03-06	2023-03-07
Bettyann Seery	203	2023-02-05	2023-02-10
Karie Yang	203	2023-09-13	2023-09-15
Wilfred Vise	207	2023-04-23	2023-04-24
Mack Simmer	301	2023-11-22	2023-11-25
Walter Holaway	301	2023-04-09	2023-04-13
Bettyann Seery	303	2023-07-28	2023-07-29
Bettyann Seery	305	2023-08-30	2023-09-01
Duane Cullison	305	2023-02-22	2023-02-24
Matthew Kobilas	307	2023-03-17	2023-03-20
*/

/* 3 - Write a query that returns all the rooms reserved for a specific guest, including the guest's name, the room(s) reserved,
	   the starting date of the reservation, and how many people were included in the reservation. (Choose a guest's name from
       the existing data.) */
SELECT
	guest.full_name,
    reservation.room_id,
    reservation.start_date,
    reservation.adults,
    reservation.children
FROM reservation
INNER JOIN guest
	ON guest.guest_id = reservation.guest_id
WHERE guest.full_name = "Matthew Kobilas";
/*
Matthew Kobilas	205	2023-06-28	2	0
Matthew Kobilas	307	2023-03-17	1	1
*/

/* 4 - Write a query that returns a list of rooms, reservation ID, and per-room cost for each
	   reservation. The results should include all rooms, whether or not there is a reservation
       associated with the room. */
SELECT
	room.room_id,
    CONCAT(reservation.guest_id, reservation.room_id) AS reservation_id,
    reservation.total_cost
FROM reservation
RIGHT JOIN room
	ON room.room_id = reservation.room_id;
/*
201	5201	199.99
202	7202	349.98
203	3203	999.95
203	5203	399.98
204	9204	184.99
205	1205	699.96
206	2206	449.97
206	12206	599.96
207	10207	174.99
208	2208	149.99
208	12208	599.96
301	2301	599.97
301	9301	799.96
302	6302	924.95
302	11302	699.96
303	3303	199.99
304	6304	184.99
305	3305	349.98
305	4305	349.98
306		
307	1307	524.97
308	2308	299.98
401	4401	1199.97
401	10401	1259.97
401	11401	1199.97
402		
*/

/* 5 - Write a query that returns all rooms with a capacity of three or more and that are
	   reserved on any date in April 2023. */
SELECT
	room.room_id
FROM room
INNER JOIN reservation
	ON room.room_id = reservation.room_id AND room.max_occupancy >= 3
WHERE
	MONTH(reservation.start_date) = 4 OR MONTH(reservation.end_date) = 4;
/*
301
*/

/* 6 - Write a query that returns all rooms with a capacity of three or more and that are reserved on any date in April 2023. */
SELECT
	guest.full_name,
    COUNT(reservation.guest_id) AS reservationCount
FROM reservation
INNER JOIN guest
	ON reservation.guest_id = guest.guest_id
GROUP BY reservation.guest_id
ORDER BY reservationCount DESC, SUBSTRING(guest.full_name, LOCATE(" ", guest.full_name)) ASC;
/*
Mack Simmer	4
Bettyann Seery	3
Duane Cullison	2
Walter Holaway	2
Matthew Kobilas	2
Aurore Lipton	2
Maritza Tilton	2
Joleen Tison	2
Wilfred Vise	2
Karie Yang	2
Zachery Luechtefeld	1
*/

/* 7 - Write a query that displays the name, address, and phone number of a guest based on their
	   phone number. (Choose a phone number from the existing data.) */
SELECT
	guest.full_name,
    guest.address,
    guest.phone
FROM guest
WHERE guest.phone = "3775070974";
/*
Aurore Lipton	762 Wild Rose Street	3775070974
*/