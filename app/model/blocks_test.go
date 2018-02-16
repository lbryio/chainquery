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

func testBlocks(t *testing.T) {
	t.Parallel()

	query := Blocks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testBlocksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = block.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Blocks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlocksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BlockSlice{block}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testBlocksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := BlockExists(tx, block.ID)
	if err != nil {
		t.Errorf("Unable to check if Block exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BlockExistsG to return true, but got false.")
	}
}
func testBlocksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	blockFound, err := FindBlock(tx, block.ID)
	if err != nil {
		t.Error(err)
	}

	if blockFound == nil {
		t.Error("want a record, got nil")
	}
}
func testBlocksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Blocks(tx).Bind(block); err != nil {
		t.Error(err)
	}
}

func testBlocksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Blocks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBlocksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blockOne := &Block{}
	blockTwo := &Block{}
	if err = randomize.Struct(seed, blockOne, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err = randomize.Struct(seed, blockTwo, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blockOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = blockTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Blocks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBlocksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	blockOne := &Block{}
	blockTwo := &Block{}
	if err = randomize.Struct(seed, blockOne, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}
	if err = randomize.Struct(seed, blockTwo, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blockOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = blockTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testBlocksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlocksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx, blockColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlockToManyBlockHashTransactions(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, transactionDBTypes, false, transactionColumnsWithDefault...)
	randomize.Struct(seed, &c, transactionDBTypes, false, transactionColumnsWithDefault...)

	b.BlockHash.Valid = true
	c.BlockHash.Valid = true
	b.BlockHash.String = a.Hash
	c.BlockHash.String = a.Hash
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	transaction, err := a.BlockHashTransactions(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range transaction {
		if v.BlockHash.String == b.BlockHash.String {
			bFound = true
		}
		if v.BlockHash.String == c.BlockHash.String {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BlockSlice{&a}
	if err = a.L.LoadBlockHashTransactions(tx, false, (*[]*Block)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BlockHashTransactions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BlockHashTransactions = nil
	if err = a.L.LoadBlockHashTransactions(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BlockHashTransactions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", transaction)
	}
}

func testBlockToManyAddOpBlockHashTransactions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c, d, e Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Transaction{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Transaction{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBlockHashTransactions(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Hash != first.BlockHash.String {
			t.Error("foreign key was wrong value", a.Hash, first.BlockHash.String)
		}
		if a.Hash != second.BlockHash.String {
			t.Error("foreign key was wrong value", a.Hash, second.BlockHash.String)
		}

		if first.R.BlockHash != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.BlockHash != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BlockHashTransactions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BlockHashTransactions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BlockHashTransactions(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBlockToManySetOpBlockHashTransactions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c, d, e Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Transaction{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetBlockHashTransactions(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BlockHashTransactions(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetBlockHashTransactions(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BlockHashTransactions(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.BlockHash.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.BlockHash.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.Hash != d.BlockHash.String {
		t.Error("foreign key was wrong value", a.Hash, d.BlockHash.String)
	}
	if a.Hash != e.BlockHash.String {
		t.Error("foreign key was wrong value", a.Hash, e.BlockHash.String)
	}

	if b.R.BlockHash != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BlockHash != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BlockHash != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.BlockHash != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.BlockHashTransactions[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.BlockHashTransactions[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBlockToManyRemoveOpBlockHashTransactions(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Block
	var b, c, d, e Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Transaction{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddBlockHashTransactions(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BlockHashTransactions(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveBlockHashTransactions(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BlockHashTransactions(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.BlockHash.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.BlockHash.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.BlockHash != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BlockHash != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BlockHash != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.BlockHash != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.BlockHashTransactions) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.BlockHashTransactions[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.BlockHashTransactions[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBlocksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = block.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testBlocksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BlockSlice{block}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testBlocksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Blocks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	blockDBTypes = map[string]string{`Bits`: `varchar`, `BlockSize`: `bigint`, `BlockTime`: `bigint`, `Chainwork`: `varchar`, `Confirmations`: `int`, `Created`: `datetime`, `Difficulty`: `decimal`, `Hash`: `varchar`, `Height`: `bigint`, `ID`: `bigint`, `MedianTime`: `bigint`, `MerkleRoot`: `varchar`, `Modified`: `datetime`, `NameClaimRoot`: `varchar`, `NextBlockHash`: `varchar`, `Nonce`: `bigint`, `PreviousBlockHash`: `varchar`, `Target`: `varchar`, `TransactionHashes`: `text`, `TransactionsProcessed`: `tinyint`, `Version`: `bigint`, `VersionHex`: `varchar`}
	_            = bytes.MinRead
)

func testBlocksUpdate(t *testing.T) {
	t.Parallel()

	if len(blockColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err = block.Update(tx); err != nil {
		t.Error(err)
	}
}

func testBlocksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(blockColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	block := &Block{}
	if err = randomize.Struct(seed, block, blockDBTypes, true, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, block, blockDBTypes, true, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(blockColumns, blockPrimaryKeyColumns) {
		fields = blockColumns
	} else {
		fields = strmangle.SetComplement(
			blockColumns,
			blockPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(block))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := BlockSlice{block}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testBlocksUpsert(t *testing.T) {
	t.Parallel()

	if len(blockColumns) == len(blockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	block := Block{}
	if err = randomize.Struct(seed, &block, blockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = block.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Block: %s", err)
	}

	count, err := Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &block, blockDBTypes, false, blockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	if err = block.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Block: %s", err)
	}

	count, err = Blocks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
