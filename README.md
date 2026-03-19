# TP0: Docker + Comunicaciones + Concurrencia

## Ejercicio 4
### Descripción
Para este ejercicio se modificó el client y el server para terminar la aplicación de forma graceful.

En el cliente, se crea un canal stop al cual se le carga la señal SIGTERM si es recibida por el proceso, para posteriormente verificar si debe hacerse el shutdown.

En el server, se implementó la función shutdown que maneja el cierre, además de implementar mensajes de log para informar el cierre de los sockets y que el shutdown tuvo éxito.

### Ejecución
Se inicia el sistema mediante

`make docker-compose-up`

Se abren los logs usando

`make docker-compose-logs`

Finalmente, para probar el shutdown se escribe en una segunda terminal

`docker compose -f docker-compose-dev.yaml stop`

### Ejemplo
Mediante las funciones mostradas anteriormente, el log para el caso de 5 clientes es:

```
client1  | 2026-03-19 03:50:33 INFO     action: config | result: success | client_id: 1 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client1  | 2026-03-19 03:50:33 INFO     action: receive_message | result: success | client_id: 1 | msg: [CLIENT 1] Message N°1
client5  | 2026-03-19 03:50:33 INFO     action: config | result: success | client_id: 5 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client5  | 2026-03-19 03:50:33 INFO     action: receive_message | result: success | client_id: 5 | msg: [CLIENT 5] Message N°1
client2  | 2026-03-19 03:50:33 INFO     action: config | result: success | client_id: 2 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client2  | 2026-03-19 03:50:33 INFO     action: receive_message | result: success | client_id: 2 | msg: [CLIENT 2] Message N°1
client4  | 2026-03-19 03:50:33 INFO     action: config | result: success | client_id: 4 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client4  | 2026-03-19 03:50:33 INFO     action: receive_message | result: success | client_id: 4 | msg: [CLIENT 4] Message N°1
client3  | 2026-03-19 03:50:33 INFO     action: config | result: success | client_id: 3 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client3  | 2026-03-19 03:50:33 INFO     action: receive_message | result: success | client_id: 3 | msg: [CLIENT 3] Message N°1
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: in_progress
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-19 03:50:33 INFO     action: receive_message | result: success | ip: 172.25.125.3 | msg: [CLIENT 3] Message N°1
server   | 2026-03-19 03:50:33 INFO     action: close_client_socket | result: success
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: in_progress
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-19 03:50:33 INFO     action: receive_message | result: success | ip: 172.25.125.4 | msg: [CLIENT 2] Message N°1
server   | 2026-03-19 03:50:33 INFO     action: close_client_socket | result: success
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: in_progress
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-19 03:50:33 INFO     action: receive_message | result: success | ip: 172.25.125.5 | msg: [CLIENT 5] Message N°1
server   | 2026-03-19 03:50:33 INFO     action: close_client_socket | result: success
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: in_progress
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-19 03:50:33 INFO     action: receive_message | result: success | ip: 172.25.125.6 | msg: [CLIENT 1] Message N°1
server   | 2026-03-19 03:50:33 INFO     action: close_client_socket | result: success
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: in_progress
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-19 03:50:33 INFO     action: receive_message | result: success | ip: 172.25.125.7 | msg: [CLIENT 4] Message N°1
server   | 2026-03-19 03:50:33 INFO     action: close_client_socket | result: success
server   | 2026-03-19 03:50:33 INFO     action: accept_connections | result: in_progress
server   | 2026-03-19 03:50:38 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-19 03:50:38 INFO     action: receive_message | result: success | ip: 172.25.125.3 | msg: [CLIENT 3] Message N°2
server   | 2026-03-19 03:50:38 INFO     action: close_client_socket | result: success
client3  | 2026-03-19 03:50:38 INFO     action: receive_message | result: success | client_id: 3 | msg: [CLIENT 3] Message N°2
server   | 2026-03-19 03:50:38 INFO     action: accept_connections | result: in_progress
server   | 2026-03-19 03:50:38 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-19 03:50:38 INFO     action: receive_message | result: success | ip: 172.25.125.4 | msg: [CLIENT 2] Message N°2
server   | 2026-03-19 03:50:38 INFO     action: close_client_socket | result: success
server   | 2026-03-19 03:50:38 INFO     action: accept_connections | result: in_progress
client2  | 2026-03-19 03:50:38 INFO     action: receive_message | result: success | client_id: 2 | msg: [CLIENT 2] Message N°2
server   | 2026-03-19 03:50:38 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-19 03:50:38 INFO     action: receive_message | result: success | ip: 172.25.125.5 | msg: [CLIENT 5] Message N°2
server   | 2026-03-19 03:50:38 INFO     action: close_client_socket | result: success
client5  | 2026-03-19 03:50:38 INFO     action: receive_message | result: success | client_id: 5 | msg: [CLIENT 5] Message N°2
server   | 2026-03-19 03:50:38 INFO     action: accept_connections | result: in_progress
server   | 2026-03-19 03:50:38 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-19 03:50:38 INFO     action: receive_message | result: success | ip: 172.25.125.6 | msg: [CLIENT 1] Message N°2
server   | 2026-03-19 03:50:38 INFO     action: close_client_socket | result: success
server   | 2026-03-19 03:50:38 INFO     action: accept_connections | result: in_progress
client1  | 2026-03-19 03:50:38 INFO     action: receive_message | result: success | client_id: 1 | msg: [CLIENT 1] Message N°2
server   | 2026-03-19 03:50:38 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-19 03:50:38 INFO     action: receive_message | result: success | ip: 172.25.125.7 | msg: [CLIENT 4] Message N°2
server   | 2026-03-19 03:50:38 INFO     action: close_client_socket | result: success
server   | 2026-03-19 03:50:38 INFO     action: accept_connections | result: in_progress
client4  | 2026-03-19 03:50:38 INFO     action: receive_message | result: success | client_id: 4 | msg: [CLIENT 4] Message N°2
client3  | 2026-03-19 03:50:43 INFO     action: shutdown | result: success
client2  | 2026-03-19 03:50:43 INFO     action: shutdown | result: success
client5  | 2026-03-19 03:50:43 INFO     action: shutdown | result: success
client1  | 2026-03-19 03:50:43 INFO     action: shutdown | result: success
client4  | 2026-03-19 03:50:43 INFO     action: shutdown | result: success
client3 exited with code 0
client2 exited with code 0
client5 exited with code 0
client1 exited with code 0
client4 exited with code 0
server   | 2026-03-19 03:50:43 INFO     action: close_server_socket | result: success
server   | 2026-03-19 03:50:43 INFO     action: shutdown | result: success
server exited with code 0
```
