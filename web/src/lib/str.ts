export function camelize(str: string) {
    return str.replace(/^\w|[A-Z]|\b\w/g, function (word, index) {
        return index === 0 ? word.toUpperCase() : word.toLowerCase();
    }).replace(/\s+/g, '');
}
export function trim(str: string, length: number) {
    if (str.length <= length) {
        return str;
    }
    return str.trim().slice(0, length) + '...';
}
