import axios, { AxiosError, AxiosRequestConfig } from "axios";
import * as api from "~/services/api"

const axios_ = axios.create({});

// axios_.interceptors.request.use((config: AxiosRequestConfig) => {
//     const headers = config.headers || {}
//     headers['Authorization'] = localStorage.getItem("access_token") || ""
//     config.headers = headers
//     return config;
// })

const ax = axios.create({})

axios_.interceptors.response.use((response) => {
    return response
}, async function (error) {
    const originalRequest = error.config;
    if (error.response.status === 401 && !originalRequest._retry) {
        originalRequest._retry = true
        const access_token = await api.refreshAccessToken(ax)
        if (access_token == "") {
            window.location.href = "/login"
            return Promise.reject(error);
        }
        // axios.defaults.headers.common['Authorization'] = access_token
        return axios_(originalRequest);
    }
    return Promise.reject(error);
});


export default axios_;