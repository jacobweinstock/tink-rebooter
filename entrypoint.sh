#!/bin/sh

cp -a /tink-reboot /host/tmp
chroot /host /bin/bash <<"EOT"
su root -c /tmp/tink-reboot
EOT