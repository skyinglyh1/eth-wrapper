pragma solidity >=0.6.0;

interface IPolyWrapper {
    function feeCollector() external view returns (address);
    function lockProxy() external view returns (address);
    function paused() external view returns (bool);

    function lock(
        address fromAsset,
        uint64 toChainId, 
        bytes calldata toAddress,
        uint amount,
        uint fee
    ) external payable;

    function speedUp(
        address fromAsset, 
        bytes calldata txHash,
        uint fee
    ) external payable;
}