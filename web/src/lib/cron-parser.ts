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

export function convertToNextExecutionTime(value: string, failMessage?: string) {
    const result = nextCronDate(value)
    if (result === null) {
        return failMessage || 'No next execution time.';
    }
    return formatDate(result)
}

export function convertHours(value: string): string {
    const result = parseCronExpression(value)
    return result.fields.hour.toString()
}

export function convertDayOfWeek(value: string): string {
    const result = parseCronExpression(value)
    const daysRaw = result.fields.dayOfWeek.toString()
    const daysOfWeek = ['sunday', 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday'];
    const days = daysRaw.split(',').map(Number);
    if (days.includes(0) && days.includes(7)) {
        days.splice(days.indexOf(0), 1);
        days.splice(days.indexOf(7), 1);
        days.push(0);
    }
    if (days.length === 7) {
        return "everyday"
    } else if (days.length === 5 && !days.some(day => day === 0 || day === 6)) {
        return "every_weekday"
    } else {
        return days.map(day => daysOfWeek[day]).join(',')
    }
}

function formatDate(dateTime: Date): string {
    const options: Intl.DateTimeFormatOptions = {
        year: 'numeric',
        month: 'numeric',
        day: 'numeric',
        hour: 'numeric',
        minute: 'numeric',
        hour12: false, // Use 24-hour format
    };

    const formatter = new Intl.DateTimeFormat('ko-KR', options);
    const formattedString = formatter.format(dateTime);

    return formattedString.replace(/(\d+)\/(\d+)\/(\d+), (\d+):(\d+)/, '$3 $4 $5');
}

