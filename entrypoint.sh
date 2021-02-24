#!/bin/sh

cp -a /tink-reboot /host/usr/sbin/tink-reboot
chroot /host <<EOT
/usr/sbin/tink-reboot -install
service tink-reboot stop
service tink-reboot start
EOT