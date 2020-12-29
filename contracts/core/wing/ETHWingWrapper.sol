// SPDX-License-Identifier: LGPL-3.0-or-later
pragma solidity ^0.5.0;

import "./../../libs/common/ZeroCopySource.sol";
import "./../../libs/common/ZeroCopySink.sol";
import "./../cross_chain_manager/interface/IEthCrossChainManager.sol";
import "./../cross_chain_manager/interface/IEthCrossChainManagerProxy.sol";
import "./../cross_chain_manager/interface/IEthWingWrapper.sol";
import "./../../libs/token/ERC20/SafeERC20.sol";
import "./ILockProxy.sol";


contract Wingwrapper  {
    using SafeERC20 for IERC20;

    // record poly lock proxy contract address
    address payable public polyLockProxy;

    //ontology user agent contract address
    bytes public userAgentContract;

    bytes public ontLockProxy;

    address public managerProxyContract;


    // record admin address
    address public admin;

    /* events */
    event PolyLockProxyUpdated(address oldProxy, address newProxy);
    event OntLockProxyUpdated(bytes oldProxy, bytes newProxy);
    event ManagerProxyUpdate(address oldProxy, address newProxy);

    event Supply(address user, address token, uint256 amount);
    event Withdraw(address user, address token, uint256 amount);
    event WithdrawWing(address user,bytes target);

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
    constructor(
     address _admin,   
     address payable _polyLockProxy,
     bytes memory _userAgentContract,
     address _managerProxyContract,
     bytes memory _ontLockproxy) public{
        admin = _admin;
        polyLockProxy = _polyLockProxy;
        userAgentContract = _userAgentContract;
        managerProxyContract = _managerProxyContract;
        ontLockProxy = _ontLockproxy;
    }

    function updatePolyProxy(address payable newProxy) public onlyAdmin {
        address oldProxy = polyLockProxy;
        polyLockProxy = newProxy;
        emit PolyLockProxyUpdated(oldProxy, newProxy);
    }

    function updateUserAgentContract(bytes memory _userAgent) public onlyAdmin {
        bytes memory oldProxy = userAgentContract;
        userAgentContract = _userAgent;
        emit OntLockProxyUpdated(oldProxy,userAgentContract);
    }

    function updateOntLockProxy(bytes memory _ontProxy) public onlyAdmin {
        bytes memory oldProxy = ontLockProxy;
        ontLockProxy = _ontProxy;
        emit OntLockProxyUpdated(oldProxy,ontLockProxy);
    }

    function updateManagerProxyContract(address _managerProxyContract) public onlyAdmin {
        address oldProxy = managerProxyContract;
        managerProxyContract = _managerProxyContract;
        emit PolyLockProxyUpdated(oldProxy, _managerProxyContract);
    }


    function supply(address token, uint256 amount,uint64 toChainId) public payable polyNotEmpty {
        // transfer from user and approve to poly
        ILockProxy lp = ILockProxy(polyLockProxy);
        require(amount != 0,"amount should be greater than 0");
        if (token == address(0)) {
            require(msg.value != 0, "transferred ether cannot be zero!");
            require(msg.value == amount, "transferred ether is not equal to amount!");
            require(lp.lock.value(amount)(token, toChainId,userAgentContract,amount),"lp lock failed");

        }else{
            IERC20 erc20Token = IERC20(token);
            // ERC20 erc20 = ERC20(token);
            address self = address(this);
            erc20Token.safeTransferFrom(msg.sender, self, amount);
            erc20Token.safeApprove(polyLockProxy, 0);
            erc20Token.safeApprove(polyLockProxy, amount);
            require(lp.lock(token, toChainId,userAgentContract,amount),"lp lock failed");

        }
        
        // //call cmcc
        IEthCrossChainManagerProxy eccmp = IEthCrossChainManagerProxy(managerProxyContract);
        address eccmAddr = eccmp.getEthCrossChainManager();
        IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
        
        bytes memory toAssetHash = lp.assetHashMap(token,toChainId);
        require(toAssetHash.length > 0,"no toAssetHash found");
        bytes memory param = _serializeSupplyParam(msg.sender,  toAssetHash,amount);
        bytes memory txData = _serializeCrosschainParam(userAgentContract,param);

        require(ontLockProxy.length!= 0, "empty illegal toProxyHash");
        require(eccm.crossChain(toChainId, ontLockProxy, "invokeWasmContract", txData), "EthCrossChainManager crossChain executed error!");

        // for cost reason, we don't record any status here
        emit Supply(msg.sender, token, amount);
    }

    function withdraw(address token, uint256 amount,uint64 toChainId) public polyNotEmpty {
 
        //call cmcc
        IEthCrossChainManagerProxy eccmp = IEthCrossChainManagerProxy(managerProxyContract);
        address eccmAddr = eccmp.getEthCrossChainManager();
        IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);

        ILockProxy lp = ILockProxy(polyLockProxy);
        bytes memory toAssetHash = lp.assetHashMap(token,toChainId);
        require(toAssetHash.length > 0,"no toAssetHash found");

        bytes memory param = _serializeWithdrawParam(msg.sender,toAssetHash,amount);
        bytes memory txData = _serializeCrosschainParam(userAgentContract,param);

        require(ontLockProxy.length != 0, "empty illegal toProxyHash");
        require(eccm.crossChain(toChainId, ontLockProxy, "invokeWasmContract", txData), "EthCrossChainManager crossChain executed error!");

        emit Withdraw(msg.sender, token, amount);
    }

    function withdrawWing(bytes memory targetAddress,uint64 toChainId ) public polyNotEmpty {
        //call poly to cross chain
        //call cmcc
        IEthCrossChainManagerProxy eccmp = IEthCrossChainManagerProxy(managerProxyContract);
        address eccmAddr = eccmp.getEthCrossChainManager();
        IEthCrossChainManager eccm = IEthCrossChainManager(eccmAddr);
        
        bytes memory param = _serializeWithdrawWingParam(msg.sender,targetAddress);
        bytes memory txData = _serializeCrosschainParam(userAgentContract,param);

        require(ontLockProxy.length!= 0, "empty illegal toProxyHash");
        require(eccm.crossChain(toChainId, ontLockProxy, "invokeWasmContract", txData), "EthCrossChainManager crossChain executed error!");

        emit WithdrawWing(msg.sender,targetAddress);
    }


    function _serializeSupplyParam(address selfaddr,bytes memory assetHash,uint256 amount) internal pure returns (bytes memory){
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
        
        string memory selfaddrstr = _addrtoHexString(selfaddr); 
        // bytes memory assethashstr = abi.encodePacked(assetHash);

        buff = abi.encodePacked(
            byte(0x00),
            byte(0x10),  //param list
            ZeroCopySink.WriteUint32(uint32(4)), // we have 4 parameters
            byte(0x01),  //1st is string "supply"
            ZeroCopySink.WriteUint32(uint32(method.length)),
            method,
            byte(0x01), //2nd is string self address
            ZeroCopySink.WriteUint32(uint32(bytes(selfaddrstr).length)),
            selfaddrstr,
            byte(0x02), //3rd is string asset hash
            assetHash
         );

         return abi.encodePacked(buff,
            byte(0x04), //4th is amount 
            ZeroCopySink.WriteUint128(amount));

    }

    function _addrtoHexString(address addr) internal pure returns(string memory) {
        bytes memory data =  abi.encodePacked(addr);
        bytes memory alphabet = "0123456789abcdef";

        bytes memory str = new bytes(2 + data.length * 2);
        str[0] = "0";
        str[1] = "x";
        for (uint i = 0; i < data.length; i++) {
            str[2+i*2] = alphabet[uint(uint8(data[i] >> 4))];
            str[3+i*2] = alphabet[uint(uint8(data[i] & 0x0f))];
        }
    return string(str);
}

    function _serializeWithdrawParam(address selfaddr,bytes memory assetHash,uint256 amount) internal pure returns (bytes memory){
        bytes memory method = bytes("withdraw");

        bytes memory buff;
        
        string memory selfaddrstr = _addrtoHexString(selfaddr); 
        // bytes memory assethashstr = abi.encodePacked(assetHash);

        buff = abi.encodePacked(
            byte(0x00),
            byte(0x10),  //param list
            ZeroCopySink.WriteUint32(uint32(4)), // we have 4 parameters
            byte(0x01),  //1st is string "withdraw"
            ZeroCopySink.WriteUint32(uint32(method.length)),
            method,
            byte(0x01), //2nd is string self address
            ZeroCopySink.WriteUint32(uint32(bytes(selfaddrstr).length)),
            selfaddrstr,
            byte(0x02), //3rd is string asset hash
            assetHash
         );

         return abi.encodePacked(buff,
            byte(0x04), //4th is amount 
            ZeroCopySink.WriteUint128(amount));

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
        string memory selfaddrstr = _addrtoHexString(selfaddr); 
        buff = abi.encodePacked(
            byte(0x00),
            byte(0x10),  //param list
            ZeroCopySink.WriteUint32(3), // we have 4 parameters
            byte(0x01),  //1st is string "withdrawWing"
            ZeroCopySink.WriteUint32(uint32(method.length)),
            method,
            byte(0x01), //2nd is string self address
            ZeroCopySink.WriteUint32(uint32(bytes(selfaddrstr).length)),
            selfaddrstr,
            byte(0x02), //3rd is string ontology address
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
