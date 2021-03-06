package database

import (
	"github.com/LemoFoundationLtd/lemochain-core/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewAssetToken(id common.Hash, code common.Hash, addr common.Address, isNil bool) *AssetToken {
	result := &AssetToken{
		Id:    id,
		Code:  code,
		Owner: addr,
	}

	if !isNil {
		result.MetaData = "profile"
	}

	return result
}

func NewAssetTokenBatch1(id common.Hash, isNil bool) *AssetToken {
	return NewAssetToken(id, common.HexToHash("0x0abcd"), common.HexToAddress("0x01234"), isNil)
}

func NewAssetTokenBatch20() []common.Hash {
	result := make([]common.Hash, 20)
	result[0] = common.HexToHash("0x023456789")
	result[1] = common.HexToHash("0x123456789")
	result[2] = common.HexToHash("0x223456789")
	result[3] = common.HexToHash("0x323456789")
	result[4] = common.HexToHash("0x423456789")
	result[5] = common.HexToHash("0x523456789")
	result[6] = common.HexToHash("0x623456789")
	result[7] = common.HexToHash("0x723456789")
	result[8] = common.HexToHash("0x823456789")
	result[9] = common.HexToHash("0x923456789")
	result[10] = common.HexToHash("0x1023456789")
	result[11] = common.HexToHash("0x1123456789")
	result[12] = common.HexToHash("0x1223456789")
	result[13] = common.HexToHash("0x1323456789")
	result[14] = common.HexToHash("0x1423456789")
	result[15] = common.HexToHash("0x1523456789")
	result[16] = common.HexToHash("0x1623456789")
	result[17] = common.HexToHash("0x1723456789")
	result[18] = common.HexToHash("0x1823456789")
	result[19] = common.HexToHash("0x1923456789")
	return result
}

func TestAssetTokenDao_Get(t *testing.T) {
	db := NewMySqlDB(DRIVER_MYSQL, HOST_MYSQL)
	defer db.Close()
	defer db.Clear()
	AssetTokenDao := NewAssetTokenDao(db)

	ids := NewAssetTokenBatch20()
	data := NewAssetTokenBatch1(ids[0], false)

	err := AssetTokenDao.Set(data)
	assert.NoError(t, err)
	result, err := AssetTokenDao.Get(ids[0])
	assert.NoError(t, err)
	assert.Equal(t, data, result)

	err = AssetTokenDao.Set(data)
	assert.NoError(t, err)
	result, err = AssetTokenDao.Get(ids[0])
	assert.NoError(t, err)
	assert.Equal(t, data, result)

	// profile is nil.
	data = NewAssetTokenBatch1(ids[0], true)
	err = AssetTokenDao.Set(data)
	assert.NoError(t, err)
	result, err = AssetTokenDao.Get(ids[0])
	assert.NoError(t, err)
	assert.Equal(t, data, result)

	err = AssetTokenDao.Set(data)
	assert.NoError(t, err)
	result, err = AssetTokenDao.Get(ids[0])
	assert.NoError(t, err)
	assert.Equal(t, data, result)
}

func TestAssetTokenDao_GetPage(t *testing.T) {
	db := NewMySqlDB(DRIVER_MYSQL, HOST_MYSQL)
	defer db.Close()
	defer db.Clear()
	AssetTokenDao := NewAssetTokenDao(db)

	ids := NewAssetTokenBatch20()
	for index := 0; index < len(ids); index++ {
		data := NewAssetTokenBatch1(ids[index], false)
		AssetTokenDao.Set(data)
	}

	data := NewAssetTokenBatch1(ids[0], false)
	result, err := AssetTokenDao.GetPage(data.Owner, 0, 5)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(result))

	result, err = AssetTokenDao.GetPage(data.Owner, 20, 1)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(result))
}

