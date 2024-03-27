-- name: CreateLocation :one
INSERT INTO locations (
    event,
    lat,
    long,
    block_no,
    apartment_name,
    apartment_number,
    street,
    city,
    country
) values (
   $1,$2, $3 ,$4, $5 ,$6 ,$7 ,$8 ,$9 
) RETURNING *;

-- name: GetLocation :one
SELECT * FROM locations 
WHERE event = $1
LIMIT 1;

-- name: UpdateLocation :one
UPDATE locations 
SET lat = $1,
    long = $2,
    block_no = $3,
    apartment_name = $4,
    apartment_number = $5,
    street = $6,
    city = $7,
    country = $8
WHERE id = $9
RETURNING *;