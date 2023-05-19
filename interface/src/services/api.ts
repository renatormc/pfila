import { ListResponse, Process } from '~/types/types';
import axios from './axios'

const parseParams = (proc: Process): Process => {
    proc.params = JSON.parse(proc.params as string)
    return proc
}

export const getResources = async <T>(url: string): Promise<T[]> => {
    const resp = await axios.get<T[]>(url);
    return resp.data;
}

export const getResourcesPaginated = async <T>(url: string, page?: number): Promise<ListResponse<T>> => {
    let fullUrl = `${url}?page=${page || 1}`
    const resp = await axios.get<ListResponse<T>>(fullUrl);
    return resp.data;
}

export const getResource = async <T>(url: string, id: number): Promise<T> => {
    let fullUrl = `${url}/${id}`
    const resp = await axios.get<T>(fullUrl);
    return resp.data;
}


export const deleteResource = async <T>(id: number, url: string): Promise<T> => {
    const resp = await axios.delete<any>(`${url}/${id}`);
    return resp.data;
}

export const updateResource = async <T>(id: number, item: T, url: string): Promise<T> => {
    const resp = await axios.put<T>(`${url}/${id}`, item);
    return resp.data;
}


export const createResource = async <T>(item: T, url: string): Promise<T> => {
    const resp = await axios.post<T>(`${url}`, item);
    return resp.data;
}

export const createProcess = async (item: Process): Promise<Process> => {
    const p = {...item}
    p.params = JSON.stringify(p.params)
    const resp = await axios.post<Process>('/api/proc', p);
    return resp.data;
}

export const updateProcess = async (id: number, item: Process): Promise<Process> => {
    const p = {...item}
    p.params = JSON.stringify(p.params)
    const resp = await axios.put<Process>(`/api/proc/${id}`, p);
    return resp.data;
}

export const deleteProcess = async (id: number): Promise<any> => {
    const resp = await axios.delete<any>(`/api/proc/${id}`);
    return resp.data;
}


export const getProcess = async (id: number): Promise<Process> => {
    const resp = await axios.get<Process>(`/api/proc/${id}`);
    return parseParams(resp.data)
}

export const getProcessess = async (): Promise<Process[]> => {
    const resp = await axios.get<Process[]>("/api/proc");
    const ps = resp.data as Process[]
    for (let index = 0; index < ps.length; index++) {
        ps[index] = parseParams(ps[index])
    }
    return ps;
}

export const queueProcess = async (id: number): Promise<Process> => {
    const resp = await axios.put<Process>(`/api/queue-proc/${id}`);
    console.log(resp.data)
    return parseParams(resp.data)
}

export const stopProcess = async (id: number): Promise<Process> => {
    const resp = await axios.put<Process>(`/api/stop-proc/${id}`);
    return parseParams(resp.data)
}

export const procConsole = async (id: number): Promise<string> => {
    const resp = await axios.get<any>(`/api/proc-console/${id}`);
    return resp.data.console;
}