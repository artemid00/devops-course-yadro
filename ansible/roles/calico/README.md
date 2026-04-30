Calico
=========

Install Calico CNI plagin via tigera operator on master node

Role Variables
--------------

`calico_version` (string) - version of calico cni plugin (default: "3.31.5")

Dependencies
------------

1. Kubectl

Example Playbook
----------------

```yml
- hosts: masters
  roles: 
    - kubectl
    - calico
```
