import React from 'react';
import { Product } from '../types';
import { Table, styled, TableRow, TableCell, TableBody, TableSortLabel, Paper, TableContainer, TablePagination, TableHead, Checkbox } from '@mui/material';
import { tableCellClasses } from '@mui/material/TableCell';

interface ProductTableProps {
    products: Product[];
    handleSort: (column: string) => void;
    sortColumn: string;
    sortDirection: 'asc' | 'desc';
    page : number;
    pageSize: number;
    total: number;
    onPageChanged: (event : unknown, page :number) => void
    onPageSizeChanged: (event: React.ChangeEvent<HTMLInputElement>) => void
}

interface HeadCell { 
    disablePadding: boolean;
    id: string;
    label: string;
    numeric: boolean;
}

const StyledTableCell = styled(TableCell)(({ theme }) => ({
    [`&.${tableCellClasses.head}`]: {
      backgroundColor: '#9197a1',
      color: theme.palette.common.white,
    },
    [`&.${tableCellClasses.head}:hover, &.${tableCellClasses.head} span, &.${tableCellClasses.head} svg`]: {
        fontWeight: "bold",
        color: `${theme.palette.common.white} !important`,
      },
    [`&.${tableCellClasses.body}`]: {
      fontSize: 14,
    },
  }));

const StyledTableRow = styled(TableRow)(({ theme }) => ({
'&:nth-of-type(odd)': {
    backgroundColor: theme.palette.action.hover,
},
// hide last border
'&:last-child td, &:last-child th': {
    border: 0,
},
}));
// non-decimal currency format
const currencyFormat = (value: number) => {
    let options: Intl.NumberFormatOptions | undefined = {
      style: 'currency',
      currency: 'USD',
      minimumFractionDigits: 2,
    };
    
    const formatter = new Intl.NumberFormat('en-US', options);
  
    return formatter.format(value);
  };

const headers : readonly HeadCell[] = [
    {
        id: "id",
        numeric: false,
        disablePadding: false,
        label: "ID"
    },
    {
        id: "name",
        numeric: false,
        disablePadding: false,
        label: "Name"
    },
    {
        id: "category",
        numeric: false,
        disablePadding: false,
        label: "Category"
    },
    {
        id: "description",
        numeric: false,
        disablePadding: false,
        label: "Description",
    },
    {
        id: "price",
        numeric: true,
        disablePadding: false,
        label: "Price"
    },
    {
        id: "stock",
        numeric: true,
        disablePadding: false,
        label: "Stock"
    },
    {
        id: "isActive",
        numeric: false,
        disablePadding: true,
        label: "Active"
    },
    {
        id: "picture",
        numeric: false,
        disablePadding: false,
        label: "Picture"
    },
    {
        id: "discount.status",
        numeric: false,
        disablePadding: false,
        label: "Has discount?"
    },
    {
        id: "discount.value",
        numeric: true,
        disablePadding: false,
        label: "Value"
    }
];
const ProductTable: React.FC<ProductTableProps> = ({ products, handleSort, sortColumn, sortDirection , page, pageSize, total, onPageSizeChanged, onPageChanged}) => { 
    return (
        <Paper sx={{ width: '100%', overflow: 'hidden' }}>
            <TableContainer sx={{ maxHeight: 700 }}>
                <Table size="medium" stickyHeader aria-label="sticky table" >
                    <TableHead>
                        <TableRow>
                            { headers.map ( header => (
                                <StyledTableCell
                                    key={ header.id }
                                    align={'left'}
                                    padding={header.disablePadding ? 'none' : 'normal'}
                                    sortDirection={sortColumn === header.id ? sortDirection : false}
                                > 
                                    <TableSortLabel
                                        active= { sortColumn === header.id }
                                        direction= { sortColumn === header.id ? sortDirection : 'asc'}
                                        onClick={ () => handleSort(header.id)}
                                    >
                                        { header.label}
                                    </TableSortLabel>
                                </StyledTableCell>
                            ))}
                        </TableRow>
                    </TableHead>
                    <TableBody sx={{ overflow: "scroll", maxHeight: "50px" }} >
                        {products?.map(product => (
                            <StyledTableRow key={product.id} hover>
                                <TableCell>{product.id}</TableCell>
                                <TableCell>{product.name}</TableCell>
                                <TableCell>{product.category}</TableCell>
                                <TableCell>{product.description}</TableCell>
                                <TableCell>{ currencyFormat (product.price)}</TableCell>
                                <TableCell>{product.stock}</TableCell>
                                <TableCell>
                                    <Checkbox checked={product.isActive} readOnly />
                                </TableCell>
                                <TableCell><a href={product.picture} target='_blank'> <img src={product.picture} alt={product.name} width={50} /></a></TableCell>
                                <TableCell>
                                    <Checkbox checked={product.discount.status} readOnly />
                                </TableCell>
                                <TableCell> { currencyFormat (product.discount.value )}</TableCell>
                            </StyledTableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
            <TablePagination 
                    rowsPerPageOptions={[ 10, 50, 100]}
                    count={ total}
                    rowsPerPage={ pageSize }
                    page={ page }
                    onPageChange={ onPageChanged }
                    onRowsPerPageChange={ onPageSizeChanged }
            />
        </Paper>

    );
}

export default ProductTable;
