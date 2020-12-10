// const _wingwrapper = require("./5_wingwrapper.js_bak");

// const LockProxy = artifacts.require("./contracts/core/lock_proxy/LockProxy");
// const eccd = artifacts.require("./contracts/core/cross_chain_manager/data/EthCrossChainData.sol");
// const cmcc = artifacts.require("./contracts/core/cross_chain_manager/logic/EthCrossChainManager.sol");
const wingWrapper = artifacts.require("./contracts/core/wing/Wingwrapper");


module.exports = function(deployer,network,accounts) {
  // deployer.deploy(LockProxy)
  //   .then(()=> console.log("lockproxy:" + LockProxy.address));
  // deployer.deploy(eccd)
  //   .then(()=> console.log("eccd:" + eccd.address));

  // deployer.deploy(cmcc,eccd.address)
  //   .then(()=> console.log("cmcc:" + cmcc.address));

  deployer.deploy(wingWrapper,accounts[0],"0xD8aE73e06552E270340b63A8bcAbf9277a1aac99","0x875c4d1767d3ec7be759ba583ee7f85fd5f65f85","0xb600c8a2e8852832B75DB9Da1A3A1c173eAb28d8","0xda460071260f6564cb37d33f3b9aeab46f845cac")
    .then(()=> console.log("wingWrapper:" + wingWrapper.address));

};
