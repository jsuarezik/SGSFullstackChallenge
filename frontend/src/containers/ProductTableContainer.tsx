import React, { useState, useEffect } from 'react';
import { Product, ApiResponse } from '../types';
import { fetchProducts } from '../services/api';
import ProductTable from '../components/ProductTable';
import { Container, TextField, Box, InputAdornment, IconButton, Typography } from '@mui/material';

import useDebounce from '../hooks/useDebounce';
import SearchIcon from '@mui/icons-material/Search';

const ProductTableContainer: React.FC = () => {
    const [products, setProducts] = useState<Product[]>([]);
    const [currentPage, setCurrentPage] = useState<number>(1);
    const [pageSize, setPageSize] = useState<number>(10);
    const [total, setTotal] = useState<number>(0);
    const [sortColumn, setSortColumn] = useState<string>('');
    const [sortDirection, setSortDirection] = useState<'asc' | 'desc'>("asc");
    const [filterQuery, setFilterQuery] = useState<string>('');
    const inputValueDebounce = useDebounce(filterQuery, 500);

    useEffect(() => {
        const loadProducts = async () => {
            try {
                const response: ApiResponse = await fetchProducts(currentPage, pageSize, sortColumn, sortDirection, inputValueDebounce);;
                setProducts(response.data);
                setTotal(response.pagination.total_count);
            } catch (error) {
                console.error("Failed to fetch products:", error)
            }
        };

        loadProducts();
    }, [currentPage, pageSize, sortColumn, sortDirection, inputValueDebounce]);

    const handleSort = (column: string) => {
        if (column === sortColumn) {
            setSortDirection(prevDirection => prevDirection === 'asc' ? 'desc': 'asc');
        } else {
            setSortColumn(column);
            setSortDirection('asc');
        }
    }

    const handleChangePage = (event: unknown, newPage: number) => {
        setCurrentPage(newPage);
    };

    const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPageSize(parseInt(event.target.value, 10));
        setCurrentPage(1);
      };

    return (
        <>  
            <Container maxWidth="xl">
                <Box justifyContent={'space-between'} display={'flex'} marginTop={ '10px'} marginBottom={ '10px' }>
                    <Typography variant="h4" textAlign={'left'}>
                        Products List
                    </Typography>
                    <TextField 
                        placeholder="Search by name and category"
                        variant="outlined"
                        value={filterQuery} 
                        onChange={e => setFilterQuery(e.target.value)}
                        InputProps={
                            {
                                endAdornment : (<InputAdornment position="end">
                                    <IconButton
                                        edge="end"
                                        >
                                        <SearchIcon />
                                        </IconButton>
                                </InputAdornment>
                            )}
                        }
                        style={ { minWidth : "300px"}}
                    />
                </Box>
                <ProductTable 
                    products={products} 
                    handleSort={handleSort} 
                    sortColumn={sortColumn} 
                    sortDirection={sortDirection}
                    page={currentPage}
                    pageSize={pageSize}
                    total={total}
                    onPageChanged={ handleChangePage }
                    onPageSizeChanged={ handleChangeRowsPerPage }
                />
            </Container>
        </>
    );
}

export default ProductTableContainer;
