module main

require aph-go-service/transport v0.0.0

require aph-go-service/datastruct v0.0.0

require aph-go-service/logging v0.0.0

require aph-go-service/service v0.0.0

replace aph-go-service/transport => ./transport

replace aph-go-service/datastruct => ./datastruct

replace aph-go-service/logging => ./logging

replace aph-go-service/service => ./service

go 1.17

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.1.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.8.1 // indirect
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/lib/pq v1.10.3 // indirect
	github.com/rs/cors v1.8.0 // indirect
	github.com/unrolled/render v1.4.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20210917161153-d61c044b1678 // indirect
	golang.org/x/text v0.3.7 // indirect
	gorm.io/driver/postgres v1.1.2 // indirect
	gorm.io/gorm v1.21.16 // indirect
)
