# How to Install and config fluent-bit 
* # install
```bash
curl https://raw.githubusercontent.com/fluent/fluent-bit/master/install.sh | sh
```
```bash
deb [signed-by=/usr/share/keyrings/fluentbit-keyring.gpg] https://packages.fluentbit.io/ubuntu/${CODENAME} ${CODENAME} main
```
```bash
sudo apt-get update
```
```bash
sudo apt-get install fluent-bit
```
```bash
sudo systemctl start fluent-bit
```
```bash
sudo systemctl status fluent-bit
```
----
# Configuration fluent-bit
`` /etc/fluent-bit/fluent-bit.conf ``
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
```bash
sudo systemctl restart fluent-bit
```
```bash
sudo systemctl status fluent-bit
```
