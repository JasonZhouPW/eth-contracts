pragma solidity >=0.5.0;

/*
 * @dev wing dao wrapper for Ethereum
 */
interface IEthWingWrapper{
    
    function supply(address token, uint256 amount,uint64 toChainId) payable external;
    
    function withdraw(address token,uint64 _amount,uint64 toChainId) external ;
    
    function withdrawWing(bytes calldata targetAddress,uint64 toChainId) external;

    function updateUserAgentContract(bytes calldata userAgent) external;
    
    function updateOntLockProxy(bytes calldata ontProxy) external;

    function ontLockProxy() external view returns (bytes memory);

    function userAgentContract() external view returns (bytes memory);

}