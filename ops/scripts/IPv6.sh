#!/bin/sh
set -ex

## In the Travis VM-based build environment, IPv6 networking is not
## enabled by default. The sysctl operations below enable IPv6.
## IPv6 is needed by some of the CoreDNS test cases.
cat /proc/net/if_inet6
sudo bash -c 'if [ `cat /proc/net/if_inet6 | wc -l` = "0" ]; then echo "Enabling IPv6" ; sysctl net.ipv6.conf.all.disable_ipv6=0 ; sysctl net.ipv6.conf.default.disable_ipv6=0 ; sysctl net.ipv6.conf.lo.disable_ipv6=0 ; fi'
cat /proc/net/if_inet6
