# API_GO
# Event Management System

A simple event management system built with Go that supports CRUD operations on events. Only authorized users can create, update, and delete events. Additionally, users can register and unregister for events, and fetch the list of users registered for a single event.

## Features

- **Create an Event**: Authorized users can create new events.
- **Read Events**: Fetch the list of all events.
- **Update an Event**: Only the user who created the event can update it.
- **Delete an Event**: Only the user who created the event can delete it.
- **Register for an Event**: Authorized users can register for an event.
- **Unregister from an Event**: Authorized users can unregister from an event.
- **Fetch Registered Users**: Fetch the list of users registered for a single event.

## Requirements

- Go 1.16+
- A database (e.g., PostgreSQ,Sql lite)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/your-username/event-management-system.git
    cd event-management-system
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```


## Usage

### Running the Server

Start the server:
```sh
go run main.go
