# How to Install Redis
## Installation redis
Install the base package.
```bash
sudo apt-get install lsb-release curl gpg
```
Add the GPG Key and Repository.
```bash
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
sudo chmod 644 /usr/share/keyrings/redis-archive-keyring.gpg
```
Add Redis software source.
```bash
echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list
```
Install Redis.
```bash
sudo apt-get update
sudo apt-get install redis
```
## Enable & Start Redis.
Enable Redis to start automatically on boot.
```bash
sudo systemctl enable redis-server
```
Start Redis service.
```bash
sudo systemctl start redis-server
```
Check Redis status.
```bash
sudo systemctl status redis-server
```
### Test Redis usage open CLI of Redis.
```bash
redis-cli
```
```bash
PING
#After that, I will answer. >> "PONG"

```

