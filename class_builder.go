package main

type ClassBuilder interface {
	Build(result *ParsedResult) string
}
