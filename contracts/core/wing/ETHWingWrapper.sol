// SPDX-License-Identifier: LGPL-3.0-or-later
pragma solidity ^0.5.0;

// import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.2.1-solc-0.7/contracts/token/ERC20/ERC20.sol";
import "./../../libs/common/ZeroCopySource.sol";
import "./../../libs/common/ZeroCopySink.sol";
import "./../cross_chain_manager/interface/IEthCrossChainManager.sol";
import "./../cross_chain_manager/interface/IEthCrossChainManagerProxy.sol";
import "./../cross_chain_manager/interface/IEthWingWrapper.sol";
import "./../../libs/token/ERC20/SafeERC20.sol";

// 1. supply & withdraw
// 2. pre-charge eth as handling fee
// 3. record request status

// TODO: polyLockProxy cannot callback this contract, relay invoke this contract
// TODO: notify formatted param to cross chain. use ZeroCopy serialization and deserialization

contract Wingwrapper  {

    // record user max index as request index
    mapping(address => uint256) public requestIndex;

    // record user request status
    enum RequestStatus{Initialized, Committed, Completed, Failed}
    enum OperationType{Empty, Supply, Withdraw} // use Empty as empty value
    struct Request {
        OperationType operationType;
        RequestStatus status;
    }

    mapping(address => mapping(uint256 => Request)) requests;

    // record poly lock proxy contract address
    address public polyLockProxy;

    //ontology user agent contract address
    bytes public userAgentContract;

    bytes public ontLockProxy;

    address public managerProxyContract;


    // record admin address
    address payable public admin;

    // pre-charge handling fee threshold
    uint256 public supplyFeeThreshold;
    uint256 public withdrawFeeThreshold;

    /* events */
    event PolyLockProxyUpdated(address oldProxy, address newProxy);

    event SupplyFeeThresholdUpdated(uint256 oldThreshold, uint256 newThreshold);
    event WithdrawFeeThresholdUpdated(uint256 oldThreshold, uint256 newThreshold);

    event Supply(address user, address token, uint256 amount);
    event Withdraw(address user, address token, uint256 amount);
    event WithdrawWing(address user,bytes target);

    event RequestUpdated(address user, uint256 requestIndex, OperationType oeprationType, RequestStatus status);

    // admin withdraw handling fee
    event WithdrawFee(uint256 amount);

    /* modifier */
    modifier onlyAdmin() {
        require(msg.sender == admin, "only admin");
        _;
    }

    modifier onlyPolyLockProxy(){
        require(msg.sender == admin, "only poly");
        _;
    }

    modifier adminNotEmpty(){
        require(admin != address(0), "admin empty");
        _;
    }

    modifier polyNotEmpty(){
        require(polyLockProxy != address(0), "poly empty");
        _;
    }

    /* function */
    constructor(address _polyLockProxy,
     bytes memory _userAgentContract,
     address _managerProxyContract,
     bytes memory _ontLockproxy) public{
        admin = msg.sender;
        polyLockProxy = _polyLockProxy;
        userAgentContract = _userAgentContract;
        managerProxyContract = _managerProxyContract;
        ontLockProxy = _ontLockproxy;
    }

    function updatePolyProxy(address newProxy) public onlyAdmin {
        address oldProxy = polyLockProxy;
        polyLockProxy = newProxy;
        emit PolyLockProxyUpdated(oldProxy, newProxy);
    }

    // function updateSupplyFeeThreshold(uint256 threshold) public onlyAdmin {
    //     uint256 oldThreshold = supplyFeeThreshold;
    //     supplyFeeThreshold = threshold;
    //     emit SupplyFeeThresholdUpdated(oldThreshold, threshold);
    // }

    // function updateWithdrawFeeThreshold(uint256 threshold) public onlyAdmin {
    //     uint256 oldThreshold = withdrawFeeThreshold;
    //     withdrawFeeThreshold = threshold;
    //     emit WithdrawFeeThresholdUpdated(oldThreshold, threshold);
    // }

    function supply(address token, uint256 amount) public polyNotEmpty {
        // require(msg.value >= supplyFeeThreshold, "fee is not enough");
        // transfer from user and approve to poly
        IERC20 erc20Token = IERC20(token);

        // ERC20 erc20 = ERC20(token);
        address self = address(this);
        erc20Token.transferFrom(self, self, amount);
        erc20Token.approve(polyLockProxy, amount);
        uint64 toChainId = 4;

        //  call poly to cross chain
        (bool success, bytes memory result) = polyLockProxy.delegatecall(abi.encodeWithSignature("lock(address,uint64,bytes,uint256)", token, toChainId,userAgentContract,amount));
        require(success,"call lock proxy failed");


        //call cmcc
        IEthCrossChainManagerProxy eccmp = IEthCrossChainManagerProxy(managerProxyContract);
        address eccmAddr = eccmp.getEthCrossChainManager();
        IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
        
        bytes memory param = _serializeSupplyParam(msg.sender,token,amount);
        bytes memory txData = _serializeCrosschainParam(userAgentContract,param);

        require(ontLockProxy.length != 0, "empty illegal toProxyHash");
        require(eccm.crossChain(toChainId, ontLockProxy, "invokeContract", txData), "EthCrossChainManager crossChain executed error!");

        // record request
        uint256 userIndex = requestIndex[msg.sender];
        requests[msg.sender][userIndex] = Request(OperationType.Supply, RequestStatus.Initialized);
        requestIndex[msg.sender]++;
        emit RequestUpdated(msg.sender, userIndex, OperationType.Supply, RequestStatus.Initialized);
        emit Supply(msg.sender, token, amount);
    }

    function withdraw(address token, uint256 amount) public polyNotEmpty {
        // require(msg.value >= withdrawFeeThreshold, "fee is not enough");
        // TODO: call poly to cross chain
        //call cmcc
        IEthCrossChainManagerProxy eccmp = IEthCrossChainManagerProxy(managerProxyContract);
        address eccmAddr = eccmp.getEthCrossChainManager();
        IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
        
        bytes memory txData = _serializeWithdrawParam(msg.sender,token,amount);

        require(ontLockProxy.length != 0, "empty illegal toProxyHash");
        require(eccm.crossChain(4, ontLockProxy, "invokeContract", txData), "EthCrossChainManager crossChain executed error!");


        // record request
        uint256 userIndex = requestIndex[msg.sender];
        requests[msg.sender][userIndex] = Request(OperationType.Withdraw, RequestStatus.Initialized);
        requestIndex[msg.sender]++;
        emit RequestUpdated(msg.sender, userIndex, OperationType.Withdraw, RequestStatus.Initialized);
        emit Withdraw(msg.sender, token, amount);
    }

    function withdrawWing(bytes memory targetAddress ) public polyNotEmpty {
        // require(msg.value >= withdrawFeeThreshold, "fee is not enough");
        // TODO: call poly to cross chain
        //call cmcc
        IEthCrossChainManagerProxy eccmp = IEthCrossChainManagerProxy(managerProxyContract);
        address eccmAddr = eccmp.getEthCrossChainManager();
        IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
        
        bytes memory txData = _serializeWithdrawWingParam(msg.sender,targetAddress);

        require(ontLockProxy.length!= 0, "empty illegal toProxyHash");
        require(eccm.crossChain(4, ontLockProxy, "invokeContract", txData), "EthCrossChainManager crossChain executed error!");


        // record request
        uint256 userIndex = requestIndex[msg.sender];
        requests[msg.sender][userIndex] = Request(OperationType.Withdraw, RequestStatus.Initialized);
        requestIndex[msg.sender]++;
        emit RequestUpdated(msg.sender, userIndex, OperationType.Withdraw, RequestStatus.Initialized);
        emit WithdrawWing(msg.sender,targetAddress);
    }



    function updateRequestStatus(address user, uint256 index, RequestStatus status) public onlyPolyLockProxy {
        Request storage request = requests[user][index];
        require(request.operationType != OperationType.Empty, "request is not existed");
        require(uint256(request.operationType) < uint256(status), "status could only forward");
        request.status = status;
        emit RequestUpdated(user, index, request.operationType, status);
    }

    // function withdrawFee() public adminNotEmpty {
    //     uint256 balance = address(this).balance;
    //     admin.transfer(balance);
    //     emit WithdrawFee(balance);
    // }

    function _serializeSupplyParam(address selfaddr,address assetHash,uint256 amount) internal pure returns (bytes memory){
        // bytes1 magicversion = byte(0x00);
        // bytes1 typeByteArray = byte(0x00);
        // bytes1 typeString = byte(0x01);
        // bytes1 typeAddress = byte(0x02);
        // bytes1 typeBool = byte(0x03);
        // bytes1 typeInt = byte(0x04);
        // bytes1 typeH256 = byte(0x05);
        // bytes1 typeList = byte(0x10);

        bytes memory method = bytes("supply");

        bytes memory buff;
        bytes memory selfaddrstr = abi.encodePacked(selfaddr); 
        bytes memory assethashstr = abi.encodePacked(assetHash);

        buff = abi.encodePacked(
            byte(0x00),
            byte(0x10),  //param list
            ZeroCopySink.WriteUint16(4), // we have 4 parameters
            byte(0x01),  //1st is string "supply"
            ZeroCopySink.WriteUint16(uint16(method.length)),
            ZeroCopySink.WriteVarBytes(method),
            byte(0x01), //2nd is string self address
            ZeroCopySink.WriteUint16(uint16(selfaddrstr.length)),
            selfaddrstr,
            byte(0x01), //3rd is string asset hash
            ZeroCopySink.WriteUint16(uint16(assethashstr.length)),
            assethashstr
         );

         return abi.encodePacked(buff,
            byte(0x04), //4th is amount 
            ZeroCopySink.WriteUint255(amount));

    }


    function _serializeWithdrawParam(address selfaddr,address assetHash,uint256 amount) internal pure returns (bytes memory){
        // bytes1 magicversion = byte(0x00);
        // bytes1 typeByteArray = byte(0x00);
        // bytes1 typeString = byte(0x01);
        // bytes1 typeAddress = byte(0x02);
        // bytes1 typeBool = byte(0x03);
        // bytes1 typeInt = byte(0x04);
        // bytes1 typeH256 = byte(0x05);
        // bytes1 typeList = byte(0x10);

        bytes memory method = bytes("withdraw");

        bytes memory buff;
        bytes memory selfaddrstr = abi.encodePacked(selfaddr); 
        bytes memory assethashstr = abi.encodePacked(assetHash);

         buff = abi.encodePacked(
            byte(0x00),
            byte(0x10),  //param list
            ZeroCopySink.WriteUint16(4), // we have 4 parameters
            byte(0x01),  //1st is string "withdraw"
            ZeroCopySink.WriteUint16(uint16(method.length)),
            ZeroCopySink.WriteVarBytes(method),
            byte(0x01), //2nd is string self address
            ZeroCopySink.WriteUint16(uint16(selfaddrstr.length)),
            selfaddrstr,
            byte(0x01), //3rd is string asset hash
            ZeroCopySink.WriteUint16(uint16(assethashstr.length)),
            assethashstr
         );

         return abi.encodePacked(buff,
            byte(0x04), //4th is amount 
            ZeroCopySink.WriteUint255(amount));
    }


    function _serializeWithdrawWingParam(address selfaddr,bytes memory targetAddress) internal pure returns (bytes memory){
        // bytes1 magicversion = byte(0x00);
        // bytes1 typeByteArray = byte(0x00);
        // bytes1 typeString = byte(0x01);
        // bytes1 typeAddress = byte(0x02);
        // bytes1 typeBool = byte(0x03);
        // bytes1 typeInt = byte(0x04);
        // bytes1 typeH256 = byte(0x05);
        // bytes1 typeList = byte(0x10);

        bytes memory method = bytes("withdrawWing");

        bytes memory buff;
        bytes memory selfaddrstr = abi.encodePacked(selfaddr); 
        buff = abi.encodePacked(
            byte(0x00),
            byte(0x10),  //param list
            ZeroCopySink.WriteUint16(3), // we have 4 parameters
            byte(0x01),  //1st is string "supply"
            ZeroCopySink.WriteUint16(uint16(method.length)),
            ZeroCopySink.WriteVarBytes(method),
            byte(0x01), //2nd is string self address
            ZeroCopySink.WriteUint16(uint16(selfaddrstr.length)),
            selfaddrstr,
            byte(0x01), //3rd is string ontology address
            ZeroCopySink.WriteUint16(uint16(targetAddress.length)),
            targetAddress
         );

        
        return buff;
    }



    function _serializeCrosschainParam(bytes memory targetContractHash, bytes memory param) internal pure returns (bytes memory) {
        bytes memory buff;
        buff = abi.encodePacked(
            ZeroCopySink.WriteVarBytes(targetContractHash),
            ZeroCopySink.WriteVarBytes(param)
            );
        return buff;
    }
}
