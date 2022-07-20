// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ExerciseDate is an object representing the database table.
type ExerciseDate struct {
	ID             int      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID         int      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	ExerciseDate   int      `boil:"exercise_date" json:"exercise_date" toml:"exercise_date" yaml:"exercise_date"`
	ExerciseDateEn string   `boil:"exercise_date_en" json:"exercise_date_en" toml:"exercise_date_en" yaml:"exercise_date_en"`
	ExerciseTime   null.Int `boil:"exercise_time" json:"exercise_time,omitempty" toml:"exercise_time" yaml:"exercise_time,omitempty"`

	R *exerciseDateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L exerciseDateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ExerciseDateColumns = struct {
	ID             string
	UserID         string
	ExerciseDate   string
	ExerciseDateEn string
	ExerciseTime   string
}{
	ID:             "id",
	UserID:         "user_id",
	ExerciseDate:   "exercise_date",
	ExerciseDateEn: "exercise_date_en",
	ExerciseTime:   "exercise_time",
}

var ExerciseDateTableColumns = struct {
	ID             string
	UserID         string
	ExerciseDate   string
	ExerciseDateEn string
	ExerciseTime   string
}{
	ID:             "exercise_dates.id",
	UserID:         "exercise_dates.user_id",
	ExerciseDate:   "exercise_dates.exercise_date",
	ExerciseDateEn: "exercise_dates.exercise_date_en",
	ExerciseTime:   "exercise_dates.exercise_time",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var ExerciseDateWhere = struct {
	ID             whereHelperint
	UserID         whereHelperint
	ExerciseDate   whereHelperint
	ExerciseDateEn whereHelperstring
	ExerciseTime   whereHelpernull_Int
}{
	ID:             whereHelperint{field: "`exercise_dates`.`id`"},
	UserID:         whereHelperint{field: "`exercise_dates`.`user_id`"},
	ExerciseDate:   whereHelperint{field: "`exercise_dates`.`exercise_date`"},
	ExerciseDateEn: whereHelperstring{field: "`exercise_dates`.`exercise_date_en`"},
	ExerciseTime:   whereHelpernull_Int{field: "`exercise_dates`.`exercise_time`"},
}

// ExerciseDateRels is where relationship names are stored.
var ExerciseDateRels = struct {
}{}

// exerciseDateR is where relationships are stored.
type exerciseDateR struct {
}

// NewStruct creates a new relationship struct
func (*exerciseDateR) NewStruct() *exerciseDateR {
	return &exerciseDateR{}
}

// exerciseDateL is where Load methods for each relationship are stored.
type exerciseDateL struct{}

var (
	exerciseDateAllColumns            = []string{"id", "user_id", "exercise_date", "exercise_date_en", "exercise_time"}
	exerciseDateColumnsWithoutDefault = []string{"user_id", "exercise_date", "exercise_date_en", "exercise_time"}
	exerciseDateColumnsWithDefault    = []string{"id"}
	exerciseDatePrimaryKeyColumns     = []string{"id"}
	exerciseDateGeneratedColumns      = []string{}
)

type (
	// ExerciseDateSlice is an alias for a slice of pointers to ExerciseDate.
	// This should almost always be used instead of []ExerciseDate.
	ExerciseDateSlice []*ExerciseDate
	// ExerciseDateHook is the signature for custom ExerciseDate hook methods
	ExerciseDateHook func(context.Context, boil.ContextExecutor, *ExerciseDate) error

	exerciseDateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	exerciseDateType                 = reflect.TypeOf(&ExerciseDate{})
	exerciseDateMapping              = queries.MakeStructMapping(exerciseDateType)
	exerciseDatePrimaryKeyMapping, _ = queries.BindMapping(exerciseDateType, exerciseDateMapping, exerciseDatePrimaryKeyColumns)
	exerciseDateInsertCacheMut       sync.RWMutex
	exerciseDateInsertCache          = make(map[string]insertCache)
	exerciseDateUpdateCacheMut       sync.RWMutex
	exerciseDateUpdateCache          = make(map[string]updateCache)
	exerciseDateUpsertCacheMut       sync.RWMutex
	exerciseDateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var exerciseDateAfterSelectHooks []ExerciseDateHook

var exerciseDateBeforeInsertHooks []ExerciseDateHook
var exerciseDateAfterInsertHooks []ExerciseDateHook

var exerciseDateBeforeUpdateHooks []ExerciseDateHook
var exerciseDateAfterUpdateHooks []ExerciseDateHook

var exerciseDateBeforeDeleteHooks []ExerciseDateHook
var exerciseDateAfterDeleteHooks []ExerciseDateHook

var exerciseDateBeforeUpsertHooks []ExerciseDateHook
var exerciseDateAfterUpsertHooks []ExerciseDateHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ExerciseDate) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseDateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ExerciseDate) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseDateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ExerciseDate) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseDateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ExerciseDate) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseDateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ExerciseDate) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseDateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ExerciseDate) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseDateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ExerciseDate) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseDateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ExerciseDate) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseDateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ExerciseDate) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range exerciseDateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddExerciseDateHook registers your hook function for all future operations.
func AddExerciseDateHook(hookPoint boil.HookPoint, exerciseDateHook ExerciseDateHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		exerciseDateAfterSelectHooks = append(exerciseDateAfterSelectHooks, exerciseDateHook)
	case boil.BeforeInsertHook:
		exerciseDateBeforeInsertHooks = append(exerciseDateBeforeInsertHooks, exerciseDateHook)
	case boil.AfterInsertHook:
		exerciseDateAfterInsertHooks = append(exerciseDateAfterInsertHooks, exerciseDateHook)
	case boil.BeforeUpdateHook:
		exerciseDateBeforeUpdateHooks = append(exerciseDateBeforeUpdateHooks, exerciseDateHook)
	case boil.AfterUpdateHook:
		exerciseDateAfterUpdateHooks = append(exerciseDateAfterUpdateHooks, exerciseDateHook)
	case boil.BeforeDeleteHook:
		exerciseDateBeforeDeleteHooks = append(exerciseDateBeforeDeleteHooks, exerciseDateHook)
	case boil.AfterDeleteHook:
		exerciseDateAfterDeleteHooks = append(exerciseDateAfterDeleteHooks, exerciseDateHook)
	case boil.BeforeUpsertHook:
		exerciseDateBeforeUpsertHooks = append(exerciseDateBeforeUpsertHooks, exerciseDateHook)
	case boil.AfterUpsertHook:
		exerciseDateAfterUpsertHooks = append(exerciseDateAfterUpsertHooks, exerciseDateHook)
	}
}

