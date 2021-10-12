import { useEffect, useState } from "react";
import { ReqResApi } from "../api/ReqResapi";
import { ResponseList, Product } from "../interfaces/Response";

export const useProducts = () => {
  const [productos, setProductos] = useState<Product[]>([]);

  useEffect(() => {
    loadProducts();
  }, []);

  const loadProducts = async () => {
    const result = await ReqResApi.get<ResponseList>("/product");
    setProductos(result.data.data);
  };

  return {
    productos,
  };
};
