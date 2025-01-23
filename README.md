# 200sh Dashboard

![200sh Logo](./public/200sh-logo.svg)

## Project Description

200sh Dashboard is a web-based application designed to provide a user-friendly interface for managing and monitoring various aspects of the 200sh platform. It includes features such as user authentication, real-time monitoring, and more.

## Installation

To install the 200sh Dashboard, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/200sh/200sh-dashboard.git
   cd 200sh-dashboard
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   npm install --prefix tailwind
   ```

3. Build the project:
   ```sh
   go build -o ./tmp/main ./cmd/
   ```

## Configuration

To configure the project, create a `.env` file in the root directory with the following content:
```sh
HANKO_API_URL="" # Replace with your project's API URL
DATABASE_NAME="dev.db" # Name of the database. Currently hardcoded to be a sqlite3 database
LOG_LEVEL="WARN" # WARN is the default level. Others: DEBUG, INFO, WARN, ERROR, OFF
DEV="false" # false is the default. Setting to true will remove some caching making live reloading better.
```

### Admin API Key
To generate an admin api key for the dashboard, run the following command:
```shell
openssl rand -base64 32
```

## Usage

To start the 200sh Dashboard, run the following command:
```sh
./tmp/main
```

The server will start on `http://localhost:3000`.

for live reloading when developing we recommend using [air](https://github.com/air-verse/air) and running the server
with the command:
```shell
air
```


## Contribution Guidelines

We welcome contributions to the 200sh Dashboard! To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Make your changes and commit them with clear commit messages.
4. Push your changes to your forked repository.
5. Create a pull request to the main repository.
