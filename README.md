# Nettica Admin

<h1><img src="./ui/src/assets/nettica.png" alt="A WireGuard control plane"></h1>

A control plane for [WireGuard](https://wireguard.com).

## Requirements

* OIDC compliant OAuth2 implementation
* MongoDB
* Mail Server credentials for sending outgoing email
* golang
* nginx
* NodeJS / Vue 2

![Screenshot](nettica-architecture.jpg)

## Features

 * Self-hosted and web-based management of WireGuard networks
 * Networks define the configuration of the hosts in the network
 * Invite people to network with email
 * Authenticate them with OAuth2
 * Generation of configuration files on demand
 * User authentication (Oauth2 OIDC)
 * Fully configure all aspects of your VPN
 * Manage hosts remotely
 * Simple
 * Lightweight
 * Secure



![Screenshot](nettica-screenshot.png)

## Running


### Directly

Install dependencies

Sample NGINX Config:

```
server {

        server_name netticavpn.com;

        root /usr/share/nettica-admin/ui/dist; index index.html; location / {
            try_files $uri $uri/ /index.html;
       }

    location /api/ {
        # app2 reverse proxy settings follow
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header Host localhost;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://127.0.0.1:8080;
    }


    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/netticavpn.com/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/netticavpn.com/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}
server {
    if ($host = netticavpn.com) {
        return 301 https://$host$request_uri;
    } # managed by Certbot



    server_name netticavpn.com;
    listen 80;
    return 404; # managed by Certbot


}
```

Example `.env` file:

```
# IP address to listen to
SERVER=0.0.0.0
# port to bind
PORT=8080
# Gin framework release mode
GIN_MODE=release

# SMTP settings to send email to clients
SMTP_HOST=smtp.sendgrid.net
SMTP_PORT=587
SMTP_USERNAME=apikey
SMTP_PASSWORD=...
SMTP_FROM=Nettica <info@netticavpn.com>

# MONGO settings
MONGODB_CONNECTION_STRING=mongodb://127.0.0.1:27017

# example with GitHub
#OAUTH2_PROVIDER_NAME=github
#OAUTH2_PROVIDER=https://github.com
#OAUTH2_CLIENT_ID=
#OAUTH2_CLIENT_SECRET=
#OAUTH2_REDIRECT_URL=

#OAUTH2_PROVIDER_NAME=oauth2oidc
#OAUTH2_PROVIDER=https://auth.netticavpn.com/
#OAUTH2_PROVIDER_URL=nettica.us.auth0.com
#OAUTH2_CLIENT_ID=...
#OAUTH2_CLIENT_SECRET=...
#OAUTH2_REDIRECT_URL=https://vpn.netticavpn.com

# Example settings for oauth2oidc-based Nettica Agent 
#OAUTH2_AGENT_PROVIDER=https://auth.nettica....
#OAUTH2_AGENT_PROVIDER_URL=
#OAUTH2_AGENT_CLIENT_ID=NativeAppClientId...
#OAUTH2_AGENT_CLIENT_SECRET=...
#OAUTH2_AGENT_AUDIENCE=guid or URL...
#OAUTH2_AGENT_REDIRECT_URL=com.nettica.agent://callback/agent
#OAUTH2_AGENT_LOGOUT_URL=https://auth.nettica..../v2/logout?client_id=NativeAppClientId...&returnTo=com.nettica.agent://callback/agent

#OAUTH2_PROVIDER_NAME=microsoft
#OAUTH2_PROVIDER=https://login.microsoftonline.com/.../v2.0
#OAUTH2_CLIENT_ID=
#OAUTH2_CLIENT_SECRET=
#OAUTH2_REDIRECT_URL=https://netticavpn.com
#OAUTH2_TENET=...

# OAuth2 provider using Microsoft's MSAL library.  Allows for a full range of Microsoft authentication
OAUTH2_PROVIDER_NAME=microsoft2
OAUTH2_PROVIDER=https://login.microsoftonline.com/common/v2.0
OAUTH2_CLIENT_ID=ApplicationID (guid)
OAUTH2_CLIENT_SECRET=...
OAUTH2_REDIRECT_URL=https://vpn.netticavpn.com
OAUTH2_TENET=... (guid)
OAUTH2_LOGOUT_URL=https://login.microsoftonline.com/{tenet guid}/oauth2/v2.0/logout

# Example Nettica Agent config for Microsoft MSAL
# When creating the App Registration in Azure Microsoft Entra ID, use "Add a Platform"
# to your Web App Registration you already created above, and choose "Mobile and Desktop Application",
# then add com.nettica.agent://callback/agent to the Redirect URLs

OAUTH2_AGENT_PROVIDER=https://login.microsoftonline.com/common/v2.0
OAUTH2_AGENT_CLIENT_ID=Application ID (guid - same as above)
OAUTH2_AGENT_CLIENT_SECRET=... (same as above)
OAUTH2_AGENT_REDIRECT_URL=com.nettica.agent://callback/agent
OAUTH2_AGENT_LOGOUT_URL=https://login.microsoftonline.com/{tenet guid}/oauth2/v2.0/logout




# valid settings: oauth2oidc, microsoft, microsoft2, basic, fake
# For google use microsoft provider
#OAUTH2_PROVIDER_NAME=microsoft2

# Basic auth requires no other parameters but OAUTH_PROVIDER_NAME
OAUTH2_PROVIDER_NAME=basic

```

Create a systemd service for the API:

```
cat  /lib/systemd/system/nettica-api.service
[Unit]
Description=Nettica API
ConditionPathExists=/usr/share/nettica-admin/cmd/nettica-api
After=network.target

[Service]
Type=simple
User=root
Group=root
LimitNOFILE=1024000

Restart=on-failure
RestartSec=10
#startLimitIntervalSec=60

WorkingDirectory=/usr/share/nettica-admin/
ExecStart=/usr/share/nettica-admin/cmd/nettica-api/nettica-api

# make sure log directory exists and owned by syslog
PermissionsStartOnly=true
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=nettica

[Install]
WantedBy=multi-user.target
```

Build the API
```
cd /usr/share/nettica-admin/cmd/nettica-api
go build
```

Enable the service:

```
sudo systemctl enable nettica-api
sudo systemctl start nettica-api
```

Install NodeJS using NVM
```
nvm use lts-latest
```

Build the frontend

```
cd ui
npm install
npm run build
```

With the given nginx config, you should now be able to use your website.  Don't forget
to get a cert using certbot

## Need Help

mailto:support@nettica.com

## License
* Released under MIT License

WireGuard® is a registered trademark of Jason A. Donenfeld.
