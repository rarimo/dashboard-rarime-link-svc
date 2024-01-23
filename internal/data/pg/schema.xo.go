// Package pg contains generated code for schema 'public'.
package pg

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"

	"github.com/google/uuid"
	"github.com/rarimo/rarime-link-svc/internal/data"
)

// Storage is the helper struct for database operations
type Storage struct {
	db *pgdb.DB
}

// New - returns new instance of storage
func New(db *pgdb.DB) *Storage {
	return &Storage{
		db,
	}
}

// DB - returns db used by Storage
func (s *Storage) DB() *pgdb.DB {
	return s.db
}

// Clone - returns new storage with clone of db
func (s *Storage) Clone() data.Storage {
	return New(s.db.Clone())
}

// Transaction begins a transaction on repo.
func (s *Storage) Transaction(tx func() error) error {
	return s.db.Transaction(tx)
} // GorpMigrationQ represents helper struct to access row of 'gorp_migrations'.
type GorpMigrationQ struct {
	db *pgdb.DB
}

// NewGorpMigrationQ  - creates new instance
func NewGorpMigrationQ(db *pgdb.DB) GorpMigrationQ {
	return GorpMigrationQ{
		db,
	}
}

// GorpMigrationQ  - creates new instance of GorpMigrationQ
func (s Storage) GorpMigrationQ() data.GorpMigrationQ {
	return NewGorpMigrationQ(s.DB())
}

var colsGorpMigration = `id, applied_at`

