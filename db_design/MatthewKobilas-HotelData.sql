USE hotel_software_guild;

INSERT INTO amenity (amenity_name)
	VALUES
		("microwave"),
        ("jacuzzi"),
        ("refrigerator"),
        ("oven");

INSERT INTO room (room_id, beds, ada_access, standard_occupancy, max_occupancy, base_price, price_per_added_guest)
	VALUES
		(201, "double", FALSE, 2, 4, 199.99, 10.00),
        (202, "double", TRUE, 2, 4, 174.99, 10.00),
        (203, "double", FALSE, 2, 4, 199.99, 10.00),
        (204, "double", TRUE, 2, 4, 174.99, 10.00),
        (205, "single", FALSE, 2, 2, 174.99, NULL),
        (206, "single", TRUE, 2, 2, 149.99, NULL),
        (207, "single", FALSE, 2, 2, 174.99, NULL),
        (208, "single", TRUE, 2, 2, 149.99, NULL),
        (301, "double", FALSE, 2, 4, 199.99, 10.00),
        (302, "double", TRUE, 2, 4, 174.99, 10.00),
        (303, "double", FALSE, 2, 4, 199.99, 10.00),
        (304, "double", TRUE, 2, 4, 174.99, 10.00),
        (305, "single", FALSE, 2, 2, 174.99, NULL),
        (306, "single", TRUE, 2, 2, 149.99, NULL),
        (307, "single", FALSE, 2, 2, 174.99, NULL),
        (308, "single", TRUE, 2, 2, 149.99, NULL),
        (401, "suite", TRUE, 3, 8, 399.99, 20.00),
        (402, "suite", TRUE, 3, 8, 399.99, 20.00);

/*	1 - Microwave
	2 - Jacuzzi
    3 - Refrigerator
    4 - Oven */
INSERT INTO room_amenity (room_id, amenity_id)
	VALUES
		(201, 1), (201, 2),
        (202, 3),
        (203, 1), (203, 2),
        (204, 3),
        (205, 1), (205, 3), (205, 4),
        (206, 1), (206, 3),
        (207, 1), (207, 2), (207, 3),
        (208, 1), (208, 3),
        (301, 1), (301, 2),
        (302, 3),
        (303, 1), (303, 2),
        (304, 3),
        (305, 1), (305, 2), (305, 3),
        (306, 1), (306, 3),
        (307, 1), (307, 2), (307, 3),
        (308, 1), (308, 3),
        (401, 1), (401, 3), (401, 4),
        (402, 1), (402, 3), (402, 4);

INSERT INTO guest (full_name, address, city, state, zip, phone)
	VALUES
		("Matthew Kobilas", "25 Hickory Ln.", "Red Bank", "NJ", "07070", "7327778888"), -- 1
        ("Mack Simmer", "379 Old Shore Street", "Council Bluffs", "IA", "51501", "2915530508"), -- 2
        ("Bettyann Seery", "750 Wintergreen Dr.", "Wasilla", "AK", "99654", "4782779632"), -- 3
        ("Duane Cullison", "9662 Foxrun Lane", "Harlingen", "TX", "78552", "3084940198"), -- 4
        ("Karie Yang", "9378 W. Augusta Ave.", "West Deptford", "NJ", "08096", "2147300298"), -- 5 
        ("Aurore Lipton", "762 Wild Rose Street", "Saginaw", "MI", "48601", "3775070974"), -- 6
        ("Zachery Luechtefeld", "7 Poplar Dr.", "Arvada", "CO", "80003", "8144852615"), -- 7
        ("Jeremiah Pendergrass", "70 Oakwood St.", "Zion", "IL", "60099", "2794910960"), -- 8
        ("Walter Holaway", "7556 Arrowhead St.", "Cumberland", "RI", "02864", "4463966785"), -- 9
        ("Wilfred Vise", "77 West Surrey Street", "Oswego", "NY", "13126", "8347271001"), -- 10
        ("Maritza Tilton", "939 Linda Rd.", "Burke", "VA", "22015", "4463516860"), -- 11
        ("Joleen Tison", "87 Queen St.", "Drexel Hill", "PA", "19026", "2318932755"); -- 12

INSERT INTO reservation (guest_id, room_id, adults, children, start_date, end_date, total_cost)
	VALUES
		(2, 308, 1, 0, "2023-02-02", "2023-02-04", 299.98),
        (3, 203, 2, 1, "2023-02-05", "2023-02-10", 999.95),
        (4, 305, 2, 0, "2023-02-22", "2023-02-24", 349.98),
        (5, 201, 2, 2, "2023-03-06", "2023-03-07", 199.99),
        (1, 307, 1, 1, "2023-03-17", "2023-03-20", 524.97),
        (6, 302, 3, 0, "2023-03-18", "2023-03-23", 924.95),
        (7, 202, 2, 2, "2023-03-29", "2023-03-31", 349.98),
        (8, 304, 2, 0, "2023-03-31", "2023-04-05", 874.95),
        (9, 301, 1, 0, "2023-04-09", "2023-04-13", 799.96),
        (10, 207, 1, 1, "2023-04-23", "2023-04-24", 174.99),
        (11, 401, 2, 4, "2023-05-30", "2023-06-02", 1199.97),
        (12, 206, 2, 0, "2023-06-10", "2023-06-14", 599.96),
        (12, 208, 1, 0, "2023-06-10", "2023-06-14", 599.96),
        (6, 304, 3, 0, "2023-06-17", "2023-06-18", 184.99),
        (1, 205, 2, 0, "2023-06-28", "2023-07-02", 699.96),
        (9, 204, 3, 1, "2023-07-13", "2023-07-14", 184.99),
        (10, 401, 4, 2, "2023-07-18", "2023-07-21", 1259.97),
        (3, 303, 2, 1, "2023-07-28", "2023-07-29", 199.99),
        (3, 305, 1, 0, "2023-08-30", "2023-09-01", 349.98),
        (2, 208, 2, 0, "2023-09-16", "2023-09-17", 149.99),
        (5, 203, 2, 2, "2023-09-13", "2023-09-15", 399.98),
        (4, 401, 2, 2, "2023-11-22", "2023-11-25",  1199.97),
        (2, 206, 2, 0, "2023-11-22", "2023-11-25", 449.97),
        (2, 301, 2, 2, "2023-11-22", "2023-11-25", 599.97),
        (11, 302, 2, 0, "2023-12-24", "2023-12-28", 699.96);

DELETE FROM reservation
WHERE guest_id = 8;
DELETE FROM guest
WHERE guest_id = 8;