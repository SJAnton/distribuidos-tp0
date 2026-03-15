# TP0: Docker + Comunicaciones + Concurrencia

## Ejercicio 1
### Descripción
La resolución de este ejercicio consiste en un script de Bash que genera un archivo con un número de clientes configurable.

El script funciona escribiendo en el archivo las características del contenedor server primero, después las de los contenedores client un número de veces n, y finalmente lo relacionado a la red.

Los contenedores client siguen el formato client1, client2, client3, etc.

### Ejecución
El script recibe por parámetro el nombre del archivo de salida y el número de clientes a generar.

`./generar-compose.sh <nombre_archivo> <n_clientes>`

### Ejemplo

Para el caso de 5 clientes, se ejecuta el comando

`./generar-compose.sh docker-compose-dev.yaml 5`

Y el respectivo archivo generado contiene lo siguiente:

```
name: tp0
services:
  server:
    container_name: server
    image: server:latest
    entrypoint: python3 /main.py
    environment:
      - PYTHONUNBUFFERED=1
      - LOGGING_LEVEL=DEBUG
    networks:
      - testing_net

  client1:
    container_name: client1
    image: client:latest
    entrypoint: /client
    environment:
      - CLI_ID=1
      - CLI_LOG_LEVEL=DEBUG
    networks:
      - testing_net
    depends_on:
      - server

  client2:
    container_name: client2
    image: client:latest
    entrypoint: /client
    environment:
      - CLI_ID=2
      - CLI_LOG_LEVEL=DEBUG
    networks:
      - testing_net
    depends_on:
      - server

  client3:
    container_name: client3
    image: client:latest
    entrypoint: /client
    environment:
      - CLI_ID=3
      - CLI_LOG_LEVEL=DEBUG
    networks:
      - testing_net
    depends_on:
      - server

  client4:
    container_name: client4
    image: client:latest
    entrypoint: /client
    environment:
      - CLI_ID=4
      - CLI_LOG_LEVEL=DEBUG
    networks:
      - testing_net
    depends_on:
      - server

  client5:
    container_name: client5
    image: client:latest
    entrypoint: /client
    environment:
      - CLI_ID=5
      - CLI_LOG_LEVEL=DEBUG
    networks:
      - testing_net
    depends_on:
      - server

networks:
  testing_net:
    ipam:
      driver: default
      config:
        - subnet: 172.25.125.0/24
```
