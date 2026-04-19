Role Name
=========
Install CRI-O

Requirements
------------

OS: Ubuntu 22.04/24.04
Root access: required

Role Variables
--------------

`crio_version` - version of cri-o (default "v1.35")

Dependencies
------------

Absent

Example Playbook
----------------

```yml
- hosts: all
  become: true
  roles:
    - role: cri-o
```

