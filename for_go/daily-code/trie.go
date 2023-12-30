package dailycode

// 前缀树

type TailTree struct {
	Val     byte
	IsLeave bool
	SubTree map[byte]*TailTree
}

func GetAllPossibleString(root *TailTree, withPreStr string) []string {
	if root == nil {
		return nil
	}
	targetNode := getTargetNode(root, withPreStr, 0)

	val := make([]string, 0)
	getNodeAllSubNode(targetNode, withPreStr, &val)
	return val
}

func getNodeAllSubNode(targetNode *TailTree, preStr string, val *[]string) {
	if targetNode == nil {
		return
	}

	if targetNode.IsLeave {
		*val = append(*val, preStr)
	}

	for key, node := range targetNode.SubTree {
		getNodeAllSubNode(node, preStr+string(key), val)
	}
}

func getTargetNode(root *TailTree, withPreStr string, currentIndex int) *TailTree {
	if root == nil || currentIndex == len(withPreStr) {
		return root
	}
	if val, ok := root.SubTree[withPreStr[currentIndex]]; ok {
		return getTargetNode(val, withPreStr, currentIndex+1)
	}
	return nil
}

func buildStringTrie(rot *TailTree, valueStr []byte, index int) {
	if rot.SubTree == nil {
		rot.SubTree = make(map[byte]*TailTree)
	}

	if _, ok := rot.SubTree[valueStr[index]]; ok {
		rot.SubTree[valueStr[index]] = &TailTree{
			Val: valueStr[index],
		}
	}
	buildStringTrie(rot.SubTree[valueStr[index]], valueStr, index+1)
}

func NewTrie(val []string) *TailTree {
	root := new(TailTree)
	for _, v := range val {
		buildStringTrie(root, []byte(v), 0)
	}
	return root
}
