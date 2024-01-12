//GET https://fapi.binance.com/fapi/v1/premiumIndex

interface BinanceTicker {
    symbol: string;
    markPrice: string;
    indexPrice: string;
    estimatedSettlePrice: string;
    lastFundingRate: string;
    interestRate: string;
    nextFundingTime: number;
    time: number;
}

export async function getBinanceTicker(symbol: string = ""): Promise<BinanceTicker> {
    const response = await fetch('https://fapi.binance.com/fapi/v1/premiumIndex?symbol=' + symbol, {
        method: 'GET',
        headers: {
            'accept': 'application/json'
        }
    });

    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }

    return await response.json();
}

