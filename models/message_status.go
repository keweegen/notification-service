// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// MessageStatus is an object representing the database table.
type MessageStatus struct {
	ID          int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	MessageID   string    `boil:"message_id" json:"message_id" toml:"message_id" yaml:"message_id"`
	Status      string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	Description string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	IsLast      bool      `boil:"is_last" json:"is_last" toml:"is_last" yaml:"is_last"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *messageStatusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L messageStatusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MessageStatusColumns = struct {
	ID          string
	MessageID   string
	Status      string
	Description string
	IsLast      string
	CreatedAt   string
}{
	ID:          "id",
	MessageID:   "message_id",
	Status:      "status",
	Description: "description",
	IsLast:      "is_last",
	CreatedAt:   "created_at",
}

var MessageStatusTableColumns = struct {
	ID          string
	MessageID   string
	Status      string
	Description string
	IsLast      string
	CreatedAt   string
}{
	ID:          "message_status.id",
	MessageID:   "message_status.message_id",
	Status:      "message_status.status",
	Description: "message_status.description",
	IsLast:      "message_status.is_last",
	CreatedAt:   "message_status.created_at",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var MessageStatusWhere = struct {
	ID          whereHelperint64
	MessageID   whereHelperstring
	Status      whereHelperstring
	Description whereHelperstring
	IsLast      whereHelperbool
	CreatedAt   whereHelpertime_Time
}{
	ID:          whereHelperint64{field: "\"message_status\".\"id\""},
	MessageID:   whereHelperstring{field: "\"message_status\".\"message_id\""},
	Status:      whereHelperstring{field: "\"message_status\".\"status\""},
	Description: whereHelperstring{field: "\"message_status\".\"description\""},
	IsLast:      whereHelperbool{field: "\"message_status\".\"is_last\""},
	CreatedAt:   whereHelpertime_Time{field: "\"message_status\".\"created_at\""},
}

// MessageStatusRels is where relationship names are stored.
var MessageStatusRels = struct {
	Message string
}{
	Message: "Message",
}

// messageStatusR is where relationships are stored.
type messageStatusR struct {
	Message *Message `boil:"Message" json:"Message" toml:"Message" yaml:"Message"`
}

// NewStruct creates a new relationship struct
func (*messageStatusR) NewStruct() *messageStatusR {
	return &messageStatusR{}
}

func (r *messageStatusR) GetMessage() *Message {
	if r == nil {
		return nil
	}
	return r.Message
}

// messageStatusL is where Load methods for each relationship are stored.
type messageStatusL struct{}

var (
	messageStatusAllColumns            = []string{"id", "message_id", "status", "description", "is_last", "created_at"}
	messageStatusColumnsWithoutDefault = []string{"message_id", "status", "description"}
	messageStatusColumnsWithDefault    = []string{"id", "is_last", "created_at"}
	messageStatusPrimaryKeyColumns     = []string{"id"}
	messageStatusGeneratedColumns      = []string{}
)

type (
	// MessageStatusSlice is an alias for a slice of pointers to MessageStatus.
	// This should almost always be used instead of []MessageStatus.
	MessageStatusSlice []*MessageStatus
	// MessageStatusHook is the signature for custom MessageStatus hook methods
	MessageStatusHook func(context.Context, boil.ContextExecutor, *MessageStatus) error

	messageStatusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	messageStatusType                 = reflect.TypeOf(&MessageStatus{})
	messageStatusMapping              = queries.MakeStructMapping(messageStatusType)
	messageStatusPrimaryKeyMapping, _ = queries.BindMapping(messageStatusType, messageStatusMapping, messageStatusPrimaryKeyColumns)
	messageStatusInsertCacheMut       sync.RWMutex
	messageStatusInsertCache          = make(map[string]insertCache)
	messageStatusUpdateCacheMut       sync.RWMutex
	messageStatusUpdateCache          = make(map[string]updateCache)
	messageStatusUpsertCacheMut       sync.RWMutex
	messageStatusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var messageStatusAfterSelectHooks []MessageStatusHook

var messageStatusBeforeInsertHooks []MessageStatusHook
var messageStatusAfterInsertHooks []MessageStatusHook

var messageStatusBeforeUpdateHooks []MessageStatusHook
var messageStatusAfterUpdateHooks []MessageStatusHook

var messageStatusBeforeDeleteHooks []MessageStatusHook
var messageStatusAfterDeleteHooks []MessageStatusHook

var messageStatusBeforeUpsertHooks []MessageStatusHook
var messageStatusAfterUpsertHooks []MessageStatusHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MessageStatus) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageStatusAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MessageStatus) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageStatusBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MessageStatus) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageStatusAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MessageStatus) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageStatusBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MessageStatus) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageStatusAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MessageStatus) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageStatusBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MessageStatus) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageStatusAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MessageStatus) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageStatusBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MessageStatus) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range messageStatusAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMessageStatusHook registers your hook function for all future operations.
func AddMessageStatusHook(hookPoint boil.HookPoint, messageStatusHook MessageStatusHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		messageStatusAfterSelectHooks = append(messageStatusAfterSelectHooks, messageStatusHook)
	case boil.BeforeInsertHook:
		messageStatusBeforeInsertHooks = append(messageStatusBeforeInsertHooks, messageStatusHook)
	case boil.AfterInsertHook:
		messageStatusAfterInsertHooks = append(messageStatusAfterInsertHooks, messageStatusHook)
	case boil.BeforeUpdateHook:
		messageStatusBeforeUpdateHooks = append(messageStatusBeforeUpdateHooks, messageStatusHook)
	case boil.AfterUpdateHook:
		messageStatusAfterUpdateHooks = append(messageStatusAfterUpdateHooks, messageStatusHook)
	case boil.BeforeDeleteHook:
		messageStatusBeforeDeleteHooks = append(messageStatusBeforeDeleteHooks, messageStatusHook)
	case boil.AfterDeleteHook:
		messageStatusAfterDeleteHooks = append(messageStatusAfterDeleteHooks, messageStatusHook)
	case boil.BeforeUpsertHook:
		messageStatusBeforeUpsertHooks = append(messageStatusBeforeUpsertHooks, messageStatusHook)
	case boil.AfterUpsertHook:
		messageStatusAfterUpsertHooks = append(messageStatusAfterUpsertHooks, messageStatusHook)
	}
}

// One returns a single messageStatus record from the query.
func (q messageStatusQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MessageStatus, error) {
	o := &MessageStatus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for message_status")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MessageStatus records from the query.
func (q messageStatusQuery) All(ctx context.Context, exec boil.ContextExecutor) (MessageStatusSlice, error) {
	var o []*MessageStatus

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MessageStatus slice")
	}

	if len(messageStatusAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MessageStatus records in the query.
func (q messageStatusQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count message_status rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q messageStatusQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if message_status exists")
	}

	return count > 0, nil
}

// Message pointed to by the foreign key.
func (o *MessageStatus) Message(mods ...qm.QueryMod) messageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.MessageID),
	}

	queryMods = append(queryMods, mods...)

	return Messages(queryMods...)
}

// LoadMessage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (messageStatusL) LoadMessage(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMessageStatus interface{}, mods queries.Applicator) error {
	var slice []*MessageStatus
	var object *MessageStatus

	if singular {
		var ok bool
		object, ok = maybeMessageStatus.(*MessageStatus)
		if !ok {
			object = new(MessageStatus)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMessageStatus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMessageStatus))
			}
		}
	} else {
		s, ok := maybeMessageStatus.(*[]*MessageStatus)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMessageStatus)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMessageStatus))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &messageStatusR{}
		}
		args = append(args, object.MessageID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &messageStatusR{}
			}

			for _, a := range args {
				if a == obj.MessageID {
					continue Outer
				}
			}

			args = append(args, obj.MessageID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`message`),
		qm.WhereIn(`message.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Message")
	}

	var resultSlice []*Message
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Message")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for message")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for message")
	}

	if len(messageStatusAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Message = foreign
		if foreign.R == nil {
			foreign.R = &messageR{}
		}
		foreign.R.MessageStatuses = append(foreign.R.MessageStatuses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MessageID == foreign.ID {
				local.R.Message = foreign
				if foreign.R == nil {
					foreign.R = &messageR{}
				}
				foreign.R.MessageStatuses = append(foreign.R.MessageStatuses, local)
				break
			}
		}
	}

	return nil
}

