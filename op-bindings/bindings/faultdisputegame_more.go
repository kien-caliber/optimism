// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const FaultDisputeGameStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"createdAt\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1019\"},{\"astId\":1001,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"status\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_enum(GameStatus)1010\"},{\"astId\":1002,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"bondManager\",\"offset\":9,\"slot\":\"0\",\"type\":\"t_contract(IBondManager)1009\"},{\"astId\":1003,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"l1Head\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_userDefinedValueType(Hash)1017\"},{\"astId\":1004,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"claimData\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_array(t_struct(ClaimData)1011_storage)dyn_storage\"},{\"astId\":1005,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"proposals\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_struct(OutputProposals)1013_storage\"},{\"astId\":1006,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"claims\",\"offset\":0,\"slot\":\"7\",\"type\":\"t_mapping(t_userDefinedValueType(ClaimHash)1015,t_bool)\"},{\"astId\":1007,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"subgames\",\"offset\":0,\"slot\":\"8\",\"type\":\"t_mapping(t_uint256,t_array(t_uint256)dyn_storage)\"},{\"astId\":1008,\"contract\":\"src/dispute/FaultDisputeGame.sol:FaultDisputeGame\",\"label\":\"subgameAtRootResolved\",\"offset\":0,\"slot\":\"9\",\"type\":\"t_bool\"}],\"types\":{\"t_array(t_struct(ClaimData)1011_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct IFaultDisputeGame.ClaimData[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(ClaimData)1011_storage\"},\"t_array(t_uint256)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint256[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(IBondManager)1009\":{\"encoding\":\"inplace\",\"label\":\"contract IBondManager\",\"numberOfBytes\":\"20\"},\"t_enum(GameStatus)1010\":{\"encoding\":\"inplace\",\"label\":\"enum GameStatus\",\"numberOfBytes\":\"1\"},\"t_mapping(t_uint256,t_array(t_uint256)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256[])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_uint256)dyn_storage\"},\"t_mapping(t_userDefinedValueType(ClaimHash)1015,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(ClaimHash =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(ClaimHash)1015\",\"value\":\"t_bool\"},\"t_struct(ClaimData)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IFaultDisputeGame.ClaimData\",\"numberOfBytes\":\"96\"},\"t_struct(OutputProposal)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IFaultDisputeGame.OutputProposal\",\"numberOfBytes\":\"64\"},\"t_struct(OutputProposals)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IFaultDisputeGame.OutputProposals\",\"numberOfBytes\":\"128\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_userDefinedValueType(Claim)1014\":{\"encoding\":\"inplace\",\"label\":\"Claim\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(ClaimHash)1015\":{\"encoding\":\"inplace\",\"label\":\"ClaimHash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Clock)1016\":{\"encoding\":\"inplace\",\"label\":\"Clock\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Hash)1017\":{\"encoding\":\"inplace\",\"label\":\"Hash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Position)1018\":{\"encoding\":\"inplace\",\"label\":\"Position\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Timestamp)1019\":{\"encoding\":\"inplace\",\"label\":\"Timestamp\",\"numberOfBytes\":\"8\"}}}"

var FaultDisputeGameStorageLayout = new(solc.StorageLayout)

