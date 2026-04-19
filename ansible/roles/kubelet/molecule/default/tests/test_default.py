from testinfra.host import Host

def test_kubelet_package_status(host: Host):
    pkg = host.package("kubelet")
    assert pkg.is_installed
    
def test_kubelet_service(host: Host):
    svc = host.service("kubelet")
    assert svc.is_enabled
