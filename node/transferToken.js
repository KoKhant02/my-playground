const { ethers } = require("ethers");

async function sendTransaction() {
    const provider = new ethers.JsonRpcProvider("https://sepolia.infura.io/v3/67feff665bfb44ba8a238eb13314f81b");
    const walletPrivateKey = "xxxxx"; 
    const wallet = new ethers.Wallet(walletPrivateKey, provider);

    // Get the gas price from the provider
    const feeData = await provider.getFeeData();
    const gasPrice = feeData.gasPrice;

    const tx = {
        to: "0x6Be0cD257bCdb753b038fD20d1197A5692B02DbF",
        value: ethers.parseUnits("0.0001", "ether"), 
        gasLimit: 21000, 
        gasPrice: gasPrice,
    };

    try {
        const txResponse = await wallet.sendTransaction(tx);
        console.log("Transaction Sent:", txResponse.hash);
    } catch (error) {
        console.error("Transaction failed:", error);
    }
}

sendTransaction();
