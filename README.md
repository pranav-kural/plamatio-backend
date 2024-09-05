# Plamatio Backend

Plamatio is a production-ready E-commerce store focused on Llama-inspired products. This e-commerce store was built as part of [my portfolio](https://www.pkural.ca/) projects for demonstration of some cutting-edge technologies in building a highly performant, scalable, and resilient backend that composes of a robust Go-lang written REST API with PostgreSQL databases, Kafka-based real-time data streaming, Redis cache system, backend system monitoring and observability through Grafana, and more.

[Plamatio Backend Documentation](https://pkural.notion.site/Plamatio-Backend-Documentation-d8c426f7851546c19df095c7fbf72282)

![Plamatio Backend Deployment - Staging](https://img.shields.io/github/deployments/pranav-kural/plamatio-backend/staging?label=staging)
![Plamatio Backend Deployment - Production](https://img.shields.io/github/deployments/pranav-kural/plamatio-backend/prod?label=production)



## Primary Features

- **Distributed REST API:** Plamatio backend consists of multiple distributed modularized REST API services, each with its Redis cache and PostgreSQL database. This distributed architecture allows for maximal scalability, reduced development and maintenance complexity, and increased productivity in fixing issues and releasing feature updates.
- **PostgreSQL Database + Redis Cache:**  Low-latency request processing pipeline architected with a robust PostgreSQL database (with special indexes) and a distributed, responsive and scalable Redis-based cache for frequently accessed API endpoints.
- **API-key Authenticated Endpoints:** Simple API-key based authentication for most of the endpoints, reducing overhead and request processing time for both frontends and the backend, without compromise on security.
- **Kafka-powered Real-time Data Streaming:** Using Confluent-based Kafka services for architecting a real-time data streaming pipeline to not only provide a scalable solution for real-time updates on data mutations but also to enhance capturing of interactions and events on any number of frontends.
- **Grafana:** Using Grafana Cloud for enchanced monitoring and observability going beyond basic logging, metrics collection and analysis, and API request tracing.
- **Encore Cloud:** Using Encore Cloud for hosting backend application, secrets management, PostgreSQL database hosting, Redis hosting, and resource provisioning.

Below image provides a high level overview of the primary components of the Plamatio Backend Infrastructure.
<p align="center">
<img alt="Plamatio Backend Core Components" src="https://github.com/user-attachments/assets/6f60a6ca-2d79-49c7-89a0-f23cead66ab3" width="500px" />
</p>

## Project Structure

Plamatio backend is built using the [Encore.go Backend Framework](https://encore.dev/docs/go).

Project is structured in a way that reduces complexity and increases productivity. Since, Encore enables you to build distributed API services, dependency between each service is minimal.

There are five key services that Plamatio Backend exposes: Products, Categories, Cart, Orders, Users.

For each of these services, there are four key folders:

- `api`: contains the core REST API endpoints defintions and code.
- `db`: contains and abstracts methods required to interact with the required table(s) in the PostgreSQL database.
- `models`: defines the data models used by the service.
- `utils`: provides any utility functions used across the code of the service.

Basic workflow when adding a new API service would look like:

1. **Define Data Models:** Define data models required by the service. For example, Products service may require a `Product` type definition to store and work with products.
2. **Define Utility Functions:** Define data validation functions for validating data for creating instances for your defined data models. For example, for Products service, you may define a method to validate data received in a request to create a new product.
3. **Define Database Methods:** Define methods to interact with the appropriate table in the PostgreSQL database for the service. You might start by defining the SQL statements, followed by defining the methods that handle the execution of these statements.
4. **Define Service API:** Define the API endpoint methods for the service. This may involve use of the database methods you defined earlier. Moreover, at this stage, you may also configure cache for storing data for frequently accessed API endpoints, and auth handler for authenticated API endpoints.


## API Specifications

Please refer to [Plamatio Backend REST API Specifications](https://pkural.notion.site/REST-API-Specifications-c3fe4301baec4f23a01a86373896ff6a) for more information on the REST API endpoints and their specifications.

## REST API Architecture

Below image presents a high-level overview of the distributed and scalable REST API architecture of the Plamatio Backend.

<p align="center">
<img alt="Plamatio Backend REST API Architecture" src="https://github.com/user-attachments/assets/c2f2b947-d067-49b9-8dca-0df2c83627ed" />
</p>

## Real-time Data Streaming

To keep a scalable amounts of frontend interfaces in sync with data mutations, Plamatio uses a Confluent-based Kafka Service architecture to listen to and stream real-time data updates.

<p align="center">
<img alt="Plamatio Backend Real-time Data Streaming" src="https://github.com/user-attachments/assets/661c10a3-1cca-4c8d-aadb-cc88429d3f18" width="700px" />
</p>

## Local Development

Feel free to use this project as basis for developing your own high-quality scalable backend systems.

Below are some sections to assist in you in this.

### Developing locally

To develop this project locally, you will need to install the [Encore CLI](https://encore.dev/docs/install).

Once you have Encore CLI installed, you can use the below command to clone this project.

```bash
encore app clone plamantio-backend-hoki [directory]
```

### Running locally

Before running your application, make sure you have Docker installed and running. It's required to locally run Encore applications with databases.

```bash
encore run
```

### Open the developer dashboard

While `encore run` is running, open [http://localhost:9400/](http://localhost:9400/) to access Encore's [local developer dashboard](https://encore.dev/docs/observability/dev-dash).

Here you can see API docs, make requests in the API explorer, and view traces of the responses.

### Deployment

Deploy your application to a staging environment in Encore cloud:

```bash
git add -A .
git commit -m 'Commit message'
git push encore
```

Then head over to the [Cloud Dashboard](https://app.encore.dev) to monitor your deployment and find your production URL.

From there you can also connect your own AWS or GCP account to use for deployment.

### Testing

```bash
encore test ./...
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Issues

If you encounter any issues or bugs while using Plamatio Backend, please report them by following these steps:

1. Check if the issue has already been reported by searching our issue tracker.
2. If the issue hasn't been reported, create a new issue and provide a detailed description of the problem.
3. Include steps to reproduce the issue and any relevant error messages or screenshots.

[Open Issue](https://github.com/pranav-kural/plamatio-backend/issues)
