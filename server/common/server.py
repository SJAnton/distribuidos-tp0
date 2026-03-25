from common import comms, utils
import signal
import socket
import logging
import threading


class Server:
    def __init__(self, port, listen_backlog):
        # Initialize server socket
        self._server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self._server_socket.bind(('', port))
        self._server_socket.listen(listen_backlog)
        self._lock = threading.Lock()
        self._known_agencies = set()
        self._done_agencies = set()
        self._done_draw = False
        self._is_running = True

    def run(self):
        """
        Dummy Server loop

        Server that accept a new connections and establishes a
        communication with a client. After client with communucation
        finishes, servers starts to accept new connections again
        """
        signal.signal(signal.SIGTERM, self.shutdown)
        while self._is_running:
            try:
                client_sock = self.__accept_new_connection()
                threading.Thread(
                    target=self.__handle_client_connection,
                    args=(client_sock,)
                ).start()
            except OSError as e:
                if not self._is_running:
                    break

    def shutdown(self, signum, frame):
        '''
        Shuts down communication between server and client
        '''
        self._is_running = False
        try:
            self._server_socket.close()
            logging.info("action: close_server_socket | result: success")
        except OSError as e:
            logging.error(f"action: close_server_socket | result: fail | error: {e}")
        logging.info("action: shutdown | result: success")

    def __handle_client_connection(self, client_sock):
        """
        Read message from a specific client socket and closes the socket

        If a problem arises in the communication with the client, the
        client socket will also be closed
        """
        size = 0
        try:
            recv_msg = comms.receive_message(client_sock)
            if recv_msg.endswith("DONE"):
                agency_id = recv_msg.split("|")[0]
                with self._lock:
                    self._done_agencies.add(int(agency_id))
                    if self._done_agencies == self._known_agencies:
                        self._done_draw = True
                        logging.info("action: sorteo | result: success")

            elif recv_msg.endswith("RESULTS"):
                agency_id = recv_msg.split("|")[0]
                with self._lock:
                    if not self._done_draw:
                        comms.send_message(client_sock, "WAIT")
                        return
                winners = utils.get_winners(int(agency_id))
                comms.send_message(client_sock, "|".join(winners))
                
            else:
                bets_str = recv_msg.split("\n")
                size = len(bets_str)
                bets = utils.list_to_bets(bets_str)
                utils.store_bets(bets)
                agency_id = bets_str[0].split("|")[0]
                with self._lock:
                    self._known_agencies.add(int(agency_id))
                logging.info(
                    f"action: apuesta_recibida | result: success | cantidad: {size}"
                )
                comms.send_message(client_sock, "OK")
        except Exception:
            logging.error(f"action: apuesta_recibida | result: fail | cantidad: {size}")
            comms.send_message(client_sock, "FAIL")
        finally:
            try:
                client_sock.close()
                logging.info("action: close_client_socket | result: success")
            except OSError as e:
                logging.error(f"action: close_client_socket | result: fail | error: {e}")

    def __accept_new_connection(self):
        """
        Accept new connections

        Function blocks until a connection to a client is made.
        Then connection created is printed and returned
        """

        # Connection arrived
        logging.info('action: accept_connections | result: in_progress')
        c, addr = self._server_socket.accept()
        logging.info(f'action: accept_connections | result: success | ip: {addr[0]}')
        return c
