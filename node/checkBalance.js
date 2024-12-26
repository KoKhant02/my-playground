const { ethers } = require("ethers");

const provider = new ethers.JsonRpcProvider("https://sepolia.infura.io/v3/67feff665bfb44ba8a238eb13314f81b");
const walletAddress = "0xFF56795E929e4d70581fd85aDd561883820A8114";

async function checkBalance() {
    const balance = await provider.getBalance(walletAddress);
    console.log("Balance:", ethers.formatEther(balance), "ETH");
}

checkBalance();
