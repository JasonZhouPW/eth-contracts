const _wingwrapper = require("./5_wingwrapper.js_bak");

const LockProxy = artifacts.require("./contracts/core/lock_proxy/LockProxy");
const eccd = artifacts.require("./contracts/core/cross_chain_manager/data/EthCrossChainData.sol");
const cmcc = artifacts.require("./contracts/core/cross_chain_manager/logic/EthCrossChainManager.sol");
const wingWrapper = artifacts.require("./contracts/core/wing/Wingwrapper");


module.exports = function(deployer,network,accounts) {
  deployer.deploy(LockProxy)
    .then(()=> console.log("lockproxy:" + LockProxy.address));
  deployer.deploy(eccd)
    .then(()=> console.log("eccd:" + eccd.address));

  deployer.deploy(cmcc,eccd.address)
    .then(()=> console.log("cmcc:" + cmcc.address));

  deployer.deploy(wingWrapper,accounts[0],LockProxy.address,"0x49206861766520313030e282ac",cmcc.address,"0x49206861766520313030e282ac")
    .then(()=> console.log("wingWrapper:" + wingWrapper.address));

};
