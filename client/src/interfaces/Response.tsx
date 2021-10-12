export interface ResponseProductList {
  data: Product[];
  status: string;
}

export interface ResponsePriceList {
  data: Price[];
  status: string;
}

export interface Product {
  ID: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: null;
  name: string;
  price: number;
}

export interface Price {
  ID: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: null;
  ProductId: string;
  price: number;
}