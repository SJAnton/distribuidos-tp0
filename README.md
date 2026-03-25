# TP0: Docker + Comunicaciones + Concurrencia

## Ejercicio 8
### Descripción
En este ejercicio se modificó el servidor para aceptar conexiones de forma paralela, mediante el uso de threads. La sincronización se lleva a cabo por medio de un mutex, que bloquea el acceso a los recursos compartidos y a las funciones para cargar y guardar apuestas, que no son thread-safe.

Cada conexión se lanza como un nuevo thread de la siguiente forma:

```
threading.Thread(
    target=self.__handle_client_connection,
    args=(client_sock,)
).start()
```

Cuando el thread debe acceder a recursos compartidos, se utiliza el mutex para bloquear al resto de threads hasta que el thread que tiene el mutex termine su operación, por ejemplo:

```
with self._lock:
    self._done_agencies.add(int(agency_id))
    if self._done_agencies == self._known_agencies:
        self._done_draw = True
        logging.info("action: sorteo | result: success")
```

Tanto las funciones store_bets como load_bets, al no ser thread-safe, son protegidas mediante el mutex (la función load_bets es llamada dentro de get_winners).

```
with self._lock:
    if not self._done_draw:
        comms.send_message(client_sock, "WAIT")
        return
    winners = utils.get_winners(int(agency_id))
```

```
with self._lock:
    utils.store_bets(bets)
    self._known_agencies.add(int(agency_id))
```

El uso de un mutex garantiza la exclusión mutua y evita la aparición de race conditions.

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
client3  | 2026-03-25 22:13:49 INFO     action: config | result: success | client_id: 3 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client3  | 2026-03-25 22:13:50 INFO     action: apuesta_enviada | result: success
client4  | 2026-03-25 22:13:50 INFO     action: config | result: success | client_id: 4 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client4  | 2026-03-25 22:13:50 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:13:49 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:49 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-25 22:13:49 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:49 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:13:49 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:13:50 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-25 22:13:50 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:50 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:13:50 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:13:50 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-25 22:13:50 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:50 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-25 22:13:50 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:50 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:13:50 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:13:50 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:13:50 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:13:50 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-25 22:13:50 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:50 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:13:50 INFO     action: close_client_socket | result: success
client5  | 2026-03-25 22:13:49 INFO     action: config | result: success | client_id: 5 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client5  | 2026-03-25 22:13:49 INFO     action: apuesta_enviada | result: success
client1  | 2026-03-25 22:13:50 INFO     action: config | result: success | client_id: 1 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client1  | 2026-03-25 22:13:50 INFO     action: apuesta_enviada | result: success
client2  | 2026-03-25 22:13:50 INFO     action: config | result: success | client_id: 2 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client2  | 2026-03-25 22:13:50 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:13:55 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-25 22:13:55 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:55 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:13:55 INFO     action: close_client_socket | result: success
client5  | 2026-03-25 22:13:55 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:13:55 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-25 22:13:55 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:55 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:13:55 INFO     action: close_client_socket | result: success
client3  | 2026-03-25 22:13:55 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:13:55 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-25 22:13:55 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:55 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-25 22:13:55 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:55 INFO     action: apuesta_recibida | result: success | cantidad: 150
client1  | 2026-03-25 22:13:55 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:13:55 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:13:55 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:13:55 INFO     action: close_client_socket | result: success
client4  | 2026-03-25 22:13:55 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:13:55 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-25 22:13:55 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:13:55 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:13:55 INFO     action: close_client_socket | result: success
client2  | 2026-03-25 22:13:55 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:14:00 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-25 22:14:00 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:00 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:14:00 INFO     action: close_client_socket | result: success
client5  | 2026-03-25 22:14:00 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:14:00 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-25 22:14:00 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:00 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:14:00 INFO     action: close_client_socket | result: success
client3  | 2026-03-25 22:14:00 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:14:00 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-25 22:14:00 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:00 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-25 22:14:00 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:14:00 INFO     action: accept_connections | result: in_progress
client4  | 2026-03-25 22:14:00 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:14:00 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:14:00 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:14:00 INFO     action: close_client_socket | result: success
client1  | 2026-03-25 22:14:00 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:14:00 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-25 22:14:00 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:00 INFO     action: apuesta_recibida | result: success | cantidad: 150
server   | 2026-03-25 22:14:00 INFO     action: close_client_socket | result: success
client2  | 2026-03-25 22:14:00 INFO     action: apuesta_enviada | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:05 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:05 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:05 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:05 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-25 22:14:05 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-25 22:14:05 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-25 22:14:05 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:05 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-25 22:14:05 INFO     action: sorteo | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-25 22:14:05 INFO     action: close_client_socket | result: success
server   | 2026-03-25 22:14:05 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:05 INFO     action: close_client_socket | result: success
client2  | 2026-03-25 22:14:05 INFO     action: consulta_ganadores | result: success | cant_ganadores: 0
client2  | 2026-03-25 22:14:05 INFO     action: loop_finished | result: success | client_id: 2
client2 exited with code 0
server   | 2026-03-25 22:14:10 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-25 22:14:10 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:10 INFO     action: close_client_socket | result: success
client5  | 2026-03-25 22:14:10 INFO     action: consulta_ganadores | result: success | cant_ganadores: 0
client5  | 2026-03-25 22:14:10 INFO     action: loop_finished | result: success | client_id: 5
server   | 2026-03-25 22:14:10 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-25 22:14:10 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:10 INFO     action: close_client_socket | result: success
client3  | 2026-03-25 22:14:10 INFO     action: consulta_ganadores | result: success | cant_ganadores: 0
client3  | 2026-03-25 22:14:10 INFO     action: loop_finished | result: success | client_id: 3
server   | 2026-03-25 22:14:10 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-25 22:14:10 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:10 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-25 22:14:10 INFO     action: accept_connections | result: in_progress
server   | 2026-03-25 22:14:10 INFO     action: close_client_socket | result: success
client1  | 2026-03-25 22:14:10 INFO     action: consulta_ganadores | result: success | cant_ganadores: 0
client1  | 2026-03-25 22:14:10 INFO     action: loop_finished | result: success | client_id: 1
client4  | 2026-03-25 22:14:10 INFO     action: consulta_ganadores | result: success | cant_ganadores: 0
client4  | 2026-03-25 22:14:10 INFO     action: loop_finished | result: success | client_id: 4
server   | 2026-03-25 22:14:10 INFO     action: close_client_socket | result: success
client5 exited with code 0
client3 exited with code 0
client1 exited with code 0
client4 exited with code 0
```
