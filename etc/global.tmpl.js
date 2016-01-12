var direct = 'DIRECT';
var httpProxy = 'PROXY {{ .server }}; DIRECT';

function FindProxyForURL(url, host) {
    if (url.substring(0, 4) == "ftp:")
        return direct;
    if (shExpMatch(url, "*.local"))
        return direct;

    if (shExpMatch(host, "*.cn")) {
        return "DIRECT";
    };

    // Private domain, send direct.
    if (isInNet(dnsResolve(host), "10.0.0.0", "255.0.0.0") ||
        isInNet(dnsResolve(host), "172.16.0.0", "255.240.0.0") ||
        isInNet(dnsResolve(host), "192.168.0.0", "255.255.0.0") ||
        isInNet(dnsResolve(host), "127.0.0.0", "255.255.255.0"))
        return direct;

    return httpProxy;
}