export type ErrorsType = { [key: string]: string[] }

export interface ListResponse<T> {
    per_page: number;
    page: number;
    offset: number;
    total_rows: number;
    total_pages: number;
    items: T[];
}

export type ProcType = "iped" | "ftkimager"

export interface Process {
    id: number;
    type: ProcType;
    name: string;
    user: string;
    createdAt: string;
    start: string;
    startWaiting: string;
    finish: string;
    status: string;
    params?: ProcParams
}

export const DEFAULT_PROCESS: Process = {
    id: 0,
    type: "iped",
    name: "",
    user: "",
    createdAt: "",
    start: "",
    startWaiting: "",
    finish: "",
    status: "ADDED",
} 

export type ProcParams = IpedParams | FtkParams

export interface IpedParams {

}
export interface FtkParams {

}



