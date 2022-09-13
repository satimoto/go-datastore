package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCommand = &cobra.Command{
	Use:   "data-migrate",
	Short: "Migrate Satimoto database",
	Long:  "Manages Satimoto database migrations using migration up/down files",
	Run: func(cmd *cobra.Command, args []string) {
		dbHost, _ := cmd.Flags().GetString("db_host")
		dbName, _ := cmd.Flags().GetString("db_name")
		dbUser, _ := cmd.Flags().GetString("db_user")
		dbPass, _ := cmd.Flags().GetString("db_pass")
		down, _ := cmd.Flags().GetBool("down")
		forceVersion, _ := cmd.Flags().GetString("force")
		steps, _ := cmd.Flags().GetString("steps")
		migrationsPath, _ := cmd.Flags().GetString("migrations_path")
		sslMode, _ := cmd.Flags().GetString("ssl_mode")

		if len(dbHost) == 0 || len(dbName) == 0 || len(dbPass) == 0 || len(dbUser) == 0 {
			log.Fatalf("Database env variables not defined")
		}	

		dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbName, sslMode)
		db, err := sql.Open("postgres", dataSource)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			log.Fatal(err)
		}

		m, err := migrate.NewWithDatabaseInstance("file://"+migrationsPath, "postgres", driver)
		if err != nil {
			log.Fatal(err)
		}

		if len(forceVersion) > 0 {
			if value, err := strconv.Atoi(forceVersion); err == nil {
				err = m.Force(value)
			}
		} else if len(steps) > 0 {
			if value, err := strconv.Atoi(steps); err == nil {
				if down {
					err = m.Steps(-value)
				} else {
					err = m.Steps(value)
				}
			}
		} else if down {
			log.Fatalf("Reverse migrations must be used with steps")
		} else {
			err = m.Up()
		}

		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

		version, dirty, err := m.Version()

		if err != nil {
			log.Printf("Migration no changes: %v", err)
		} else {
			log.Printf("Migration complete: %v (%v)", version, dirty)
		}
	},
}

func main() {
	viper.AutomaticEnv()

	migrateCommand.Flags().StringP("db_user", "u", viper.GetString("DB_USER"), "Database user")
	migrateCommand.Flags().StringP("db_pass", "P", viper.GetString("DB_PASS"), "Database password")
	migrateCommand.Flags().StringP("db_host", "H", viper.GetString("DB_HOST"), "Database host:port")
	migrateCommand.Flags().StringP("db_name", "n", viper.GetString("DB_NAME"), "Database name")
	migrateCommand.Flags().StringP("migrations_path", "m", viper.GetString("MIGRATIONS_PATH"), "Migration files path")
	migrateCommand.Flags().StringP("ssl_mode", "s", "disable", "Database SSL mode")
	migrateCommand.Flags().StringP("force", "F", "", "Force a migration version")
	migrateCommand.Flags().StringP("steps", "S", "", "Steps to migrate")
	migrateCommand.Flags().BoolP("down", "D", false, "Reverse migrations")
	migrateCommand.Execute()
}
