

export interface Discount {
    status: boolean;
    value: number;
}

export interface Product {
    id: string;
    category: string;
    description: string;
    discount: Discount;
    isActive: boolean;
    name: string;
    picture: string;
    price: number;
    stock: number;
}

export interface Pagination {
    current: number;
    page_size: number;
    total_pages: number;
    total_count: number;
}

export interface ApiResponse {
    status: number;
    message: string;
    data: Product[];
    pagination: Pagination;
}
