env "local" {
    src = "file://schema/schema.sql"
    url = getenv("DATABASE_URL")
}
