# How to install Zeek.
https://docs.zeek.org/en/current/install.html
-----------
# Configuration zeek. 
default configuration file is located at /opt/zeek/etc/networks.cfg. You can edit it using the nano editor.
```bash
nano /opt/zeek/etc/networks.cfg
```
Add your internal network as shown below:
```bash
10.0.0.0/24
```
edit the Zeek node.cfg configuration file.
```bash
nano /opt/zeek/etc/node.cfg
```
```bash
[zeek]
type=standalone
host=localhost
interface=enp0s3
LogDir = /opt/zeek/logs
SyslogOutput = true
```
Save and close the file. 

apply the above configurations using the following command.
```bash
zeekctl deploy
```
By default, Zeek stores all log files at /opt/zeek/logs/current/.

To see Zeek log files, run the following command.
```bash
ls -l /opt/zeek/logs/current/
```
