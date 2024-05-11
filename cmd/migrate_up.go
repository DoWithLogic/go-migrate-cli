package cmd

import (
	"fmt"

	"github.com/DoWithLogic/go-migrate-cli/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"

	sqlLib "github.com/go-sql-driver/mysql"
)

var userDB, passwordDB, hostDB, nameDB, path string

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate to v1 command",
	Long:  `Command to install version 1 of our application`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Running migrate up command")

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
			cmd.Printf("opening file error: %v \n", err)
			return
		}

		// Create a new migration with MySQL driver instance
		m, err := migrate.NewWithInstance("file", fileSource, nameDB, dbDriver)
		if err != nil {
			cmd.Printf("migrate error: %v \n", err)
			return
		}

		// Perform the migration downgrade
		if err = m.Up(); err != nil {
			cmd.Println(err)
		}

		cmd.Println("Migrate up done with success")
	},
}

func init() {
	migrateUpCmd.Flags().StringVarP(&nameDB, "name", "N", "", "Database Name")
	migrateUpCmd.Flags().StringVarP(&hostDB, "host", "H", "", "Database Host")
	migrateUpCmd.Flags().StringVarP(&userDB, "user", "U", "", "Database User")
	migrateUpCmd.Flags().StringVarP(&passwordDB, "password", "P", "", "Database Password")
	migrateUpCmd.Flags().StringVarP(&path, "path", "", "", "migration path")

	migrateCmd.AddCommand(migrateUpCmd)
}
