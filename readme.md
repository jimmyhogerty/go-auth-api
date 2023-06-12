WIP~~~~~~~~~~~~~~~~
To connect to DB:

in one terminal: docker compose up
in another terminal:
docker exec -it lenslocked-db-1 /usr/bin/psql -U baloo -d lenslocked

to see Adminer visual Postgres DB, visit:
http://localhost:3333/?server=db&username=baloo

Note you must switch "System" from SQL to PostgreSQL.

DEPS
Gorilla/csrf
Gorilla/csrf for cross-site request forgery (CSRF) protection.
Github: https://github.com/gorilla/csrf
