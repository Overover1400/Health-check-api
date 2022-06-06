wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

#whatch for .go file changes
compileDaemon --build="go build -o main main.go"