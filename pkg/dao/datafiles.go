package dao

import (
	"context"
	"time"

	"github.com/narendra121/data-dir-db/pkg/model"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllDatafiles is a function to get a slice of record(s) from datafiles table in the postgres database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllDatafiles(ctx context.Context, page, pagesize int, order string) (results []*model.Datafiles, totalRows int64, err error) {

	resultOrm := DB.Model(&model.Datafiles{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetDatafiles is a function to get a single record from the datafiles table in the postgres database
// error - ErrNotFound, db Find error
func GetDatafiles(ctx context.Context, argID int32) (record *model.Datafiles, err error) {
	record = &model.Datafiles{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddDatafiles is a function to add a single record to datafiles table in the postgres database
// error - ErrInsertFailed, db save call failed
func AddDatafiles(ctx context.Context, record *model.Datafiles) (result *model.Datafiles, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateDatafiles is a function to update a single record from datafiles table in the postgres database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateDatafiles(ctx context.Context, argID int32, updated *model.Datafiles) (result *model.Datafiles, RowsAffected int64, err error) {

	result = &model.Datafiles{}
	db := DB.First(result, argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteDatafiles is a function to delete a single record from datafiles table in the postgres database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteDatafiles(ctx context.Context, argID int32) (rowsAffected int64, err error) {

	record := &model.Datafiles{}
	db := DB.First(record, argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
