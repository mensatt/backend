// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mensatt/backend/internal/database/ent/migrate"

	"github.com/mensatt/backend/internal/database/ent/dish"
	"github.com/mensatt/backend/internal/database/ent/dishalias"
	"github.com/mensatt/backend/internal/database/ent/image"
	"github.com/mensatt/backend/internal/database/ent/location"
	"github.com/mensatt/backend/internal/database/ent/occurrence"
	"github.com/mensatt/backend/internal/database/ent/review"
	"github.com/mensatt/backend/internal/database/ent/tag"
	"github.com/mensatt/backend/internal/database/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Dish is the client for interacting with the Dish builders.
	Dish *DishClient
	// DishAlias is the client for interacting with the DishAlias builders.
	DishAlias *DishAliasClient
	// Image is the client for interacting with the Image builders.
	Image *ImageClient
	// Location is the client for interacting with the Location builders.
	Location *LocationClient
	// Occurrence is the client for interacting with the Occurrence builders.
	Occurrence *OccurrenceClient
	// Review is the client for interacting with the Review builders.
	Review *ReviewClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.Dish = NewDishClient(c.config)
	c.DishAlias = NewDishAliasClient(c.config)
	c.Image = NewImageClient(c.config)
	c.Location = NewLocationClient(c.config)
	c.Occurrence = NewOccurrenceClient(c.config)
	c.Review = NewReviewClient(c.config)
	c.Tag = NewTagClient(c.config)
	c.User = NewUserClient(c.config)
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Dish:       NewDishClient(cfg),
		DishAlias:  NewDishAliasClient(cfg),
		Image:      NewImageClient(cfg),
		Location:   NewLocationClient(cfg),
		Occurrence: NewOccurrenceClient(cfg),
		Review:     NewReviewClient(cfg),
		Tag:        NewTagClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:        ctx,
		config:     cfg,
		Dish:       NewDishClient(cfg),
		DishAlias:  NewDishAliasClient(cfg),
		Image:      NewImageClient(cfg),
		Location:   NewLocationClient(cfg),
		Occurrence: NewOccurrenceClient(cfg),
		Review:     NewReviewClient(cfg),
		Tag:        NewTagClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Dish.
//		Query().
//		Count(ctx)
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
	c.Dish.Use(hooks...)
	c.DishAlias.Use(hooks...)
	c.Image.Use(hooks...)
	c.Location.Use(hooks...)
	c.Occurrence.Use(hooks...)
	c.Review.Use(hooks...)
	c.Tag.Use(hooks...)
	c.User.Use(hooks...)
}

// DishClient is a client for the Dish schema.
type DishClient struct {
	config
}

