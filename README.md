# Linkip

Linkip is a simple application to keep updated a dns record where the ip associated is dynamic. 
To achieve this, it provides commands to execute in a host computer and make the update to the DNS provider where the 
record has been created.

### Usage
Search for IP changes and update DNS record
```
linkip sync
```
To avoid prompt for update confirmation
```` 
linkip sync --update yes
````
Show information about the last execution
````bash
linkip status
````

## Integrations
**Public IP API**
* [Ipify](https://www.ipify.org/)

**DNS Providers**
* DigitalOcean