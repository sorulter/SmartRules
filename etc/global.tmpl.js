var direct = 'DIRECT';
var httpProxy = 'PROXY {{ .server }}; DIRECT';

function FindProxyForURL(url, host) {
    if (url.substring(0, 4) == "ftp:")
        return direct;
    if (host.indexOf(".local", host.length - 6) !== -1) {
        return direct;
    }

    return httpProxy;
}