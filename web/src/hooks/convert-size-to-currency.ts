import React from "react";
import {round} from "@floating-ui/utils";
import {convertNumberToCurrencyStr} from "@/lib/number";
import {getUpbitTicker} from "@/lib/upbit";
import {getBinanceTicker} from "@/lib/binance";

export function useConvertSizeToCurrency() {
    const [convertedTotal, setConvertedTotal] = React.useState("");

    const fetchConvertedTotal = (symbol: string, currency: string, size: number) => {
        switch (currency) {
            case "KRW":
                convertKRWToCurrencyStr(symbol, currency, size).then(setConvertedTotal);
                break;
            case "USDT":
                convertUSDTToCurrencyStr(symbol, currency, size).then(setConvertedTotal);
                break;
            default:
                setConvertedTotal("");
        }
    }

    return {
        fetchConvertedTotal,
        convertedTotal
    }
}

async function convertKRWToCurrencyStr(symbol: string, currency: string, size: number) {
    const data = await getUpbitTicker(currency + "-" + symbol);

    if (!data || data.length === 0) {
        return ""
    }

    const tradePrice = data[0].trade_price;
    const convertedTotal = round(tradePrice * size)
    return convertNumberToCurrencyStr(convertedTotal, 0)
}

async function convertUSDTToCurrencyStr(symbol: string, currency: string, size: number) {
    const data = await getBinanceTicker(symbol + currency);

    if (!data) {
        return ""
    }

    const markPrice = parseFloat(data.markPrice);
    const convertedTotal = round(markPrice * size)
    return convertNumberToCurrencyStr(convertedTotal, 0)
}
