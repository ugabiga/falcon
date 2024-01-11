export function convertNumberToCryptoSize(value: number, symbol: string): string {
    const decimalPlaces = 5
    return new Intl.NumberFormat('en-US', {
        minimumFractionDigits: decimalPlaces,
    }).format(value) + ' ' + symbol
}

export function convertNumberToCurrencyStr(value: number, minimumFractionDigits: number = 2): string {
    return new Intl.NumberFormat('en-US', {
        minimumFractionDigits: minimumFractionDigits,
    }).format(value)
}
