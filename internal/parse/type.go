package parse

import (
	"github.com/mavryk-network/mvgo/contract/ast"
	m "github.com/mavryk-network/mvgo/micheline"
)

func (p *parser) buildTypeStructs(t *m.Typedef) (*ast.Struct, error) {
	// Preserve option so generated code uses bind.Option[T]
	if t.Optional {
		inner, err := p.buildTypeStructs(&m.Typedef{Name: "", Type: t.Type, Args: t.Args})
		if err != nil {
			return nil, err
		}
		return &ast.Struct{
			Name:          t.Name,
			MichelineType: "option",
			Type:          inner,
		}, nil
	}
	// Builtin types
	if op, err := m.ParseOpCode(t.Type); err == nil {
		opstr := op.String()
		switch op {
		case m.T_NAT,
			m.T_INT,
			m.T_STRING,
			m.T_BOOL,
			m.T_BYTES,
			m.T_UNIT,
			m.T_TIMESTAMP,
			m.T_ADDRESS,
			m.T_MUMAV,
			m.T_KEY,
			m.T_KEY_HASH,
			m.T_SIGNATURE,
			m.T_CHAIN_ID,
			m.T_OPERATION,
			m.T_CONTRACT:
			return &ast.Struct{
				Name:          t.Name,
				MichelineType: opstr,
			}, nil
		case m.T_BIG_MAP,
			m.T_MAP,
			m.T_LAMBDA:
			type1, err := p.buildTypeStructs(&t.Args[0])
			if err != nil {
				return nil, err
			}
			type2, err := p.buildTypeStructs(&t.Args[1])
			if err != nil {
				return nil, err
			}
			switch op {
			case m.T_BIG_MAP, m.T_MAP:
				return &ast.Struct{
					Name:          t.Name,
					MichelineType: opstr,
					Key:           type1,
					Value:         type2,
				}, nil
			case m.T_LAMBDA:
				return &ast.Struct{
					Name:          t.Name,
					MichelineType: opstr,
					ParamType:     type1,
					ReturnType:    type2,
				}, nil
			}
		}
	}
	// container type
	type1, err := p.buildTypeStructs(&t.Args[0])
	if err != nil {
		return nil, err
	}
	switch t.Type {
	case m.TypeStruct:
		return p.buildStruct(t)
	case m.TypeUnion:
		type2, err := p.buildTypeStructs(&t.Args[1])
		if err != nil {
			return nil, err
		}
		return &ast.Struct{
			MichelineType: "union",
			LeftType:      type1,
			RightType:     type2,
		}, nil
	case "list":
		return &ast.Struct{
			MichelineType: "list",
			Type:          type1,
		}, nil
	case "set":
		return &ast.Struct{
			MichelineType: "set",
			Type:          type1,
		}, nil
	}
	return nil, nil
}

// pathsRelativeToStruct normalizes field paths to be relative to the struct root
// by stripping the longest common prefix from all paths. This is necessary when
// generating structs for list/set element types or union branches, where paths
// from the type tree include the parent container's path (e.g., list index 0).
//
// For example, a list element with three fields might have absolute paths:
//
//	[0, 0], [0, 1, 0], [0, 1, 1]
//
// The common prefix [0] represents the list index. After normalization:
//
//	[0], [1, 0], [1, 1]
//
// These relative paths are used by MarshalParamsPath to build a tree rooted at
// the element itself, not at the parent container.
//
// If no common prefix exists, paths are returned unchanged.
// If a path becomes empty after stripping the prefix, it's replaced with [0].
func pathsRelativeToStruct(paths [][]int) [][]int {
	if len(paths) == 0 {
		return paths
	}
	// Find longest common prefix
	prefixLen := 0
	for prefixLen < len(paths[0]) {
		cur := paths[0][prefixLen]
		same := true
		for _, p := range paths {
			if prefixLen >= len(p) || p[prefixLen] != cur {
				same = false
				break
			}
		}
		if !same {
			break
		}
		prefixLen++
	}
	if prefixLen == 0 {
		return paths
	}
	out := make([][]int, len(paths))
	for i, p := range paths {
		rel := p[prefixLen:]
		if len(rel) == 0 {
			rel = []int{0}
		}
		out[i] = rel
	}
	return out
}

func (p *parser) buildStruct(t *m.Typedef) (*ast.Struct, error) {
	fieldTypes := make([]*ast.Struct, 0, len(t.Args))
	path := make([][]int, 0, len(t.Args))
	for _, a := range t.Args {
		typ, err := p.buildTypeStructs(&a)
		if err != nil {
			return nil, err
		}
		name := a.Name
		if startsWithInt(name) {
			name = "field" + name
		}
		fieldTypes = append(fieldTypes, &ast.Struct{Name: name, Type: typ})
		path = append(path, a.Path)
	}
	// Make paths relative to this struct so MarshalParamsPath builds the right tree
	// (e.g. list item type had paths [0,0],[0,1,0],[0,1,1] -> [0],[1,0],[1,1])
	path = pathsRelativeToStruct(path)
	st := &ast.Struct{
		MichelineType: "struct",
		Fields:        fieldTypes,
		Path:          path,
	}
	// Without annotation, structs gets a
	// @-prefixed auto generated name.
	// We want to remove it, so we get our auto-generated name.
	if len(t.Name) > 0 && t.Name[0] != '@' {
		st.Name = t.Name
	}
	cachedStruct, err := p.registerStruct(st)
	if err != nil {
		return nil, err
	}
	if cachedStruct != nil {
		return cachedStruct, nil
	}
	return st, nil
}

func (p *parser) registerStruct(newStruct *ast.Struct) (*ast.Struct, error) {
	if found, ok := p.cache.IsCached(newStruct); ok {
		return found, nil
	}
	p.structs = append(p.structs, newStruct)
	return nil, p.cache.CacheStruct(newStruct)
}
