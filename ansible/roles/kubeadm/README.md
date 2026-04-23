Kubeadm
=========

Install Kubeadm from official Kubernetes repo

Requirements
------------
OS: Ubuntu 22.04/24.04
Root access: required

Role Variables
--------------

`k8s_version` (string) - version of Kubernetes therefore kubeadm (default "1.35")

Dependencies
------------

1. kubelet or/and kubectl
2. CRI (e.g. CRI-O) for kubelet

Example Playbook
----------------

```yaml
- hosts: all
  become: true
  roles:
    - role: cri-o
    - role: kubelet
    - role: kubeadm
```
