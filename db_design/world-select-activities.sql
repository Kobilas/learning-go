USE world;

-- Activity 1
SELECT * FROM city
LIMIT 10;
SELECT * FROM country
LIMIT 10;
SELECT * FROM countrylanguage
LIMIT 10;

-- Activity 2
SELECT * FROM city
WHERE Population < 10000
ORDER BY Population DESC;

-- Activity 3
SELECT
	country.region,
    country.`Name` AS countryName,
    city.`Name` AS cityName
FROM city
JOIN country
	ON city.CountryCode = country.`Code`
GROUP BY country.region, country.`Name`
ORDER BY country.region ASC, country.`Name` ASC, city.`Name` ASC;

-- Activity 4
SELECT
	country.`Name`,
    countrylanguage.`Language`,
    countrylanguage.Percentage
FROM country
INNER JOIN countrylanguage
	ON country.`Code` = countrylanguage.CountryCode
WHERE countrylanguage.`Language` LIKE "%french%"
ORDER BY countrylanguage.Percentage DESC;

-- Activity 5
SELECT
	country.`Name`,
    country.Continent,
    country.Population
FROM country
WHERE ISNULL(country.IndepYear)
ORDER BY country.`Name` ASC;

-- Activity 6
SELECT
	country.`Name`,
    country.Continent,
    countrylanguage.`Language`,
    countrylanguage.Percentage
FROM country
LEFT JOIN countrylanguage
	ON country.`Code` = countrylanguage.CountryCode
ORDER BY country.`Name` ASC, countrylanguage.Percentage DESC;

-- Activity 7
SELECT
	country.`Name`,
	country.Continent
FROM country
JOIN countrylanguage
	ON country.`Code` = countrylanguage.CountryCode
WHERE ISNULL(countrylanguage.`Language`)
ORDER BY country.Continent DESC, country.`Name` DESC;

-- Activity 8
SELECT
	country.`Name`,
    SUM(city.Population) AS cityPopulation
FROM city
JOIN country
	ON city.CountryCode = country.`Code`
GROUP BY country.`Name`
ORDER BY cityPopulation ASC;

-- Activity 9
SELECT 
	country.Continent,
    AVG(city.Population) AS avgCityPopulation
FROM city
JOIN country
	ON city.CountryCode = country.`Code`
GROUP BY country.Continent
ORDER BY avgCityPopulation DESC;

-- Activity 10
SELECT
	country.`Name`,
    country.GNP
FROM country
ORDER BY country.GNP DESC
LIMIT 10;

-- Activity 11
SELECT
	city.`Name` AS cityName,
    country.`Name` AS countryName,
    city.Population AS cityPopulation,
    countrylanguage.`Language` AS countryLanguage
FROM city
JOIN country
	ON city.ID = country.Capital
JOIN countrylanguage
	ON country.`Code` = countrylanguage.CountryCode
WHERE countrylanguage.IsOfficial = TRUE
ORDER BY cityName ASC;

-- Activity 12
SELECT
	country.`Name` AS countryName,
    city.`Name` AS cityName
FROM country
JOIN city
	ON country.Capital = city.ID
ORDER BY country.`Name` ASC;