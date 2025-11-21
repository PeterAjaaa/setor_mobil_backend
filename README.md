# Setor-Mobil Backend

## ERD

![Setor Mobil Backend ERD](./Setor-Mobil-Backend-ERD.png "Setor Mobil Backend ERD")

**WHENEVER THERE'S _BOTH_ MOTORCYCLE_ID AND CAR_ID IN THE SAME TABLE, THEIR RELATION ARE EXCLUSIVE TO EACH OTHER**

**MEANING THAT BOTH FIELDS _CAN NOT_ BE POPULATED SIMULTANEOUSLY**

**BUT THERE MUST BE _EXACTLY ONE_ NON-NULL FIELD**

## TODO

- Better logging (middleware logging, instead of per-function logging, perhaps?)
- GraphQL for fetching data to fetch only the data needed (e.g. overfetching on Car DTO)
- Make better DB scheme (e.g. allow multiple car/bike per order, merge cars/bikes table as vehicles if separation isn't needed)
- Refactor login function, it's close enough that I think it can be turned into a generic function, usable for both admin and user login
- Protect certain routes with a new middleware, making it admin-only. Currently user JWT token can let them access admin endpoint. For now, this is fine, mitigated by the frontend being two different applications. But in the future, this is a bad security hole. Important to fix.
- Use OpenAPI/Swagger to generate code from specification, for REST API
- Change seeding logic to only once if exist.
- Change how DTOs are used now. Maybe something that automatically fills the field by some kind of rule or something (e.g. Cars returning nil when everything works properly. Turns out it's DTO lagging behind and isn't setup to fill properly)
- Change 'Rating' field in Car and Motorcycle model to a uint8, where this 'Rating' is calculated automatically using background worker, the trigger being new entry being added in the Rating table. The 'Rating' field value in Car and Motorcycle would be calculated by getting the average value for each of the Car or Motorcycle.
- Zero-downtime deployment with blue-green strategy (maybe with Docker Rollout, or something.)
- Cloudflare Turnstile protecting routes like register and login
- Add timeout to login attempt to prevent brute-force
- Implement unit testing for each of the functions
- Refactor services functions, since it looks like some of them share the same characteristics for handling certain operations, no matter the context
- Add realtime data update to clients, maybe with WebSocket or similar
- Add healthcheck endpoint for API
- Add metrics endpoint for observability
- Set service fee and insurance from backend

## Flow for new feature addition

1. Define tables in `models/`
2. Refactor other tables to resolve relationship, if needed
3. Implement `Model` interface, and other interfaces if applicable.
4. Edit the migration to add the newly made table
5. Define the DTO needed for said feature
6. Implement the services needed
7. Set the route and the function called on said route
8. Add seed data, and call it
