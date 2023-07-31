# Configure VyOS
## Create instance
```
onp vyos NETWORK=onp1
```
## Configure network interfaces
```
configure
set interfaces ethernet eth0 address 192.168.124.2/24
set service ssh
commit
save
exit
```
## SSH to VyOS
```
ssh -o StrictHostKeyChecking=no vyos@192.168.124.2
```
## Configure BGP
```
configure
set protocols bgp 64512 parameters router-id 192.168.124.2
set protocols bgp 64512 neighbor 192.168.124.3 address-family ipv4-unicast
set protocols bgp 64512 neighbor 192.168.124.3 remote-as 64512
commit
save
exit
```
