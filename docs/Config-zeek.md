# How to Install and Configure Zeek
## Installation zeek.
```bash
sudo apt-get update
sudo apt-get install -y --no-install-recommends g++ cmake make libpcap-dev
```
Add the Zeek software source.
```bash
sudo apt-get update
```
```bash
echo 'deb https://download.opensuse.org/repositories/security:/zeek/xUbuntu_22.04/ /' | sudo tee /etc/apt/sources.list.d/security:zeek.list
curl -fsSL https://download.opensuse.org/repositories/security:zeek/xUbuntu_22.04/Release.key | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/security_zeek.gpg > /dev/null
sudo apt update
sudo apt install zeek
```
Add Zeek to PATH.
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
## Zeek Configuration 
Set network open networks.cfg file and add the IP range you are using.
```bash
nano /opt/zeek/etc/networks.cfg
```
Example.
```bash
10.0.0.0/8
```
Set node configuration open node.cfg file and specify the interface to use.
```bash
nano /opt/zeek/etc/node.cfg
```
Example.
```bash
#[zeek]
#type=standalone
#host=localhost
#interface=eth0 # or ens33 or whatever name suits you
```
Deploy Zeek with zeekctl.
```bash
zeekctl deploy
```
View the generated log.
`` /opt/zeek/logs/current ``
```bash
root@ubuntu-VirtualBox:/opt/zeek/logs/current# ls
capture_loss.log  conn.log  dhcp.log  dns.log  http.log  notice.log  ntp.log  ssl.log  stats.log  stderr.log  stdout.log  telemetry.log  weird.log
```
## Convert Log to JSON
Open the file local.zeek.
`` /opt/zeek/share/zeek/site/local.zeek ``
```bash
nano /opt/zeek/share/zeek/site/local.zeek 
```
Add the last line.
```bash
@load policy/tuning/json-logs
```
Re-deploy for the settings to take effect.
```bash
zeekctl deploy
```
After this, Zeek will save all .logs in JSON format.

## Log Directory
Zeek will store the logs in
```bash
/opt/zeek/logs/current/
```
Check Zeek status.
```bash
zeekctl status
```
View logs in real-time.
```bash
tail -f /opt/zeek/logs/current/conn.log
```
