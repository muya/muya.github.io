---
layout: post
title: Extracting Safaricom MPESA B2C API Certificates
author: Fred Muya
---

When doing integration to the Safaricom MPESA B2C API, the most difficult (or rather least understood) part of the process is the SSL integration. Mutual authentication, certificates, etc.

This post will explain how to extract the individual certificates from the `.p7b` certificate file provided by Safaricom. This is useful, especially when you need all the certificates converted to PEM format

The `.p7b` certificate provided by Safaricom is a combination of several certificates (certificate chain). The instructions here will ensure you are able to extract all the certificates in the chain without a problem.

This post assumes you're using a *nix OS (because what else ;), and that you have openssl installed.

In this case, the file provided by Safaricom is called `safaricom-b2c-cert.p7b`

First, convert the cert from DER to PEM format

{% highlight bash %}
$ openssl pkcs7 -inform der -in safaricom-b2c-cert.p7b -out safaricom-b2c-cert-readable.p7b
{% endhighlight %}

The contents of `safaricom-b2c-cert-readable.p7b` should look something like this:
{% highlight bash %}
-----BEGIN PKCS7-----
a base64 encoded string
-----END PKCS7-----
{% endhighlight %}

Next, extract the certificates by running the following command:
{% highlight bash %}
$ openssl pkcs7 -print_certs  -in safaricom-b2c-cert-readable.p7b  -out safaricom-b2c-cert-readable.cer
{% endhighlight %}

The `safaricom-b2c-cert-readable.cer` should have several sections within it, looking something like:
{% highlight bash %}
subject=/DC=NET/DC=LABSAFCOM/CN=LABSAFCOM-PKIISSUE03-CA
issuer=/DC=NET/DC=LABSAFCOM/CN=LABSAFCOM-PKIPOLICY-CA
----BEGIN CERTIFICATE-----
a base64 encoded string
-----END CERTIFICATE-----
{% endhighlight %}

Each of the sections represent a specific certificate in the chain.

If you put each section in a separate file, then you can decode them separately to see their details using:
{% highlight bash %}
$ openssl x509 -text -in separate-file.cer  -noout
{% endhighlight %}

It's easier to use the `.cer` files in your PHP, Java or Python application.

I hope this helps someone.

If you need help doing integration to the Safaricom MPESA B2C API, feel free to [reach out](mailto:kingkonig@gmail.com){:target="_blank"}

You can check out my [LinkedIn profile](https://linkedin.com/in/fredmuya){:target="_blank"} to see previous integrations I've done.
