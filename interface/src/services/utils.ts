import moment from "moment"

export function strToDateTime(date: string): number {
    let date2 = date
    if (!date.includes("T")) {
        date2 = `${date}T12:00:00`
    }
    return Date.parse(date2)
}

export function formatDate(date: string, format?: string): string {
    format = format || "DD/MM/YYYY"
    const date2 = strToDateTime(date)
    const res = moment(date2)
    return res.isValid() ? res.format(format) : ""
}

export function formatDateTime(date: string, format?: string): string {
    format = format || "DD/MM/YYYY HH:mm:ss"
    const date2 = strToDateTime(date)
    const res = moment(date2)
    return res.isValid() ? res.format(format) : ""
}