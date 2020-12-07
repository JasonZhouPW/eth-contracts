const { expectRevert, expectEvent, constants, BN } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const Wingwrapper = artifacts.require('/home/zhoupw/work/solidity/polynetwork/eth_contracts/contracts/core/wing/ETHWingWrapper.sol');

contract('EthCrossChain', (accounts) => {

    before(async function () {
        this.wingWrapper = await Wingwrapper.new(accounts[0],
            web3.utils.utf8ToHex("sampletest"),
            "0xb3EF0C7ee74342Cf1fEa20C474BC7D948B735238",
            web3.utils.utf8ToHex("sampletest"),
            { from: accounts[0], value: web3.utils.toWei('0', 'ether'), gas: 10000000, gasPrice: 50 });
        
    });

    describe('testupdatePolyProxy', function () {
        it('this.wingWrapper.updatePolyProxy(polyProxyAddress) correctly', async function () {
            console.log("we are here 1");

            const {logs} = await this.wingWrapper.updatePolyProxy("0xb3EF0C7ee74342Cf1fEa20C474BC7D948B735238");
            console.log("we are here");
            expectEvent.inLogs(logs, 'PolyLockProxyUpdated', {
                oldProxy:"0xcd0c19b8d40CE00F4a1175e02Eb7C748b7C8479E",
                newProxy:"0xb3EF0C7ee74342Cf1fEa20C474BC7D948B735238",
            });
        });
    });

});