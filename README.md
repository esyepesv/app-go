# Calculator API

This is a simple Go application that provides a web server with the following functionality:

- It serves a greeting message on the root path (`/`).
- It calculates the square and cube of a given number provided as a path parameter.

## Getting Started

### Prerequisites

- [Go](https://golang.org/) (version 1.22 or later) must be installed on your system.

### Installation

1. Clone this repository:

    ```shell
    git clone <repository-url>
    ```

2. Navigate to the project directory:

    ```shell
    cd <project-directory>
    ```

### Usage

1. Run the application:

    ```shell
    go run .
    ```

2. The server will start on `http://localhost:5000`.

## Endpoints

- `GET /` - Responds with "Hola Mundo".
- `GET /cuadrado/<number>` - Responds with the square of the given number in the format "Cuadrado de \<number\> es \<result\>".
- `GET /cubo/<number>` - Responds with the cube of the given number in the format "Cubo de \<number\> es \<result\>".
