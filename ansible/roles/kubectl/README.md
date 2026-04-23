Kubectl
=========

Install Kubectl from official Kubernetes repo

Requirements
------------
OS: Ubuntu 22.04/24.04
Root access: required

Role Variables
--------------

`k8s_version` (string) - version of Kubernetes therefore kubeadm (default "1.35")

Dependencies
------------

Absent

Example Playbook
----------------

```yaml
- hosts: masters
  become: true
  roles:
    - role: kubectl
```
