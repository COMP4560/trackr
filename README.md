<p align="center">
    <img src='logo.svg?raw=true' width='50%'>
</p>

---

A platform for makers, hobbyists, students, and professionals to easily store and visualize data recorded from internet-connected devices like [Arduinos] and [Raspberry Pis].

- [Motivation](#motivation)
- [Features](#features)
- [Building and Running](#building-and-running)
  - [Backend](#backend)
  - [Frontend](#frontend)
- [Tests](#tests)
- [Architecture](#architecture)
- [Licensing](#licensing)

## Motivation

There exists a community of makers, students, hobbyists, and professionals who create and design systems that are capable of collecting sensor data like humidity, temperature, and atmospheric pressure.

These makers wish to store and visualize their sensor data, but there exist barriers in place like managing databases, web servers, etc.

As such, Trackr sets out to make it easier for these makers to store and visualize their data.

## Features

- Fully open-source codebase.
- Docker image builds for frontend/backend, along with a configured Docker architecture which includes a MySQL database and Nginx load balancer.
- Graph, table, and pie chart visualizations of sensor data.
- Easy-to-use REST API that enables the storing and retrieval of sensor data.
- Aggregation of sensor data by month, days, and hours.
- Compatible with Raspberry Pis and Arduino devices.
- Slick and responsive web-interface.
- Alerting system using a Discord bot that can be installed on any Discord server.

## Building and Running

Currently, Trackr consists of two primary components: the frontend and the backend.
The front end is the website that users interact with, whereas the back end is the REST API that responds to requests made by users through the front end.

Additionally, the production architecture for Docker uses Nginx and a container-based MySQL database.
All required services for Trackr can be initialized as containers easily using the Docker setup commands.

### Environment Variables Configuration

The current backend and production frontend environment variables are configured to work with Trackr's deployment. 
To develop locally, the backend and production frontend environment variables will have to be modified to reference `localhost` instead of `wryneck.cs.umanitoba.ca`.

To accomplish this, modify `backend/.env` and change any instances of `wryneck.cs.umanitoba.ca` to `localhost`.
The Docker deployment will now work locally.

Note that the production frontend environment variables must be modified to work with Docker locally as that is the environment variable file that is used in the frontend Docker image.
The development environment variable file that exists in the frontend is only used when building the frontend without Docker.

This configuration can be useful to resolve issues with the production Docker setup, however if a debugger is required then starting the frontend and backend without the use of Docker is recommended.

#### Non-Docker Environment Variable Configuration

If Docker is not being used to start the frontend and backend, the MySQL database will also not exist by default.
Either a different MySQL database must be used, the Docker MySQL database must be used, 
or the `DB_TYPE` and `DB_CONNECTION_STRING` must be changed to use the SQLite database in the backend environment variable file. 

The required `DB_TYPE` and `DB_CONNECTION_STRING` to connect to SQLite is already defined in the backend environment variable file, 
however it is commented out. Simply comment out the current `DB_TYPE` and `DB_CONNECTION_STRING` variables which are configured for MySQL, 
and uncomment the SQLite versions. Trackr will then use an in-memory SQLite database instead of the MySQL database.

### Building Trackr using Docker

To build Trackr using the Docker configuration, Docker must be installed and running on your machine.
Execute the following commands to start Trackr using Docker:

1.  `docker-compose build`
2. `docker-compose up --scale server=X`

Note that X is the number of Trackr server instances to start, for example use X=10 to start ten Trackr instances:

- `docker-compose up --scale server=10`

The following command can be used to remove all Docker images, which can be used to clear up space:
-  `docker system prune -a --volumes`


### Non-Docker Backend

**Note:** When building or running, the Go server executable will use the `.env` file for its environment variables on runtime. Make sure this file exists when running the server executable.

#### Building

To build the backend, make sure you have [Go] installed and GCC installed.

1. From the root directory, navigate to the `backend/` directory.
2. Type `make`

Your build files will be located in the `bin/` directory.

#### Running

In addition to building, you can also run the backend in a development environment.
To run the backend, make sure you have [Go] installed and GCC installed.

1. From the root directory, navigate to the `backend/` directory.
2. Type `make run`
3. Press `Control + C` on your keyboard to exit.

### Non-Docker Frontend

#### Building

To build the frontend, make sure you have [NodeJS] installed.

**Note:** When building, React will use the `.env.production` file for its environment variables. Make sure this file exists when building&mdash;this file is only used when building.

1. From the root directory, navigate to the `frontend/` directory.
2. If it is your first time building the front end, type `make `clean`, otherwise, skip this step.
3. Type `make`

Your build files will be located in the `build/` directory.

#### Running

In addition to building, you can also run the frontend in a development environment.
To run the frontend, make sure you have [NodeJS] installed.

**Note:** When running, React will use the `.env.development` file for its environment variables. Make sure this file exists when running the frontend.

1. From the root directory, navigate to the `frontend/` directory.
2. Type `make run`
3. Press `Control + C` on your keyboard to exit.

## Tests

At the moment, Trackr only has tests available for the backend component. There are two types of tests: database service implementation tests and API controller endpoint tests.

To run the tests:

1. From the root directory, navigate to the `backend/` directory.
2. Type `make test`

## Architecture

The following is the core Trackr architecture. When making use of Docker, Nginx also exists between the web application and the backend server(s) which distributes requests to multiple backend Trackr servers.

![image](https://user-images.githubusercontent.com/63835313/214208475-0d327caf-95e1-4ed3-88d4-df5f27dffdc8.png)

## Licensing

For more information, see ...

[raspberry pis]: https://www.raspberrypi.org/
[arduinos]: https://www.arduino.cc/
[nodejs]: https://nodejs.org/en/
[go]: https://go.dev/
