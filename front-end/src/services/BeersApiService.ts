import axios from 'axios';

export const BeersApiService = axios.create({
    baseURL: "https://localhost:8080/api/v1/beers",
    // baseURL: process.env.VUEAPP_BEERS_API_URL,
})