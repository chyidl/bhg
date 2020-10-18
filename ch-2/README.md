Chapter 2 - TCP, SCANNERS, AND PROXIES 
======================================

* Understanding the TCP Handshake
```
the port is open:
  a three-way handshake takes place.
    1. First, the client sends a sync packet, which signals the beginning of a comunication 
    2. Second, the server then  responds with a syn-ack, or acknowledgment of the syn packet it received 
    3. Third, the client to finish with an ack, or acknowledgment of the server response. 

the port is closed:
  1. First, the client sends a sync packet, which signals the beginning of a communication 
  2. Second, the server responds with a rst packet instead of a sync-ack 

the port is Filtered Port:
  1. First, the client sends a sync packet, which signals the beginning of a communication 
  the client will typically receive no response from the server.
```

* Bypassing Firewalls with Port Forwarding
> You can use port forwarding to exploit several restrictive network configurations.
```
```
