# Using golang to connect to an Oracle Database

This is a small snippet to connect to an Oracle Database using GO 1.13.6. It includes a Dockerfile to run this image.



##Pre requisite:
- Oracle requests for "Oracle Instant" installed in the host machine. Reference: https://www.oracle.com/database/technologies/instant-client/downloads.html. This Oracle Instant Dockerfile is included in the folder `./oracle-instant` of the repository. It is needed to run the main Dockerfile.
- An Oracle Database running

## Environment variables
- `DB_URL`. The database URL in the following format: `{host}:{port}/{service_name}`
- `DB_USER`. Username to connect to database service.
- `DB_PASS`. Password to connect to database service.
- `QUERY`. The query to run in the database.


