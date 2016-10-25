### ProxyList — Print Usable Proxy Servers

> In computer networks, a proxy server is a server (a computer system or an application) that acts as an intermediary for requests from clients seeking resources from other servers. A client connects to the proxy server, requesting some service, such as a file, connection, web page, or other resource available from a different server and the proxy server evaluates the request as a way to simplify and control its complexity. Proxies were invented to add structure and encapsulation to distributed systems. Today, most proxies are web proxies, facilitating access to content on the World Wide Web and providing anonymity.
>
> — https://en.wikipedia.org/wiki/Proxy_server

### Installation

```shell
go get -u github.com/cixtor/proxylist
```

### Features

- [x] `-csv` - Export data using CSV format to standard output,
- [x] `-json` - Export data using JSON format to standard output,
- [x] `-table` - Export data using ASCII format to standard output,
- [x] `-sort=connection` - Sort data in descending order by the connection,
- [x] `-sort=anonymity` - Sort data in ascending order by the anonymity,
- [x] `-sort=protocol` - Sort data in ascending order by the protocol,
- [x] `-sort=speed` - Sort data in descending order by the speed,
- [x] `-sort=port` - Sort data in ascending order by the port,

![ProxyList](screenshot.png)
