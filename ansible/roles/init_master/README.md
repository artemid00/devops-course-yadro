Init Kubeadm
=========

Init controle plane role via kubeadm (generate join command and copy admin conf in /root/.kube dir)

Dependencies
------------

1. Prepare
2. kubelet
3. kubeadm

Example Playbook
----------------

```yaml
- hosts: masters
  roles:
    - prepare
    - kubelet
    - kubeadm
    - init_kubeadm

```
