package main

/*	Author: 
	Jenil K. Thakker

	Overview: 
	To run this code in the terminal, please type: "go run MerkleTreeI.go"
	If there is a json file to be parsed, the following command can be used
	"go run MerkleTreeI.go < file.json" (fmt.Scanf() function would be used)

	A simple implementation of a merkle tree.
	This code implements the functions assigned in this API and tests them for verification.
	For simplicity, main() tests the functions rather than following an object-oriented approach.

	MNode is a structure I've added to the API, to help me compute the Merkle Tree.
*/

import("fmt"
	   "crypto/sha256"
)

/*Hashable - anything that can provide it's hash */
type Hashable interface {
    GetHash(node *MNode) string
    GetHashBytes() []byte
}

/*GetHash - Returns the hash of a Node */
func GetHash(node *MNode) string {
	if node.left != nil && node.right != nil {                 //checking if the node values are not nil
        node.hash = Hash(node.left.hash + node.right.hash)    // Another way of hashing this is using the MHash function

    } else if node.left != nil && node.right == nil {
        node.hash = node.left.hash
	}
	
	return node.hash
}

/*GetHashBytes - Returns the hash in terms of an array of bytes */
func GetHashBytes() string {
	sum := sha256.Sum256([]byte(""))     //Used crypto/sha256 API
	return fmt.Sprintf("%x", sum) 
}

/*MerkleTreeI - a merkle tree interface required for constructing and providing verification */
type MerkleTreeI interface {
    //API to create a tree from leaf nodes
    ComputeTree(hashes []Hashable) (MNode)
    GetRoot() string
    GetTree() []string

    // API for verification when the leaf node is known
    GetPath(node MNode) bool              // Server needs to provide this
    VerifyPath(hash Hashable, path *MTPath) bool //This is only required by a client but useful for testing

    /* API for random verification when the leaf node is uknown
    (verification of the data to hash used as leaf node is outside this API) */
    GetPathByIndex(idx int) *MTPath
}

/*ComputeTree - Computes a merkle tree using the defined hash function to hash the nodes */
func ComputeTree(hashes []Hashable) (MNode){

	list := MNode{isLeaf: false, hash: "", left: nil, right: nil}

	for i, n := range hashes {
		if n != nil {
			fmt.Printf("> %p : %+v\n", hashes[i], hashes[i])
			//list = append(list, isLeaf: true, hash: "", left: n.left, right: n.right)
		}
	}

	return ComputeTree(list)
}

/*GetRoot - Returns the root of the merkle tree */
  func (node MNode) GetRoot() string {
	return  node.hash
}

/*GetTree - Returns the computed tree (String of hashed nodes from root to leaf) */
func GetTree() []string {
	list := MNode{isLeaf: false, hash: "", left: nil, right: nil}
	return fmt.Println(ComputeTree(list))
}

/*GetPath - Returns the path from the root to the specified leaf node (hashed) */
func GetPath(node MNode) bool {
	i := 0
    height := node.GetHeight()

    for i = 0; i < height && node.right != nil; i++ {
        node = *node.right
    }

    return height == i

}


/*VerifyPath - Verifies the path by comparing hashes */
//VerifyPath(hash Hashable, path *MTPath) bool {
	

//}

/*GetPathByIndex - iterate and return the path through the node index*/
func GetPathByIndex(idx int) *MTPath {
	height := node.GetHeight()

    if idx > height {
        return nil
    }
    return node.getNodesByLevel(idx), nil
}

func (node MNode) GetHeight() (height int) {
    for height = 0; !node.isLeaf && node.left != nil; height++ {
        node = *node.left
    }
    return
}

func (node MNode) getNodesByLevel(level int) []string {
    if level <= 0 {
        return []string{node.hash}

    } else {
        if node.right != nil {
            return append(node.left.getNodesByLevel(level-1), node.right.getNodesByLevel(level-1)...)
        }
        return node.left.getNodesByLevel(level - 1)
    }
}

/*MTPath - The merkle tree path*/
type MTPath struct {
    Nodes     []string `json:"nodes"`
    LeafIndex int      `json:"leaf_index"`
}

/*MNode - The merkle tree node*/
type MNode struct {
    isLeaf     bool
    hash       string
    left       *MNode
    right      *MNode
}

/*Hash - the hashing used for the merkle tree construction */
func Hash(text string) string {
	sum := sha256.Sum256([]byte(text))     //Used the Go lang package documentation (https://golang.org/pkg/crypto/sha256/)
	return fmt.Sprintf("%x", sum)      
}

/*MHash - merkle hashing of a pair of child hashes */
func MHash(h1 string, h2 string) string {
    return Hash(h1 + h2)
}

/*Main - Using the main function with this API to test and verify the defined functions */
func main() {

	//Testing the Hash() function through MHash
	fmt.Println("The merkle hash of a pair of child nodes is: ", MHash("merkle", "tree"))

    //Creating a sample array of nodes for testing
	Nodes := []string{"merkle", "tree", "computation"}
	fmt.Println(Nodes)

}