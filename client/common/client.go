package common

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("log")

// ClientConfig Configuration used by the client
type ClientConfig struct {
	ID            string
	ServerAddress string
	LoopAmount    int
	LoopPeriod    time.Duration
}

// Client Entity that encapsulates how
type Client struct {
	config ClientConfig
	conn   net.Conn
}

// NewClient Initializes a new client receiving the configuration
// as a parameter
func NewClient(config ClientConfig) *Client {
	client := &Client{
		config: config,
	}
	return client
}

// CreateClientSocket Initializes client socket. In case of
// failure, error is printed in stdout/stderr and exit 1
// is returned
func (c *Client) createClientSocket() error {
	conn, err := net.Dial("tcp", c.config.ServerAddress)
	if err != nil {
		log.Criticalf(
			"action: connect | result: fail | client_id: %v | error: %v",
			c.config.ID,
			err,
		)
	}
	c.conn = conn
	return nil
}

// StartClientLoop Send messages to the client until some time threshold is met
func (c *Client) StartClientLoop() {
	// There is an autoincremental msgID to identify every message sent
	// Messages if the message amount threshold has not been surpassed

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	bet := NewBet(
		c.config.ID,
		os.Getenv("NOMBRE"),
		os.Getenv("APELLIDO"),
		os.Getenv("DOCUMENTO"),
		os.Getenv("NACIMIENTO"),
		os.Getenv("NUMERO"),
	)

	for msgID := 1; msgID <= c.config.LoopAmount; msgID++ {
		select {
		case <-stop:
			log.Infof("action: shutdown | result: success")
			return
		default:
			// Create the connection the server in every loop iteration. Send an
			c.createClientSocket()
			sendMessage(c.conn, []byte(bet.MakeMessage()))
			msg, err := receiveMessage(c.conn)
			c.conn.Close()

			if err != nil || msg == "FAIL" {
				log.Errorf("action: apuesta_enviada | result: fail | dni: %v | numero: %v",
					bet.Id,
					bet.Number,
				)
				return
			}

			log.Infof("action: apuesta_enviada | result: success | dni: %v | numero: %v",
				bet.Id,
				bet.Number,
			)

			// Wait a time between sending one message and the next one
			time.Sleep(c.config.LoopPeriod)

		}
	}
	log.Infof("action: loop_finished | result: success | client_id: %v", c.config.ID)
}
