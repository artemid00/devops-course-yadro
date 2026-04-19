from testinfra.host import Host

def test_crio_package(host: Host):
    pkg = host.package("cri-o")
    assert pkg.is_installed

def test_crio_service(host: Host):
    svc = host.service("cri-o")
    assert svc.is_running
    assert svc.is_enabled

def test_crio_socket(host: Host):
    assert host.file("/var/run/crio/crio.sock").is_socket
