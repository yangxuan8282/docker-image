#!/bin/bash

set -x

trap "stop; exit 0;" SIGTERM SIGINT

stop()
{
  # We're here because we've seen SIGTERM, likely via a Docker stop command or similar
  # Let's shutdown cleanly
  echo "SIGTERM caught, terminating NFS process(es)..."
  /usr/sbin/exportfs -ua
  pid1=$(pidof rpc.nfsd)
  pid2=$(pidof rpc.mountd)
  pid3=$(pidof rpc.statd)
  kill -TERM $pid1 $pid2 $pid3 > /dev/null 2>&1
  echo "Terminated."
  exit
}

mkdir -p /nfsshare
echo "/nfsshare   *(rw,fsid=0,insecure,no_root_squash,no_subtree_check,sync)" > /etc/exports


# Fixed nlockmgr port
echo 'fs.nfs.nlm_tcpport=32768' >> /etc/sysctl.conf
echo 'fs.nfs.nlm_udpport=32768' >> /etc/sysctl.conf
sysctl -p > /dev/null

mount -t nfsd nfds /proc/fs/nfsd

rpcbind -w
rpc.nfsd -N 2 -V 3 -N 4 -N 4.1 8
exportfs -arfv

rpc.statd -p 32765 -o 32766
rpc.mountd -N 2 -V 3 -N 4 -N 4.1 -p 32767 -F
