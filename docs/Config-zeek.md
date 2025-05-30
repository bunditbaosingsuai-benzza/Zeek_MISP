# How to Install zeek and config zeek
* # install zeek
```bash
apt-get update
```
```bash
apt-get install -y --no-install-recommends g++ cmake make libpcap-dev
```
```bash
sudo apt-get update
```
```bash
echo 'deb https://download.opensuse.org/repositories/security:/zeek/xUbuntu_22.04/ /' | sudo tee /etc/apt/sources.list.d/security:zeek.list
curl -fsSL https://download.opensuse.org/repositories/security:zeek/xUbuntu_22.04/Release.key | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/security_zeek.gpg > /dev/null
sudo apt update
sudo apt install zeek
```
```bash
echo "export PATH=$PATH:/opt/zeek/bin" >> ~/.bashrc
```
```bash
source ~/.bashrc
```
```bash
zeek --version
```
----
# Configuration zeek 
`` /opt/zeek/etc/networks.cfg. ``
```bash
nano /opt/zeek/etc/networks.cfg
```
```bash
10.0.0.0/8
```
```bash
nano /opt/zeek/etc/node.cfg
```
```bash
#[zeek]
#type=standalone
#host=localhost
#interface=< interface name >
```
```bash
zeekctl deploy
```
`` /opt/zeek/logs/current ``
```bash
root@ubuntu-VirtualBox:/opt/zeek/logs/current# ls
capture_loss.log  conn.log  dhcp.log  dns.log  http.log  notice.log  ntp.log  ssl.log  stats.log  stderr.log  stdout.log  telemetry.log  weird.log
```
* convert .log > json

`` /opt/zeek/share/zeek/site/local.zeek ``
```bash
nano /opt/zeek/share/zeek/site/local.zeek 
```
```bash
@load policy/tuning/json-logs
```
```bash
zeekctl deploy
```