// SetMessage of the messageStatus to the related item.
// Sets o.R.Message to related.
// Adds o to related.R.MessageStatuses.
func (o *MessageStatus) SetMessage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Message) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"message_status\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"message_id"}),
		strmangle.WhereClause("\"", "\"", 2, messageStatusPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.MessageID = related.ID
	if o.R == nil {
		o.R = &messageStatusR{
			Message: related,
		}
	} else {
		o.R.Message = related
	}

	if related.R == nil {
		related.R = &messageR{
			MessageStatuses: MessageStatusSlice{o},
		}
	} else {
		related.R.MessageStatuses = append(related.R.MessageStatuses, o)
	}

	return nil
}

// MessageStatuses retrieves all the records using an executor.
func MessageStatuses(mods ...qm.QueryMod) messageStatusQuery {
	mods = append(mods, qm.From("\"message_status\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"message_status\".*"})
	}

	return messageStatusQuery{q}
}

// FindMessageStatus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMessageStatus(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*MessageStatus, error) {
	messageStatusObj := &MessageStatus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"message_status\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, messageStatusObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from message_status")
	}

	if err = messageStatusObj.doAfterSelectHooks(ctx, exec); err != nil {
		return messageStatusObj, err
	}

	return messageStatusObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MessageStatus) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no message_status provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(messageStatusColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	messageStatusInsertCacheMut.RLock()
	cache, cached := messageStatusInsertCache[key]
	messageStatusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			messageStatusAllColumns,
			messageStatusColumnsWithDefault,
			messageStatusColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(messageStatusType, messageStatusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(messageStatusType, messageStatusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"message_status\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"message_status\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into message_status")
	}

	if !cached {
		messageStatusInsertCacheMut.Lock()
		messageStatusInsertCache[key] = cache
		messageStatusInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MessageStatus.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MessageStatus) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	messageStatusUpdateCacheMut.RLock()
	cache, cached := messageStatusUpdateCache[key]
	messageStatusUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			messageStatusAllColumns,
			messageStatusPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update message_status, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"message_status\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, messageStatusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(messageStatusType, messageStatusMapping, append(wl, messageStatusPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update message_status row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for message_status")
	}

	if !cached {
		messageStatusUpdateCacheMut.Lock()
		messageStatusUpdateCache[key] = cache
		messageStatusUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q messageStatusQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for message_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for message_status")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MessageStatusSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messageStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"message_status\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, messageStatusPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in messageStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all messageStatus")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MessageStatus) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no message_status provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(messageStatusColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	messageStatusUpsertCacheMut.RLock()
	cache, cached := messageStatusUpsertCache[key]
	messageStatusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			messageStatusAllColumns,
			messageStatusColumnsWithDefault,
			messageStatusColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			messageStatusAllColumns,
			messageStatusPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert message_status, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(messageStatusPrimaryKeyColumns))
			copy(conflict, messageStatusPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"message_status\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(messageStatusType, messageStatusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(messageStatusType, messageStatusMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert message_status")
	}

	if !cached {
		messageStatusUpsertCacheMut.Lock()
		messageStatusUpsertCache[key] = cache
		messageStatusUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MessageStatus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MessageStatus) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MessageStatus provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), messageStatusPrimaryKeyMapping)
	sql := "DELETE FROM \"message_status\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from message_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for message_status")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q messageStatusQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no messageStatusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from message_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for message_status")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MessageStatusSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(messageStatusBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messageStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"message_status\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messageStatusPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from messageStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for message_status")
	}

	if len(messageStatusAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *MessageStatus) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMessageStatus(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MessageStatusSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MessageStatusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), messageStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"message_status\".* FROM \"message_status\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, messageStatusPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MessageStatusSlice")
	}

	*o = slice

	return nil
}

// MessageStatusExists checks if the MessageStatus row exists.
func MessageStatusExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"message_status\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if message_status exists")
	}

	return exists, nil
}