func TestAssetTokenDao_GetPageWithTotal(t *testing.T) {
	db := NewMySqlDB(DRIVER_MYSQL, HOST_MYSQL)
	defer db.Close()
	defer db.Clear()
	AssetTokenDao := NewAssetTokenDao(db)

	ids := NewAssetTokenBatch20()
	for index := 0; index < len(ids); index++ {
		data := NewAssetTokenBatch1(ids[index], false)
		AssetTokenDao.Set(data)
	}

	data := NewAssetTokenBatch1(ids[0], false)
	result, total, err := AssetTokenDao.GetPageWithTotal(data.Owner, 0, 5)
	assert.NoError(t, err)
	assert.Equal(t, 20, total)
	assert.Equal(t, 5, len(result))

	result, total, err = AssetTokenDao.GetPageWithTotal(data.Owner, 20, 1)
	assert.NoError(t, err)
	assert.Equal(t, 20, total)
	assert.Equal(t, 0, len(result))
}

func TestAssetTokenDao_GetPageByCode(t *testing.T) {
	db := NewMySqlDB(DRIVER_MYSQL, HOST_MYSQL)
	defer db.Close()
	defer db.Clear()
	AssetTokenDao := NewAssetTokenDao(db)

	ids := NewAssetTokenBatch20()
	for index := 0; index < len(ids); index++ {
		data := NewAssetTokenBatch1(ids[index], false)
		AssetTokenDao.Set(data)
	}

	data := NewAssetTokenBatch1(ids[0], false)
	result, err := AssetTokenDao.GetPageByCode(data.Code, 0, 5)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(result))

	result, err = AssetTokenDao.GetPageByCode(data.Code, 20, 1)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(result))
}

func TestAssetTokenDao_GetPageByCodeWithTotal(t *testing.T) {
	db := NewMySqlDB(DRIVER_MYSQL, HOST_MYSQL)
	defer db.Close()
	defer db.Clear()
	AssetTokenDao := NewAssetTokenDao(db)

	ids := NewAssetTokenBatch20()
	for index := 0; index < len(ids); index++ {
		data := NewAssetTokenBatch1(ids[index], false)
		AssetTokenDao.Set(data)
	}

	data := NewAssetTokenBatch1(ids[0], false)
	result, total, err := AssetTokenDao.GetPageByCodeWithTotal(data.Code, 0, 5)
	assert.NoError(t, err)
	assert.Equal(t, 20, total)
	assert.Equal(t, 5, len(result))

	result, total, err = AssetTokenDao.GetPageByCodeWithTotal(data.Code, 20, 1)
	assert.NoError(t, err)
	assert.Equal(t, 20, total)
	assert.Equal(t, 0, len(result))
}

func TestAssetTokenDao_NotExist(t *testing.T) {
	db := NewMySqlDB(DRIVER_MYSQL, HOST_MYSQL)
	defer db.Close()
	defer db.Clear()
	AssetTokenDao := NewAssetTokenDao(db)

	result, err := AssetTokenDao.Get(common.HexToHash("0x01"))
	assert.Equal(t, err, ErrNotExist)
	assert.Nil(t, result)
}

func TestAssetTokenDao_ArgInvalid(t *testing.T) {
	db := NewMySqlDB(DRIVER_MYSQL, HOST_MYSQL)
	defer db.Close()
	defer db.Clear()
	AssetTokenDao := NewAssetTokenDao(db)

	result1, err := AssetTokenDao.Get(common.Hash{})
	assert.Equal(t, err, ErrArgInvalid)
	assert.Nil(t, result1)

	err = AssetTokenDao.Set(nil)
	assert.Equal(t, err, ErrArgInvalid)

	result2, err := AssetTokenDao.GetPage(common.Address{}, -1, 0)
	assert.Equal(t, err, ErrArgInvalid)
	assert.Nil(t, result2)

	result2, total, err := AssetTokenDao.GetPageWithTotal(common.Address{}, -1, 0)
	assert.Equal(t, err, ErrArgInvalid)
	assert.Equal(t, -1, total)
	assert.Nil(t, result2)

	result2, err = AssetTokenDao.GetPageByCode(common.Hash{}, -1, 0)
	assert.Equal(t, err, ErrArgInvalid)
	assert.Nil(t, result2)

	result2, total, err = AssetTokenDao.GetPageByCodeWithTotal(common.Hash{}, -1, 0)
	assert.Equal(t, err, ErrArgInvalid)
	assert.Equal(t, -1, total)
	assert.Nil(t, result2)
}