// One returns a single exerciseDate record from the query.
func (q exerciseDateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ExerciseDate, error) {
	o := &ExerciseDate{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for exercise_dates")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ExerciseDate records from the query.
func (q exerciseDateQuery) All(ctx context.Context, exec boil.ContextExecutor) (ExerciseDateSlice, error) {
	var o []*ExerciseDate

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ExerciseDate slice")
	}

	if len(exerciseDateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ExerciseDate records in the query.
func (q exerciseDateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count exercise_dates rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q exerciseDateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if exercise_dates exists")
	}

	return count > 0, nil
}

// ExerciseDates retrieves all the records using an executor.
func ExerciseDates(mods ...qm.QueryMod) exerciseDateQuery {
	mods = append(mods, qm.From("`exercise_dates`"))
	return exerciseDateQuery{NewQuery(mods...)}
}

// FindExerciseDate retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindExerciseDate(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ExerciseDate, error) {
	exerciseDateObj := &ExerciseDate{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `exercise_dates` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, exerciseDateObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from exercise_dates")
	}

	if err = exerciseDateObj.doAfterSelectHooks(ctx, exec); err != nil {
		return exerciseDateObj, err
	}

	return exerciseDateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ExerciseDate) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no exercise_dates provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(exerciseDateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	exerciseDateInsertCacheMut.RLock()
	cache, cached := exerciseDateInsertCache[key]
	exerciseDateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			exerciseDateAllColumns,
			exerciseDateColumnsWithDefault,
			exerciseDateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(exerciseDateType, exerciseDateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(exerciseDateType, exerciseDateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `exercise_dates` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `exercise_dates` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `exercise_dates` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, exerciseDatePrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into exercise_dates")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == exerciseDateMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for exercise_dates")
	}

CacheNoHooks:
	if !cached {
		exerciseDateInsertCacheMut.Lock()
		exerciseDateInsertCache[key] = cache
		exerciseDateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ExerciseDate.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ExerciseDate) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	exerciseDateUpdateCacheMut.RLock()
	cache, cached := exerciseDateUpdateCache[key]
	exerciseDateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			exerciseDateAllColumns,
			exerciseDatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update exercise_dates, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `exercise_dates` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, exerciseDatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(exerciseDateType, exerciseDateMapping, append(wl, exerciseDatePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update exercise_dates row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for exercise_dates")
	}

	if !cached {
		exerciseDateUpdateCacheMut.Lock()
		exerciseDateUpdateCache[key] = cache
		exerciseDateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q exerciseDateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for exercise_dates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for exercise_dates")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ExerciseDateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exerciseDatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `exercise_dates` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, exerciseDatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in exerciseDate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all exerciseDate")
	}
	return rowsAff, nil
}

var mySQLExerciseDateUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ExerciseDate) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no exercise_dates provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(exerciseDateColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLExerciseDateUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	exerciseDateUpsertCacheMut.RLock()
	cache, cached := exerciseDateUpsertCache[key]
	exerciseDateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			exerciseDateAllColumns,
			exerciseDateColumnsWithDefault,
			exerciseDateColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			exerciseDateAllColumns,
			exerciseDatePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert exercise_dates, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`exercise_dates`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `exercise_dates` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(exerciseDateType, exerciseDateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(exerciseDateType, exerciseDateMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for exercise_dates")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == exerciseDateMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(exerciseDateType, exerciseDateMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for exercise_dates")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for exercise_dates")
	}

CacheNoHooks:
	if !cached {
		exerciseDateUpsertCacheMut.Lock()
		exerciseDateUpsertCache[key] = cache
		exerciseDateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ExerciseDate record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ExerciseDate) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ExerciseDate provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), exerciseDatePrimaryKeyMapping)
	sql := "DELETE FROM `exercise_dates` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from exercise_dates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for exercise_dates")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q exerciseDateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no exerciseDateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from exercise_dates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for exercise_dates")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ExerciseDateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(exerciseDateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exerciseDatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `exercise_dates` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, exerciseDatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from exerciseDate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for exercise_dates")
	}

	if len(exerciseDateAfterDeleteHooks) != 0 {
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
func (o *ExerciseDate) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindExerciseDate(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ExerciseDateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ExerciseDateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), exerciseDatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `exercise_dates`.* FROM `exercise_dates` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, exerciseDatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ExerciseDateSlice")
	}

	*o = slice

	return nil
}

// ExerciseDateExists checks if the ExerciseDate row exists.
func ExerciseDateExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `exercise_dates` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if exercise_dates exists")
	}

	return exists, nil
}
