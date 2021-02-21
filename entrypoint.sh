#!/bin/sh

cp -a /tink-reboot /host/tmp
chroot /host <<"EOT"
su root -c /tmp/tink-reboot
EOT