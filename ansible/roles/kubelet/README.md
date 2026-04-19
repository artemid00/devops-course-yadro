Kubelet 
=========

Install kubelet from official Kubernetes repo

Requirements
------------

OS: Ubuntu 22.04/24.04
Root access: required


Role Variables
--------------

`k8s_version` - version of Kubernetes therefore kubelet (default "1.35")

Dependencies
------------

Container runtime (e.g. CRI-O)

Example Playbook
----------------

```yaml
- hosts: kubernetes_nodes
  become: true
  roles:
    - role: cri-o
    - role: kubelet
```
