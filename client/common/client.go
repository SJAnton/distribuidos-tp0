package common

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/op/go-logging"
)

const batchSize = 8192

var log = logging.MustGetLogger("log")

// ClientConfig Configuration used by the client
type ClientConfig struct {
	ID             string
	ServerAddress  string
	LoopPeriod     time.Duration
	BatchMaxAmount int
	BatchMaxSize   int
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

	records, err := read_csv_file("./agency.csv")
	if err != nil {
		log.Criticalf("action: read_csv_file | result: fail | error: %v", err)
		return
	}
	batches := GetBatches(records, batchSize, c.config)
	for i := 0; i < len(batches); i += 1 {
		select {
		case <-stop:
			log.Infof("action: shutdown | result: success")
			return
		default:
			// Create the connection the server in every loop iteration. Send an
			c.createClientSocket()
			sendMessage(c.conn, []byte(batches[i]))
			msg, err := receiveMessage(c.conn)
			c.conn.Close()

			if err != nil || msg == "FAIL" {
				log.Errorf("action: apuesta_enviada | result: fail")
				return
			}

			log.Infof("action: apuesta_enviada | result: success")

			// Wait a time between sending one message and the next one
			time.Sleep(c.config.LoopPeriod)

		}
	}
	log.Infof("action: loop_finished | result: success | client_id: %v", c.config.ID)
}
