package store

func (store *Store) RunMigration() error {
	err := store.migrator.ApplyMigrations(store.db)

	return err
}
