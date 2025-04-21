# Simple HTTP Uptime Monitor

This Go application provides a basic blackbox monitoring solution for HTTP services. It periodically probes defined targets and presents their uptime history through a simple web UI.

## Features (MVP)

* **HTTP Probing:** Monitors the availability of HTTP endpoints by sending requests.
* **Variable Intervals:** Allows you to configure different probing intervals for each target.
* **YAML Configuration:** Defines the target services and their monitoring settings using a straightforward YAML file.
* **Web UI:** Provides a user-friendly web interface to view the historical status of each monitored target.
* **Status Indicators:** Clearly indicates the status of each target as "Up", "Pending", or "Down".

## Getting Started

### Prerequisites

* **Go:** Version 1.24.4.

### Installation

1.  Clone the repository:
    ```bash
    git clone [https://github.com/RykoL/uptime-probe](https://github.com/RykoL/uptime-probe)
    cd uptime-probe
    ```

2.  Build the Go application.
    ```bash
    go build -o uptime-monitor .
    ```

### Configuration

1.  Create a `config.yaml` file in the same directory as the executable.

2.  Define your target services in the `config.yaml` file using the following structure:

    ```yaml
    targets:
      - name: My Website
        url: [https://www.example.com](https://www.example.com)
        interval: 60 # Interval in seconds
      - name: My API
        url: [https://api.example.com/health](https://api.example.com/health)
        interval: 30
      - name: Another Service
        url: http://localhost:8080
        interval: 120
    ```

    * `name`: A descriptive name for the target.
    * `url`: The HTTP URL to probe.
    * `interval`: The probing interval in seconds.

### Running the Application

1.  Execute the built binary from your terminal:
    ```bash
    ./uptime-monitor
    ```

2.  The application will start monitoring the defined targets in the background.

3.  Open your web browser and navigate to `http://localhost:8080` (or the address the application indicates) to view the uptime history.

## Web UI

The web UI provides a simple table displaying the status history for each configured target. For each target, you will typically see:

* **Name:** The name you defined in the `config.yaml`.
* **Status:** The latest status of the target ("Up", "Pending", or "Down").
* **History:** A chronological list or visual representation of recent status checks.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues for bugs and feature requests on the [GitHub repository](https://github.com/RykoL/uptime-probe).

## License

(Optional) If your project has a license, specify it here. For example:

[MIT License](LICENSE)