package cons

const ConnectionString = "root:root@tcp(localhost:3306)/rs_db?parseTime=true&loc=UTC"

const PropertyInsertQuery = `INSERT INTO rs_db.properties
								(description,
								type,
								orientation,
								courtyard_size,
								bedrooms,
								living_room_size,
								kitchen_size,
								bedrooms_sizes,
								bathrooms,
								showers,
								size,
								construction_year,
								padron,
								building_name,
								apartments_per_floor,
								floors,
								terrace_size,
								balcony_size,
								expenses,
								amenities,
								created_date,
								address_street,
								address_street_number,
								address_apartment_number,
								address_neighborhood,
								address_city,
								address_country,
								address_postal_code,
								address_latitude,
								address_longitude,
								elevators)
								VALUES
								(?,?,?,?,?,?,?,?,?,?,
								?,?,?,?,?,?,?,?,?,?,
								?,?,?,?,?,?,?,?,?,?,?);`

// const PropertyInsertQuery = `INSERT INTO rs_db.properties
// 								(type,
// 								bedrooms,
// 								bathrooms,
// 								showers,
// 								size,
// 								expenses,
// 								floors,
// 								courtyard_size,
// 								kitchen_size,
// 								address_street,
// 								address_neighborhood,
// 								address_city,
// 								address_country,
// 								address_postal_code,
// 								address_latitude,
// 								address_longitude)
// 								VALUES
// 								(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
