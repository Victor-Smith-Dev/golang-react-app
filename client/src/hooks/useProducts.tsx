import { useEffect, useState } from "react";
import { ReqResApi } from "../api/ReqResapi";
import { ResponseProductList, Product } from "../interfaces/Response";

export const useProducts = () => {
  const [productos, setProductos] = useState<Product[]>([]);

  useEffect(() => {
    loadProducts();
  }, []);

  const loadProducts = async () => {
    const result = await ReqResApi.get<ResponseProductList>("/product");
    setProductos(result.data.data);
  };

  return {
    productos,
  };
};
