package lite

import (
	"errors"
	"log"

	"github.com/WaynePluto/go-lite/utils"
)

// 路由节点，比如 /user/:id 有三个节点，分别是 /, user, :id
type Node struct {
	name      string           // 节点名称，比如 user, 当为动态路径节点时，part为 ":"
	path      string           // 节点路径, 从根节点到此节点的完整路径, 比如 /user/:id, 用来查找中间件
	paramName string           // 动态参数名，比如 id
	children  map[string]*Node // 子节点
}

// 根据路径插入路径中所有的节点, 返回末级节点
func (root *Node) insert(path string) (*Node, error) {
	if root.children == nil {
		root.children = map[string]*Node{}
	}
	parts := utils.SplitPathToParts(path)
	curNode := root
	for i := 0; i < len(parts); i++ {
		// 需要插入的节点名称
		curPart := parts[i]
		if curPart == "" {
			return nil, errors.New("path error: " + path)
		}
		if curNode.name == curPart {
			if i == len(parts)-1 {
				return curNode, nil
			} else {
				continue
			}
		}

		// 查找当前节点的子节点
		searchName, paramName := curPart, ""
		if searchName[0] == ':' {
			searchName, paramName = ":", curPart[1:]
		}
		child, ok := curNode.children[searchName]
		if ok {
			if child.paramName != paramName {
				// 错误的动态路由节点定义
				return nil, errors.New("Duplicate path param node definition:" + paramName)
			}
		} else {
			// 插入新节点
			child = &Node{
				name:      searchName,
				paramName: paramName,
				path:      curPart,
				children:  map[string]*Node{},
			}
			if curNode.path == "/" {
				child.path = "/" + child.path
			} else {
				child.path = curNode.path + "/" + child.path
			}

			curNode.children[searchName] = child
		}
		curNode = child
	}
	return curNode, nil
}

type matchData struct {
	node   *Node
	params map[string]string
}

// 路由匹配，需要匹配 出最后的节点、所有中间件、所有路径参数
func (root *Node) match(path string) (data *matchData, matched bool) {
	data = &matchData{params: map[string]string{}}
	matched = false
	parts := utils.SplitPathToParts(path)
	curNode := root
	for i := 0; i < len(parts); i++ {
		part := parts[i]
		// 判断当前节点
		if curNode.name == part {
			if i == len(parts)-1 {
				data.node = curNode
				matched = true
				return
			} else {
				continue
			}
		}
		// 查找当前节点的子节点
		child, ok := curNode.children[part]
		if ok {
			curNode = child
			if i == len(parts)-1 {
				data.node = curNode
				matched = true
			}
		} else {
			// 没有匹配到静态节点，匹配动态节点
			child, ok = curNode.children[":"]
			if ok {
				data.params[child.paramName] = part
				curNode = child
				if i == len(parts)-1 {
					data.node = curNode
					matched = true
				}
			} else {
				log.Printf("match wrong part：%v\n", part)
				break
			}
		}
	}
	return
}
