# TP0: Docker + Comunicaciones + Concurrencia

## Ejercicio 6
### Descripción

Cada batch contiene varias apuestas serializadas en el mismo formato definido en el ejercicio anterior, separadas por el carácter \n. Es decir, el mensaje tiene la forma:

`[apuesta_1]\n[apuesta_2]\n...\n[apuesta_n]`

Donde cada apuesta tiene el formato:

`[agencia|nombre|apellido|documento|fecha_nacimiento|número]`

Los batches se generan a partir de la lectura de los archivos CSV en la carpeta .data, y su tamaño no supera los 8 KB (8192 bytes). Además, están también limitados por el campo maxAmount, definido en el archivo config.yaml del cliente. Es decir, el programa corta automáticamente un batch cuando llega a una de estas limitaciones, y comienza a generar el próximo.

El protocolo de comunicación es el mismo del ejercicio 5, ya que no se modificaron las funciones de envío y recepción de mensajes tanto del cliente como del servidor.

Si hay un fallo, el servidor envía FAIL al cliente y loguea el error y la cantidad de apuestas en el batch.

### Ejecución
Se crean n clientes genéricos corriendo el script

`./generar-compose.sh <archivo> <n_clientes>`

Se inicia el sistema mediante

`make docker-compose-up`

Se abren los logs usando

`make docker-compose-logs`

### Ejemplo
Para el ejemplo se usan 3 clientes genéricos y maxAmount = 150. Se verifica por medio del log que esta es la cantidad que recibe el servidor.

```
client1  | 2026-03-23 04:14:21 INFO     action: config | result: success | client_id: 1 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client1  | 2026-03-23 04:14:21 INFO     action: apuesta_enviada | result: success
server   | 2026-03-23 04:14:21 INFO     action: accept_connections | result: in_progress
server   | 2026-03-23 04:14:21 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-23 04:14:21 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-23 04:14:21 INFO     action: close_client_socket | result: success
server   | 2026-03-23 04:14:21 INFO     action: accept_connections | result: in_progress
server   | 2026-03-23 04:14:21 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-23 04:14:21 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-23 04:14:21 INFO     action: close_client_socket | result: success
server   | 2026-03-23 04:14:21 INFO     action: accept_connections | result: in_progress
server   | 2026-03-23 04:14:21 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-23 04:14:21 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-23 04:14:21 INFO     action: close_client_socket | result: success
server   | 2026-03-23 04:14:21 INFO     action: accept_connections | result: in_progress
client3  | 2026-03-23 04:14:21 INFO     action: config | result: success | client_id: 3 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client3  | 2026-03-23 04:14:21 INFO     action: apuesta_enviada | result: success
client2  | 2026-03-23 04:14:21 INFO     action: config | result: success | client_id: 2 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client2  | 2026-03-23 04:14:21 INFO     action: apuesta_enviada | result: success
server   | 2026-03-23 04:14:26 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-23 04:14:26 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-23 04:14:26 INFO     action: close_client_socket | result: success
server   | 2026-03-23 04:14:26 INFO     action: accept_connections | result: in_progress
client2  | 2026-03-23 04:14:26 INFO     action: apuesta_enviada | result: success
server   | 2026-03-23 04:14:26 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-23 04:14:26 INFO     action: apuesta_recibida | result: success | cantidad: 150
client3  | 2026-03-23 04:14:26 INFO     action: apuesta_enviada | result: success
server   | 2026-03-23 04:14:26 INFO     action: close_client_socket | result: success
server   | 2026-03-23 04:14:26 INFO     action: accept_connections | result: in_progress
server   | 2026-03-23 04:14:26 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-23 04:14:26 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-23 04:14:26 INFO     action: close_client_socket | result: success
server   | 2026-03-23 04:14:26 INFO     action: accept_connections | result: in_progress
...
```

En el caso en el que se usa un maxAmount alto, los batches crecen hasta llegar como mucho a 8 KB de tamaño. Con un cliente y maxAmount = 1000:

```
client1  | 2026-03-23 04:18:10 INFO     action: config | result: success | client_id: 1 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client1  | 2026-03-23 04:18:10 INFO     action: apuesta_enviada | result: success
server   | 2026-03-23 04:18:10 INFO     action: accept_connections | result: in_progress
server   | 2026-03-23 04:18:10 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-23 04:18:10 INFO     action: apuesta_recibida | result: success | cantidad: 170
server   | 2026-03-23 04:18:10 INFO     action: close_client_socket | result: success
server   | 2026-03-23 04:18:10 INFO     action: accept_connections | result: in_progress
server   | 2026-03-23 04:18:15 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-23 04:18:15 INFO     action: apuesta_recibida | result: success | cantidad: 177
server   | 2026-03-23 04:18:15 INFO     action: close_client_socket | result: success
server   | 2026-03-23 04:18:15 INFO     action: accept_connections | result: in_progress
client1  | 2026-03-23 04:18:15 INFO     action: apuesta_enviada | result: success
server   | 2026-03-23 04:18:20 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-23 04:18:20 INFO     action: apuesta_recibida | result: success | cantidad: 174
server   | 2026-03-23 04:18:20 INFO     action: close_client_socket | result: success
server   | 2026-03-23 04:18:20 INFO     action: accept_connections | result: in_progress
client1  | 2026-03-23 04:18:20 INFO     action: apuesta_enviada | result: success
...
```
