# TP0: Docker + Comunicaciones + Concurrencia

## Ejercicio 7
### Descripción
En este ejercicio se implementa el sorteo de la lotería, por medio del envío de mensajes DONE y RESULTS de los clientes al servidor.

El mensaje DONE está compuesto por el número de la agencia y la palabra DONE, separado por un delimitador "|". 

`[agencia|DONE]`

El servidor registra a las agencias activas a medida que recibe sus mensajes en batches, por lo que cuando recibe DONE de todas, hace el sorteo y determina a los ganadores, escribiendo en el log el mensaje

`action: sorteo | result: success`

Para recibir a los ganadores, su agencia correspondiente  envía un mensaje RESULTS en el mismo formato que DONE, este es

`[agencia|RESULTS]`

Tras recibir el mensaje RESULTS, el servidor devuelve WAIT si todavía no recibió el mensaje DONE de todas las agencias, o los ganadores de la lotería correspondiente a esas agencias si recibió los DONE, según el formato

`[documento_1|documento_2| ... |documento_n]`

Tras lo cual el cliente imprime el mensaje de éxito

`action: consulta_ganadores | result: success | cant_ganadores: ${CANT}`

O el correspondiente mensaje de error si no fue posible.

### Ejecución
Se crean n clientes genéricos corriendo el script

`./generar-compose.sh <archivo> <n_clientes>`

Se inicia el sistema mediante

`make docker-compose-up`

Se abren los logs usando

`make docker-compose-logs`

### Ejemplo
En el siguiente ejemplo se usan 5 clientes (agencias), y se modificó el código para finalizar con la tercera iteración, de forma tal que se vea el resultado del sorteo.

```
client5  | 2026-03-24 21:24:37 INFO     action: config | result: success | client_id: 5 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client5  | 2026-03-24 21:24:37 INFO     action: apuesta_enviada | result: success
client2  | 2026-03-24 21:24:37 INFO     action: config | result: success | client_id: 2 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client2  | 2026-03-24 21:24:37 INFO     action: apuesta_enviada | result: success
client1  | 2026-03-24 21:24:36 INFO     action: config | result: success | client_id: 1 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client1  | 2026-03-24 21:24:37 INFO     action: apuesta_enviada | result: success
client4  | 2026-03-24 21:24:37 INFO     action: config | result: success | client_id: 4 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client4  | 2026-03-24 21:24:37 INFO     action: apuesta_enviada | result: success
client3  | 2026-03-24 21:24:37 INFO     action: config | result: success | client_id: 3 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client3  | 2026-03-24 21:24:37 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:36 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:37 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-24 21:24:37 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:37 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:37 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:37 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-24 21:24:37 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:37 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:37 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:37 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-24 21:24:37 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:37 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:37 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:37 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-24 21:24:37 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:37 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:37 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:37 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-24 21:24:37 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:37 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:37 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:42 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-24 21:24:42 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:42 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:42 INFO     action: accept_connections | result: in_progress
client1  | 2026-03-24 21:24:42 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:42 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-24 21:24:42 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:42 INFO     action: close_client_socket | result: success
client3  | 2026-03-24 21:24:42 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:42 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:42 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-24 21:24:42 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:42 INFO     action: close_client_socket | result: success
client5  | 2026-03-24 21:24:42 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:42 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:42 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-24 21:24:42 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:42 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:42 INFO     action: accept_connections | result: in_progress
client2  | 2026-03-24 21:24:42 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:42 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-24 21:24:42 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:42 INFO     action: close_client_socket | result: success
client4  | 2026-03-24 21:24:42 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:42 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:47 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-24 21:24:47 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:47 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:47 INFO     action: accept_connections | result: in_progress
client1  | 2026-03-24 21:24:47 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:47 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-24 21:24:47 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:47 INFO     action: close_client_socket | result: success
client3  | 2026-03-24 21:24:47 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:47 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:47 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-24 21:24:47 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:47 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:47 INFO     action: accept_connections | result: in_progress
client5  | 2026-03-24 21:24:47 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:47 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-24 21:24:47 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:47 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:47 INFO     action: accept_connections | result: in_progress
client2  | 2026-03-24 21:24:47 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:47 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-24 21:24:47 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-24 21:24:47 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:47 INFO     action: accept_connections | result: in_progress
client4  | 2026-03-24 21:24:47 INFO     action: apuesta_enviada | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-24 21:24:52 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-24 21:24:52 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-24 21:24:52 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-24 21:24:52 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-24 21:24:52 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-24 21:24:52 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-24 21:24:52 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-24 21:24:52 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-24 21:24:52 INFO     action: sorteo | result: success
server   | 2026-03-24 21:24:52 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-24 21:24:52 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:52 INFO     action: accept_connections | result: in_progress
client4  | 2026-03-24 21:24:52 INFO     action: consulta_ganadores | result: success | cant_ganadores: 0
client4  | 2026-03-24 21:24:52 INFO     action: loop_finished | result: success | client_id: 4
client4 exited with code 0
server   | 2026-03-24 21:24:57 INFO     action: accept_connections | result: success | ip: 172.25.125.3
client1  | 2026-03-24 21:24:57 INFO     action: consulta_ganadores | result: success | cant_ganadores: 0
client1  | 2026-03-24 21:24:57 INFO     action: loop_finished | result: success | client_id: 1
server   | 2026-03-24 21:24:57 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:57 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:57 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-24 21:24:57 INFO     action: close_client_socket | result: success
client3  | 2026-03-24 21:24:57 INFO     action: consulta_ganadores | result: success | cant_ganadores: 0
client3  | 2026-03-24 21:24:57 INFO     action: loop_finished | result: success | client_id: 3
server   | 2026-03-24 21:24:57 INFO     action: accept_connections | result: in_progress
server   | 2026-03-24 21:24:57 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-24 21:24:57 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:57 INFO     action: accept_connections | result: in_progress
client5  | 2026-03-24 21:24:57 INFO     action: consulta_ganadores | result: success | cant_ganadores: 0
client5  | 2026-03-24 21:24:57 INFO     action: loop_finished | result: success | client_id: 5
server   | 2026-03-24 21:24:57 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-24 21:24:57 INFO     action: close_client_socket | result: success
server   | 2026-03-24 21:24:57 INFO     action: accept_connections | result: in_progress
client2  | 2026-03-24 21:24:57 INFO     action: consulta_ganadores | result: success | cant_ganadores: 0
client2  | 2026-03-24 21:24:57 INFO     action: loop_finished | result: success | client_id: 2
client1 exited with code 0
client5 exited with code 0
client3 exited with code 0
client2 exited with code 0
```
