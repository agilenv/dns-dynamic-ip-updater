# Linkip

Linkip is a simple application to keep updated a dns record where the associated IP is dynamic, can be used in a cron expression to schedule executions periodically.

Example
```
*/5 * * * * linkip sync --provider digitalocean --auto-update --env-file=/link/to/.env  
```

### Usage
List available dns providers
```
linkip list providers
```
Search for IP changes and update DNS record
```
linkip sync --provider <provider> --env-file=/link/to/.env
```
To avoid prompt for update confirmation
``` 
linkip sync --provider <provider> --auto-update --env-file=/link/to/.env
```
Show information about the last execution
```bash
linkip status
```

## Integrations
**Public IP API**
* [Ipify](https://www.ipify.org/)

**DNS Providers**
* DigitalOcean