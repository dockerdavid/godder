
# Kickoff

Servicio para revisar el disco y consultas lentas

## Paso 1
### Crear la carpeta desde donde se va a llamar el archivo
```bash
cd /opt
mkdir godder
cd godder
```

## Paso 2
### Clonar el release
```bash
wget https://github.com/dockerdavid/godder/releases/download/${RELEASE_VERSION}/${RELEASE_TAG}
```

## Paso 3
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
    disk_unit: GB
    alert_threshold: 1
  sql:
    query_unit: s
    slow_query_time: 100
    databases:
      - name: test
        type: test
        host: test
        port: test
        user: test
        password: test
        slow_query_time: 10

```

## Paso 4
### Ejecutar la instalación
```bash
./godder -install
```

## Paso 5
### Verificar la instalación
```bash
systemctl status godder.service
```