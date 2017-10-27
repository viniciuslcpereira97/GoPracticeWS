package models

// TODO - implement interface
// Model contract interface
type Model interface {
    Create(model interface {})
    Find(conditions interface {}, fields interface {}) (interface {})
    Where(conditions interface {}, fields interface {}) ([] interface {})
}