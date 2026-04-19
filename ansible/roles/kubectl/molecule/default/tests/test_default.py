from testinfra.host import Host

def test_kubectl_package_installed(host: Host):
    pkg = host.package("kubectl")
    assert pkg.is_installed

def test_kubectl_binary_exists(host: Host):
    assert host.file("/usr/bin/kubectl").exists

def test_kubectl_client_version(host: Host):
    cmd = host.run("kubectl version --client")
    assert cmd.rc == 0

