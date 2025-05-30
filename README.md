# Zeek-MISP Threat Detection System.

This project is designed to receive logs from **Zeek**, compare them against **MISP** IOCs (Indicators of Compromise), and send matched threat data to **OpenSearch** for visualization. The system is built with *Go*, runs inside a *Docker container*, and uses *Redis* for data caching.

This project is designed to receive log data from **Zeek**, a tool used for network monitoring and logging network events. The received logs are compared against Indicators of Compromise (IOCs) stored in **MISP (Malware Information Sharing Platform)** to identify "threats" or "potentially malicious events." When a match is found with the specified IOC, the system sends the detected threat data to **OpenSearch** for visualization and further analysis.

---

## The main components.

* **Zeek**: Zeek acts as a tool for detecting and logging network traffic activities. It helps the system identify potential threats from incoming data, such as external attacks, unauthorized data access, or the use of prohibited protocols.

* **MISP (Malware Information Sharing Platform)**: MISP is used to store and share Indicators of Compromise (IOCs) that are used to detect various threats. The system pulls IOC data from MISP to compare it against the logs received from Zeek.

* **Redis**: Redis is used for data caching to improve processing efficiency. By storing frequently accessed data in memory, Redis allows faster access, improving the performance of data processing and comparison.

* **Fluent Bit**: Fluent Bit is an efficient log management and forwarding tool. It is used to collect logs from Zeek and forward them to other systems like Redis or OpenSearch. Fluent Bit facilitates data extraction from multiple sources (Multiple Inputs) and handles tasks such as filtering or transforming the data for further processing. Additionally, Fluent Bit works with Zeek and Redis to send matched IOC data to OpenSearch quickly and efficiently.

* **Go**: Go is the programming language used to develop this project due to its fast processing capabilities and support for concurrent operations. Go is well-suited for systems that handle large volumes of data and require real-time processing.

* **Docker**: The system operates within Docker containers, which helps isolate the project from the underlying operating system. This allows for easier configuration and management of different environments, ensuring standardized execution for testing and development purposes.

---

## Getting Started

### Step 1 install these tools.

1. [ Install and config Fluent-bit](docs/Config-fluent-bit.md)
2. [ Inatall and Config zeek](docs/Config-zeek.md)
3. [ Install redis](docs/install-redis.md)

###  Step 2 Create ``.env`` File and ``alert.log`` 

Create a ``.env`` file in the Zeek-MISP folder.
```bash
MISP_URL=https://<your-misp-ip>/attributes/restSearch.json
MISP_API_KEY=your_misp_api_key
```
Create a ``alert.log``  file in the Zeek-MISP folder.
```bash
touch /Desktop/alert.log
```
### Step 3 Build docker.
```bash
docker build -t < name images > .
```
```bash
docker run -d \
  --name < name container > \
  --network host \
  -v "$HOME/Desktop/alert.log:/alert.log" \
  -v /etc/localtime:/etc/localtime:ro \
  -v $(pwd)/.env:/app/.env \
  < name images >
```
--------




