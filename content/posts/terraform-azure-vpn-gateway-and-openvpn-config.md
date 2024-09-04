---
author: gripdev
category:
  - uncategorized
date: "2020-07-20T18:41:36+00:00"
guid: http://blog.gripdev.xyz/?p=1332
title: Terraform, Azure VPN Gateway and OpenVPN Config
url: /2020/07/20/terraform-azure-vpn-gateway-and-openvpn-config/

---
21 June 2021: Updated to resolve bug where wrong private key was passed to client\_certs. This was exposed by service change.

There is a good guide to generating the necessary certificates and manually editing the `openvpn` config you can download from the portal [in the official docs.](https://docs.microsoft.com/en-us/azure/vpn-gateway/vpn-gateway-certificates-point-to-site-linux)

Being a sucker for punishment I wondered if I could automate the process (mainly because I always forget the openssl commands) and secondly it was to be run by someone else (not me) so I wanted it to be as simple as possible.

Warning: Please take time to understand the limitations of this certificate generation before production use (TF State file containing CA private keys) and review the limitations in the TLS provider for Terraform.

So how would this work? First choice is Terraform for the automation. Luckily that [has a provider for generating certs!](https://registry.terraform.io/providers/hashicorp/tls/latest/docs) So the flow goes like this:

1\. Create the Root CA   
2\. Use that to generate a client cert  
3\. Output the client cert for use in the openvpn config file  
4\. Inject the CA into the Azure VPN configuration and create it  
5\. Run a script to fetch the Azure VPN OpenVPN configuration file (as this contains the Pre-shared key we don't set) then inject the client cert we outputted from the Terraform.   
6\. Connect

One gotcha - The TLS provider only outputs standard pems. Azure VPN requires the CA to be in a specific format outputted by the following command.

`openssl x509 -in caCert.pem -outform der | base64 -w0 > caCert.der`

As a result there is a `null_resource` and `local_file` resource to handle this translation.

So all you should need is Terraform, Bash, OpenSSL and Azure CLI - not perfect but it's the best I could do! (This is currently a working draft - please use with caution).

https://gist.github.com/lawrencegripper/294ae913feff17315f7bb08d382a577d
