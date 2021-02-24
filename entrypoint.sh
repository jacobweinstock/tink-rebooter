#!/bin/sh

cp -a /tink-reboot /host/usr/sbin/tink-reboot
chroot /host <<EOT
/usr/sbin/tink-reboot -install
/usr/sbin/tink-reboot -start
/usr/sbin/tink-reboot -run
EOT