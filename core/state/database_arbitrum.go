package state

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
)

func (db *cachingDB) ActivatedAsm(targetName string, moduleHash common.Hash) ([]byte, error) {
	cacheKey := activatedAsmCacheKey{moduleHash, targetName}
	if asm, _ := db.activatedAsmCache.Get(cacheKey); len(asm) > 0 {
		return asm, nil
	}
	if asm := rawdb.ReadActivatedAsm(db.wasmdb, targetName, moduleHash); len(asm) > 0 {
		db.activatedAsmCache.Add(cacheKey, asm)
		return asm, nil
	}
	return nil, errors.New("not found")
}