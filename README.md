# Enako [![BrandonWade](https://circleci.com/gh/BrandonWade/enako.svg?style=shield)](https://github.com/BrandonWade/enako)

A tool for recording & managing expenses to help with budgeting. Used to learn about modern React, Docker, Kubernetes, web security, and more.

## Getting Started

1. Run `cp .env.example .env`
2. Provide values for the environment variables
3. Run `docker-compose up --build -d`
4. Run `docker restart enako-api` (Note: the API container fails to connect to the DB container as it takes awhile to start)
5. Navigate to `localhost:8100`

## Testing

Tests are run automatically in CircleCI when pushing commits to master.

-   To generate or update test fakes, run `cd api && go generate ./...`
-   To run the backend test suite, run `ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --compilers=2 --untilItFails`

## Credit

Icons are from the [Entypo+ icon set](http://entypo.com/) by Daniel Bruce.
