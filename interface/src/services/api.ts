import { AxiosInstance } from 'axios';
import { ResourceItem, ResourceName } from '~/types/resources';
import { LoginResponse, RefreshTokenResponse } from '~/types/responses';
import { ListResponse } from '~/types/types';
import axios from './axios'

export const getResources = async <T extends ResourceItem>(url: ResourceName, page?: number): Promise<ListResponse<T>> => {
    let fullUrl = `/api/${url}?page=${page || 1}`
    const resp = await axios.get<ListResponse<T>>(fullUrl);
    return resp.data;
}

export const getResource = async <T extends ResourceItem>(url: ResourceName, id: number): Promise<T> => {
    let fullUrl = `/api/${url}/${id}`
    const resp = await axios.get<T>(fullUrl);
    return resp.data;
}


export const deleteResource = async <T extends ResourceItem>(item: T, url: ResourceName): Promise<T> => {
    console.log(typeof item)
    const resp = await axios.delete<any>(`/api/${url}/${item.id}`);
    return resp.data;
}

export const updateResource = async <T extends ResourceItem>(item: T, url: ResourceName): Promise<T> => {
    const resp = await axios.put<T>(`/api/${url}/${item.id}`, item);
    return resp.data;
}


export const createResource = async <T extends ResourceItem>(item: T, url: ResourceName): Promise<T> => {
    const resp = await axios.post<T>(`/api/${url}`, item);
    return resp.data;
}

export const login = async (username: string, password: string): Promise<LoginResponse> => {

    const resp = await axios.post<LoginResponse>("/api/login", { username, password });
    return resp.data;
}

export const refreshAccessToken = async (ax: AxiosInstance): Promise<string> => {
    const token = localStorage.getItem("refresh_token")
    if(token == null){
        return ""
    }
    try {
        const resp = await ax.post<RefreshTokenResponse>("/api/refresh-token", { refresh_token: token })
        localStorage.setItem("access_token", resp.data.access_token)
        return resp.data.access_token
    } catch (error) {
        console.log(error)
        return ""
    }
}

// export const caseUp = async (identifier: string): Promise<any> => {
//     const resp = await axios.get<any>("/api/caseup/" + identifier);
//     return resp.data;
// }

export const caseDown = async (identifier: string): Promise<any> => {
    const resp = await axios.get<any>("/api/casedown/" + identifier);
    return resp.data;
}