var FaultDisputeGameDeployedBin = "0x6080604052600436106101b75760003560e01c80638980e0cc116100ec578063c55cd0c71161008a578063d8cc1a3c11610064578063d8cc1a3c146106c1578063e66825b2146106e1578063fa24f74314610701578063fdffbb281461072557600080fd5b8063c55cd0c714610629578063c6f0308c1461063c578063cf09e0d0146106a057600080fd5b8063bbdc02db116100c6578063bbdc02db1461052d578063bcef3b551461056b578063c0c3a092146105a8578063c31b29ce146105dc57600080fd5b80638980e0cc146104a45780638b85902b146104b957806392931298146104f957600080fd5b8063529184c911610159578063609d333411610133578063609d333414610451578063632247ea146104665780636361506d146104795780638129fc1c1461048f57600080fd5b8063529184c91461033757806354fd4d501461036b57806355ef20e6146103c157600080fd5b8063298c900511610195578063298c90051461024f57806335fef5671461028f578063363cc427146102a45780634778efe81461030357600080fd5b8063200d2ed2146101bc578063266198f9146101f85780632810e1d61461023a575b600080fd5b3480156101c857600080fd5b506000546101e29068010000000000000000900460ff1681565b6040516101ef91906122c8565b60405180910390f35b34801561020457600080fd5b5061022c7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016101ef565b34801561024657600080fd5b506101e2610738565b34801561025b57600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036040013561022c565b6102a261029d366004612309565b61089c565b005b3480156102b057600080fd5b506000546102de906901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101ef565b34801561030f57600080fd5b5061022c7f000000000000000000000000000000000000000000000000000000000000000081565b34801561034357600080fd5b506102de7f000000000000000000000000000000000000000000000000000000000000000081565b34801561037757600080fd5b506103b46040518060400160405280600681526020017f302e302e3132000000000000000000000000000000000000000000000000000081525081565b6040516101ef9190612396565b3480156103cd57600080fd5b5060408051606080820183526003546fffffffffffffffffffffffffffffffff808216845270010000000000000000000000000000000091829004811660208086019190915260045485870152855193840186526005548083168552929092041690820152600654928101929092526104439182565b6040516101ef9291906123b0565b34801561045d57600080fd5b506103b46108ac565b6102a2610474366004612419565b6108ba565b34801561048557600080fd5b5061022c60015481565b34801561049b57600080fd5b506102a2610ed0565b3480156104b057600080fd5b5060025461022c565b3480156104c557600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036020013561022c565b34801561050557600080fd5b506102de7f000000000000000000000000000000000000000000000000000000000000000081565b34801561053957600080fd5b5060405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101ef565b34801561057757600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033561022c565b3480156105b457600080fd5b506102de7f000000000000000000000000000000000000000000000000000000000000000081565b3480156105e857600080fd5b506106107f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff90911681526020016101ef565b6102a2610637366004612309565b6114d4565b34801561064857600080fd5b5061065c61065736600461244e565b6114e0565b6040805163ffffffff90961686529315156020860152928401919091526fffffffffffffffffffffffffffffffff908116606084015216608082015260a0016101ef565b3480156106ac57600080fd5b506000546106109067ffffffffffffffff1681565b3480156106cd57600080fd5b506102a26106dc3660046124b0565b611551565b3480156106ed57600080fd5b506102a26106fc36600461253a565b611a7d565b34801561070d57600080fd5b50610716611c42565b6040516101ef93929190612566565b6102a261073336600461244e565b611c9f565b60008060005468010000000000000000900460ff16600281111561075e5761075e612299565b14610795576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60095460ff166107d1576040517f9a07664600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60026000815481106107e5576107e5612591565b6000918252602090912060039091020154640100000000900460ff1661080c57600261080f565b60015b6000805491925082917fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff166801000000000000000083600281111561085657610856612299565b02179055600281111561086b5761086b612299565b6040517f5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da6090600090a290565b905090565b6108a8828260006108ba565b5050565b606061089760206040611fd1565b6000805468010000000000000000900460ff1660028111156108de576108de612299565b14610915576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82158015610921575080155b15610958576040517fa42637bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006002848154811061096d5761096d612591565b600091825260208083206040805160a081018252600394909402909101805463ffffffff808216865264010000000090910460ff16151593850193909352600181015491840191909152600201546fffffffffffffffffffffffffffffffff80821660608501819052700100000000000000000000000000000000909204166080840152919350610a019190859061206816565b90507f0000000000000000000000000000000000000000000000000000000000000000610ac0826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff161115610b02576040517f56f57b2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160009063ffffffff90811614610b62576002836000015163ffffffff1681548110610b3157610b31612591565b906000526020600020906003020160020160109054906101000a90046fffffffffffffffffffffffffffffffff1690505b608083015160009067ffffffffffffffff1667ffffffffffffffff1642610b9b846fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff16610baf91906125ef565b610bb99190612607565b9050677fffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000060011c1667ffffffffffffffff82161115610c2c576040517f3381d11400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604082901b42176000888152608086901b6fffffffffffffffffffffffffffffffff8b1617602052604081209192509060008181526007602052604090205490915060ff1615610caa576040517f80497e3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260076020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155815160a08101835263ffffffff808f1682529381018581529281018d81526fffffffffffffffffffffffffffffffff808c16606084019081528982166080850190815260028054808801825599819052945160039099027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace8101805498511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009099169a909916999099179690961790965590517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf8701559351925184167001000000000000000000000000000000000292909316919091177f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad09093019290925580548b908110610e2257610e22612591565b6000918252602080832060039092029091018054931515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff909416939093179092558a8152600890915260409020600254610e8690600190612607565b8154600181018355600092835260208320015560405133918a918c917f9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be91a4505050505050505050565b367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90033560001a6001811480610f10575060ff81166002145b610f76576040517ff40239db000000000000000000000000000000000000000000000000000000008152367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900335600482015260240160405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000164267ffffffffffffffff161781556040805160a08101825263ffffffff81526020810192909252600291908101610ffb7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c90033590565b815260016020820152604001426fffffffffffffffffffffffffffffffff9081169091528254600181810185556000948552602080862085516003909402018054918601511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000090921663ffffffff909416939093171782556040840151908201556060830151608090930151821670010000000000000000000000000000000002929091169190911760029091015573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016637f00642061112460207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c9003013590565b6040518263ffffffff1660e01b815260040161114291815260200190565b602060405180830381865afa15801561115f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611183919061261e565b9050600073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663a25ae5576111ce600185612607565b6040518263ffffffff1660e01b81526004016111ec91815260200190565b606060405180830381865afa158015611209573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061122d9190612686565b6040517fa25ae5570000000000000000000000000000000000000000000000000000000081526004810184905290915060009073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063a25ae55790602401606060405180830381865afa1580156112be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112e29190612686565b9050600073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000166399d548aa367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003604001356040518263ffffffff1660e01b815260040161136e91815260200190565b6040805180830381865afa15801561138a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113ae9190612712565b905081602001516fffffffffffffffffffffffffffffffff16816020015167ffffffffffffffff161161140d576040517f13809ba500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160a081018252908190810180611428600189612607565b6fffffffffffffffffffffffffffffffff908116825260408881015182166020808501919091529851928101929092529183528051606081018252978216885285810151821688880152945187860152908501959095528051805181860151908716700100000000000000000000000000000000918816820217600355908401516004559084015180519481015194861694909516029290921760055591909101516006555160015550565b6108a8828260016108ba565b600281815481106114f057600080fd5b600091825260209091206003909102018054600182015460029092015463ffffffff8216935064010000000090910460ff1691906fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041685565b6000805468010000000000000000900460ff16600281111561157557611575612299565b146115ac576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600287815481106115c1576115c1612591565b6000918252602082206003919091020160028101549092506fffffffffffffffffffffffffffffffff16908715821760011b90506116207f000000000000000000000000000000000000000000000000000000000000000060016125ef565b6116bc826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff16146116fd576040517f5f53dd9800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080891561178057611721836fffffffffffffffffffffffffffffffff16612070565b67ffffffffffffffff16156117545761174b61173e600186612799565b865463ffffffff16612116565b60010154611776565b7f00000000000000000000000000000000000000000000000000000000000000005b915084905061179a565b8460010154915061179784600161173e91906127ca565b90505b600882901b60088a8a6040516117b19291906127fe565b6040518091039020901b146117f2576040517f696550ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081600101547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e14ced328c8c8c8c60006040518663ffffffff1660e01b815260040161185b959493929190612857565b6020604051808303816000875af115801561187a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061189e919061261e565b600284810154929091149250600091611949906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b6119e5886fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b6119ef9190612891565b6119f991906128b2565b67ffffffffffffffff161590508115158103611a41576040517ffb4e40dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505084547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff166401000000001790945550505050505050505050565b6000805468010000000000000000900460ff166002811115611aa157611aa1612299565b14611ad8576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637dc0d1d06040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b699190612900565b7f52f0f3ad00000000000000000000000000000000000000000000000000000000601c8190526020869052604085905290915060008560018114611bd55760028114611bdf5760038114611be95760048114611bf35760058114611c035763ff137e656000526004601cfd5b6001549150611c0a565b6004549150611c0a565b6006549150611c0a565b60035460801c60c01b9150611c0a565b4660c01b91505b50606052600160038611811b6005031b60805260a083905260008060a4601c82865af1611c3b573d6000803e3d6000fd5b5050505050565b7f0000000000000000000000000000000000000000000000000000000000000000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003356060611c986108ac565b9050909192565b6000805468010000000000000000900460ff166002811115611cc357611cc3612299565b14611cfa576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060028281548110611d0f57611d0f612591565b60009182526020909120600260039092020190810154909150677fffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000060011c1690611d7f90700100000000000000000000000000000000900467ffffffffffffffff1642612607565b6002830154611daf9190700100000000000000000000000000000000900460401c67ffffffffffffffff166125ef565b11611de6576040517ff2440b5300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260086020526040902082158015611e04575060095460ff165b15611e3b576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8054158015611e4957508215155b15611e80576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000805b8254811015611f4e576000838281548110611ea157611ea1612591565b6000918252602080832090910154808352600890915260409091205490915015611ef7576040517f9a07664600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060028281548110611f0c57611f0c612591565b600091825260209091206003909102018054909150640100000000900460ff16611f3b57600193505050611f4e565b505080611f4790612936565b9050611e84565b5082547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff16640100000000821515021783556000848152600860205260408120611f979161225f565b83600003611fcb57600980547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555b50505050565b6060600061200884367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c90036125ef565b90508267ffffffffffffffff1667ffffffffffffffff81111561202d5761202d612637565b6040519080825280601f01601f191660200182016040528015612057576020820181803683370190505b509150828160208401375092915050565b151760011b90565b6000806120fd837e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b600167ffffffffffffffff919091161b90920392915050565b600080612134846fffffffffffffffffffffffffffffffff166121b3565b90506002838154811061214957612149612591565b906000526020600020906003020191505b60028201546fffffffffffffffffffffffffffffffff8281169116146121ac57815460028054909163ffffffff1690811061219757612197612591565b9060005260206000209060030201915061215a565b5092915050565b60008119600183011681612247827e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169390931c8015179392505050565b508054600082559060005260206000209081019061227d9190612280565b50565b5b808211156122955760008155600101612281565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020810160038310612303577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b6000806040838503121561231c57600080fd5b50508035926020909101359150565b6000815180845260005b8181101561235157602081850181015186830182015201612335565b81811115612363576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006123a9602083018461232b565b9392505050565b82516fffffffffffffffffffffffffffffffff90811682526020808501518216818401526040808601518185015284518316606085015290840151909116608083015282015160a082015260c081016123a9565b8035801515811461241457600080fd5b919050565b60008060006060848603121561242e57600080fd5b833592506020840135915061244560408501612404565b90509250925092565b60006020828403121561246057600080fd5b5035919050565b60008083601f84011261247957600080fd5b50813567ffffffffffffffff81111561249157600080fd5b6020830191508360208285010111156124a957600080fd5b9250929050565b600080600080600080608087890312156124c957600080fd5b863595506124d960208801612404565b9450604087013567ffffffffffffffff808211156124f657600080fd5b6125028a838b01612467565b9096509450606089013591508082111561251b57600080fd5b5061252889828a01612467565b979a9699509497509295939492505050565b60008060006060848603121561254f57600080fd5b505081359360208301359350604090920135919050565b60ff84168152826020820152606060408201526000612588606083018461232b565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115612602576126026125c0565b500190565b600082821015612619576126196125c0565b500390565b60006020828403121561263057600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b80516fffffffffffffffffffffffffffffffff8116811461241457600080fd5b60006060828403121561269857600080fd5b6040516060810181811067ffffffffffffffff821117156126e2577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052825181526126f560208401612666565b602082015261270660408401612666565b60408201529392505050565b60006040828403121561272457600080fd5b6040516040810167ffffffffffffffff828210818311171561276f577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b816040528451835260208501519150808216821461278c57600080fd5b5060208201529392505050565b60006fffffffffffffffffffffffffffffffff838116908316818110156127c2576127c26125c0565b039392505050565b60006fffffffffffffffffffffffffffffffff8083168185168083038211156127f5576127f56125c0565b01949350505050565b8183823760009101908152919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60608152600061286b60608301878961280e565b828103602084015261287e81868861280e565b9150508260408301529695505050505050565b600067ffffffffffffffff838116908316818110156127c2576127c26125c0565b600067ffffffffffffffff808416806128f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910692915050565b60006020828403121561291257600080fd5b815173ffffffffffffffffffffffffffffffffffffffff811681146123a957600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612967576129676125c0565b506001019056fea164736f6c634300080f000a"


