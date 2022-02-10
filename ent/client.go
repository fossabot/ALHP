// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"ALHP.GO/ent/migrate"

	"ALHP.GO/ent/dbpackage"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// DbPackage is the client for interacting with the DbPackage builders.
	DbPackage *DbPackageClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.DbPackage = NewDbPackageClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		DbPackage: NewDbPackageClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:    cfg,
		DbPackage: NewDbPackageClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		DbPackage.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.DbPackage.Use(hooks...)
}

// DbPackageClient is a client for the DbPackage schema.
type DbPackageClient struct {
	config
}

// NewDbPackageClient returns a client for the DbPackage from the given config.
func NewDbPackageClient(c config) *DbPackageClient {
	return &DbPackageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dbpackage.Hooks(f(g(h())))`.
func (c *DbPackageClient) Use(hooks ...Hook) {
	c.hooks.DbPackage = append(c.hooks.DbPackage, hooks...)
}

// Create returns a create builder for DbPackage.
func (c *DbPackageClient) Create() *DbPackageCreate {
	mutation := newDbPackageMutation(c.config, OpCreate)
	return &DbPackageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DbPackage entities.
func (c *DbPackageClient) CreateBulk(builders ...*DbPackageCreate) *DbPackageCreateBulk {
	return &DbPackageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DbPackage.
func (c *DbPackageClient) Update() *DbPackageUpdate {
	mutation := newDbPackageMutation(c.config, OpUpdate)
	return &DbPackageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DbPackageClient) UpdateOne(dp *DbPackage) *DbPackageUpdateOne {
	mutation := newDbPackageMutation(c.config, OpUpdateOne, withDbPackage(dp))
	return &DbPackageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DbPackageClient) UpdateOneID(id int) *DbPackageUpdateOne {
	mutation := newDbPackageMutation(c.config, OpUpdateOne, withDbPackageID(id))
	return &DbPackageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DbPackage.
func (c *DbPackageClient) Delete() *DbPackageDelete {
	mutation := newDbPackageMutation(c.config, OpDelete)
	return &DbPackageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DbPackageClient) DeleteOne(dp *DbPackage) *DbPackageDeleteOne {
	return c.DeleteOneID(dp.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DbPackageClient) DeleteOneID(id int) *DbPackageDeleteOne {
	builder := c.Delete().Where(dbpackage.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DbPackageDeleteOne{builder}
}

// Query returns a query builder for DbPackage.
func (c *DbPackageClient) Query() *DbPackageQuery {
	return &DbPackageQuery{
		config: c.config,
	}
}

// Get returns a DbPackage entity by its id.
func (c *DbPackageClient) Get(ctx context.Context, id int) (*DbPackage, error) {
	return c.Query().Where(dbpackage.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DbPackageClient) GetX(ctx context.Context, id int) *DbPackage {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DbPackageClient) Hooks() []Hook {
	return c.hooks.DbPackage
}
