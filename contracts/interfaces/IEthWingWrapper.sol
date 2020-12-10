pragma solidity >=0.5.0;

/*
 * @dev wing dao wrapper for Ethereum
 */
interface IEthWingWrapper{
    function supply(address token, uint256 amount, bytes calldata ontAddr) external;
    function withdraw(bytes calldata _assetContract,uint64 _amount) external ;
    function withdrawWing(bytes calldata targetAddress) external ;

}