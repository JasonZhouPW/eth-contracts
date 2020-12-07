const LockProxy = artifacts.require("./contracts/core/lock_proxy/LockProxy");

module.exports = function(deployer) {
  deployer.deploy(LockProxy);
};
