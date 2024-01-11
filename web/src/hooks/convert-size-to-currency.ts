import React from "react";
import {round} from "@floating-ui/utils";
import {convertNumberToCurrencyStr} from "@/lib/number";
import {getTicker} from "@/lib/upbit";

export function useConvertSizeToCurrency() {
    const [convertedTotal, setConvertedTotal] = React.useState("");

    const fetchConvertedTotal = (symbol: string, currency: string, size: number) => {
        switch (currency) {
            case "KRW":
                convertKRWToCurrencyStr(symbol, currency, size).then(setConvertedTotal);
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
    const data = await getTicker(currency + "-" + symbol);

    if (!data || data.length === 0) {
        return ""
    }

    const tradePrice = data[0].trade_price;
    const convertedTotal = round(tradePrice * size)
    return convertNumberToCurrencyStr(convertedTotal, 0)
}