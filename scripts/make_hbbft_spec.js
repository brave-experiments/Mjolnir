// for now this is mostly a copy of make_spec.js (which builds for AuRa).
// It's supposed to gradually transition to hbbft only, as the contracts and Parity integration are completed

const fs = require('fs');
const path = require('path');
const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider("https://dai.poa.network"));
const utils = require('./utils/utils');

const VALIDATOR_SET_CONTRACT = '0x1000000000000000000000000000000000000001';
const BLOCK_REWARD_CONTRACT = '0x2000000000000000000000000000000000000001';
const RANDOM_CONTRACT = '0x3000000000000000000000000000000000000001';
const STAKING_CONTRACT = '0x1100000000000000000000000000000000000001';
const PERMISSION_CONTRACT = '0x4000000000000000000000000000000000000001';
const CERTIFIER_CONTRACT = '0x5000000000000000000000000000000000000001';

main();

async function main() {
  const networkName = process.env.NETWORK_NAME;
  const networkID = process.env.NETWORK_ID;
  const owner = process.env.OWNER.trim();
  let initialValidators = process.env.INITIAL_VALIDATORS.split(',');
  for (let i = 0; i < initialValidators.length; i++) {
    initialValidators[i] = initialValidators[i].trim();
  }
  let stakingAddresses = process.env.STAKING_ADDRESSES.split(',');
  for (let i = 0; i < stakingAddresses.length; i++) {
    stakingAddresses[i] = stakingAddresses[i].trim();
  }
  const firstValidatorIsUnremovable = process.env.FIRST_VALIDATOR_IS_UNREMOVABLE === 'true';
  const stakingEpochDuration = process.env.STAKING_EPOCH_DURATION;
  const stakeWithdrawDisallowPeriod = process.env.STAKE_WITHDRAW_DISALLOW_PERIOD;
  const collectRoundLength = process.env.COLLECT_ROUND_LENGTH;
  const erc20Restricted = process.env.ERC20_RESTRICTED === 'true';

  const contracts = [
    'AdminUpgradeabilityProxy',
    'BlockRewardAuRa',
    'Certifier',
    'InitializerAuRa',
    'RandomAuRa',
    'Registry',
    'StakingAuRa',
    'TxPermission',
    'ValidatorSetAuRa',
    'KeyGenHistory'
  ];

  let spec = JSON.parse(fs.readFileSync(path.join(__dirname, '..', 'templates', 'hbbft_dev.json'), 'UTF-8'));

  spec.name = networkName;
  spec.params.networkID = networkID;

  let contractsCompiled = {};
  for (let i = 0; i < contracts.length; i++) {
    const contractName = contracts[i];
    let realContractName = contractName;
    let dir = 'contracts/';

    if (contractName == 'AdminUpgradeabilityProxy') {
      dir = 'contracts/upgradeability/';
    } else if (contractName == 'StakingAuRa' && erc20Restricted) {
      realContractName = 'StakingAuRaCoins';
      dir = 'contracts/base/';
    } else if (contractName == 'BlockRewardAuRa' && erc20Restricted) {
      realContractName = 'BlockRewardAuRaCoins';
      dir = 'contracts/base/';
    }

    console.log(`Compiling ${contractName}...`);
    const compiled = await compile(
      path.join(__dirname, '..', dir),
      realContractName
    );
    contractsCompiled[contractName] = compiled;
  }

  const storageProxyCompiled = contractsCompiled['AdminUpgradeabilityProxy'];
  let contract = new web3.eth.Contract(storageProxyCompiled.abi);
  let deploy;

  // Build ValidatorSetAuRa contract
  deploy = await contract.deploy({data: '0x' + storageProxyCompiled.bytecode, arguments: [
      '0x1000000000000000000000000000000000000000', // implementation address
      owner,
      []
    ]});
  spec.engine.hbbft.params.validators.multi = {
    "0": {
      "contract": VALIDATOR_SET_CONTRACT
    }
  };
  spec.accounts[VALIDATOR_SET_CONTRACT] = {
    balance: '0',
    constructor: await deploy.encodeABI()
  };
  spec.accounts['0x1000000000000000000000000000000000000000'] = {
    balance: '0',
    constructor: '0x' + contractsCompiled['ValidatorSetAuRa'].bytecode
  };

  // Build StakingAuRa contract
  deploy = await contract.deploy({data: '0x' + storageProxyCompiled.bytecode, arguments: [
      '0x1100000000000000000000000000000000000000', // implementation address
      owner,
      []
    ]});
  spec.accounts[STAKING_CONTRACT] = {
    balance: '0',
    constructor: await deploy.encodeABI()
  };
  spec.accounts['0x1100000000000000000000000000000000000000'] = {
    balance: '0',
    constructor: '0x' + contractsCompiled['StakingAuRa'].bytecode
  };

  // Build BlockRewardAuRa contract
  deploy = await contract.deploy({data: '0x' + storageProxyCompiled.bytecode, arguments: [
      '0x2000000000000000000000000000000000000000', // implementation address
      owner,
      []
    ]});
  spec.accounts[BLOCK_REWARD_CONTRACT] = {
    balance: '0',
    constructor: await deploy.encodeABI()
  };
  spec.engine.hbbft.params.blockRewardContractAddress = BLOCK_REWARD_CONTRACT;
  spec.engine.hbbft.params.blockRewardContractTransition = 0;
  spec.accounts['0x2000000000000000000000000000000000000000'] = {
    balance: '0',
    constructor: '0x' + contractsCompiled['BlockRewardAuRa'].bytecode
  };

  // Build RandomAuRa contract
  deploy = await contract.deploy({data: '0x' + storageProxyCompiled.bytecode, arguments: [
      '0x3000000000000000000000000000000000000000', // implementation address
      owner,
      []
    ]});
  spec.accounts[RANDOM_CONTRACT] = {
    balance: '0',
    constructor: await deploy.encodeABI()
  };
  spec.accounts['0x3000000000000000000000000000000000000000'] = {
    balance: '0',
    constructor: '0x' + contractsCompiled['RandomAuRa'].bytecode
  };
  spec.engine.hbbft.params.randomnessContractAddress = RANDOM_CONTRACT;

  // Build TxPermission contract
  deploy = await contract.deploy({data: '0x' + storageProxyCompiled.bytecode, arguments: [
      '0x4000000000000000000000000000000000000000', // implementation address
      owner,
      []
    ]});
  spec.accounts[PERMISSION_CONTRACT] = {
    balance: '0',
    constructor: await deploy.encodeABI()
  };
  spec.params.transactionPermissionContract = PERMISSION_CONTRACT;
  spec.accounts['0x4000000000000000000000000000000000000000'] = {
    balance: '0',
    constructor: '0x' + contractsCompiled['TxPermission'].bytecode
  };

  // Build Certifier contract
  deploy = await contract.deploy({data: '0x' + storageProxyCompiled.bytecode, arguments: [
      '0x5000000000000000000000000000000000000000', // implementation address
      owner,
      []
    ]});
  spec.accounts[CERTIFIER_CONTRACT] = {
    balance: '0',
    constructor: await deploy.encodeABI()
  };
  spec.accounts['0x5000000000000000000000000000000000000000'] = {
    balance: '0',
    constructor: '0x' + contractsCompiled['Certifier'].bytecode
  };

  // Build Registry contract
  contract = new web3.eth.Contract(contractsCompiled['Registry'].abi);
  deploy = await contract.deploy({data: '0x' + contractsCompiled['Registry'].bytecode, arguments: [
      CERTIFIER_CONTRACT,
      owner
    ]});
  spec.accounts['0x6000000000000000000000000000000000000000'] = {
    balance: '0',
    constructor: await deploy.encodeABI()
  };
  spec.params.registrar = '0x6000000000000000000000000000000000000000';

  // Build InitializerAuRa contract
  contract = new web3.eth.Contract(contractsCompiled['InitializerAuRa'].abi);
  deploy = await contract.deploy({data: '0x' + contractsCompiled['InitializerAuRa'].bytecode, arguments: [
      [ // _contracts
        VALIDATOR_SET_CONTRACT,
        BLOCK_REWARD_CONTRACT,
        RANDOM_CONTRACT,
        STAKING_CONTRACT,
        PERMISSION_CONTRACT,
        CERTIFIER_CONTRACT
      ],
      owner, // _owner
      initialValidators, // _miningAddresses
      stakingAddresses, // _stakingAddresses
      firstValidatorIsUnremovable, // _firstValidatorIsUnremovable
      web3.utils.toWei('1', 'ether'), // _delegatorMinStake
      web3.utils.toWei('1', 'ether'), // _candidateMinStake
      stakingEpochDuration, // _stakingEpochDuration
      0, // _stakingEpochStartBlock
      stakeWithdrawDisallowPeriod, // _stakeWithdrawDisallowPeriod
      collectRoundLength // _collectRoundLength
    ]});
  spec.accounts['0x7000000000000000000000000000000000000000'] = {
    balance: '0',
    constructor: await deploy.encodeABI()
  };

  // Build KeyGenHistory contract
  contract = new web3.eth.Contract(contractsCompiled['KeyGenHistory'].abi);
  deploy = await contract.deploy({data: '0x' + contractsCompiled['KeyGenHistory'].bytecode, arguments: [
      ['0x896997c606a0abe1080f2c5535219cbd1c6d81d6'],
      [[0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,181,129,31,84,186,242,5,151,59,35,196,140,106,29,40,112,142,156,132,158,47,223,253,185,227,249,190,96,5,99,239,213,127,29,136,115,71,164,202,44,6,171,131,251,147,159,54,49,1,0,0,0,0,0,0,0,153,0,0,0,0,0,0,0,4,177,133,61,18,58,222,74,65,5,126,253,181,113,165,43,141,56,226,132,208,218,197,119,179,128,30,162,251,23,33,73,38,120,246,223,233,11,104,60,154,241,182,147,219,81,45,134,239,69,169,198,188,152,95,254,170,108,60,166,107,254,204,195,170,234,154,134,26,91,9,139,174,178,248,60,65,196,218,46,163,218,72,1,98,12,109,186,152,148,159,121,254,34,112,51,70,121,51,167,35,240,5,134,197,125,252,3,213,84,70,176,160,36,73,140,104,92,117,184,80,26,240,106,230,241,26,79,46,241,195,20,106,12,186,49,254,168,233,25,179,96,62,104,118,153,95,53,127,160,237,246,41]],
      [[[0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,145,0,0,0,0,0,0,0,4,239,1,112,13,13,251,103,186,212,78,44,47,250,221,84,118,88,7,64,206,186,11,2,8,204,140,106,179,52,251,237,19,53,74,187,217,134,94,66,68,89,42,85,207,155,220,101,223,51,199,37,38,203,132,13,77,78,114,53,219,114,93,21,25,164,12,43,252,160,16,23,111,79,230,121,95,223,174,211,172,231,0,52,25,49,152,79,128,39,117,216,85,201,237,242,151,219,149,214,77,233,145,47,10,184,175,162,174,237,177,131,45,126,231,32,147,227,170,125,133,36,123,164,232,129,135,196,136,186,45,73,226,179,169,147,42,41,140,202,191,12,73,146,2]]]
    ]});
  spec.accounts['0x8000000000000000000000000000000000000000'] = {
    balance: '0',
    constructor: await deploy.encodeABI()
  };

  console.log('Saving hbbft_spec.json file ...');
  fs.writeFileSync(path.join(__dirname, '..', 'hbbft_spec.json'), JSON.stringify(spec, null, '  '), 'UTF-8');
  console.log('Done');
}

async function compile(dir, contractName) {
  const compiled = await utils.compile(dir, contractName);
  const abiFile = `abis/${contractName}.json`;
  console.log(`saving abi to ${abiFile}`);
  fs.writeFileSync(abiFile, JSON.stringify(compiled.abi, null, 2), 'UTF-8');
  return {abi: compiled.abi, bytecode: compiled.evm.bytecode.object};
}

