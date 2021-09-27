package main

import "testing"

func TestInitDB(t *testing.T) {
	InitDB()
}

func TestQueryRow(t *testing.T) {
	InitDB()
	QueryRow(2)
}

func TestQueryMultiRow(t *testing.T) {
	InitDB()
	QueryMultiRow(0)
}

func TestInsert(t *testing.T) {
	InitDB()
	Insert()
}

func TestUpdate(t *testing.T) {
	InitDB()
	Update()
}

func TestDelete(t *testing.T) {
	InitDB()
	Delete()
}

func TestTransaction(t *testing.T) {
	InitDB()
	Transaction()
}
