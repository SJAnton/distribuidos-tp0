# TP0: Docker + Comunicaciones + Concurrencia

## Ejercicio 5
### Descripción
El protocolo de transmisión implementado en este ejercicio es el TCP, ya que garantiza la entrega confiable y ordenada de datos entre los clientes y el servidor.

El envío de mensajes en ambos lados consiste en enviar un mensaje de 2 bytes en formato big endian con el largo en bytes del mensaje principal, para después enviar el mensaje principal y manejar correctamente los short-reads, ya que el receptor sabe cuánto debe recibir. A la vez, el mensaje principal se envía por medio de un ciclo que no finaliza hasta que se haya enviado en su totalidad o sucedido un error, por lo que los short-writes son manejados por la capa de comunicación.

Del lado del cliente, las variables de entorno son encapsuladas por un objeto Bet, que representa la capa de dominio

```
type Bet struct {
	Agency  int
	Name    string
	Surname string
	Id      int
	Dob     string
	Number  int
}
```

Donde la creación del mensaje se hace por medio de la función MakeMessage(), que usa el caracter "|" como delimitador, dando como resultado el formato:

`[agencia|nombre|apellido|documento|fecha_nacimiento|número]`

La capa de comunicación está compuesta por las funciones sendMessage y receiveMessage del cliente, y send_message y receive_message del server, en los archivos comms.go y comms.py respectivamente. Las funciones de envío de mensaje envían los mensajes serializados en bytes en formato UTF-8, los cuales se reconvierten en strings al llegar al receptor.

El cliente envía:

`[largo][mensaje]`

El servidor recibe el mensaje, lo procesa separando cada uno de los campos por el delimitador, y almacena la información con la función store_bet(). Luego, envía:

`[largo][señal]`

La señal puede ser OK o FAIL, dependiendo de si el servidor tuvo éxito al guardar la apuesta. En cualquiera de los dos casos se imprime en el log si fue un éxito o un fallo.

### Ejecución
Se crean n clientes genéricos corriendo el script

`./generar-compose.sh <archivo> <n_clientes>`

Se inicia el sistema mediante

`make docker-compose-up`

Se abren los logs usando

`make docker-compose-logs`

### Ejemplo
Para el ejemplo se usan 5 clientes genéricos y se verifica que del lado del cliente y el servidor se escriban en el log los mensajes esperados registrando que la apuesta fue enviada y almacenada respectivamente.

