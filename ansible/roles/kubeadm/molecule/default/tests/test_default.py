from testinfra.host import Host

def test_kubeadm_package_installed(host: Host):
    pkg = host.package("kubeadm")
    assert pkg.is_installed

def test_kubeadm_binary_exists(host: Host):
    assert host.file("/usr/bin/kubeadm").exists

def test_kubeadm_version_works(host: Host):
    cmd = host.run("kubeadm version")
    assert cmd.rc == 0

