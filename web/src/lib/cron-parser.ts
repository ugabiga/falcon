import cronParser from "cron-parser";

export function parseCronExpression(expression: string) {
    try {
        return cronParser.parseExpression(expression);
    } catch (error) {
        console.error(`Error parsing cron expression: ${expression}`, error);
        throw error;
    }
}

export function nextCronDate(expression: string) {
    const cronDate = parseCronExpression(expression);
    return cronDate.next().toDate();
}