```
client1  | 2026-03-21 18:53:13 INFO     action: config | result: success | client_id: 1 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client1  | 2026-03-21 18:53:13 INFO     action: apuesta_enviada | result: success | dni: 30000001 | numero: 5001
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: in_progress
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-21 18:53:13 INFO     action: apuesta_almacenada | result: success | dni: 30000002 | numero: 5002
server   | 2026-03-21 18:53:13 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: in_progress
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-21 18:53:13 INFO     action: apuesta_almacenada | result: success | dni: 30000003 | numero: 5003
server   | 2026-03-21 18:53:13 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: in_progress
client3  | 2026-03-21 18:53:13 INFO     action: config | result: success | client_id: 3 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client3  | 2026-03-21 18:53:13 INFO     action: apuesta_enviada | result: success | dni: 30000003 | numero: 5003
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: success | ip: 172.25.125.5
client2  | 2026-03-21 18:53:13 INFO     action: config | result: success | client_id: 2 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client2  | 2026-03-21 18:53:13 INFO     action: apuesta_enviada | result: success | dni: 30000002 | numero: 5002
server   | 2026-03-21 18:53:13 INFO     action: apuesta_almacenada | result: success | dni: 30000001 | numero: 5001
server   | 2026-03-21 18:53:13 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: in_progress
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-21 18:53:13 INFO     action: apuesta_almacenada | result: success | dni: 30000005 | numero: 5005
server   | 2026-03-21 18:53:13 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: in_progress
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-21 18:53:13 INFO     action: apuesta_almacenada | result: success | dni: 30000004 | numero: 5004
server   | 2026-03-21 18:53:13 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:13 INFO     action: accept_connections | result: in_progress
client4  | 2026-03-21 18:53:13 INFO     action: config | result: success | client_id: 4 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client4  | 2026-03-21 18:53:13 INFO     action: apuesta_enviada | result: success | dni: 30000004 | numero: 5004
client5  | 2026-03-21 18:53:13 INFO     action: config | result: success | client_id: 5 | server_address: server:12345 | loop_amount: 5 | loop_period: 5s | log_level: INFO
client5  | 2026-03-21 18:53:13 INFO     action: apuesta_enviada | result: success | dni: 30000005 | numero: 5005
server   | 2026-03-21 18:53:18 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-21 18:53:18 INFO     action: apuesta_almacenada | result: success | dni: 30000002 | numero: 5002
server   | 2026-03-21 18:53:18 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:18 INFO     action: accept_connections | result: in_progress
client2  | 2026-03-21 18:53:18 INFO     action: apuesta_enviada | result: success | dni: 30000002 | numero: 5002
server   | 2026-03-21 18:53:18 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-21 18:53:18 INFO     action: apuesta_almacenada | result: success | dni: 30000003 | numero: 5003
server   | 2026-03-21 18:53:18 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:18 INFO     action: accept_connections | result: in_progress
client3  | 2026-03-21 18:53:18 INFO     action: apuesta_enviada | result: success | dni: 30000003 | numero: 5003
server   | 2026-03-21 18:53:18 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-21 18:53:18 INFO     action: apuesta_almacenada | result: success | dni: 30000001 | numero: 5001
server   | 2026-03-21 18:53:18 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:18 INFO     action: accept_connections | result: in_progress
client1  | 2026-03-21 18:53:18 INFO     action: apuesta_enviada | result: success | dni: 30000001 | numero: 5001
server   | 2026-03-21 18:53:18 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-21 18:53:18 INFO     action: apuesta_almacenada | result: success | dni: 30000005 | numero: 5005
client5  | 2026-03-21 18:53:18 INFO     action: apuesta_enviada | result: success | dni: 30000005 | numero: 5005
server   | 2026-03-21 18:53:18 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:18 INFO     action: accept_connections | result: in_progress
server   | 2026-03-21 18:53:18 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-21 18:53:18 INFO     action: apuesta_almacenada | result: success | dni: 30000004 | numero: 5004
server   | 2026-03-21 18:53:18 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:18 INFO     action: accept_connections | result: in_progress
client4  | 2026-03-21 18:53:18 INFO     action: apuesta_enviada | result: success | dni: 30000004 | numero: 5004
server   | 2026-03-21 18:53:23 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-21 18:53:23 INFO     action: apuesta_almacenada | result: success | dni: 30000002 | numero: 5002
server   | 2026-03-21 18:53:23 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:23 INFO     action: accept_connections | result: in_progress
client2  | 2026-03-21 18:53:23 INFO     action: apuesta_enviada | result: success | dni: 30000002 | numero: 5002
server   | 2026-03-21 18:53:23 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-21 18:53:23 INFO     action: apuesta_almacenada | result: success | dni: 30000003 | numero: 5003
server   | 2026-03-21 18:53:23 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:23 INFO     action: accept_connections | result: in_progress
client3  | 2026-03-21 18:53:23 INFO     action: apuesta_enviada | result: success | dni: 30000003 | numero: 5003
server   | 2026-03-21 18:53:23 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-21 18:53:23 INFO     action: apuesta_almacenada | result: success | dni: 30000001 | numero: 5001
server   | 2026-03-21 18:53:23 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:23 INFO     action: accept_connections | result: in_progress
client1  | 2026-03-21 18:53:23 INFO     action: apuesta_enviada | result: success | dni: 30000001 | numero: 5001
server   | 2026-03-21 18:53:23 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-21 18:53:23 INFO     action: apuesta_almacenada | result: success | dni: 30000005 | numero: 5005
server   | 2026-03-21 18:53:23 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:23 INFO     action: accept_connections | result: in_progress
client5  | 2026-03-21 18:53:23 INFO     action: apuesta_enviada | result: success | dni: 30000005 | numero: 5005
server   | 2026-03-21 18:53:23 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-21 18:53:23 INFO     action: apuesta_almacenada | result: success | dni: 30000004 | numero: 5004
server   | 2026-03-21 18:53:23 INFO     action: close_client_socket | result: success
client4  | 2026-03-21 18:53:23 INFO     action: apuesta_enviada | result: success | dni: 30000004 | numero: 5004
server   | 2026-03-21 18:53:23 INFO     action: accept_connections | result: in_progress
server   | 2026-03-21 18:53:28 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-21 18:53:28 INFO     action: apuesta_almacenada | result: success | dni: 30000002 | numero: 5002
server   | 2026-03-21 18:53:28 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:28 INFO     action: accept_connections | result: in_progress
client2  | 2026-03-21 18:53:28 INFO     action: apuesta_enviada | result: success | dni: 30000002 | numero: 5002
server   | 2026-03-21 18:53:28 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-21 18:53:28 INFO     action: apuesta_almacenada | result: success | dni: 30000003 | numero: 5003
server   | 2026-03-21 18:53:28 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:28 INFO     action: accept_connections | result: in_progress
client3  | 2026-03-21 18:53:28 INFO     action: apuesta_enviada | result: success | dni: 30000003 | numero: 5003
server   | 2026-03-21 18:53:28 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-21 18:53:28 INFO     action: apuesta_almacenada | result: success | dni: 30000001 | numero: 5001
server   | 2026-03-21 18:53:28 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:28 INFO     action: accept_connections | result: in_progress
client1  | 2026-03-21 18:53:28 INFO     action: apuesta_enviada | result: success | dni: 30000001 | numero: 5001
server   | 2026-03-21 18:53:28 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-21 18:53:28 INFO     action: apuesta_almacenada | result: success | dni: 30000005 | numero: 5005
server   | 2026-03-21 18:53:28 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:28 INFO     action: accept_connections | result: in_progress
client5  | 2026-03-21 18:53:28 INFO     action: apuesta_enviada | result: success | dni: 30000005 | numero: 5005
server   | 2026-03-21 18:53:28 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-21 18:53:28 INFO     action: apuesta_almacenada | result: success | dni: 30000004 | numero: 5004
server   | 2026-03-21 18:53:28 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:28 INFO     action: accept_connections | result: in_progress
client4  | 2026-03-21 18:53:28 INFO     action: apuesta_enviada | result: success | dni: 30000004 | numero: 5004
server   | 2026-03-21 18:53:33 INFO     action: accept_connections | result: success | ip: 172.25.125.3
server   | 2026-03-21 18:53:33 INFO     action: apuesta_almacenada | result: success | dni: 30000002 | numero: 5002
server   | 2026-03-21 18:53:33 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:33 INFO     action: accept_connections | result: in_progress
client2  | 2026-03-21 18:53:33 INFO     action: apuesta_enviada | result: success | dni: 30000002 | numero: 5002
server   | 2026-03-21 18:53:33 INFO     action: accept_connections | result: success | ip: 172.25.125.4
server   | 2026-03-21 18:53:33 INFO     action: apuesta_almacenada | result: success | dni: 30000003 | numero: 5003
server   | 2026-03-21 18:53:33 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:33 INFO     action: accept_connections | result: in_progress
client3  | 2026-03-21 18:53:33 INFO     action: apuesta_enviada | result: success | dni: 30000003 | numero: 5003
server   | 2026-03-21 18:53:33 INFO     action: accept_connections | result: success | ip: 172.25.125.5
server   | 2026-03-21 18:53:33 INFO     action: apuesta_almacenada | result: success | dni: 30000001 | numero: 5001
server   | 2026-03-21 18:53:33 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:33 INFO     action: accept_connections | result: in_progress
client1  | 2026-03-21 18:53:33 INFO     action: apuesta_enviada | result: success | dni: 30000001 | numero: 5001
server   | 2026-03-21 18:53:33 INFO     action: accept_connections | result: success | ip: 172.25.125.6
server   | 2026-03-21 18:53:33 INFO     action: apuesta_almacenada | result: success | dni: 30000005 | numero: 5005
server   | 2026-03-21 18:53:33 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:33 INFO     action: accept_connections | result: in_progress
client5  | 2026-03-21 18:53:33 INFO     action: apuesta_enviada | result: success | dni: 30000005 | numero: 5005
server   | 2026-03-21 18:53:33 INFO     action: accept_connections | result: success | ip: 172.25.125.7
server   | 2026-03-21 18:53:33 INFO     action: apuesta_almacenada | result: success | dni: 30000004 | numero: 5004
server   | 2026-03-21 18:53:33 INFO     action: close_client_socket | result: success
server   | 2026-03-21 18:53:33 INFO     action: accept_connections | result: in_progress
client4  | 2026-03-21 18:53:33 INFO     action: apuesta_enviada | result: success | dni: 30000004 | numero: 5004
client2  | 2026-03-21 18:53:38 INFO     action: loop_finished | result: success | client_id: 2
client3  | 2026-03-21 18:53:38 INFO     action: loop_finished | result: success | client_id: 3
client1  | 2026-03-21 18:53:38 INFO     action: loop_finished | result: success | client_id: 1
client5  | 2026-03-21 18:53:38 INFO     action: loop_finished | result: success | client_id: 5
client4  | 2026-03-21 18:53:38 INFO     action: loop_finished | result: success | client_id: 4
client2 exited with code 0
client3 exited with code 0
client1 exited with code 0
client5 exited with code 0
client4 exited with code 0
```