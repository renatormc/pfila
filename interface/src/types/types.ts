export type ErrorsType = { [key: string]: string[] }

export interface ListResponse<T> 
{ 
    per_page: number;
    page: number;
    offset: number;
    total_rows: number;
    total_pages: number;
    items:T[];
}



