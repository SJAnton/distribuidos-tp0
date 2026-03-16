# TP0: Docker + Comunicaciones + Concurrencia

## Ejercicio 2
### Descripción
Los cambios en el archivo de configuración se preservan mediante el uso de volumes. Para eso, se modificó el script que genera el archivo de configuración de Docker Compose, incluyendo ahora la declaración volume para el servidor y los clientes.

El resultado es que los archivos de configuración se montan dentro de los containers del servidor y los clientes, y no es necesario reconstruir la imagen al modificarlos.

### Ejecución
El script recibe por parámetro el nombre del archivo de salida y el número de clientes a generar.

`./generar-compose.sh <nombre_archivo> <n_clientes>`

### Ejemplo
El archivo de configuración YAML de Docker Compose queda de la siguiente manera, con la declaración volumes indicando la ruta de los archivos a preservar y la ruta del volume montado dentro del container.

```
name: tp0
services:
  server:
    container_name: server
    image: server:latest
    entrypoint: python3 /main.py
    environment:
      - PYTHONUNBUFFERED=1
    networks:
      - testing_net
    volumes:
      - ./server/config.ini:/config.ini

  client1:
    container_name: client1
    image: client:latest
    entrypoint: /client
    environment:
      - CLI_ID=1
    networks:
      - testing_net
    depends_on:
      - server
    volumes:
      - ./client/config.yaml:/config.yaml

networks:
  testing_net:
    ipam:
      driver: default
      config:
        - subnet: 172.25.125.0/24
```
