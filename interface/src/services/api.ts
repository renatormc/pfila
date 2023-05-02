import { AxiosInstance } from 'axios';
import { ListResponse } from '~/types/types';
import axios from './axios'

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



