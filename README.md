# tink-rebooter
Reboot Action for Tinkerbell

```yaml
actions:
- name: "reboot"
  image: ghcr.io/jacobweinstock/tink-rebooter:v0.3.1
  timeout: 90
  pid: host
  volumes:
  - /:/host
```
