const HDWalletProvider = require("@truffle/hdwallet-provider");

module.exports = {
  compilers: {
    solc: {
      version: "0.6.12",
      settings: {
        optimizer: {
          enabled: true,
          runs: 10000000,
        },
      },
    },
  },
  networks: {
    development: {
      host: "127.0.0.1",
      port: 8545,
      network_id: "*", // Match any network id
    },
    local_testnet: {
      host: "ganache",
      port: 8545,
      network_id: "*", // Match any network id
    },
    mainnet: {
      provider: infuraProvider("mainnet"),
      network_id: 1,
    },
    ropsten: {
      provider: infuraProvider("ropsten"),
      network_id: 3,
    },
  },
  mocha: {
    timeout: 100000, // prevents tests from failing when pc is under heavy load
    reporter: "Spec",
  },
  plugins: ["solidity-coverage"],
};

function infuraProvider(network) {
  return () => {
    return new HDWalletProvider(
      "61a66f7441f1c7150a34bc050be12a680cce9720403d472233ee2fff8ade8951",
      `https://${network}.infura.io/v3/ba2038b90830474080b13f29f1213564`
    );
  };
}