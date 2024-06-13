package postgres

import (
	"context"
	"fmt"
	rDomain "github.com/Mx1q/ppo_repo/domain"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

const PageSize = 30

type measurementRepository struct {
	db *pgxpool.Pool
}

func NewMeasrementRepository(db *pgxpool.Pool) rDomain.IMeasurementRepository {
	return &measurementRepository{
		db: db,
	}
}

func (r *measurementRepository) Create(ctx context.Context, measurement *domain.Measurement) error {
	query := `insert into saladRecipes.measurement(name, grams)
		values ($1, $2)`

	_, err := r.db.Exec(
		ctx,
		query,
		measurement.Name,
		measurement.Grams)
	if err != nil {
		return fmt.Errorf("creating measurement: %w", err)
	}
	return nil
}

func (r *measurementRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Measurement, error) {
	query := `select id, name, grams
		from saladRecipes.measurement
		where id = $1`

	measurement := new(domain.Measurement)
	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&measurement.ID,
		&measurement.Name,
		&measurement.Grams)

	if err != nil {
		return nil, fmt.Errorf("getting measurement by id: %w", err)
	}
	return measurement, nil
}

func (r *measurementRepository) GetByRecipeId(ctx context.Context,
	ingredientId uuid.UUID, recipeId uuid.UUID) (*domain.Measurement, int, error) {
	query := `with mId as (
    	select measurement as id, amount
    	from saladRecipes.recipeIngredient
    	where recipeId = $1 and ingredientId = $2
	)
	select id, name, grams, (select amount from mId)
	from saladRecipes.measurement
	where id = (select id from mId)`

	measurement := new(domain.Measurement)
	amount := 0
	err := r.db.QueryRow(
		ctx,
		query,
		recipeId,     // uuidToString(recipeId),
		ingredientId, // uuidToString(ingredientId),
	).Scan(
		&measurement.ID,
		&measurement.Name,
		&measurement.Grams,
		&amount)

	if err != nil {
		return nil, 0, fmt.Errorf("getting measurement by recipe and ingredient: %w", err)
	}
	return measurement, amount, nil
}

func (r *measurementRepository) GetAll(ctx context.Context) ([]*domain.Measurement, error) {
	query := `select id, name, grams
		from saladRecipes.measurement`

	rows, err := r.db.Query(
		ctx,
		query,
	)

	measurements := make([]*domain.Measurement, 0)
	for rows.Next() {
		tmp := new(domain.Measurement)
		err = rows.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Grams,
		)
		measurements = append(measurements, tmp)
		if err != nil {
			return nil, fmt.Errorf("scanning measurements: %w", err)
		}
	}
	return measurements, nil
}

func (r *measurementRepository) Update(ctx context.Context, measurement *domain.Measurement) error {
	query := `update saladRecipes.measurement
		set
			name = $1,
			grams = $2
		where id = $3`

	_, err := r.db.Exec(
		ctx,
		query,
		measurement.Name,
		measurement.Grams,
		measurement.ID,
	)
	if err != nil {
		return fmt.Errorf("updating measurement: %w", err)
	}
	return nil
}

func (r *measurementRepository) DeleteById(ctx context.Context, id uuid.UUID) error {
	query := `delete from saladRecipes.measurement
       where id = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		id)
	if err != nil {
		return fmt.Errorf("deleting measurement by id: %w", err)
	}
	return nil
}

func (r *measurementRepository) UpdateLink(ctx context.Context, linkId uuid.UUID, measurementId uuid.UUID, amount int) error {
	query := `update saladRecipes.recipeIngredient
	set
    	measurement = $1,
    	amount = $2
	where id = $3`

	_, err := r.db.Exec(
		ctx,
		query,
		measurementId,
		amount,
		linkId)
	if err != nil {
		return fmt.Errorf("updating measurement: %w", err)
	}
	return nil
}
