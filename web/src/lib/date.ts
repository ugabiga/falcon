import { DateTime } from 'luxon';

export function datetimeFromString(datetimeString: string): string {
    let date = DateTime.fromISO(datetimeString, { zone: 'UTC' });
    let seoulDate = date.setZone('Asia/Seoul');

    return seoulDate.toFormat('yyyy-MM-dd HH:mm:ss');
}