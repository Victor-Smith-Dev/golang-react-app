import { useEffect, useState } from "react"
import { ReqResApi } from "../api/ReqResapi";
import { ResponsePriceList, Price } from "../interfaces/Response";


export const usePricesHistoryScreen = (id: string | undefined) => {
  const [prices, setPrices] = useState<Price[]>([]);

  useEffect(() => {  
    loadPrices();
  }, [ id ]);

  const loadPrices = async () => {
    const result = await ReqResApi.get<ResponsePriceList>(`/product/${id}/prices`);
    setPrices(result.data.data);
  };

  return {
    prices,
  };
}