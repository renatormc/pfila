export type ErrorsType = { [key: string]: string[] }

export type ErrorsTypeObj<T> = {
    [key in keyof T]: string[]
}


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
    created_at: string;
    start: string;
    start_waiting: string;
    finish: string;
    status: string;
    params: ProcParams
}

export type ProcParams = IpedParams | FtkParams | string

export interface IpedParams {
    destination: string;
    sources: string[];
    portable: boolean;
    profile: string;
}

export const DEFAULT_IPED_PARAMS: IpedParams = {
    destination: "",
    sources: [""],
    portable: true,
    profile: ""
}

export interface FtkParams {
    disk: string;
    destination: string;
    verify: boolean;
    format: string;
}

export const DEFAULT_FTK_PARAMS: FtkParams = {
    disk: "",
    destination: "",
    verify: true,
    format: "e01"
}

export function getDefaultProcess(ptype: ProcType): Process {
    let pars: ProcParams
    switch (ptype) {
        case "iped":
            pars = { ...DEFAULT_IPED_PARAMS }
            break
        case "ftkimager":
            pars = { ...DEFAULT_FTK_PARAMS }
            break
    }
    
    const p = {
        id: 0,
        type: ptype,
        name: "",
        user: "",
        created_at: "",
        start: "",
        start_waiting: "",
        finish: "",
        status: "ADDED",
        params: pars
    }
    return p
}



