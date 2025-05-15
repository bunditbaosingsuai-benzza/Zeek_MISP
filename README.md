# Zeek-MISP Threat Detection System

This project is designed to receive logs from **Zeek**, compare them against **MISP** IOCs (Indicators of Compromise), and send matched threat data to **OpenSearch** for visualization. The system is built with **Go**, runs inside a **Docker container**, and uses **Redis** for data caching.

---

## System Architecture

- **Zeek** generates network security logs.
- **Fluent Bit** sends logs over **TCP port 5050** to a **Go application**.
- The Go app:
  - Fetches MISP IOCs every 5 minutes.
  - Stores new IOCs in **Redis** (avoiding duplicates).
  - Compares incoming logs against IOC values.
  - If matched:
    - Displays alerts in the console.
    - Writes alerts to a local file `alert.log`.
    - Sends alerts to **OpenSearch** for dashboard analytics.

Everything is containerized via **Docker** for portability and consistency.

---

## ⚙️ Getting Started

### 1.  Create `.env` File and 

In the root directory of the project, create a file named `.env` with the following content:

.env
MISP_URL=https://<your-misp-ip>/attributes/restSearch.json
MISP_API_KEY=your_misp_api_key

creat to /Desktop/alert.log

##  Build docker

```bash

docker build -t < name images > .

docker run -d \
  --name < name container > \
  --network host \
  -v "$HOME/Desktop/alert.log:/alert.log" \
  -v /etc/localtime:/etc/localtime:ro \
  -v $(pwd)/.env:/app/.env \
  < name images > 

