# Zeek + MISP Threat Detection System

##  Project Overview

This project integrates **Zeek**, **MISP**, **Fluent Bit**, **Go**, **Redis**, and **OpenSearch** to create a lightweight threat detection pipeline. It monitors network logs, compares them with threat intelligence (MISP), and visualizes the results via OpenSearch Dashboards.

---

## 🔧 How It Works

1. **Zeek** captures network traffic and generates security logs.
2. **Fluent Bit** forwards Zeek logs via **TCP port 5050** to a **Go application**.
3. The **Go application**:
   - Periodically fetches MISP threat data every 5 minutes.
   - Stores and updates IOC (IP, domain, URL, hash...) in **Redis**.
   - Compares incoming logs with MISP data.
   - If a match is found, the system:
     - Displays an alert in the console.
     - Sends the alert to **OpenSearch** for indexing and dashboard visualization.
4. Everything is containerized using **Docker** for portability.

---

##  Features

-  Real-time detection of malicious indicators in Zeek logs.
-  IOC comparison using MISP threat intelligence.
-  No duplicate API fetching — only new IOC data is pulled.
-  Dashboard-ready logs for OpenSearch.
-  Containerized Go app with support for `.env` configuration.

---

## 🐳 Docker Instructions

```bash
creat to path/to/alert.log
docker build -t misp-log-app .
docker run -d \
  --name misp-container \
  --network host \
  -v "$HOME/path/alert.log:/alert.log" \
  -v /etc/localtime:/etc/localtime:ro \
  -v $(pwd)/.env:/app/.env \
  misp-log-app
