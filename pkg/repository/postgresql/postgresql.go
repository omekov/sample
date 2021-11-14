package postgresql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type Manager struct {
	conn                *pgx.Conn
	connConfig          *pgx.ConnConfig
	logger              *logrus.Logger
	keepAlivePollPeriod int
	maxConnectAttempts  int
}

type Repositories struct {
	User Userer
}

func NewManager(uri string, logrus *logrus.Logger) *Manager {
	pgxConfig, err := pgx.ParseConfig(uri)
	if err != nil {
		logrus.Warnf("ParseConfig. %s", err.Error())
	}

	return &Manager{
		connConfig: pgxConfig,
		logger:     logrus,
	}
}

func (m *Manager) connection(ctx context.Context) (err error) {
	m.conn, err = pgx.ConnectConfig(ctx, m.connConfig)
	if err != nil {
		return err
	}

	return m.ping(ctx)
}

func (m *Manager) ping(ctx context.Context) error {
	return m.conn.Ping(ctx)
}

func (m *Manager) GetConnection(ctx context.Context) *pgx.Conn {
	if m.conn == nil {
		if err := m.connection(ctx); err != nil {
			m.logger.Warnf("connection. %s", err)
		}
	} else {
		if err := m.ping(ctx); err != nil {
			attempt := 0
			ticker := time.NewTicker(time.Duration(m.keepAlivePollPeriod) * time.Second)
			defer ticker.Stop()
			for range ticker.C {
				if attempt >= m.maxConnectAttempts {
					m.logger.Warnf("connection failed after %d attempt\n", attempt)
				}
				attempt++
				m.logger.Infof("reconnecting...")

				if err := m.connection(ctx); err != nil {
					m.logger.Infof("connection was lost. Error: %s. Waiting for %d sec...\n", err, m.keepAlivePollPeriod)
				}
			}
		}
	}

	return m.conn
}

func NewRepositories(conn *pgx.Conn) *Repositories {
	return &Repositories{
		User: NewUser(conn),
	}
}