// InsertCtx inserts a GorpMigration to the database.
func (q GorpMigrationQ) InsertCtx(ctx context.Context, gm *data.GorpMigration) error {
	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO public.gorp_migrations (` +
		`id, applied_at` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, gm.ID, gm.AppliedAt)
	return errors.Wrap(err, "failed to execute insert query")
}

// Insert insert a GorpMigration to the database.
func (q GorpMigrationQ) Insert(gm *data.GorpMigration) error {
	return q.InsertCtx(context.Background(), gm)
}

// UpdateCtx updates a GorpMigration in the database.
func (q GorpMigrationQ) UpdateCtx(ctx context.Context, gm *data.GorpMigration) error {
	// update with composite primary key
	sqlstr := `UPDATE public.gorp_migrations SET ` +
		`applied_at = $1 ` +
		`WHERE id = $2`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, gm.AppliedAt, gm.ID)
	return errors.Wrap(err, "failed to execute update")
}

// Update updates a GorpMigration in the database.
func (q GorpMigrationQ) Update(gm *data.GorpMigration) error {
	return q.UpdateCtx(context.Background(), gm)
}

// UpsertCtx performs an upsert for GorpMigration.
func (q GorpMigrationQ) UpsertCtx(ctx context.Context, gm *data.GorpMigration) error {
	// upsert
	sqlstr := `INSERT INTO public.gorp_migrations (` +
		`id, applied_at` +
		`) VALUES (` +
		`$1, $2` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`applied_at = EXCLUDED.applied_at `
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, gm.ID, gm.AppliedAt); err != nil {
		return errors.Wrap(err, "failed to execute upsert stmt")
	}
	return nil
}

// Upsert performs an upsert for GorpMigration.
func (q GorpMigrationQ) Upsert(gm *data.GorpMigration) error {
	return q.UpsertCtx(context.Background(), gm)
}

// DeleteCtx deletes the GorpMigration from the database.
func (q GorpMigrationQ) DeleteCtx(ctx context.Context, gm *data.GorpMigration) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.gorp_migrations ` +
		`WHERE id = $1`
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, gm.ID); err != nil {
		return errors.Wrap(err, "failed to exec delete stmt")
	}
	return nil
}

// Delete deletes the GorpMigration from the database.
func (q GorpMigrationQ) Delete(gm *data.GorpMigration) error {
	return q.DeleteCtx(context.Background(), gm)
} // LinkQ represents helper struct to access row of 'links'.
type LinkQ struct {
	db *pgdb.DB
}

// NewLinkQ  - creates new instance
func NewLinkQ(db *pgdb.DB) LinkQ {
	return LinkQ{
		db,
	}
}

// LinkQ  - creates new instance of LinkQ
func (s Storage) LinkQ() data.LinkQ {
	return NewLinkQ(s.DB())
}

var colsLink = `id, user_id, created_at, name`

// InsertCtx inserts a Link to the database.
func (q LinkQ) InsertCtx(ctx context.Context, l *data.Link) error {
	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO public.links (` +
		`id, user_id, created_at, name` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, l.ID, l.UserID, l.CreatedAt, l.Name)
	return errors.Wrap(err, "failed to execute insert query")
}

// Insert insert a Link to the database.
func (q LinkQ) Insert(l *data.Link) error {
	return q.InsertCtx(context.Background(), l)
}

// UpdateCtx updates a Link in the database.
func (q LinkQ) UpdateCtx(ctx context.Context, l *data.Link) error {
	// update with composite primary key
	sqlstr := `UPDATE public.links SET ` +
		`user_id = $1, name = $2 ` +
		`WHERE id = $3`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, l.UserID, l.Name, l.ID)
	return errors.Wrap(err, "failed to execute update")
}

// Update updates a Link in the database.
func (q LinkQ) Update(l *data.Link) error {
	return q.UpdateCtx(context.Background(), l)
}

// UpsertCtx performs an upsert for Link.
func (q LinkQ) UpsertCtx(ctx context.Context, l *data.Link) error {
	// upsert
	sqlstr := `INSERT INTO public.links (` +
		`id, user_id, created_at, name` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`user_id = EXCLUDED.user_id, name = EXCLUDED.name `
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, l.ID, l.UserID, l.CreatedAt, l.Name); err != nil {
		return errors.Wrap(err, "failed to execute upsert stmt")
	}
	return nil
}

// Upsert performs an upsert for Link.
func (q LinkQ) Upsert(l *data.Link) error {
	return q.UpsertCtx(context.Background(), l)
}

// DeleteCtx deletes the Link from the database.
func (q LinkQ) DeleteCtx(ctx context.Context, l *data.Link) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.links ` +
		`WHERE id = $1`
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, l.ID); err != nil {
		return errors.Wrap(err, "failed to exec delete stmt")
	}
	return nil
}

// Delete deletes the Link from the database.
func (q LinkQ) Delete(l *data.Link) error {
	return q.DeleteCtx(context.Background(), l)
} // LinksToProofQ represents helper struct to access row of 'links_to_proofs'.
type LinksToProofQ struct {
	db *pgdb.DB
}

// NewLinksToProofQ  - creates new instance
func NewLinksToProofQ(db *pgdb.DB) LinksToProofQ {
	return LinksToProofQ{
		db,
	}
}

// LinksToProofQ  - creates new instance of LinksToProofQ
func (s Storage) LinksToProofQ() data.LinksToProofQ {
	return NewLinksToProofQ(s.DB())
}

var colsLinksToProof = `link_id, proof_id`

// InsertCtx inserts a LinksToProof to the database.
func (q LinksToProofQ) InsertCtx(ctx context.Context, ltp *data.LinksToProof) error {
	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO public.links_to_proofs (` +
		`link_id, proof_id` +
		`) VALUES (` +
		`$1, $2` +
		`)`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, ltp.LinkID, ltp.ProofID)
	return errors.Wrap(err, "failed to execute insert query")
}

// Insert insert a LinksToProof to the database.
func (q LinksToProofQ) Insert(ltp *data.LinksToProof) error {
	return q.InsertCtx(context.Background(), ltp)
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// DeleteCtx deletes the LinksToProof from the database.
func (q LinksToProofQ) DeleteCtx(ctx context.Context, ltp *data.LinksToProof) error {
	// delete with composite primary key
	sqlstr := `DELETE FROM public.links_to_proofs ` +
		`WHERE link_id = $1 AND proof_id = $2`
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, ltp.LinkID, ltp.ProofID); err != nil {
		return errors.Wrap(err, "failed to exec delete stmt")
	}
	return nil
}

// Delete deletes the LinksToProof from the database.
func (q LinksToProofQ) Delete(ltp *data.LinksToProof) error {
	return q.DeleteCtx(context.Background(), ltp)
} // ProofQ represents helper struct to access row of 'proofs'.
type ProofQ struct {
	db *pgdb.DB
}

// NewProofQ  - creates new instance
func NewProofQ(db *pgdb.DB) ProofQ {
	return ProofQ{
		db,
	}
}

// ProofQ  - creates new instance of ProofQ
func (s Storage) ProofQ() data.ProofQ {
	return NewProofQ(s.DB())
}

var colsProof = `id, creator, created_at, proof, org_id, type, schema_url`

// InsertCtx inserts a Proof to the database.
func (q ProofQ) InsertCtx(ctx context.Context, p *data.Proof) error {
	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO public.proofs (` +
		`id, creator, created_at, proof, org_id, type, schema_url` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, p.ID, p.Creator, p.CreatedAt, p.Proof, p.OrgID, p.Type, p.SchemaURL)
	return errors.Wrap(err, "failed to execute insert query")
}

// Insert insert a Proof to the database.
func (q ProofQ) Insert(p *data.Proof) error {
	return q.InsertCtx(context.Background(), p)
}

// UpdateCtx updates a Proof in the database.
func (q ProofQ) UpdateCtx(ctx context.Context, p *data.Proof) error {
	// update with composite primary key
	sqlstr := `UPDATE public.proofs SET ` +
		`creator = $1, proof = $2, org_id = $3, type = $4, schema_url = $5 ` +
		`WHERE id = $6`
	// run
	err := q.db.ExecRawContext(ctx, sqlstr, p.Creator, p.Proof, p.OrgID, p.Type, p.SchemaURL, p.ID)
	return errors.Wrap(err, "failed to execute update")
}

// Update updates a Proof in the database.
func (q ProofQ) Update(p *data.Proof) error {
	return q.UpdateCtx(context.Background(), p)
}

// UpsertCtx performs an upsert for Proof.
func (q ProofQ) UpsertCtx(ctx context.Context, p *data.Proof) error {
	// upsert
	sqlstr := `INSERT INTO public.proofs (` +
		`id, creator, created_at, proof, org_id, type, schema_url` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`creator = EXCLUDED.creator, proof = EXCLUDED.proof, org_id = EXCLUDED.org_id, type = EXCLUDED.type, schema_url = EXCLUDED.schema_url `
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, p.ID, p.Creator, p.CreatedAt, p.Proof, p.OrgID, p.Type, p.SchemaURL); err != nil {
		return errors.Wrap(err, "failed to execute upsert stmt")
	}
	return nil
}

