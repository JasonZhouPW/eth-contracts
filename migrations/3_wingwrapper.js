const wingWrapper = artifacts.require("./contracts/core/wing/Wingwrapper");

module.exports = function(deployer,network,accounts) {
  deployer.deploy(wingWrapper,accounts[0],"0x49206861766520313030e282ac","0xb3EF0C7ee74342Cf1fEa20C474BC7D948B735238","0x49206861766520313030e282ac");
  // deployer.deploy(wingWrapper);
};
