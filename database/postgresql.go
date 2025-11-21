package database

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var PostgresDB *sql.DB

func ConnectPostgres() {
    dsn := os.Getenv("POSTGRES_DSN")
    if dsn == "" {
        dsn = "host=localhost user=postgres password=2255 dbname=prestasi_mahasiswa port=5432 sslmode=disable"
    }

    var err error
    PostgresDB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("❌ Gagal koneksi ke PostgreSQL:", err)
    }

    if err = PostgresDB.Ping(); err != nil {
        log.Fatal("❌ Gagal ping PostgreSQL:", err)
    }

    log.Println("✅ PostgreSQL connected successfully")
}
