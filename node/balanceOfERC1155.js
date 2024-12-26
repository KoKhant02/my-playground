const { Web3 } = require('web3');

const provider = 'https://sepolia.infura.io/v3/67feff665bfb44ba8a238eb13314f81b';
const web3 = new Web3(new Web3.providers.HttpProvider(provider));

const contractAddress = '0x5cDDD02594fcB102662737C45ea51f91ac77e766';

const erc1155Abi = [
  {
    constant: true,
    inputs: [
      { name: "account", type: "address" },
      { name: "id", type: "uint256" }
    ],
    name: "balanceOf",
    outputs: [{ name: "", type: "uint256" }],
    type: "function"
  }
];

const erc1155Contract = new web3.eth.Contract(erc1155Abi, contractAddress);

/**
 * Fetch the ERC1155 token balance for a given address and token ID.
 * @param {string} walletAddress
 * @param {number} tokenId
 * @returns {Promise<string>}
 */
const getTokenBalance = async (walletAddress, tokenId) => {
  try {
    const balance = await erc1155Contract.methods.balanceOf(walletAddress, tokenId).call();
    console.log(`Token Balance for ID ${tokenId}:`, balance);
    return balance;
  } catch (error) {
    console.error('Error fetching balance:', error);
    return null;
  }
};

const walletAddress = '0x0794a4cad4FbA3EAFe0f77Ea56e95E5064ba4965';
const tokenId = 1;

getTokenBalance(walletAddress, tokenId).then(balance => {
  console.log(`Final Balance for Token ID ${tokenId}: ${balance}`);
});