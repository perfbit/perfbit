import axiosInstance from './axios';

export const signup = async (username: string, password: string) => {
    return axiosInstance.post('/signup', { username, password });
};

export const login = async (username: string, password: string) => {
    return axiosInstance.post('/login', { username, password });
};

export const verify = async (username: string, code: string) => {
    return axiosInstance.post('/verify', { username, code });
};

export const refresh = async (refreshToken: string) => {
    return axiosInstance.post('/refresh', { refreshToken });
};
