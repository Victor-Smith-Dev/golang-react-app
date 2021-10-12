import React from "react";
import { usePricesHistoryScreen } from "../hooks/usePricesHistoryScreen";
import {
  useRouteMatch
} from "react-router-dom";
import { Price } from "../interfaces/Response";

interface MatchParams {
  id: string ;
}

export const HistoryPricesScreen = () => {
  const match = useRouteMatch<MatchParams>('/historyPrices/:id')
  const priceId = match?.params.id!;
  const {prices} = usePricesHistoryScreen(priceId);


  const renderItem = ({ ProductId, price, CreatedAt }: Price) => {
    return (
      <tr key={ProductId.toString()}>
        <td>{ProductId}</td>
        <td>{price}</td>
        <td>{CreatedAt}</td>
      </tr>
    );
  };

  return (
    <>
      <div>
        <h3>Historial de precios del id: <small>{priceId}</small></h3>
          {
            prices.length > 0 && <div style={{
              height: '300px',
              overflowY: 'scroll',
            }}>
              <h3>Precios: ({prices.length})</h3>
              <table className="table">
                <thead>
                  <tr>
                    <th>Producto Id</th>
                    <th>Precio</th>
                    <th>Fecha</th>
                  </tr>
                </thead>
                <tbody>{prices.map((price: Price) => renderItem(price))}</tbody>
              </table>
            </div>
          }
          {
            prices.length === 0 && <div> Sin historial</div> 
          }
      </div>
    </>
  );
}