Get the pb7 obtained from safaricom
currently file has weird characters in it

openssl pkcs7 -inform der -in MAMAMIKE_vpn.p7b -out MAMAMIKE_vpn_readable.pb7

then now you can convert the updated file to cer
openssl pkcs7 -print_certs -in MAMAMIKE_vpn_readable.p7b -out MAMAMIKE_vpn.cer

then debug
openssl s_client -connect 196.201.214.136:18423 -CAfile /etc/pki/tls/certs/B2C/testbroker.safaricom.com_X509_cert.crt -key /etc/pki/tls/certs/B2C/server.key -cert /etc/pki/tls/certs/B2C/MAMAMIKE_vpn.cer -debug

curl
