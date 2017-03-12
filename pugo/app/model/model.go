// Package model declares interfaces for model objects
package model

import (
	"html/template"
	"time"

	"github.com/go-xiaohei/pugo/app/model/index"
)

type (
	// Target is object that will compiled
	Target interface {
		SrcFile() string
		DstFile() string
	}
	// Node is object of a node in site
	Node interface {
		URL() string
		CreateTime() time.Time
		UpdateTime() time.Time
	}
	// Content is object of node with content output
	Content interface {
		Target
		Node
		Content() []byte
		ContentHTML() template.HTML
		ContentLength() int
	}
	// Index is object of node that can get index for content
	Index interface {
		Index() []*index.Index
	}
)
