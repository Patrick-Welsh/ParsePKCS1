# ParsePKCS1
Script to decode a PEM block from a generated PKCS1 RSA private key, and parse the key

## Before running this script, you will need to generate an RSA key in PKCS1 format, which can be done by running the following command. The script will only work correctly if the key is in this format.
  ssh-keygen -m PEM -t rsa -b 4096 -C "" -f $HOME/ParsePKCS1/rsa
  
## The way to determine if a key is in PKCS1 format is if you see the blocks 
  -----BEGIN RSA PRIVATE KEY-----
  -----END RSA PRIVATE KEY-----
  
 ## If your blocks look the the following, your key format will be in PKCS8
  -----BEGIN PRIVATE KEY-----
  -----END PRIVATE KEY-----
