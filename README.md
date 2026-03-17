# TP0: Docker + Comunicaciones + Concurrencia

## Ejercicio 3
### Descripción
La solución consiste en un script que envía un "Hello world!" al echo server haciendo uso de la red utilizada por los containers server y client para comunicarse.

Este script crea un contenedor temporal que envía el mensaje mediante netcat y se elimina automáticamente.

### Ejecución
El script se ejecuta con el comando

`./validar-echo-server.sh`

La IP del servidor, el puerto y el nombre de la red pueden editarse, por defecto están como

```
server_ip=server
server_port=12345
network=tp0_testing_net
```

### Ejemplo
Se verifica el funcionamiento del servidor mediante el script creado. Si tiene éxito, se imprimirá el mensaje correspondiente.

```
./validar-echo-server.sh
action: test_echo_server | result: success
```

Si no fuera posible conectarse, el script imprimirá que falló al conectarse.

```
./validar-echo-server.sh
action: test_echo_server | result: fail
```
