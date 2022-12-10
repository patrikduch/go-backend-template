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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Movie is an object representing the database table.
type Movie struct {
	ID          int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title       null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	ReleaseDate null.Time   `boil:"release_date" json:"release_date,omitempty" toml:"release_date" yaml:"release_date,omitempty"`
	Runtime     null.Int    `boil:"runtime" json:"runtime,omitempty" toml:"runtime" yaml:"runtime,omitempty"`
	MpaaRating  null.String `boil:"mpaa_rating" json:"mpaa_rating,omitempty" toml:"mpaa_rating" yaml:"mpaa_rating,omitempty"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Image       null.String `boil:"image" json:"image,omitempty" toml:"image" yaml:"image,omitempty"`
	CreatedAt   null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt   null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *movieR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L movieL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MovieColumns = struct {
	ID          string
	Title       string
	ReleaseDate string
	Runtime     string
	MpaaRating  string
	Description string
	Image       string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Title:       "title",
	ReleaseDate: "release_date",
	Runtime:     "runtime",
	MpaaRating:  "mpaa_rating",
	Description: "description",
	Image:       "image",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var MovieTableColumns = struct {
	ID          string
	Title       string
	ReleaseDate string
	Runtime     string
	MpaaRating  string
	Description string
	Image       string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "movies.id",
	Title:       "movies.title",
	ReleaseDate: "movies.release_date",
	Runtime:     "movies.runtime",
	MpaaRating:  "movies.mpaa_rating",
	Description: "movies.description",
	Image:       "movies.image",
	CreatedAt:   "movies.created_at",
	UpdatedAt:   "movies.updated_at",
}

// Generated where

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
func (w whereHelpernull_Int) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var MovieWhere = struct {
	ID          whereHelperint
	Title       whereHelpernull_String
	ReleaseDate whereHelpernull_Time
	Runtime     whereHelpernull_Int
	MpaaRating  whereHelpernull_String
	Description whereHelpernull_String
	Image       whereHelpernull_String
	CreatedAt   whereHelpernull_Time
	UpdatedAt   whereHelpernull_Time
}{
	ID:          whereHelperint{field: "\"movies\".\"id\""},
	Title:       whereHelpernull_String{field: "\"movies\".\"title\""},
	ReleaseDate: whereHelpernull_Time{field: "\"movies\".\"release_date\""},
	Runtime:     whereHelpernull_Int{field: "\"movies\".\"runtime\""},
	MpaaRating:  whereHelpernull_String{field: "\"movies\".\"mpaa_rating\""},
	Description: whereHelpernull_String{field: "\"movies\".\"description\""},
	Image:       whereHelpernull_String{field: "\"movies\".\"image\""},
	CreatedAt:   whereHelpernull_Time{field: "\"movies\".\"created_at\""},
	UpdatedAt:   whereHelpernull_Time{field: "\"movies\".\"updated_at\""},
}

// MovieRels is where relationship names are stored.
var MovieRels = struct {
	MoviesGenres string
}{
	MoviesGenres: "MoviesGenres",
}

// movieR is where relationships are stored.
type movieR struct {
	MoviesGenres MoviesGenreSlice `boil:"MoviesGenres" json:"MoviesGenres" toml:"MoviesGenres" yaml:"MoviesGenres"`
}

// NewStruct creates a new relationship struct
func (*movieR) NewStruct() *movieR {
	return &movieR{}
}

func (r *movieR) GetMoviesGenres() MoviesGenreSlice {
	if r == nil {
		return nil
	}
	return r.MoviesGenres
}

// movieL is where Load methods for each relationship are stored.
type movieL struct{}

var (
	movieAllColumns            = []string{"id", "title", "release_date", "runtime", "mpaa_rating", "description", "image", "created_at", "updated_at"}
	movieColumnsWithoutDefault = []string{}
	movieColumnsWithDefault    = []string{"id", "title", "release_date", "runtime", "mpaa_rating", "description", "image", "created_at", "updated_at"}
	moviePrimaryKeyColumns     = []string{"id"}
	movieGeneratedColumns      = []string{"id"}
)

type (
	// MovieSlice is an alias for a slice of pointers to Movie.
	// This should almost always be used instead of []Movie.
	MovieSlice []*Movie
	// MovieHook is the signature for custom Movie hook methods
	MovieHook func(context.Context, boil.ContextExecutor, *Movie) error

	movieQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	movieType                 = reflect.TypeOf(&Movie{})
	movieMapping              = queries.MakeStructMapping(movieType)
	moviePrimaryKeyMapping, _ = queries.BindMapping(movieType, movieMapping, moviePrimaryKeyColumns)
	movieInsertCacheMut       sync.RWMutex
	movieInsertCache          = make(map[string]insertCache)
	movieUpdateCacheMut       sync.RWMutex
	movieUpdateCache          = make(map[string]updateCache)
	movieUpsertCacheMut       sync.RWMutex
	movieUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var movieAfterSelectHooks []MovieHook

var movieBeforeInsertHooks []MovieHook
var movieAfterInsertHooks []MovieHook

var movieBeforeUpdateHooks []MovieHook
var movieAfterUpdateHooks []MovieHook

var movieBeforeDeleteHooks []MovieHook
var movieAfterDeleteHooks []MovieHook

var movieBeforeUpsertHooks []MovieHook
var movieAfterUpsertHooks []MovieHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Movie) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Movie) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Movie) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Movie) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Movie) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Movie) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Movie) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Movie) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Movie) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range movieAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMovieHook registers your hook function for all future operations.
func AddMovieHook(hookPoint boil.HookPoint, movieHook MovieHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		movieAfterSelectHooks = append(movieAfterSelectHooks, movieHook)
	case boil.BeforeInsertHook:
		movieBeforeInsertHooks = append(movieBeforeInsertHooks, movieHook)
	case boil.AfterInsertHook:
		movieAfterInsertHooks = append(movieAfterInsertHooks, movieHook)
	case boil.BeforeUpdateHook:
		movieBeforeUpdateHooks = append(movieBeforeUpdateHooks, movieHook)
	case boil.AfterUpdateHook:
		movieAfterUpdateHooks = append(movieAfterUpdateHooks, movieHook)
	case boil.BeforeDeleteHook:
		movieBeforeDeleteHooks = append(movieBeforeDeleteHooks, movieHook)
	case boil.AfterDeleteHook:
		movieAfterDeleteHooks = append(movieAfterDeleteHooks, movieHook)
	case boil.BeforeUpsertHook:
		movieBeforeUpsertHooks = append(movieBeforeUpsertHooks, movieHook)
	case boil.AfterUpsertHook:
		movieAfterUpsertHooks = append(movieAfterUpsertHooks, movieHook)
	}
}

// One returns a single movie record from the query.
func (q movieQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Movie, error) {
	o := &Movie{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for movies")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Movie records from the query.
func (q movieQuery) All(ctx context.Context, exec boil.ContextExecutor) (MovieSlice, error) {
	var o []*Movie

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Movie slice")
	}

	if len(movieAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Movie records in the query.
func (q movieQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count movies rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q movieQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if movies exists")
	}

	return count > 0, nil
}

// MoviesGenres retrieves all the movies_genre's MoviesGenres with an executor.
func (o *Movie) MoviesGenres(mods ...qm.QueryMod) moviesGenreQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"movies_genres\".\"movie_id\"=?", o.ID),
	)

	return MoviesGenres(queryMods...)
}

// LoadMoviesGenres allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (movieL) LoadMoviesGenres(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMovie interface{}, mods queries.Applicator) error {
	var slice []*Movie
	var object *Movie

	if singular {
		var ok bool
		object, ok = maybeMovie.(*Movie)
		if !ok {
			object = new(Movie)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeMovie)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeMovie))
			}
		}
	} else {
		s, ok := maybeMovie.(*[]*Movie)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeMovie)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeMovie))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &movieR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &movieR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`movies_genres`),
		qm.WhereIn(`movies_genres.movie_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load movies_genres")
	}

	var resultSlice []*MoviesGenre
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice movies_genres")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on movies_genres")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for movies_genres")
	}

	if len(moviesGenreAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.MoviesGenres = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &moviesGenreR{}
			}
			foreign.R.Movie = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.MovieID) {
				local.R.MoviesGenres = append(local.R.MoviesGenres, foreign)
				if foreign.R == nil {
					foreign.R = &moviesGenreR{}
				}
				foreign.R.Movie = local
				break
			}
		}
	}

	return nil
}

