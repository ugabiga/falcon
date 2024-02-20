interface Exchanges {
    value: string;
    supportCurrencies: string[];
    isEnable: boolean;
}

export const ExchangeList: Exchanges[] = [
    {
        value: "upbit",
        supportCurrencies: ["KRW"],
        isEnable: true
    },
    {
        value: "binance_futures",
        supportCurrencies: ["USDT"],
        isEnable: true
    },
    {
        value: "binance_spot",
        supportCurrencies: ["USDT"],
        isEnable: true
    }
];
