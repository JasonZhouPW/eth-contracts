const { expectRevert, expectEvent, constants, BN } = require('@openzeppelin/test-helpers');
const { expect } = require('chai');
const Wingwrapper = artifacts.require('/home/zhoupw/work/solidity/polynetwork/eth_contracts/contracts/core/wing/ETHWingWrapper.sol');
const EthCrossChainManager = artifacts.require('./../../contracts/core/.0/CrossChainManager/logic/EthCrossChainManager');
const EthCrossChainData = artifacts.require('./../../contracts/core/.0/CrossChainManager/data/EthCrossChainData');
const EthCrossChainManagerProxy = artifacts.require('./../../contracts/core/.0/CrossChainManager/upgrade/EthCrossChainManagerProxy');
const NewEthCrossChainManager = artifacts.require('./../../contracts/core/.0/CrossChainManager/logic/NewEthCrossChainManager');
const LockProxy = artifacts.require('/home/zhoupw/work/solidity/polynetwork/eth_contracts/contracts/core/lock_proxy/LockProxy.sol');
contract('EthCrossChain', (accounts) => {

    before(async function () {
        console.log("we are here 0");
        this.ECCD = await EthCrossChainData.new({ from: accounts[0], value: web3.utils.toWei('0', 'ether'), gas: 10000000, gasPrice: 50 });
        console.log("this.ECCD.address = ", this.ECCD.address);
        this.ECCM = await EthCrossChainManager.new(this.ECCD.address, { from: accounts[0], value: web3.utils.toWei('0', 'ether'), gas: 200000000, gasPrice: 50 });
        console.log("this.ECCM.address........... = ", this.ECCM.address);
        this.ECCMP = await EthCrossChainManagerProxy.new(this.ECCM.address, { from: accounts[0], value: web3.utils.toWei('0', 'ether'), gas: 60000000, gasPrice: 50 });
        console.log("this.ECCMP.address........... = ", this.ECCMP.address);
        this.LockProxy = await LockProxy.new({ from: accounts[0], value: web3.utils.toWei('0', 'ether'), gas: 20000000, gasPrice: 50 })
        console.log("this.LockProxy.address........... = ", this.LockProxy.address);
        this.wingWrapper = await Wingwrapper.new(accounts[0],
            this.LockProxy.address,
            web3.utils.utf8ToHex("sampletest"),
            this.ECCD.address,
            web3.utils.utf8ToHex("sampletest"),
            { from: accounts[0], value: web3.utils.toWei('0', 'ether'), gas: 10000000, gasPrice: 50 });
    });

    describe('testupdatePolyProxy', function () {
        it('this.wingWrapper.updatePolyProxy(polyProxyAddress) correctly', async function () {
            const {logs} = await this.wingWrapper.updatePolyProxy(this.LockProxy.address);
            expectEvent.inLogs(logs, 'PolyLockProxyUpdated', {
                oldProxy:this.LockProxy.address,
                newProxy:this.LockProxy.address,
            });
        });
        it('this.Lockproxy.setManagerProxy',async function() {
            const {logs} = await this.LockProxy.setManagerProxy(this.ECCMP.address);
            expectEvent.inLogs(logs,'SetManagerProxyEvent', {
                manager:this.ECCMP.address
            });
        });
        it('this.Lockproxy.bindAssetHash',async function(){
            const {logs} = await this.LockProxy.bindAssetHash("0x0000000000000000000000000000000000000000",
                                                "4",
                                                "0xdf19600d334bb13c6a9e3e9777aa8ec6ed6a4a79");
            expectEvent.inLogs(logs,'BindAssetEvent', {
                fromAssetHash:"0x0000000000000000000000000000000000000000",
                toChainId:"4"
            });
            console.log(await this.LockProxy.assetHashMap("0x0000000000000000000000000000000000000000","4"));

        });
        it('this.Lockproxy.lock', async function(){
            const {logs} = await this.LockProxy.lock("0x0000000000000000000000000000000000000000","4","0x0000000000000000000000000000000000000000","1",
                            { from: accounts[0], value: "1", gas: 200000000, gasPrice: 50 });
            expectEvent.inLogs(logs,'LockEvent',{
                fromAssetHash:"0x0000000000000000000000000000000000000000",
                fromAddress:accounts[0],
                toChainId:"4"
                });
        });

        it('this.wingWrapper.supply eth',async function(){
            console.log(await this.LockProxy.assetHashMap("0x0000000000000000000000000000000000000000","4"));

            const {logs} = await this.wingWrapper.supply("0x0000000000000000000000000000000000000000",
                                                    new web3.utils.BN('1000000000000000000'),
                                                    "4",
                                                    { from: accounts[0], value: web3.utils.toWei('1', 'ether'), gas: 200000000, gasPrice: 50 });
            expectEvent.inLogs(logs,'Supply',{
                user:accounts[0],
                token:'0x0000000000000000000000000000000000000000',
                amount:new web3.utils.BN('1000000000000000000')});
                console.log(logs);
            // expectEvent.inLogs(logs,'LockEvent',{
            //     fromAssetHash:"0x0000000000000000000000000000000000000000",
            //     fromAddress:accounts[0],
            //     toChainId:"4"
            //     });
        });
    });

});