PREPARE
=========

Role that prepare environment for kubernetes components installation

Disables swap, loads kernel modules and sets sysctl parameters

Requirements
------------

ansible.posix collection

Example Playbook
----------------

```yaml

- hosts: all
  roles:
    - prepare
```
