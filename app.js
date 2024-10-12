const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');

const STVContractAddress = 'STVContract';
const channelName = 'mychannel';
const chaincodeName = 'stv';

async function main() {
  try {
    const ccpPath = path.resolve(__dirname, '..', 'connection.json');
    const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);

    const identity = await wallet.get('appUser');
    if (!identity) {
      console.log('An identity for the user "appUser" does not exist in the wallet');
      return;
    }

    const gateway = new Gateway();
    await gateway.connect(ccp, {
      wallet,
      identity: 'appUser',
      discovery: { enabled: true, asLocalhost: true },
    });

    const network = await gateway.getNetwork(channelName);
    const contract = network.getContract(chaincodeName, STVContractAddress);

    const candidates = ['Alice', 'Bob', 'Charlie'];
    const quorum = 100;

    console.log('Submitting transaction to initialize the contract...');
    await contract.submitTransaction('STVContract', JSON.stringify(candidates), quorum);
    console.log('Transaction has been submitted');

    console.log('Submitting votes...');
    const vote1 = [0, 1, 2]; // Alice > Bob > Charlie
    const vote2 = [1, 0, 2]; // Bob > Alice > Charlie

    await contract.submitTransaction('vote', JSON.stringify(vote1));
    await contract.submitTransaction('vote', JSON.stringify(vote2));
    console.log('Votes have been submitted');

    console.log('Calculating STV tally...');
    const result = await contract.evaluateTransaction('calculateSTVTally');
    console.log(`STV result: ${result.toString()}`);

    await gateway.disconnect();

  } catch (error) {
    console.error(`Failed to submit transaction: ${error}`);
    process.exit(1);
  }
}

main();