// NewDishClient returns a client for the Dish from the given config.
func NewDishClient(c config) *DishClient {
	return &DishClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dish.Hooks(f(g(h())))`.
func (c *DishClient) Use(hooks ...Hook) {
	c.hooks.Dish = append(c.hooks.Dish, hooks...)
}

// Create returns a builder for creating a Dish entity.
func (c *DishClient) Create() *DishCreate {
	mutation := newDishMutation(c.config, OpCreate)
	return &DishCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Dish entities.
func (c *DishClient) CreateBulk(builders ...*DishCreate) *DishCreateBulk {
	return &DishCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Dish.
func (c *DishClient) Update() *DishUpdate {
	mutation := newDishMutation(c.config, OpUpdate)
	return &DishUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DishClient) UpdateOne(d *Dish) *DishUpdateOne {
	mutation := newDishMutation(c.config, OpUpdateOne, withDish(d))
	return &DishUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DishClient) UpdateOneID(id uuid.UUID) *DishUpdateOne {
	mutation := newDishMutation(c.config, OpUpdateOne, withDishID(id))
	return &DishUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Dish.
func (c *DishClient) Delete() *DishDelete {
	mutation := newDishMutation(c.config, OpDelete)
	return &DishDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DishClient) DeleteOne(d *Dish) *DishDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DishClient) DeleteOneID(id uuid.UUID) *DishDeleteOne {
	builder := c.Delete().Where(dish.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DishDeleteOne{builder}
}

// Query returns a query builder for Dish.
func (c *DishClient) Query() *DishQuery {
	return &DishQuery{
		config: c.config,
	}
}

// Get returns a Dish entity by its id.
func (c *DishClient) Get(ctx context.Context, id uuid.UUID) (*Dish, error) {
	return c.Query().Where(dish.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DishClient) GetX(ctx context.Context, id uuid.UUID) *Dish {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDishOccurrences queries the dish_occurrences edge of a Dish.
func (c *DishClient) QueryDishOccurrences(d *Dish) *OccurrenceQuery {
	query := &OccurrenceQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dish.Table, dish.FieldID, id),
			sqlgraph.To(occurrence.Table, occurrence.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dish.DishOccurrencesTable, dish.DishOccurrencesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAliases queries the aliases edge of a Dish.
func (c *DishClient) QueryAliases(d *Dish) *DishAliasQuery {
	query := &DishAliasQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dish.Table, dish.FieldID, id),
			sqlgraph.To(dishalias.Table, dishalias.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dish.AliasesTable, dish.AliasesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySideDishOccurrence queries the side_dish_occurrence edge of a Dish.
func (c *DishClient) QuerySideDishOccurrence(d *Dish) *OccurrenceQuery {
	query := &OccurrenceQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dish.Table, dish.FieldID, id),
			sqlgraph.To(occurrence.Table, occurrence.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, dish.SideDishOccurrenceTable, dish.SideDishOccurrencePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DishClient) Hooks() []Hook {
	return c.hooks.Dish
}

// DishAliasClient is a client for the DishAlias schema.
type DishAliasClient struct {
	config
}

// NewDishAliasClient returns a client for the DishAlias from the given config.
func NewDishAliasClient(c config) *DishAliasClient {
	return &DishAliasClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dishalias.Hooks(f(g(h())))`.
func (c *DishAliasClient) Use(hooks ...Hook) {
	c.hooks.DishAlias = append(c.hooks.DishAlias, hooks...)
}

// Create returns a builder for creating a DishAlias entity.
func (c *DishAliasClient) Create() *DishAliasCreate {
	mutation := newDishAliasMutation(c.config, OpCreate)
	return &DishAliasCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DishAlias entities.
func (c *DishAliasClient) CreateBulk(builders ...*DishAliasCreate) *DishAliasCreateBulk {
	return &DishAliasCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DishAlias.
func (c *DishAliasClient) Update() *DishAliasUpdate {
	mutation := newDishAliasMutation(c.config, OpUpdate)
	return &DishAliasUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DishAliasClient) UpdateOne(da *DishAlias) *DishAliasUpdateOne {
	mutation := newDishAliasMutation(c.config, OpUpdateOne, withDishAlias(da))
	return &DishAliasUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DishAliasClient) UpdateOneID(id string) *DishAliasUpdateOne {
	mutation := newDishAliasMutation(c.config, OpUpdateOne, withDishAliasID(id))
	return &DishAliasUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DishAlias.
func (c *DishAliasClient) Delete() *DishAliasDelete {
	mutation := newDishAliasMutation(c.config, OpDelete)
	return &DishAliasDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DishAliasClient) DeleteOne(da *DishAlias) *DishAliasDeleteOne {
	return c.DeleteOneID(da.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DishAliasClient) DeleteOneID(id string) *DishAliasDeleteOne {
	builder := c.Delete().Where(dishalias.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DishAliasDeleteOne{builder}
}

// Query returns a query builder for DishAlias.
func (c *DishAliasClient) Query() *DishAliasQuery {
	return &DishAliasQuery{
		config: c.config,
	}
}

// Get returns a DishAlias entity by its id.
func (c *DishAliasClient) Get(ctx context.Context, id string) (*DishAlias, error) {
	return c.Query().Where(dishalias.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DishAliasClient) GetX(ctx context.Context, id string) *DishAlias {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDish queries the dish edge of a DishAlias.
func (c *DishAliasClient) QueryDish(da *DishAlias) *DishQuery {
	query := &DishQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := da.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dishalias.Table, dishalias.FieldID, id),
			sqlgraph.To(dish.Table, dish.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dishalias.DishTable, dishalias.DishColumn),
		)
		fromV = sqlgraph.Neighbors(da.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DishAliasClient) Hooks() []Hook {
	return c.hooks.DishAlias
}

// ImageClient is a client for the Image schema.
type ImageClient struct {
	config
}

// NewImageClient returns a client for the Image from the given config.
func NewImageClient(c config) *ImageClient {
	return &ImageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `image.Hooks(f(g(h())))`.
func (c *ImageClient) Use(hooks ...Hook) {
	c.hooks.Image = append(c.hooks.Image, hooks...)
}

// Create returns a builder for creating a Image entity.
func (c *ImageClient) Create() *ImageCreate {
	mutation := newImageMutation(c.config, OpCreate)
	return &ImageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Image entities.
func (c *ImageClient) CreateBulk(builders ...*ImageCreate) *ImageCreateBulk {
	return &ImageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Image.
func (c *ImageClient) Update() *ImageUpdate {
	mutation := newImageMutation(c.config, OpUpdate)
	return &ImageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ImageClient) UpdateOne(i *Image) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImage(i))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ImageClient) UpdateOneID(id uuid.UUID) *ImageUpdateOne {
	mutation := newImageMutation(c.config, OpUpdateOne, withImageID(id))
	return &ImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Image.
func (c *ImageClient) Delete() *ImageDelete {
	mutation := newImageMutation(c.config, OpDelete)
	return &ImageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ImageClient) DeleteOne(i *Image) *ImageDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ImageClient) DeleteOneID(id uuid.UUID) *ImageDeleteOne {
	builder := c.Delete().Where(image.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ImageDeleteOne{builder}
}

// Query returns a query builder for Image.
func (c *ImageClient) Query() *ImageQuery {
	return &ImageQuery{
		config: c.config,
	}
}

// Get returns a Image entity by its id.
func (c *ImageClient) Get(ctx context.Context, id uuid.UUID) (*Image, error) {
	return c.Query().Where(image.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ImageClient) GetX(ctx context.Context, id uuid.UUID) *Image {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryReview queries the review edge of a Image.
func (c *ImageClient) QueryReview(i *Image) *ReviewQuery {
	query := &ReviewQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(image.Table, image.FieldID, id),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, image.ReviewTable, image.ReviewColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ImageClient) Hooks() []Hook {
	return c.hooks.Image
}

// LocationClient is a client for the Location schema.
type LocationClient struct {
	config
}

// NewLocationClient returns a client for the Location from the given config.
func NewLocationClient(c config) *LocationClient {
	return &LocationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `location.Hooks(f(g(h())))`.
func (c *LocationClient) Use(hooks ...Hook) {
	c.hooks.Location = append(c.hooks.Location, hooks...)
}

// Create returns a builder for creating a Location entity.
func (c *LocationClient) Create() *LocationCreate {
	mutation := newLocationMutation(c.config, OpCreate)
	return &LocationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Location entities.
func (c *LocationClient) CreateBulk(builders ...*LocationCreate) *LocationCreateBulk {
	return &LocationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Location.
func (c *LocationClient) Update() *LocationUpdate {
	mutation := newLocationMutation(c.config, OpUpdate)
	return &LocationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LocationClient) UpdateOne(l *Location) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocation(l))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LocationClient) UpdateOneID(id uuid.UUID) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocationID(id))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Location.
func (c *LocationClient) Delete() *LocationDelete {
	mutation := newLocationMutation(c.config, OpDelete)
	return &LocationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LocationClient) DeleteOne(l *Location) *LocationDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LocationClient) DeleteOneID(id uuid.UUID) *LocationDeleteOne {
	builder := c.Delete().Where(location.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LocationDeleteOne{builder}
}

// Query returns a query builder for Location.
func (c *LocationClient) Query() *LocationQuery {
	return &LocationQuery{
		config: c.config,
	}
}

// Get returns a Location entity by its id.
func (c *LocationClient) Get(ctx context.Context, id uuid.UUID) (*Location, error) {
	return c.Query().Where(location.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LocationClient) GetX(ctx context.Context, id uuid.UUID) *Location {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOccurrences queries the occurrences edge of a Location.
func (c *LocationClient) QueryOccurrences(l *Location) *OccurrenceQuery {
	query := &OccurrenceQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(location.Table, location.FieldID, id),
			sqlgraph.To(occurrence.Table, occurrence.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, location.OccurrencesTable, location.OccurrencesColumn),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LocationClient) Hooks() []Hook {
	return c.hooks.Location
}

// OccurrenceClient is a client for the Occurrence schema.
type OccurrenceClient struct {
	config
}

// NewOccurrenceClient returns a client for the Occurrence from the given config.
func NewOccurrenceClient(c config) *OccurrenceClient {
	return &OccurrenceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `occurrence.Hooks(f(g(h())))`.
func (c *OccurrenceClient) Use(hooks ...Hook) {
	c.hooks.Occurrence = append(c.hooks.Occurrence, hooks...)
}

// Create returns a builder for creating a Occurrence entity.
func (c *OccurrenceClient) Create() *OccurrenceCreate {
	mutation := newOccurrenceMutation(c.config, OpCreate)
	return &OccurrenceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Occurrence entities.
func (c *OccurrenceClient) CreateBulk(builders ...*OccurrenceCreate) *OccurrenceCreateBulk {
	return &OccurrenceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Occurrence.
func (c *OccurrenceClient) Update() *OccurrenceUpdate {
	mutation := newOccurrenceMutation(c.config, OpUpdate)
	return &OccurrenceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OccurrenceClient) UpdateOne(o *Occurrence) *OccurrenceUpdateOne {
	mutation := newOccurrenceMutation(c.config, OpUpdateOne, withOccurrence(o))
	return &OccurrenceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OccurrenceClient) UpdateOneID(id uuid.UUID) *OccurrenceUpdateOne {
	mutation := newOccurrenceMutation(c.config, OpUpdateOne, withOccurrenceID(id))
	return &OccurrenceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Occurrence.
func (c *OccurrenceClient) Delete() *OccurrenceDelete {
	mutation := newOccurrenceMutation(c.config, OpDelete)
	return &OccurrenceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OccurrenceClient) DeleteOne(o *Occurrence) *OccurrenceDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OccurrenceClient) DeleteOneID(id uuid.UUID) *OccurrenceDeleteOne {
	builder := c.Delete().Where(occurrence.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OccurrenceDeleteOne{builder}
}

// Query returns a query builder for Occurrence.
func (c *OccurrenceClient) Query() *OccurrenceQuery {
	return &OccurrenceQuery{
		config: c.config,
	}
}

// Get returns a Occurrence entity by its id.
func (c *OccurrenceClient) Get(ctx context.Context, id uuid.UUID) (*Occurrence, error) {
	return c.Query().Where(occurrence.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OccurrenceClient) GetX(ctx context.Context, id uuid.UUID) *Occurrence {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryLocation queries the location edge of a Occurrence.
func (c *OccurrenceClient) QueryLocation(o *Occurrence) *LocationQuery {
	query := &LocationQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(occurrence.Table, occurrence.FieldID, id),
			sqlgraph.To(location.Table, location.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, occurrence.LocationTable, occurrence.LocationColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDish queries the dish edge of a Occurrence.
func (c *OccurrenceClient) QueryDish(o *Occurrence) *DishQuery {
	query := &DishQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(occurrence.Table, occurrence.FieldID, id),
			sqlgraph.To(dish.Table, dish.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, occurrence.DishTable, occurrence.DishColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTags queries the tags edge of a Occurrence.
func (c *OccurrenceClient) QueryTags(o *Occurrence) *TagQuery {
	query := &TagQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(occurrence.Table, occurrence.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, occurrence.TagsTable, occurrence.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySideDishes queries the side_dishes edge of a Occurrence.
func (c *OccurrenceClient) QuerySideDishes(o *Occurrence) *DishQuery {
	query := &DishQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(occurrence.Table, occurrence.FieldID, id),
			sqlgraph.To(dish.Table, dish.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, occurrence.SideDishesTable, occurrence.SideDishesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReviews queries the reviews edge of a Occurrence.
func (c *OccurrenceClient) QueryReviews(o *Occurrence) *ReviewQuery {
	query := &ReviewQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(occurrence.Table, occurrence.FieldID, id),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, occurrence.ReviewsTable, occurrence.ReviewsColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OccurrenceClient) Hooks() []Hook {
	return c.hooks.Occurrence
}

// ReviewClient is a client for the Review schema.
type ReviewClient struct {
	config
}

// NewReviewClient returns a client for the Review from the given config.
func NewReviewClient(c config) *ReviewClient {
	return &ReviewClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `review.Hooks(f(g(h())))`.
func (c *ReviewClient) Use(hooks ...Hook) {
	c.hooks.Review = append(c.hooks.Review, hooks...)
}

// Create returns a builder for creating a Review entity.
func (c *ReviewClient) Create() *ReviewCreate {
	mutation := newReviewMutation(c.config, OpCreate)
	return &ReviewCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Review entities.
func (c *ReviewClient) CreateBulk(builders ...*ReviewCreate) *ReviewCreateBulk {
	return &ReviewCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Review.
func (c *ReviewClient) Update() *ReviewUpdate {
	mutation := newReviewMutation(c.config, OpUpdate)
	return &ReviewUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReviewClient) UpdateOne(r *Review) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReview(r))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReviewClient) UpdateOneID(id uuid.UUID) *ReviewUpdateOne {
	mutation := newReviewMutation(c.config, OpUpdateOne, withReviewID(id))
	return &ReviewUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Review.
func (c *ReviewClient) Delete() *ReviewDelete {
	mutation := newReviewMutation(c.config, OpDelete)
	return &ReviewDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ReviewClient) DeleteOne(r *Review) *ReviewDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ReviewClient) DeleteOneID(id uuid.UUID) *ReviewDeleteOne {
	builder := c.Delete().Where(review.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReviewDeleteOne{builder}
}

// Query returns a query builder for Review.
func (c *ReviewClient) Query() *ReviewQuery {
	return &ReviewQuery{
		config: c.config,
	}
}

// Get returns a Review entity by its id.
func (c *ReviewClient) Get(ctx context.Context, id uuid.UUID) (*Review, error) {
	return c.Query().Where(review.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReviewClient) GetX(ctx context.Context, id uuid.UUID) *Review {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOccurrence queries the occurrence edge of a Review.
func (c *ReviewClient) QueryOccurrence(r *Review) *OccurrenceQuery {
	query := &OccurrenceQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, id),
			sqlgraph.To(occurrence.Table, occurrence.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, review.OccurrenceTable, review.OccurrenceColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryImages queries the images edge of a Review.
func (c *ReviewClient) QueryImages(r *Review) *ImageQuery {
	query := &ImageQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(review.Table, review.FieldID, id),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, review.ImagesTable, review.ImagesColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ReviewClient) Hooks() []Hook {
	return c.hooks.Review
}

// TagClient is a client for the Tag schema.
type TagClient struct {
	config
}

// NewTagClient returns a client for the Tag from the given config.
func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tag.Hooks(f(g(h())))`.
func (c *TagClient) Use(hooks ...Hook) {
	c.hooks.Tag = append(c.hooks.Tag, hooks...)
}

// Create returns a builder for creating a Tag entity.
func (c *TagClient) Create() *TagCreate {
	mutation := newTagMutation(c.config, OpCreate)
	return &TagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tag entities.
func (c *TagClient) CreateBulk(builders ...*TagCreate) *TagCreateBulk {
	return &TagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tag.
func (c *TagClient) Update() *TagUpdate {
	mutation := newTagMutation(c.config, OpUpdate)
	return &TagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagClient) UpdateOne(t *Tag) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTag(t))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TagClient) UpdateOneID(id string) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTagID(id))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tag.
func (c *TagClient) Delete() *TagDelete {
	mutation := newTagMutation(c.config, OpDelete)
	return &TagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TagClient) DeleteOne(t *Tag) *TagDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TagClient) DeleteOneID(id string) *TagDeleteOne {
	builder := c.Delete().Where(tag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagDeleteOne{builder}
}

// Query returns a query builder for Tag.
func (c *TagClient) Query() *TagQuery {
	return &TagQuery{
		config: c.config,
	}
}

// Get returns a Tag entity by its id.
func (c *TagClient) Get(ctx context.Context, id string) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id string) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOccurrence queries the occurrence edge of a Tag.
func (c *TagClient) QueryOccurrence(t *Tag) *OccurrenceQuery {
	query := &OccurrenceQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(occurrence.Table, occurrence.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tag.OccurrenceTable, tag.OccurrencePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TagClient) Hooks() []Hook {
	return c.hooks.Tag
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
