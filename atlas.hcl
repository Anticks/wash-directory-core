env "local" {
  url = "postgres://user:password@localhost:5432/directory_core?sslmode=disable"

  migration {
    dir = "file://migrations"
  }
}
