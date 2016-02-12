package redis

import (
	"time"

	"github.com/google/cayley/graph/memstore/b"
	"github.com/sayden/cayley/graph"
	"github.com/sayden/cayley/quad"
)

type QuadStore struct {
	nextID     int64
	nextQuadID int64
	idMap      map[string]int64
	revIDMap   map[int64]string
	log        []LogEntry
	size       int64
	index      QuadDirectionIndex
	// vip_index map[string]map[int64]map[string]map[int64]*b.Tree
}

type QuadDirectionIndex struct {
	index [4]map[int64]*b.Tree
}

type LogEntry struct {
	ID        int64
	Quad      quad.Quad
	Action    graph.Procedure
	Timestamp time.Time
	DeletedBy int64
}

func NewQuadDirectionIndex() QuadDirectionIndex {
	return QuadDirectionIndex{[...]map[int64]*b.Tree{
		quad.Subject - 1:   make(map[int64]*b.Tree),
		quad.Predicate - 1: make(map[int64]*b.Tree),
		quad.Object - 1:    make(map[int64]*b.Tree),
		quad.Label - 1:     make(map[int64]*b.Tree),
	}}
}

func newQuadStore() *QuadStore {
	return &QuadStore{
		idMap:    make(map[string]int64),
		revIDMap: make(map[int64]string),

		// Sentinel null entry so indices start at 1
		log: make([]LogEntry, 1, 200),

		index:      NewQuadDirectionIndex(),
		nextID:     1,
		nextQuadID: 1,
	}
}
