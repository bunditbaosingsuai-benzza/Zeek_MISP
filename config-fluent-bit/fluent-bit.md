# How to install Fluent-bit.
https://docs.fluentbit.io/manual/installation/linux/ubuntu
---------------------
# Configuration Fluent-bit
The default configuration file is written to:
```bash
/etc/fluent-bit/fluent-bit.conf
```
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

[INPUT]
    Name        tail
    Path        /home/zeeh-host/Desktop/*.log
    Tag         alert
    Parser      json

[OUTPUT]
    Name              es
    Match             alert*
    Host              localhost
    Port              9200
    Index             alert-index
    HTTP_User         <Your ID>
    HTTP_Passwd       <Your Pasword>
    tls               On
    tls.verify        Off
    Suppress_Type_Name On
    Logstash_Format   On
    Logstash_Prefix   alert
    Time_Key          timestamp
    Time_Key_Format   %Y-%m-%dT%H:%M:%S%z
    Include_Tag_Key   On
    Tag_Key           _flb-key

```
