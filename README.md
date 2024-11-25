
# Kickoff

Servicio para revisar el disco y consultas lentas

## Paso 1
### Crear archivo de configuración, debe llamarse ```config.yml```, hay un archivo de ejemplo llamado ```template-config.yml```
```bash
godder:
  email:
    host: email.host
    port: 587
    from: godder-email@example.com
    password: password
    to: email@example.com
  disk:
    name: my-server
    disk_unit: GB
    alert_threshold: 1
  sql:
    query_unit: s
    slow_query_time: 180
    databases:
      - name: test
        host: test
        port: test
        user: test
        password: test
      - name: test2
        host: test2
        port: test2
        user: test2
        password: test2
      ...
```
