const wingWrapper = artifacts.require("Wingwrapper");


module.exports = function(deployer,accounts) {
    deployer.deploy(
        wingWrapper,
        accounts[0],
        "0xD8aE73e06552E270340b63A8bcAbf9277a1aac99",
        "0x875c4d1767d3ec7be759ba583ee7f85fd5f65f85",
        "0xb600c8a2e8852832B75DB9Da1A3A1c173eAb28d8",
        "0xda460071260f6564cb37d33f3b9aeab46f845cac"
    ).then(()=> console.log("wingWrapper:" + wingWrapper.address));


};