// Upsert performs an upsert for Proof.
func (q ProofQ) Upsert(p *data.Proof) error {
	return q.UpsertCtx(context.Background(), p)
}

// DeleteCtx deletes the Proof from the database.
func (q ProofQ) DeleteCtx(ctx context.Context, p *data.Proof) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.proofs ` +
		`WHERE id = $1`
	// run
	if err := q.db.ExecRawContext(ctx, sqlstr, p.ID); err != nil {
		return errors.Wrap(err, "failed to exec delete stmt")
	}
	return nil
}

// Delete deletes the Proof from the database.
func (q ProofQ) Delete(p *data.Proof) error {
	return q.DeleteCtx(context.Background(), p)
}

// GorpMigrationByIDCtx retrieves a row from 'public.gorp_migrations' as a GorpMigration.
//
// Generated from index 'gorp_migrations_pkey'.
func (q GorpMigrationQ) GorpMigrationByIDCtx(ctx context.Context, id string, isForUpdate bool) (*data.GorpMigration, error) {
	// query
	sqlstr := `SELECT ` +
		`id, applied_at ` +
		`FROM public.gorp_migrations ` +
		`WHERE id = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.GorpMigration
	err := q.db.GetRawContext(ctx, &res, sqlstr, id)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// GorpMigrationByID retrieves a row from 'public.gorp_migrations' as a GorpMigration.
