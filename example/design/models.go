package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("GoaCellar", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		Model("Account", func() {
			RendersTo(Account)
			HasMany("Bottles", "Bottle")
		})
		Model("Bottle", func() {
			RendersTo(Bottle)
			BelongsTo("Account")
		})
	})
})
