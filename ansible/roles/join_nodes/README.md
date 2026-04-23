Join node
=========

A role that join node in cluster by join command from controle plane node. 
Checks if node already in cluster

Dependencies
------------

Need cluster and join command from controle plane node:

1. init_master

Example Playbook
----------------
```yml
    - hosts: masters
      roles:
        - init_master

    - hosts: all
      roles:
         - join_node
```
