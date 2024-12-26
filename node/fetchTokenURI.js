const { ethers } = require("ethers");

const providerURL = "https://sepolia.infura.io/v3/67feff665bfb44ba8a238eb13314f81b";  // Example: https://sepolia.infura.io/v3/your_project_id

const contractAddress = "0xD0BD2fd822192F195c6F13c7F0a3765EcB596E32";

const contractABI = [
  "function tokenURI(uint256 tokenId) public view returns (string memory)"
];

const tokenId = 0;

async function fetchTokenURI() {
  const provider = new ethers.JsonRpcProvider(providerURL);
  const contract = new ethers.Contract(contractAddress, contractABI, provider);

  try {
    const uri = await contract.tokenURI(tokenId);
    console.log(`Token URI for Token ID ${tokenId}: ${uri}`);
  } catch (error) {
    console.error("Error fetching token URI:", error);
  }
}

fetchTokenURI();
