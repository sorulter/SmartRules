var direct = 'DIRECT';
var httpProxy = 'PROXY {{ .server }}; DIRECT';
var domains = {{.list}};

function FindProxyForURL(url, host) {
    if (url.substring(0, 4) == "ftp:") return direct;

    if (host.indexOf(".local", host.length - 6) !== -1) return direct;

    if (isPlainHostName(host)) return direct;

    if (shExpMatch(host, "*.cn")) return direct;

    var resolved_ip = dnsResolve(host);
    if (isInNet(resolved_ip, "10.0.0.0", "255.0.0.0") ||
        isInNet(resolved_ip, "172.16.0.0", "255.240.0.0") ||
        isInNet(resolved_ip, "192.168.0.0", "255.255.0.0") ||
        isInNet(resolved_ip, "127.0.0.0", "255.255.255.0"))
        return direct;

    var pos;
    do {
        if (domains.hasOwnProperty(host)) {
            return domains[host] ? httpProxy : direct;
        }
        pos = host.indexOf(".") + 1;
        host = host.slice(pos);
    } while (pos > 1)
    return direct;

}