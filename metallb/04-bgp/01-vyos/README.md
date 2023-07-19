# Configure VyOS
## Create instance
```
onp vyos NETWORK=onp2
```
## Configure network interfaces
```
configure
set interfaces ethernet eth0 address 192.168.125.10/24
commit
save
exit
```
## Configure BGP
```
configure
set protocols bgp 64512 parameters router-id 192.168.125.10
set protocols bgp 64512 neighbor 192.168.125.100 address-family ipv4-unicast
set protocols bgp 64512 neighbor 192.168.125.100 remote-as 64512
commit
save
exit
```
