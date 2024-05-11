package cmd

import (
	"fmt"

	"github.com/DoWithLogic/go-migrate-cli/database"

	sqlLib "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "migrate from v2 to v1",
	Long:  `Command to downgrade database from v2 to v1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migrate down command")

		cfg := sqlLib.NewConfig()
		cfg.User = userDB
		cfg.Passwd = passwordDB
		cfg.Addr = hostDB
		cfg.DBName = nameDB
		cfg.MultiStatements = true

		db := database.Open(cfg.FormatDSN())

		// Create a MySQL driver instance
		dbDriver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			fmt.Printf("instance error: %v \n", err)
			return
		}

		// Open migration source from files
		fileSource, err := (&file.File{}).Open(fmt.Sprintf("file://%s", path))
		if err != nil {
			fmt.Printf("opening file error: %v \n", err)
			return
		}

		// Create a new migration with MySQL driver instance
		m, err := migrate.NewWithInstance("file", fileSource, nameDB, dbDriver)
		if err != nil {
			fmt.Printf("migrate error: %v \n", err)
			return
		}

		// Perform the migration downgrade
		if err = m.Down(); err != nil {
			fmt.Printf("migrate down error: %v \n", err)
			return
		}

		fmt.Println("Migrate down done with success")
	},
}

func init() {
	migrateDownCmd.Flags().StringVarP(&nameDB, "name", "N", "", "Database Name")
	migrateDownCmd.Flags().StringVarP(&hostDB, "host", "H", "", "Database Host")
	migrateDownCmd.Flags().StringVarP(&userDB, "user", "U", "", "Database User")
	migrateDownCmd.Flags().StringVarP(&passwordDB, "password", "P", "", "Database Password")
	migrateDownCmd.Flags().StringVarP(&path, "path", "", "", "migration path")

	migrateCmd.AddCommand(migrateDownCmd)
}
