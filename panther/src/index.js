const ganache = require("ganache-core");
const ethers = require("ethers")
const url = "http://localhost:8545";
const provider = new ethers.providers.JsonRpcProvider(url);
var server


provider.on('block', function(blockNumber) {
    server.close();
    console.log("Firing off new local ganache fork @ block #", blockNumber)
    server = ganache.server({fork:"http://0.0.0.0:8545"});
    server.listen(1337);
});

main = () => {
    console.log("Firing off local ganache fork")
    server = ganache.server({fork:"http://0.0.0.0:8545"});
    server.listen(1337);
}

main()