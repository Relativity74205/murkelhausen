```bash
ansible-playbook -i inventory.yaml playbook.yaml -K
```

Interesting options:
- `--check`
- `--verbose`

## Links

- Docker with ansible: https://www.digitalocean.com/community/tutorials/how-to-use-ansible-to-install-and-set-up-docker-on-ubuntu-20-04
- Ansible Roles: https://www.digitalocean.com/community/tutorials/how-to-use-ansible-roles-to-abstract-your-infrastructure-environment
- Docker container module: https://docs.ansible.com/ansible/2.9/modules/docker_container_module.html