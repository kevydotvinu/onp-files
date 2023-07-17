# Configure VyOS
## Enable SSH
```
configure
set service ssh
commit
save
exit
```
## Enable network interfaces
```
configure
set interfaces ethernet eth0 address 192.168.124.10/24
set interfaces ethernet eth1 address 192.168.125.10/24
commit
save
exit
```
## Enable DHCP server
```
configure
set service dhcp-server shared-network-name LAN1 subnet 192.168.125.0/24 range 0 start 192.168.125.100
set service dhcp-server shared-network-name LAN1 subnet 192.168.125.0/24 range 0 stop 192.168.125.200
set service dhcp-server shared-network-name LAN1 subnet 192.168.125.0/24 default-router 192.168.125.10
set service dhcp-server shared-network-name LAN1 subnet 192.168.125.0/24 dns-server 192.168.125.10
set service dhcp-server shared-network-name LAN1 subnet 192.168.125.0/24 domain-name vyos.example.local
commit
save
exit
```
## Configure BGP
```
configure
set protocols bgp 64512 parameters router-id 192.168.125.10
set protocols bgp 64512 neighbor 192.168.126.100 remote-as 64512
commit
save
exit
```