var FaultDisputeGameImmutableReferencesJSON = "{\"84196\":[{\"start\":522,\"length\":32},{\"start\":5974,\"length\":32}],\"84199\":[{\"start\":789,\"length\":32},{\"start\":2565,\"length\":32},{\"start\":5626,\"length\":32}],\"84203\":[{\"start\":1518,\"length\":32},{\"start\":3014,\"length\":32},{\"start\":7475,\"length\":32}],\"84207\":[{\"start\":1291,\"length\":32},{\"start\":6139,\"length\":32},{\"start\":6876,\"length\":32}],\"84211\":[{\"start\":1466,\"length\":32},{\"start\":4298,\"length\":32},{\"start\":4510,\"length\":32},{\"start\":4727,\"length\":32}],\"84215\":[{\"start\":841,\"length\":32},{\"start\":4861,\"length\":32}],\"84219\":[{\"start\":1345,\"length\":32},{\"start\":7236,\"length\":32}]}"

func init() {
	if err := json.Unmarshal([]byte(FaultDisputeGameStorageLayoutJSON), FaultDisputeGameStorageLayout); err != nil {
		panic(err)
	}

	layouts["FaultDisputeGame"] = FaultDisputeGameStorageLayout
	deployedBytecodes["FaultDisputeGame"] = FaultDisputeGameDeployedBin
	immutableReferences["FaultDisputeGame"] = FaultDisputeGameImmutableReferencesJSON
}
