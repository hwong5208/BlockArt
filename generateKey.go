//
package main

import (
	"fmt"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"
)

//
//
//
//import (
//    "math/rand"
//	"fmt"
//	"time"
//
//	//"os"
//	"net"
//	"crypto/ecdsa"
//	"sync"
//	"net/rpc"
//	"math/big"
//	"log"
//)
//type MinerInfo struct {
//	Address net.Addr
//	Key     ecdsa.PublicKey
//}
//
//type CanvasSettings struct {
//	CanvasXMax uint32
//	CanvasYMax uint32
//}
//
//type MinerNetSettings struct {
//	GenesisBlockHash       string
//	MinNumMinerConnections uint8
//	InkPerOpBlock          uint32
//	InkPerNoOpBlock        uint32
//	HeartBeat              uint32
//	PoWDifficultyOpBlock   uint8
//	PoWDifficultyNoOpBlock uint8
//	CanvasSettings         CanvasSettings
//}
//
//type Miner struct {
//	privKey                 *ecdsa.PrivateKey
//	publicKey               ecdsa.PublicKey
//	peers                   sync.Map
//	peersCount              int64
//	laddr                   string
//	currentMsgSeqNum        uint64
//	setting                 *MinerNetSettings
//	sconn                   *rpc.Client
//	othersMsgSeqNum         map[string]uint64
//	blockChain              map[string]*BlockInfo
//	lastBlockOnLongestChain *BlockInfo
//	//map signature to transaction
//	curTxns  map[string]Transaction
//	txnMutex *sync.Mutex
//}
//type BlockInfo struct {
//	hash  string
//	block *Block
//	depth uint64
//	//hashes of childBlock
//	childBlocks []string
//	//map public key to string
//	//state map[string]State
//	inkRemaining map[string]uint64
//	//shapehash to Shape
//	shapes map[string]Shape
//	//signatures seen so far
//	signatures map[string]bool
//}
//
//type Shape struct {
//	//public key of owner
//	owner string
//	ink   uint64
//}
//
////actual block sent in the network
//type Block struct {
//	PrevHash    string
//	Txns        []Transaction
//	MinerPubKey ecdsa.PublicKey
//	Nonce       uint32
//}
//
//type Transaction struct {
//	txnType string
//	op      string
//	//op sig
//	sig    ecdsaSignature
//	pubKey ecdsa.PublicKey
//	sigStr string
//	//for testing
//}
//
//type ecdsaSignature struct {
//	r, s *big.Int
//}
//
//// Returns true if two signatures are equal
//func (sig ecdsaSignature) equal(other ecdsaSignature) bool {
//	return sig.s.Cmp(other.s) == 0 && sig.r.Cmp(other.r) == 0
//}
//
//// convert sig to string
//func (sig ecdsaSignature) toString() string {
//	return sig.s.String() + sig.r.String()
//}
//
//type Peer struct {
//	pconn *rpc.Client
//}
//
////include sender information
//type BlockMessage struct {
//	Block        Block
//	SeqNum       uint64
//	OriginSender string
//	Sender       string
//}
//type TransactionMessage struct {
//	Transaction  Transaction
//	SeqNum       uint64
//	OriginSender string
//	Sender       string
//}
//
//type Message struct {
//	Msg          string
//	SeqNum       uint64
//	OriginSender string
//	Sender       string
//}
//
//func handleError(err error) {
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//
//func main() {
//	rand.Seed(time.Now().UnixNano())
//
//	fmt.Println(rand.Int()%2)
//
//	newMiner:=new(Miner)
//	newMiner.blockChain = make(map[string]*BlockInfo)
//
//	blockInfo7 := new(BlockInfo)
//	blockInfo7.block = new(Block)
//	blockInfo7.hash = "7"
//	blockInfo7.block.PrevHash ="3"
//	txn7 := [2]Transaction{{op:"7.1",sigStr:"7.1sig"},{op:"7.2",sigStr:"7.2sig"}}
//	blockInfo7.block.Txns =txn7[:]
//	newMiner.blockChain["7"]=blockInfo7
//
//	blockInfo6 := new(BlockInfo)
//	blockInfo6.hash = "6"
//	blockInfo6.block = new(Block)
//	blockInfo6.block.PrevHash ="3"
//	txn6 := [2]Transaction{{op:"6.1",sigStr:"6.1sig"},{op:"6.2",sigStr:"6.2sig"}}
//	blockInfo6.block.Txns =txn6[:]
//	newMiner.blockChain["6"]=blockInfo6
//
//	blockInfo3 := new(BlockInfo)
//	blockInfo3.hash = "3"
//	blockInfo3.block = new(Block)
//	blockInfo3.block.PrevHash ="1"
//	txn3 := [2]Transaction{{op:"3.1",sigStr:"3.1sig"},{op:"3.2",sigStr:"3.2sig"}}
//	blockInfo3.block.Txns =txn3[:]
//	c :=[2]string{"6","7"}
//	blockInfo3.childBlocks=c[:]
//	newMiner.blockChain["3"]=blockInfo3
//
//
//	blockInfo4 := new(BlockInfo)
//	blockInfo4.block = new(Block)
//	blockInfo4.hash = "4"
//	blockInfo4.block.PrevHash ="2"
//	txn4 := [2]Transaction{{op:"4.1",sigStr:"7.1sig"},{op:"4.2",sigStr:"4.2sig"}}
//	blockInfo4.block.Txns =txn4[:]
//	newMiner.blockChain["4"]=blockInfo4
//
//	blockInfo5 := new(BlockInfo)
//	blockInfo5.block = new(Block)
//	blockInfo5.hash = "5"
//	blockInfo5.block.PrevHash ="2"
//	txn5 := [2]Transaction{{op:"5.1",sigStr:"5.1sig"},{op:"5.2",sigStr:"5.2sig"}}
//	blockInfo5.block.Txns =txn5[:]
//	newMiner.blockChain["5"]=blockInfo5
//
//
//	blockInfo2 := new(BlockInfo)
//	blockInfo2.block = new(Block)
//	blockInfo2.hash = "2"
//	blockInfo2.block.PrevHash ="1"
//	txn2 := [2]Transaction{{op:"2.1",sigStr:"2.1sig"},{op:"2.2",sigStr:"2.2sig"}}
//	blockInfo2.block.Txns =txn2[:]
//	d :=[2]string{"4","5"}
//	blockInfo2.childBlocks=d[:]
//	newMiner.blockChain["2"]=blockInfo2
//
//	blockInfo1 := new(BlockInfo)
//	blockInfo1.block = new(Block)
//	blockInfo1.hash = "1"
//	blockInfo1.block.PrevHash =""
//	txn1 := [2]Transaction{{op:"1.1",sigStr:"1.1sig"},{op:"1.2",sigStr:"1.2sig"}}
//	blockInfo1.block.Txns =txn1[:]
//	e :=[2]string{"2","3"}
//
//	blockInfo1.childBlocks=e[:]
//	newMiner.blockChain["1"]=blockInfo1
//
//	fmt.Println(findCommonParentTest(blockInfo7,"4",newMiner))
//
//
//}
//func findCommonParentTest(me *BlockInfo, siblingHash string, m *Miner) map[string]Transaction {
//	parentHash := me.block.PrevHash
//	parentBlock := m.blockChain[parentHash]
//	skipHash := me.hash
//	for _, child := range parentBlock.childBlocks {
//		if child == skipHash {
//			continue
//		}
//		childBlock := m.blockChain[child]
//		tempMap := findSiblingTest(childBlock, siblingHash, m,skipHash)
//		if tempMap != nil {
//			for _,t := range me.block.Txns{
//				_,ok:=tempMap[t.sigStr]
//				if ok{
//					delete(tempMap, t.sigStr)
//				}
//			}
//			return tempMap
//		}
//	}
///*	if parentHash == m.setting.GenesisBlockHash {
//		fmt.Println("find Common parent, reach genesis block, blockchain disconnect")
//		os.Exit(2)
//	}*/
//
//	tempMap :=findCommonParentTest(parentBlock,siblingHash,m)
//	if tempMap!= nil{
//		for _,t := range me.block.Txns{
//			_,ok:=tempMap[t.sigStr]
//			if ok{
//				delete(tempMap, t.sigStr)
//			}
//		}
//		return tempMap
//	}
//
//	return nil
//
//}
//func findSiblingTest(block *BlockInfo, siblingHash string, m *Miner,skip string) map[string]Transaction {
//	if block.hash == siblingHash {
//		tempMap := make(map[string]Transaction)
//		for _,t := range block.block.Txns{
//			tempMap[t.sigStr] = t
//		}
//		return tempMap
//	}
//	for _, child := range block.childBlocks {
//		if child == skip {
//			continue
//		}
//		childBlock := m.blockChain[child]
//		tempMap := findSiblingTest(childBlock, siblingHash, m,skip)
//		if tempMap != nil {
//			for _,t := range block.block.Txns{
//				tempMap[t.sigStr] = t
//			}
//			return tempMap
//		}
//	}
//	return nil
//}
//
//

func pubKeyToString(key ecdsa.PublicKey) string {
	fmt.Println("pub key to string ",string(elliptic.Marshal(key.Curve, key.X, key.Y)))
	return string(elliptic.Marshal(key.Curve, key.X, key.Y))
}
func main(){

	for i:= 1;i<21;i++{
		istring := fmt.Sprintf("%v",i)
		priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		fmt.Println(err)
		privBytes,err := x509.MarshalECPrivateKey(priv)
		fmt.Println(err)
		fmt.Println(pubKeyToString(priv.PublicKey))
		file, err := os.OpenFile("priv" + istring, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		file.Seek(0, 0)
		file.Write(privBytes)
		fmt.Println("key byte len:",len(privBytes))
	}



}
