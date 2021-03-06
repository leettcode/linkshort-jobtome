package cmd

import (
	"context"
	"fmt"
	"log"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/sqltool"

	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"

	"github.com/alexgtn/go-linkshort/infra/sqlite"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "generate-migration",
	Short: "Generate migration from schema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generating new migration")
		// Load the graph.
		graph, err := entc.LoadGraph("./ent/schema", &gen.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		tbls, err := graph.Tables()
		if err != nil {
			log.Fatalln(err)
		}
		// Create a local migration directory.
		d, err := migrate.NewLocalDir("migrations")
		if err != nil {
			log.Fatalln(err)
		}
		client := sqlite.Open(cfg.DatabaseURL)

		// Inspect it and compare it with the graph.
		m, err := schema.NewMigrate(client, schema.WithDir(d),
			schema.WithFormatter(sqltool.GolangMigrateFormatter),
			schema.WithSumFile(),
		)
		if err != nil {
			log.Fatalln(err)
		}
		if err := m.Diff(context.Background(), tbls...); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