// AddMoviesGenres adds the given related objects to the existing relationships
// of the movie, optionally inserting them as new records.
// Appends related to o.R.MoviesGenres.
// Sets related.R.Movie appropriately.
func (o *Movie) AddMoviesGenres(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*MoviesGenre) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.MovieID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"movies_genres\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"movie_id"}),
				strmangle.WhereClause("\"", "\"", 2, moviesGenrePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.MovieID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &movieR{
			MoviesGenres: related,
		}
	} else {
		o.R.MoviesGenres = append(o.R.MoviesGenres, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &moviesGenreR{
				Movie: o,
			}
		} else {
			rel.R.Movie = o
		}
	}
	return nil
}

// SetMoviesGenres removes all previously related items of the
// movie replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Movie's MoviesGenres accordingly.
// Replaces o.R.MoviesGenres with related.
// Sets related.R.Movie's MoviesGenres accordingly.
func (o *Movie) SetMoviesGenres(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*MoviesGenre) error {
	query := "update \"movies_genres\" set \"movie_id\" = null where \"movie_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.MoviesGenres {
			queries.SetScanner(&rel.MovieID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Movie = nil
		}
		o.R.MoviesGenres = nil
	}

	return o.AddMoviesGenres(ctx, exec, insert, related...)
}

// RemoveMoviesGenres relationships from objects passed in.
// Removes related items from R.MoviesGenres (uses pointer comparison, removal does not keep order)
// Sets related.R.Movie.
func (o *Movie) RemoveMoviesGenres(ctx context.Context, exec boil.ContextExecutor, related ...*MoviesGenre) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.MovieID, nil)
		if rel.R != nil {
			rel.R.Movie = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("movie_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.MoviesGenres {
			if rel != ri {
				continue
			}

			ln := len(o.R.MoviesGenres)
			if ln > 1 && i < ln-1 {
				o.R.MoviesGenres[i] = o.R.MoviesGenres[ln-1]
			}
			o.R.MoviesGenres = o.R.MoviesGenres[:ln-1]
			break
		}
	}

	return nil
}

