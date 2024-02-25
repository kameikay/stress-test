# Stress Test CLI

## Description
This is a CLI system written in Go to perform load testing on a web service. It allows you to specify a service URL, the total number of requests, and the concurrency level to test the web service.

## Usage
Make sure you have Go installed on your system. To run the project, follow these steps:

1. Clone the repository to your system:
    ```
    git clone https://github.com/kameikay/stress-test.git
    ```

2. Navigate to the project directory:
    ```
    cd your-repository
    ```

3. Build the project:
    ```
    go build
    ```

4. Execute the project, providing the service URL, the total number of requests, and the concurrency level:
    ```
    ./stress-test --url=http://example.com --requests=1000 --concurrency=10
    ```

Alternatively, you can run the project using Docker. Follow these steps:

1. Pull the Docker image from Docker Hub:
    ```
    docker pull kameikay/stress-test:latest
    ```

2. Run the Docker container, providing the necessary parameters:
    ```
    docker run kameikay/stress-test:latest --url=http://example.com --requests=1000 --concurrency=10
    ```

## Parameters

- `--url`: The URL of the service to be tested.
- `--requests`: The total number of requests.
- `--concurrency`: The number of simultaneous calls.

## Report
After running the tests, the system generates a report with the following information:
- Total time spent on execution
- Total number of requests made
- Number of requests with HTTP status 200
- Distribution of other HTTP status codes (such as 404, 500, etc.)

## Testing
The project includes unit tests to ensure the correct functionality of the features.

To run the tests, use the following command:

```
go test ./...
```

## Dependencies
Make sure to install this dependency before running the project:

```
go mod tidy
```

## Contributing
Contributions are welcome! If you find an issue or have any suggestions to improve this project, feel free to open an issue or submit a pull request.

