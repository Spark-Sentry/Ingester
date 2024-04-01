🚀 Data Ingester 📈

Welcome to the Data Ingester project, a high-performance and scalable solution designed to ingest real-time data into InfluxDB. This ingester efficiently processes messages from Kafka, transforms them, and securely stores them in InfluxDB for further analysis and monitoring.

## Features 🌟

- **Real-time Data Ingestion**: Seamlessly ingest data from Kafka topics into InfluxDB.
- **Scalable**: Designed to handle large volumes of data efficiently.
- **Configurable**: Easily configure Kafka and InfluxDB connections.
- **Monitoring**: Built-in support for application monitoring with Air for hot reloading during development.
- **Secure**: Ensures data is processed and stored securely.

## Getting Started 🚀

Follow these steps to get your Data Ingester up and running.

### Prerequisites

- Docker
- Go 1.16+
- Kafka running instance
- InfluxDB running instance

### Installation

1. **Clone the repository**

```bash
git clone https://github.com/your-username/data-ingester.git
cd data-ingester
```

2. **Setup InfluxDB and Kafka**

Ensure you have running instances of InfluxDB and Kafka. You can use Docker Compose configurations provided in the project for local development environments.

3. **Configure `.env`**

Duplicate `.env.example` to `.env` and adjust the configuration to match your Kafka and InfluxDB settings.

4. **Run with Air for Hot Reloading (Optional)**

```bash
air -d
```

### Usage

Start the Data Ingester by running:

```bash
go run cmd/ingester/main.go
```

The ingester will start consuming messages from the configured Kafka topic and ingest them into InfluxDB.

## Configuration 🛠

Refer to the `.env` file for all configurable options. Make sure to set the following:

- Kafka broker address
- Kafka topic name
- InfluxDB URL, token, and bucket name

## Contributing 🤝

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License 📝

Distributed under the MIT License. See `LICENSE` for more information.

## Contact 📬

Agostin Jean-baptiste - - Jbagostin@gmail.com

Project Link: [https://github.com/your-username/data-ingester](https://github.com/your-username/data-ingester)
