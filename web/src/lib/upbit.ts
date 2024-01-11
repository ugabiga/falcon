interface Ticker {
    market: string;
    tradeDate: string;
    tradeTime: string;
    tradeDateKst: string;
    tradeTimeKst: string;
    tradeTimestamp: number;
    openingPrice: number;
    highPrice: number;
    lowPrice: number;
    trade_price: number;
    prevClosingPrice: number;
    change: string;
    changePrice: number;
    changeRate: number;
    signedChangePrice: number;
    signedChangeRate: number;
    tradeVolume: number;
    accTradePrice: number;
    accTradePrice24H: number;
    accTradeVolume: number;
    accTradeVolume24H: number;
    highest52WeekPrice: number;
    highest52WeekDate: string;
    lowest52WeekPrice: number;
    lowest52WeekDate: string;
    timestamp: number;
}

export async function getTicker(symbol: string = "BTC"): Promise<Ticker[]> {
    const response = await fetch('https://api.upbit.com/v1/ticker?markets=' + symbol, {
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

