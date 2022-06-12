#!/bin/sh
NAMES="Guillermo Sandra Ada"


for current in $HOSTS; do
cat | ssh $current <<-EOF
for current in $NAMES ; do 
echo Hello $current >> /etc/motd
done
EOF


done