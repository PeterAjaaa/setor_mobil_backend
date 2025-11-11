# Setor-Mobil Backend

## TODO

- Better logging (middleware logging, instead of per-function logging, perhaps?)
- GraphQL for fetching data to fetch only the data needed (e.g. overfetching on Car DTO)
- Make better DB scheme (e.g. allow multiple car/bike per order, merge cars/bikes table as vehicles if separation isn't needed)
- Refactor login function, it's close enough that I think it can be turned into a generic function, usable for both admin and user login
- Protect certain routes with a new middleware, making it admin-only. Currently user JWT token can let them access admin endpoint. For now, this is fine, mitigated by the frontend being two different applications. But in the future, this is a bad security hole. Important to fix.
- Use OpenAPI/Swagger to generate code from specification, for REST
