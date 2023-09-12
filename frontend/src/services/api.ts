import axios from 'axios';
import { ApiResponse } from '../types';

const BASE_URL = 'http://localhost:8000';  // Replace with your actual API endpoint

export const fetchProducts = async (
    page?: number,
    size? : number,
    sortBy?: string,
    sortOrder?: string,
    q?: string
): Promise<ApiResponse> => {
    try {
        const response = await axios.get(`${BASE_URL}/products`, {
            params: {
                page,
                size,
                sortBy,
                sortOrder: sortOrder === 'asc' ? 1 : -1,
                q
            }
        });
        return response.data;
    } catch (error) {
        throw new Error(`Failed to fetch products: ${error}`);
    }
};
