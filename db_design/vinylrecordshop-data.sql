-- Matthew Kobilas
-- 06/09/20
USE vinylrecordshop;

INSERT INTO album (albumTitle, releaseDate, price, label) VALUES
	("Imagine", "1971-09-09", 9.99, "Apple"),
    ("2525 (Exordium & Terminus)", "1969-07-01", 25.99, "RCA"),
    ("No One's Gonna Change Our World", "1969-12-12", 39.35, "Regal Starline"),
    ("Moondance Studio Album", "1969-08-01", 14.99, "Warner Bros"),
    ("Clouds", "1969-05-01", 9.99, "Reprise"),
    ("Sounds of Silence Studio Album", "1966-01-17", 9.99, "Columbia"),
    ("Abbey Road", "1969-01-10", 12.99, "Apple"),
    ("Smiley Smile", "1967-09-18", 5.99, "Capitol");
SELECT * FROM album;

DELETE FROM album
WHERE albumId = 5;
INSERT INTO album (albumTitle, releaseDate, price, label)
VALUES ("Clouds", "1969-05-01", 9.99, "Reprise");
SET SQL_SAFE_UPDATES = 0;
UPDATE album
	SET albumId = 5
WHERE albumTitle = "Clouds";
UPDATE album
	SET albumId = 9
WHERE albumTitle = "Smiley Smile";
SET SQL_SAFE_UPDATES = 1;