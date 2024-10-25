import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8080/v1',
    timeout: 1000,
});

export default api;
