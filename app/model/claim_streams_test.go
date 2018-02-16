// This file is generated by SQLBoiler (https://github.com/lbryio/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package model

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/lbryio/sqlboiler/boil"
	"github.com/lbryio/sqlboiler/randomize"
	"github.com/lbryio/sqlboiler/strmangle"
)

func testClaimStreams(t *testing.T) {
	t.Parallel()

	query := ClaimStreams(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testClaimStreamsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = claimStream.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := ClaimStreams(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClaimStreamsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ClaimStreams(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := ClaimStreams(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClaimStreamsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ClaimStreamSlice{claimStream}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := ClaimStreams(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testClaimStreamsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ClaimStreamExists(tx, claimStream.ID)
	if err != nil {
		t.Errorf("Unable to check if ClaimStream exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClaimStreamExistsG to return true, but got false.")
	}
}
func testClaimStreamsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	claimStreamFound, err := FindClaimStream(tx, claimStream.ID)
	if err != nil {
		t.Error(err)
	}

	if claimStreamFound == nil {
		t.Error("want a record, got nil")
	}
}
func testClaimStreamsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ClaimStreams(tx).Bind(claimStream); err != nil {
		t.Error(err)
	}
}

func testClaimStreamsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := ClaimStreams(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClaimStreamsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStreamOne := &ClaimStream{}
	claimStreamTwo := &ClaimStream{}
	if err = randomize.Struct(seed, claimStreamOne, claimStreamDBTypes, false, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}
	if err = randomize.Struct(seed, claimStreamTwo, claimStreamDBTypes, false, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStreamOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = claimStreamTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := ClaimStreams(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClaimStreamsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	claimStreamOne := &ClaimStream{}
	claimStreamTwo := &ClaimStream{}
	if err = randomize.Struct(seed, claimStreamOne, claimStreamDBTypes, false, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}
	if err = randomize.Struct(seed, claimStreamTwo, claimStreamDBTypes, false, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStreamOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = claimStreamTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ClaimStreams(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testClaimStreamsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ClaimStreams(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClaimStreamsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx, claimStreamColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := ClaimStreams(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClaimStreamToOneClaimUsingID(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local ClaimStream
	var foreign Claim

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, claimStreamDBTypes, false, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, claimDBTypes, false, claimColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Claim struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ID(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ClaimStreamSlice{&local}
	if err = local.L.LoadID(tx, false, (*[]*ClaimStream)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.ID == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ID = nil
	if err = local.L.LoadID(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.ID == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testClaimStreamToOneSetOpClaimUsingID(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a ClaimStream
	var b, c Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, claimStreamDBTypes, false, strmangle.SetComplement(claimStreamPrimaryKeyColumns, claimStreamColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Claim{&b, &c} {
		err = a.SetID(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ID != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.IDClaimStream != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ID != x.ID {
			t.Error("foreign key was wrong value", a.ID)
		}

		if exists, err := ClaimStreamExists(tx, a.ID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testClaimStreamsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = claimStream.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testClaimStreamsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ClaimStreamSlice{claimStream}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testClaimStreamsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := ClaimStreams(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	claimStreamDBTypes = map[string]string{`ID`: `bigint`, `Stream`: `mediumtext`}
	_                  = bytes.MinRead
)

func testClaimStreamsUpdate(t *testing.T) {
	t.Parallel()

	if len(claimStreamColumns) == len(claimStreamPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ClaimStreams(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	if err = claimStream.Update(tx); err != nil {
		t.Error(err)
	}
}

func testClaimStreamsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(claimStreamColumns) == len(claimStreamPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	claimStream := &ClaimStream{}
	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := ClaimStreams(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, claimStream, claimStreamDBTypes, true, claimStreamPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(claimStreamColumns, claimStreamPrimaryKeyColumns) {
		fields = claimStreamColumns
	} else {
		fields = strmangle.SetComplement(
			claimStreamColumns,
			claimStreamPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(claimStream))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ClaimStreamSlice{claimStream}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testClaimStreamsUpsert(t *testing.T) {
	t.Parallel()

	if len(claimStreamColumns) == len(claimStreamPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	claimStream := ClaimStream{}
	if err = randomize.Struct(seed, &claimStream, claimStreamDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = claimStream.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert ClaimStream: %s", err)
	}

	count, err := ClaimStreams(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &claimStream, claimStreamDBTypes, false, claimStreamPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClaimStream struct: %s", err)
	}

	if err = claimStream.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert ClaimStream: %s", err)
	}

	count, err = ClaimStreams(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