// Movies retrieves all the records using an executor.
func Movies(mods ...qm.QueryMod) movieQuery {
	mods = append(mods, qm.From("\"movies\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"movies\".*"})
	}

	return movieQuery{q}
}

// FindMovie retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMovie(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Movie, error) {
	movieObj := &Movie{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"movies\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, movieObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from movies")
	}

	if err = movieObj.doAfterSelectHooks(ctx, exec); err != nil {
		return movieObj, err
	}

	return movieObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Movie) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no movies provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(movieColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	movieInsertCacheMut.RLock()
	cache, cached := movieInsertCache[key]
	movieInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			movieAllColumns,
			movieColumnsWithDefault,
			movieColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, movieGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(movieType, movieMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(movieType, movieMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"movies\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"movies\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into movies")
	}

	if !cached {
		movieInsertCacheMut.Lock()
		movieInsertCache[key] = cache
		movieInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Movie.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Movie) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	movieUpdateCacheMut.RLock()
	cache, cached := movieUpdateCache[key]
	movieUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			movieAllColumns,
			moviePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, movieGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update movies, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"movies\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, moviePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(movieType, movieMapping, append(wl, moviePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update movies row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for movies")
	}

	if !cached {
		movieUpdateCacheMut.Lock()
		movieUpdateCache[key] = cache
		movieUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q movieQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for movies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for movies")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MovieSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), moviePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"movies\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, moviePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in movie slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all movie")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Movie) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no movies provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(movieColumnsWithDefault, o)

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

	movieUpsertCacheMut.RLock()
	cache, cached := movieUpsertCache[key]
	movieUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			movieAllColumns,
			movieColumnsWithDefault,
			movieColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			movieAllColumns,
			moviePrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, movieGeneratedColumns)
		update = strmangle.SetComplement(update, movieGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert movies, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(moviePrimaryKeyColumns))
			copy(conflict, moviePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"movies\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(movieType, movieMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(movieType, movieMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert movies")
	}

	if !cached {
		movieUpsertCacheMut.Lock()
		movieUpsertCache[key] = cache
		movieUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Movie record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Movie) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Movie provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), moviePrimaryKeyMapping)
	sql := "DELETE FROM \"movies\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from movies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for movies")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q movieQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no movieQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from movies")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for movies")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MovieSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(movieBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), moviePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"movies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, moviePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from movie slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for movies")
	}

	if len(movieAfterDeleteHooks) != 0 {
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
func (o *Movie) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMovie(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MovieSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MovieSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), moviePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"movies\".* FROM \"movies\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, moviePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MovieSlice")
	}

	*o = slice

	return nil
}

// MovieExists checks if the Movie row exists.
func MovieExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"movies\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if movies exists")
	}

	return exists, nil
}
