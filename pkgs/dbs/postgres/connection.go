package postgres

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Connection struct {
	Host                        string
	Port                        string
	Database                    string
	User                        string
	Password                    string
	SSLMode                     SSLMode
	SSLCertAuthorityCertificate string
	SSLPublicCertificate        string
	SSLPrivateKey               string
	FallbackConnections         []FallbackConnection
	MaxOpenConnections          int
	MaxIdleConnections          int
	ConnectionMaxIdleTime       time.Duration
	ConnectionMaxLifeTime       time.Duration
	ConnectionTimeout           time.Duration
	ConnectionOptions           []ConnectionOption
}

// CombineInstance combines the host and port of the connection
// along with the host and port of any fallback connections
// and returns a comma-separated string of all the combined instances.
func (c Connection) CombineInstance() string {
	// Create a slice to store the combined instances
	instances := []string{fmt.Sprintf("%s:%s", c.Host, c.Port)}

	// Iterate over the fallback connections
	for _, v := range c.FallbackConnections {
		// Append the host and port of each fallback connection to the slice
		instances = append(instances, fmt.Sprintf("%s:%s", v.Host, v.Port))
	}

	// Join all the instances with a comma and remove any leading/trailing commas
	return strings.Trim(strings.Join(instances, ","), ",")
}

// ToPostgresConnectionString returns the PostgresSQL connection string based on the Connection struct.
func (c Connection) ToPostgresConnectionString() string {
	// Escape các giá trị để xử lý ký tự đặc biệt
	user := url.QueryEscape(c.User)
	password := url.QueryEscape(c.Password)
	host := url.QueryEscape(c.Host)
	database := url.QueryEscape(c.Database)

	// Xây dựng chuỗi kết nối cơ bản
	var builder strings.Builder
	fmt.Fprintf(&builder, "postgresql://%s:%s@%s/%s", user, password, host, database)

	// Thêm tham số đầu tiên (SSLMode)
	if c.SSLMode != "" {
		builder.WriteString(fmt.Sprintf("?sslmode=%s", c.SSLMode))
	} else {
		builder.WriteString("?sslmode=require") // Giá trị mặc định nếu không được cung cấp
	}

	// Nếu SSLMode không phải "disable", thêm các tham số SSL (nếu có)
	if c.SSLMode != "disable" {
		if c.SSLCertAuthorityCertificate != "" {
			builder.WriteString(fmt.Sprintf("&sslrootcert=%s", url.QueryEscape(c.SSLCertAuthorityCertificate)))
		}
		if c.SSLPublicCertificate != "" {
			builder.WriteString(fmt.Sprintf("&sslcert=%s", url.QueryEscape(c.SSLPublicCertificate)))
		}
		if c.SSLPrivateKey != "" {
			builder.WriteString(fmt.Sprintf("&sslkey=%s", url.QueryEscape(c.SSLPrivateKey)))
		}
	}

	// Trả về chuỗi kết nối hoàn chỉnh
	return builder.String()
}

type FallbackConnection struct {
	Host string
	Port string
}

type ConnectionOption func(*Connection)

// SetConnection returns a ConnectionOption function that sets the host and port of a Connection struct.
//
// Parameters:
//   - host: the host address to set
//   - port: the port number to set
//
// Returns:
//   - ConnectionOption: a function that sets the host and port of a Connection struct
func SetConnection(host string, port string) ConnectionOption {
	return func(c *Connection) {
		c.Host = host
		c.Port = port
	}
}

// SetFallbackConnection is a function that returns a ConnectionOption.
// It sets the fallback connection with the given host and port.
// If both the host and port are provided, the fallback connection is added to the list of fallback connections.
// The returned ConnectionOption function can be used to modify a Connection object.
// The modified Connection object will have the fallback connection added if the host and port are provided.
func SetFallbackConnection(host string, port string) ConnectionOption {
	return func(c *Connection) {
		// Check if both host and port are provided
		if host != "" && port != "" {
			// Add the fallback connection to the list of fallback connections
			c.FallbackConnections = append(c.FallbackConnections, FallbackConnection{
				Host: host,
				Port: port,
			})
		}
	}
}

// SetSSL is a function that returns a ConnectionOption function.
// The returned function sets the SSL mode and certificates for a Connection.
// It takes in the SSL mode, CA certificate, public certificate, and private key as arguments.
func SetSSL(mode SSLMode, caCertificate, publicCertificate, privateKey string) ConnectionOption {
	return func(c *Connection) {
		c.SSLMode = mode
		c.SSLCertAuthorityCertificate = caCertificate
		c.SSLPublicCertificate = publicCertificate
		c.SSLPrivateKey = privateKey
	}
}

func SetMaxOpenConnections(max int) ConnectionOption {
	return func(c *Connection) {
		c.MaxOpenConnections = max
	}
}

func SetMaxIdleConnections(max int) ConnectionOption {
	return func(c *Connection) {
		c.MaxIdleConnections = max
	}
}

func SetConnectionMaxIdleTime(max time.Duration) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionMaxIdleTime = max
	}
}

func SetConnectionMaxLifeTime(max time.Duration) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionMaxLifeTime = max
	}
}

func SetConnectionTimeout(max time.Duration) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionTimeout = max
	}
}

func SetLoginCredentials(user, password string) ConnectionOption {
	return func(c *Connection) {
		c.User = user
		c.Password = password
	}
}

func SetDatabase(database string) ConnectionOption {
	return func(c *Connection) {
		c.Database = database
	}
}

func AddChainConnectionOptions(opts ...ConnectionOption) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionOptions = opts
	}
}
