# SGS FULLSTACK CHALLENGE 
Fullstack app to fetch products from a mongoDB collection

## How to Install

- Backend :
    - Clone this repo
    - Run `go mod tidy` to install dependencies
    - Fill the `.env.template` file with your configurations, rename it to `.env`
    - Run `go run main.go`
- Frontend :
    - Go to the `frontend` directory of this project
    - Run `npm install` to install dependencies
    - Run `npm start` to start the front end client

## Notes

- If the DB collection is empty we run a small seeder to populate it with random data
