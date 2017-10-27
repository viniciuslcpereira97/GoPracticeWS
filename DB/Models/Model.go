package models

// Model contract interface
type Model interface {
    Create(model interface {})
    Find(conditions interface {}, fields interface {})
}