export function convertNumberToCryptoSize(value: number, symbol: string): string {
    const decimalPlaces = 5
    return new Intl.NumberFormat('en-US', {
        minimumFractionDigits: decimalPlaces,
    }).format(value) + ' ' + symbol
}
