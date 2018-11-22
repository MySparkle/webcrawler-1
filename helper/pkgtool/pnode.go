package pkgtool

import (
	"path/filepath"
	"strings"
)

// 代码包节点字典
var pkgNodeCache map[string]*PkgNode

func init() {
	pkgNodeCache = make(map[string]*PkgNode)
}

type PkgNode struct {
	srcDir        string
	importPath    string
	importers     []*PkgNode
	importedNodes []*PkgNode
	growed        bool
}

// SrcDir 源目录的路径
func (node *PkgNode) SrcDir() string {
	return node.srcDir
}

// ImportPath 源码包导入路径
func (node *PkgNode) ImportPath() string {
	return node.importPath
}

// AddImporter 添加导入当前代码包的代码包的节点
func (node *PkgNode) AddImporter(pn *PkgNode) {
	node.importers = append(node.importers, pn)
}

func (node *PkgNode) AddImportedNode(pn *PkgNode) {
	node.importedNodes = append(node.importedNodes, pn)
}

// Importers 获取导入当前代码包的所有代码包节点
func (node *PkgNode) Importers() []*PkgNode {
	importers := make([]*PkgNode, len(node.importers))
	copy(importers, node.importers)
	return importers
}

// ImportedNodes 获取当前代码包导入的所有代码包节点
func (node *PkgNode) ImportedNodes() []*PkgNode {
	importedNodes := make([]*PkgNode, len(node.importedNodes))
	copy(importedNodes, node.importedNodes)
	return importedNodes
}

// IsLeaf 判断当前节点是否为叶子节点
func (node *PkgNode) IsLeaf() bool {
	if len(node.importedNodes) == 0 {
		return true
	}
	return false
}

// Grow 让当前代码包节点沿着代码包的导入关系向下生长
func (node *PkgNode) Grow() error {
	if node.growed {
		return nil
	}
	importPaths, err := getImportsFromPackage(node.importPath, false)
	if err != nil {
		return err
	}
	pathLen := len(importPaths)
	if pathLen == 0 {
		node.growed = true
		return nil
	}

	// fmt.Printf("PN: %v, IPs: %v\n", node, importPaths)
	subPns := make([]*PkgNode, pathLen)
	var count int32
	for i, importPath := range importPaths {
		if importPath == node.importPath {
			continue
		}
		pn, ok := pkgNodeCache[importPath]
		if !ok {
			pn = NewPkgNode(importPath)
			pkgNodeCache[importPath] = pn
		}
		subPns[i] = pn
		count++
	}
	subPns = subPns[:count]
	for _, subPn := range subPns {
		subPn.AddImporter(node)
		node.AddImportedNode(subPn)
		err = subPn.Grow()
		if err != nil {
			return err
		}
	}
	node.growed = true
	return nil
}

func NewPkgNode(importPath string) *PkgNode {
	packageAbsPath := getAbsPathOfPackage(importPath)
	var srcDir string
	importDir := filepath.FromSlash(importPath)
	if strings.HasPrefix(packageAbsPath, importDir) {
		srcDir = packageAbsPath[:strings.LastIndex(packageAbsPath, importDir)]
	}
	return &PkgNode{
		srcDir:        srcDir,
		importPath:    importPath,
		importers:     make([]*PkgNode, 0),
		importedNodes: make([]*PkgNode, 0),
	}
}
