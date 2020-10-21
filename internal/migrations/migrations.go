package migrations

import (
	"fmt"
	"sort"

	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

func Migrate(db *gorm.DB) error {
	return gormigrate.New(db, gormigrate.DefaultOptions, all()).Migrate()
}

func all() []*gormigrate.Migration {
	allMigrations := []*gormigrate.Migration{
		changeOverridesToText(),
	}

	sort.SliceStable(allMigrations, func(i, j int) bool { return allMigrations[i].ID < allMigrations[j].ID })

	return allMigrations
}

// migrateToBefore is a helper function for migration tests
// It runs all the migrations before the given one to simplify setting up a valid test scenario
func migrateToBefore(db *gorm.DB, migrationID string) error {
	allMigrations := all()

	id := sort.Search(len(allMigrations), func(i int) bool { return allMigrations[i].ID >= migrationID })
	if id == len(allMigrations) || allMigrations[id].ID != migrationID {
		return fmt.Errorf("Failed to find migration %s in migration list", migrationID)
	}

	toRun := allMigrations[0:id]
	if len(toRun) > 0 {
		return gormigrate.New(db, gormigrate.DefaultOptions, allMigrations[0:id]).Migrate()
	}

	return nil
}

// migrateTo runs all migrations up to and including migrationID
func migrateTo(db *gorm.DB, migratoinID string) error {
	gm := gormigrate.New(db, gormigrate.DefaultOptions, all())
	return gm.MigrateTo(migratoinID)
}
