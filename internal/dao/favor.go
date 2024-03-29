// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo/internal/dao/internal"
)

// internalFavorDao is internal type for wrapping internal DAO implements.
type internalFavorDao = *internal.FavorDao

// favorDao is the data access object for table favor.
// You can define custom methods on it to extend its functionality as you wish.
type favorDao struct {
	internalFavorDao
}

var (
	// Favor is globally public accessible object for table favor operations.
	Favor = favorDao{
		internal.NewFavorDao(),
	}
)

// Fill with you ideas below.