//
// Generated from index 'gorp_migrations_pkey'.
func (q GorpMigrationQ) GorpMigrationByID(id string, isForUpdate bool) (*data.GorpMigration, error) {
	return q.GorpMigrationByIDCtx(context.Background(), id, isForUpdate)
}

// LinkByNameCtx retrieves a row from 'public.links' as a Link.
//
// Generated from index 'links_name_key'.
func (q LinkQ) LinkByNameCtx(ctx context.Context, name sql.NullString, isForUpdate bool) (*data.Link, error) {
	// query
	sqlstr := `SELECT ` +
		`id, user_id, created_at, name ` +
		`FROM public.links ` +
		`WHERE name = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.Link
	err := q.db.GetRawContext(ctx, &res, sqlstr, name)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// LinkByName retrieves a row from 'public.links' as a Link.
//
// Generated from index 'links_name_key'.
func (q LinkQ) LinkByName(name sql.NullString, isForUpdate bool) (*data.Link, error) {
	return q.LinkByNameCtx(context.Background(), name, isForUpdate)
}

// LinkByIDCtx retrieves a row from 'public.links' as a Link.
//
// Generated from index 'links_pkey'.
func (q LinkQ) LinkByIDCtx(ctx context.Context, id uuid.UUID, isForUpdate bool) (*data.Link, error) {
	// query
	sqlstr := `SELECT ` +
		`id, user_id, created_at, name ` +
		`FROM public.links ` +
		`WHERE id = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.Link
	err := q.db.GetRawContext(ctx, &res, sqlstr, id)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// LinkByID retrieves a row from 'public.links' as a Link.
//
// Generated from index 'links_pkey'.
func (q LinkQ) LinkByID(id uuid.UUID, isForUpdate bool) (*data.Link, error) {
	return q.LinkByIDCtx(context.Background(), id, isForUpdate)
}

// LinksToProofByLinkIDProofIDCtx retrieves a row from 'public.links_to_proofs' as a LinksToProof.
//
// Generated from index 'links_to_proofs_pkey'.
func (q LinksToProofQ) LinksToProofByLinkIDProofIDCtx(ctx context.Context, linkID, proofID uuid.UUID, isForUpdate bool) (*data.LinksToProof, error) {
	// query
	sqlstr := `SELECT ` +
		`link_id, proof_id ` +
		`FROM public.links_to_proofs ` +
		`WHERE link_id = $1 AND proof_id = $2`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.LinksToProof
	err := q.db.GetRawContext(ctx, &res, sqlstr, linkID, proofID)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// LinksToProofByLinkIDProofID retrieves a row from 'public.links_to_proofs' as a LinksToProof.
//
// Generated from index 'links_to_proofs_pkey'.
func (q LinksToProofQ) LinksToProofByLinkIDProofID(linkID, proofID uuid.UUID, isForUpdate bool) (*data.LinksToProof, error) {
	return q.LinksToProofByLinkIDProofIDCtx(context.Background(), linkID, proofID, isForUpdate)
}

// ProofByIDCtx retrieves a row from 'public.proofs' as a Proof.
//
// Generated from index 'proofs_pkey'.
func (q ProofQ) ProofByIDCtx(ctx context.Context, id uuid.UUID, isForUpdate bool) (*data.Proof, error) {
	// query
	sqlstr := `SELECT ` +
		`id, creator, created_at, proof, org_id, type, schema_url ` +
		`FROM public.proofs ` +
		`WHERE id = $1`
	// run
	if isForUpdate {
		sqlstr += " for update"
	}
	var res data.Proof
	err := q.db.GetRawContext(ctx, &res, sqlstr, id)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to exec select")
	}

	return &res, nil
}

// ProofByID retrieves a row from 'public.proofs' as a Proof.
//
// Generated from index 'proofs_pkey'.
func (q ProofQ) ProofByID(id uuid.UUID, isForUpdate bool) (*data.Proof, error) {
	return q.ProofByIDCtx(context.Background(), id, isForUpdate)
}
