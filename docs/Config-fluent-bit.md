# How to Install and Configure Fluent Bit
# Installation
Install Fluent Bit on Ubuntu system using the following command.  
```bash
curl https://raw.githubusercontent.com/fluent/fluent-bit/master/install.sh | sh
```
Add Fluent Bit package source.
```bash
deb [signed-by=/usr/share/keyrings/fluentbit-keyring.gpg] https://packages.fluentbit.io/ubuntu/${CODENAME} ${CODENAME} main
```
Update packages and install.
```bash
sudo apt-get update
sudo apt-get install fluent-bit
sudo systemctl start fluent-bit
sudo systemctl status fluent-bit
```

----
# Configuration fluent-bit
The main Fluent Bit configuration file is located at.
`` /etc/fluent-bit/fluent-bit.conf ``
Example of setting to read log from Zeek and output as TCP.
```bash
[INPUT]
    name             tail
    Path             /opt/zeek/logs/current/*.log
    Tag              zeek*

[OUTPUT]
    Name            tcp
    Match           zeek*
    host            127.0.0.1
    Port            5050
    Format          json_lines
```
After editing the file, restart Fluent Bit.
```bash
sudo systemctl restart fluent-bit
```
```bash
sudo systemctl status fluent-bit
```
Verify that Fluent Bit is working and receiving the desired log.
```bash
sudo journalctl -u fluent-bit -f
```
