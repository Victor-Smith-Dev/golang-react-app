export interface ResponseList {
  data: Product[];
  status: string;
}

export interface Product {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: null;
  name: string;
  price: number;
}
